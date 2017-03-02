package gonoc

import (
    "html/template"
    //"net/http"
    "fmt"
)

func DeferRecover() {
    defer func() {
        e := recover()
        if e != nil {
            fmt.Println("出现问题：", e)
        }
    }()
}    

func CheckError(err error) {
    if err != nil {
        panic(err.Error())
    }
}

func CheckCustomError(value string) {
    panic(value)
}

var tpl = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<title>404错误</title>

<style type="text/css">
body,code,dd,div,dl,dt,fieldset,form,h1,h2,h3,h4,h5,h6,input,legend,li,ol,p,pre,td,textarea,th,ul{margin:0;padding:0}
body{font:14px/1.5 'Microsoft YaHei','微软雅黑',Helvetica,Sans-serif;min-width:1200px;background:#f0f1f3;}
:focus{outline:0}
h1,h2,h3,h4,h5,h6,strong{font-weight:700}
a{color:#428bca;text-decoration:none}
a:hover{text-decoration:underline}
.error-page{background:#f0f1f3;padding:80px 0 180px}
.error-page-container{position:relative;z-index:1}
.error-page-main{position:relative;background:#f9f9f9;margin:0 auto;width:617px;-ms-box-sizing:border-box;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box;padding:50px 50px 70px}
.error-page-main:before{content:'';display:block;background:url(img/error.png);height:7px;position:absolute;top:-7px;width:100%;left:0}
.error-page-main h3{font-size:24px;font-weight:400;border-bottom:1px solid #d0d0d0}
.error-page-main h3 strong{font-size:54px;font-weight:400;margin-right:20px}
.error-page-main h4{font-size:20px;font-weight:400;color:#333}
.error-page-actions{font-size:0;z-index:100}
.error-page-actions div{font-size:14px;display:inline-block;padding:30px 0 0 10px;width:50%;-ms-box-sizing:border-box;-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box;color:#838383}
.error-page-actions ol{list-style:decimal;padding-left:20px}
.error-page-actions li{line-height:2.5em}
.error-page-actions:before{content:'';display:block;position:absolute;z-index:-1;bottom:17px;left:50px;width:200px;height:10px;-moz-box-shadow:4px 5px 31px 11px #999;-webkit-box-shadow:4px 5px 31px 11px #999;box-shadow:4px 5px 31px 11px #999;-moz-transform:rotate(-4deg);-webkit-transform:rotate(-4deg);-ms-transform:rotate(-4deg);-o-transform:rotate(-4deg);transform:rotate(-4deg)}
.error-page-actions:after{content:'';display:block;position:absolute;z-index:-1;bottom:17px;right:50px;width:200px;height:10px;-moz-box-shadow:4px 5px 31px 11px #999;-webkit-box-shadow:4px 5px 31px 11px #999;box-shadow:4px 5px 31px 11px #999;-moz-transform:rotate(4deg);-webkit-transform:rotate(4deg);-ms-transform:rotate(4deg);-o-transform:rotate(4deg);transform:rotate(4deg)}
</style>

</head>
<body>

<div class="error-page">
    <div class="error-page-container">
        <div class="error-page-main">
            <h3>
                <strong>404</strong>无法打开页面
            </h3>
            <div class="error-page-actions">
                <div>
                    <h4>可能原因：</h4>
                    <ol>
                        <li>网络信号差</li>
                        <li>找不到请求的页面</li>
                        <li>输入的网址不正确</li>
                    </ol>
                </div>
                <div>
                    <h4>可以尝试：</h4>
                    <ul>
                        <li><a href="http://www.17sucai.com/">返回首页</a></li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>

</body>
</html>
`

// render default application error page with error and stack string.
func ShowErr() {
    t, _ := template.New("test").Parse(tpl)
    // data := map[string]string{
    //     "AppError":      fmt.Sprintf("%s:%v", BConfig.AppName, err),
    //     "RequestMethod": ctx.Input.Method(),
    //     "RequestURL":    ctx.Input.URI(),
    //     "RemoteAddr":    ctx.Input.IP(),
    //     "Stack":         stack,
    //     "BeegoVersion":  VERSION,
    //     "GoVersion":     runtime.Version(),
    // }
    //ctx.ResponseWriter.WriteHeader(500)
    t.Execute(Requests.W, nil)
}

// show 401 unauthorized error.
// func unauthorized(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(401),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested can't be authorized." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br>The credentials you supplied are incorrect" +
//         "<br>There are errors in the website address" +
//         "</ul>")
//     t.Execute(rw, data)
// }

// // show 402 Payment Required
// func paymentRequired(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(402),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested Payment Required." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br>The credentials you supplied are incorrect" +
//         "<br>There are errors in the website address" +
//         "</ul>")
//     t.Execute(rw, data)
// }

// // show 403 forbidden error.
// func forbidden(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(403),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested is forbidden." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br>Your address may be blocked" +
//         "<br>The site may be disabled" +
//         "<br>You need to log in" +
//         "</ul>")
//     t.Execute(rw, data)
// }

// // show 404 notfound error.
// func notFound(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(404),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested has flown the coop." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br>The page has moved" +
//         "<br>The page no longer exists" +
//         "<br>You were looking for your puppy and got lost" +
//         "<br>You like 404 pages" +
//         "</ul>")
//     t.Execute(rw, data)
// }

// // show 405 Method Not Allowed
// func methodNotAllowed(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(405),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The method you have requested Not Allowed." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br>The method specified in the Request-Line is not allowed for the resource identified by the Request-URI" +
//         "<br>The response MUST include an Allow header containing a list of valid methods for the requested resource." +
//         "</ul>")
//     t.Execute(rw, data)
// }

// // show 500 internal server error.
// func internalServerError(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(500),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested is down right now." +
//         "<br><br><ul>" +
//         "<br>Please try again later and report the error to the website administrator" +
//         "<br></ul>")
//     t.Execute(rw, data)
// }

// // show 501 Not Implemented.
// func notImplemented(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(504),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested is Not Implemented." +
//         "<br><br><ul>" +
//         "<br>Please try again later and report the error to the website administrator" +
//         "<br></ul>")
//     t.Execute(rw, data)
// }

// // show 502 Bad Gateway.
// func badGateway(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(502),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested is down right now." +
//         "<br><br><ul>" +
//         "<br>The server, while acting as a gateway or proxy, received an invalid response from the upstream server it accessed in attempting to fulfill the request." +
//         "<br>Please try again later and report the error to the website administrator" +
//         "<br></ul>")
//     t.Execute(rw, data)
// }

// // show 503 service unavailable error.
// func serviceUnavailable(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(503),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested is unavailable." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br><br>The page is overloaded" +
//         "<br>Please try again later." +
//         "</ul>")
//     t.Execute(rw, data)
// }

// // show 504 Gateway Timeout.
// func gatewayTimeout(rw http.ResponseWriter, r *http.Request) {
//     t, _ := template.New("beegoerrortemp").Parse(errtpl)
//     data := map[string]interface{}{
//         "Title":        http.StatusText(504),
//         "BeegoVersion": VERSION,
//     }
//     data["Content"] = template.HTML("<br>The page you have requested is unavailable." +
//         "<br>Perhaps you are here because:" +
//         "<br><br><ul>" +
//         "<br><br>The server, while acting as a gateway or proxy, did not receive a timely response from the upstream server specified by the URI." +
//         "<br>Please try again later." +
//         "</ul>")
//     t.Execute(rw, data)
// }

