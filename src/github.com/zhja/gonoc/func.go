package gonoc

import (
    //"fmt"
    "strings"
    "time"
    "os"
)

//string
//字符串截取
func Substr(str string, start, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0

    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length

    if start > end {
        start, end = end, start
    }

    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}

//首字母大写 小写
func StringFC(str string) (s string) {
    strSlice := strings.Split(str, "")
    for key, val := range strSlice {
        if key == 0 {
            strSlice[key] = strings.ToUpper(val)
        }
    }
    s = strings.Join(strSlice, "")
    return
}

//判断slice中值是否存在(一维)
func ExistsSV(slices []string, value string) bool {
    var t = false
    for _, val := range slices {
        if val == value {
            t = true
            break
        }
    }
    return t
}

//判断slice中值是否存在(二维)
func ExistsSVT(slices [][]string, value string) (t bool, keys int) {
    for key, val := range slices {
        if val[0] == value {
            t = true
            keys = key
            break
        }
    }
    return
}

//获取当前时间
func MqTime() string {
    //t := time.Now()
    //return t.Format("2006-01-02 15:04:02")
    timestamp := time.Now().Unix()
    //格式化为字符串,tm为Time类型
    tm := time.Unix(timestamp, 0)
    return tm.Format("02/01/2006 15:04:05 PM")
}

func MqUnix() int {
    //t := time.Now()
    //return t.Format("2006-01-02 15:04:02")
    timestamp := time.Now().Unix()
    return int(timestamp)
}

func WriteFile(file string, note string) {
    userFile := file
    fout,err := os.Create(userFile)
    defer fout.Close()
    if err != nil {
        //fmt.Println(userFile,err)
        return
    }
    fout.WriteString(note + "\r\n")
    //fout.Write([]byte("Just a test!\r\n"))
}