package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	//VERSION tool version
	VERSION = "1.1.10"

	// URL default wx webhook address
	// URL = "http://in.qyapi.weixin.qq.com/cgi-bin/webhook/send?key="
	URL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key="

	//h help
	h bool
	//v,V version
	v, V bool
	//u wx webhook address
	u string
	//k wx webhook address key
	k string
	//m msg to send
	m string
)

/*
struct MSG
markdown msg
*/
type MSG struct {
	// msg type
	Msgtype string `json:"msgtype"`

	// markdown msg content
	Info Markdown `json:"markdown"`
}

/*
struct Markdown
markdown content
*/
type Markdown struct {
	// content detail info ,markdown type
	Content string `json:"content"`
}

/*
@ConvertJson
Convert msg to json format with markdown
send msg to wx
*/
func ConvertJson(content string) (jsonStr []byte) {
	mymsg := MSG{
		Msgtype: "markdown",
		Info: Markdown{
			Content: content,
		},
	}
	// convert json to []byte
	myjsonStr, _ := json.Marshal(mymsg)
	return myjsonStr
}

/*
@wxMsg
post http req to wx webhook
*/
func wxMsg(key string, msg string) {

	url := URL + key
	payload := strings.NewReader(msg)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(res)
	log.Println(string(body))

}

/*
@usage
edit help msg
*/
func usage() {
	fmt.Println(os.Args[0], `version:`, VERSION, `
Usage: `, os.Args[0], `[OPTIONS]

Options:
	`)
	flag.PrintDefaults()
}

/*
@init
init flag
*/
func init() {
	// set help and version
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	// wx msg webhook addr
	flag.StringVar(&u, "u", URL, "wx msg webhook addr")
	// wx msg webhook addr key
	flag.StringVar(&k, "k", "", "wx msg webhook addr key")
	// msg to send
	flag.StringVar(&m, "m", "", "msg to send")

	//set defalut Usage to usage
	flag.Usage = usage

}

/*
@main
*/
func main() {
	flag.Parse()
	if h || v || V {
		flag.Usage()
		os.Exit(0)
	}
	if len(m) > 0 {

		var cstZone = time.FixedZone("CST", 8*3600) // 8 timezone
		timestamp := time.Now().In(cstZone).Format("2006-01-02 15:04:05")

		info := timestamp + "\n" + m
		if len(k) == 0 {
			fmt.Println("k must be given")
			os.Exit(0)
		}
		fmt.Println(k, string(ConvertJson(info)))
		wxMsg(k, string(ConvertJson(info)))
	} else {
		flag.Usage()
	}

}
