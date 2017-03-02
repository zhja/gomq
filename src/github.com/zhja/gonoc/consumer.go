package gonoc

import (
	"io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "encoding/json"
    "time"
    "strconv"
    "os/exec"
    //"reflect"
)

var MQC Consumer

type Consumer struct {
}



//进程的创建、消费、守护、调度等
//任务状态 0=接口返回错误值 1=处理成功 2=待发送 3=接口未调用成功 4=队列中没有数据
func (this *Consumer) ProcessOne(key string) {
	t_time := MqUnix()
    l_time, _ := MQ.Lrange(key + "Monitor", 0, 0)
    var ll_time int
    for _, t := range l_time {
		ll_time, _ = strconv.Atoi(string(t.([]byte)))
    }
    if t_time - ll_time >= 30 {
    	cmd := exec.Command("/bin/sh", "-c", "nohup go run mqone.go -key " + key + " 1> runtime/logs/mqone.out 2> runtime/logs/mqone.err &")
		_, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
    }
}

func (this *Consumer) BranchOne(key string) {
	defer func() {
        e := recover()
        if e != nil {
            _ = NocEmail("MQ异常提醒", e.(string))
        }
    }()

    //判断进程是否存在
    //当前时间
	for _, v := range MQW.Config {
		if v["name"] == key {
			b, _ := strconv.Atoi(v["sleep"])
			ch := make(chan string)
		    go func() {
			    for {
			    	time.Sleep(time.Second * time.Duration(b))
			    	status := this.Execute(key)
					if status != 4 {
			    		ch <- key
			    	}
			    }
		    }()

		    tick := time.NewTimer(time.Second * 30)
		    ForStart:
		    select {
		        case msg := <-ch:
		            fmt.Println(msg)
		            tick.Reset(time.Second * 30)
		            goto ForStart
		        case <-tick.C:
		            fmt.Println("channel timeout close " + key)
		            close(ch)
		            goto ForEnd
		    }
		    ForEnd:
		    fmt.Println("done")
    	}
    }
}

func (this *Consumer) branch(key string, sleep string) *chan string { // 每个分支开出一个goroutine做计算并把计算结果流入各自信道
    ch := make(chan string, 10)
    go func() {
        b, _ := strconv.Atoi(sleep)
	    for {
	    	time.Sleep(time.Second * time.Duration(b))
	    	ch <- key
	    }
    }()
    return &ch
}

//所有
func (this *Consumer) Process() {
	defer func() {
        e := recover()
        if e != nil {
            _ = NocEmail("MQ异常提醒", e.(string))
        }
    }()

	//获取队列分类，每个分类创建一个进程
  	ch := make(chan string, 10)//, len(MQW.Config)
  	chs := make([]chan string, 0)
	for _, v := range MQW.Config {
		chs = append(chs, *this.branch(v["name"], v["sleep"]))
    }

    for i := 0 ; i < len(chs); i++ {//select会尝试着依次取出各个信道的数据
	    go func(key chan string) { // 开一个goroutine监视各个信道数据输出并收集数据到信道ch
	        for {
	            select {
	            	case v1 := <- key:
	            		ch <- v1
	            		_ = this.Execute(v1)
	            }
	        }
		}(chs[i])
	}

    for {
    	select {
    		case msg := <-ch:
        		fmt.Println(msg)
    	}
	}
}


func (this *Consumer) Execute(key string) (status int) {
	t := time.Now()
	data_info,_ := MQ.Pop(key)
	if data_info == nil {
		//队列中没有数据
		status = 4
	} else {
		data := *data_info
		//获取接口详细信息
		api_info := MQW.Config[data["type_name"]]

		param := this.PostParams(data["data"])

		//调用接口:
		if api_info["api"] != "" {
			i := 0;
			ApiReturnStart:
			rs, err := this.Post(api_info["api"], param)
			//返回值格式，json {"status":1, "msg":"回写接口的参数，无回写返回空"} {"status":0, "msg":"错误信息"}
			if err != nil {
		        CheckError(err)
		        //生成error case
		    }else{
				var apiReturn map[string]interface{}
		        json.Unmarshal(rs, &apiReturn)
		        if len(apiReturn) > 0 {
					status_id := -1
					if apiReturn["status"].(bool) {
						status_id = 1
					}
					data_id, _ := strconv.Atoi(data["id"])
			        msg, _ := json.Marshal(apiReturn["msg"])

			        if status_id == -1 {
			   			//失败以后尝试链接3次
			   			if i < 3 {
			   				time.Sleep(time.Second * 3)
			   				i++
			   				goto ApiReturnStart
			   			}
			        	//保存mysql回调数据
			        	var update_mq_list mq_list
			        	update_mq_list.Id = data_id
						update_mq_list.Callback_data = string(msg)
						update_mq_list.Update_time = t.Format("2006-01-02 15:04:02")
						update_mq_list.Status_id = status_id
						Db.Save(&update_mq_list)

			        	status = -1
			        	_ = NocEmail("MQ提醒", "Error:" + string(msg) + ", URL:" + api_info["api"])
			        	//生成error case
			        	//不会循环调用测试
			            //fmt.Printf("接口返回result字段是:\r\n%v", apiReturn["msg"])
			        } else if status_id == 1 { 
			        	status = 1
			        	//回调成功以后数据从未处理列表中剪切到完成列表中
			        	//获取数据信息
			        	//保存完成表数据mq_list_done
			        	var add_mq_list_done mq_list_done
						add_mq_list_done.Type_name = data["type_name"]
						add_mq_list_done.Data = data["data"]
						add_mq_list_done.Create_time = data["create_time"]
						add_mq_list_done.Update_time = t.Format("2006-01-02 15:04:02")
						add_mq_list_done.Status_id = status_id
						add_mq_list_done.Callback_data = string(msg)
						Db.Save(&add_mq_list_done)
			        	//删除未处理表数据
			        	var delete_mq_list mq_list
			        	delete_mq_list.Id = data_id
						Db.Delete(&delete_mq_list)

			        	// if api_info["callback_api"] != "" {
			        	// 	_, _ = this.Post(api_info["callback_api"], apiReturn["msg"])
			        	// }
						//callback_api不为空，调用回调函数
						//.........下一版本支持
			        }
			    }
		    }
		}
		MQ.SimplePop(key + "Monitor")
		MQ.Lpush(key + "Monitor", strconv.Itoa(MqUnix()))
	}
	//status = 1

	return
}

func (this *Consumer) Get(apiURL string, params url.Values) (rs[]byte, err error) {
    var Url *url.URL
    Url, err = url.Parse(apiURL)
    if err != nil{
        fmt.Printf("解析url错误:\r\n%v",err)
        return nil,err
    }
    //中文参数进行URLEncode
    Url.RawQuery = params.Encode()
    resp, err := http.Get(Url.String())
    if err != nil {
        fmt.Println("err:",err)
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
  
func (this *Consumer) Post(apiURL string, params url.Values) (rs[]byte, err error) {
    resp, err := http.PostForm(apiURL, params)
    if err != nil {
    	_ = NocEmail("MQ异常提醒", "Post Error:" + err.Error() + ", URL:" + apiURL)
        return nil, err
    }
    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}

func (this *Consumer) PostParams(data string) url.Values {
    param := url.Values{}
    param.Set("data", data)
    return param
}