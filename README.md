# send msg to wx by webhook

Supports:

- markdown () MarkDown type msg

## Installation


``` shell
git clone https://github.com/yeziyitao/wxmsg.git
make
```

## Quickstart

``` shell
yezi$ ./wxmsg -m "hello world" -k Your-wx-webhook-key
Your-wx-webhook-key {"msgtype":"markdown","markdown":{"content":"2020-05-16 11:12:06\nhello world"}}
2020/05/16 11:12:06 &{200 OK 200 HTTP/2.0 2 0 map[Content-Length:[27] Content-Type:[application/json; charset=UTF-8] Date:[Sat, 16 May 2020 03:12:06 GMT] Error-Code:[0] Error-Msg:[ok] Server:[nginx]] {0xc000122160} 27 [] false false map[] 0xc0000fa000 0xc000309b80}
2020/05/16 11:12:06 {"errcode":0,"errmsg":"ok"}
``` 
``` shell
yezi$ cat info.mk 
实时新增用户反馈<font color="warning">1321例</font>，请相关同事注意from file
> 类型:<font color="comment">用户反馈</font>[详情](http://work.weixin.qq.com/api/doc)
> 普通用户反馈:<font color="comment">1117例</font>
> VIP用户反馈:<font color="comment">115例</font>
``` 
``` shell
yezi$ ./wxmsg -m "`cat info.mk`" -k Your-wx-webhook-key
Your-wx-webhook-key {"msgtype":"markdown","markdown":{"content":"2020-05-16 11:12:58\n实时新增用户反馈\u003cfont color=\"warning\"\u003e1321例\u003c/font\u003e，请相关同事注意from file\n\u003e 类型:\u003cfont color=\"comment\"\u003e用户反馈\u003c/font\u003e[详情](http://work.weixin.qq.com/api/doc)\n\u003e 普通用户反馈:\u003cfont color=\"comment\"\u003e1117例\u003c/font\u003e\n\u003e VIP用户反馈:\u003cfont color=\"comment\"\u003e115例\u003c/font\u003e"}}
2020/05/16 11:12:58 &{200 OK 200 HTTP/2.0 2 0 map[Content-Length:[27] Content-Type:[application/json; charset=UTF-8] Date:[Sat, 16 May 2020 03:12:58 GMT] Error-Code:[0] Error-Msg:[ok] Server:[nginx]] {0xc00026fce0} 27 [] false false map[] 0xc00018a000 0xc000411810}
2020/05/16 11:12:58 {"errcode":0,"errmsg":"ok"}
``` 
