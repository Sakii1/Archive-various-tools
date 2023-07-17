package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var room int64 = 6154037 //输入房间号  这也找不到建议速速remake  可修改
var number1 int = 3      //有三个同样的弹幕  将会复制此弹幕发送   可修改
var switchh = 1          //是否开启重复模式？ 1为开启  2为关闭(可以发送之前发送的弹幕,不过在人少的直播间会显的很奇怪)  可修改

var time1 = 5  //多少秒检测一次弹幕   可修改  单位秒
var time2 = 30 //发送弹幕后休息的时间   可修改   单位秒

var cookie1 string = "SESSDATA=51c522ae%2C1704897927%2Ccd00f972qQEpgWiY6wRVb_XTBWdnBZAz7w8BKafhCW7DjqVp0L9ds8_SGZ3WJ9Vs3VVLgqeMFiDVswAAYgA;bili_jct=6837b577c021aaa648a3d624acee9c5e;DedeUserID=3494379283024098;DedeUserID__ckMd5=f557baad743bcd7b;sid=fbhblv61"

//cookie获取教程 ：  https://blog.csdn.net/weixin_53891182/article/details/125846559  ,别把Cookie：复制进来

//因个人爱好  以下内容不会跳过重复： ？,呃呃,好好好.   可自行在下方删除或添加

var quchong []string

func main() {
	var csrf string
	// ps:chatGPT真好用~  ↓

	splits := strings.Split(cookie1, "; ") // 按照分号和空格分割字符串

	for _, s := range splits {
		if strings.HasPrefix(s, "bili_jct=") { // 判断是否以 "bili_jct=" 开头
			value := strings.TrimPrefix(s, "bili_jct=") // 去掉前缀获取值
			//fmt.Println(value)
			csrf = value
			break // 找到后立即跳出循环
		}
	}

	//fmt.Println(csrf)
	for {
		fmt.Println("检测中...")
		danmu(csrf)
		time.Sleep(time.Duration(time1) * time.Second)

	}

	//fasong(csrf)
}

//wggg.Add(2)
//
//danmu()
//
//wggg.Wait()

func danmu(csrf string) {
	var ddanmu []string
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.live.bilibili.com/ajax/msg?roomid=%v", room), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.live.bilibili.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	//miao, err := jsonparser.GetString(bodyText, "data")
	for i := 0; i < 10; i++ {
		text := gjson.Parse(string(bodyText)).Get(fmt.Sprintf("data.room.%v.text", i)).String()
		//uid := gjson.Parse(string(bodyText)).Get(fmt.Sprintf("data.room.%v.uid", i)).Int()
		//ts := gjson.Parse(string(bodyText)).Get(fmt.Sprintf("data.room.%v.check_info.ts", i)).Int()
		//nickname := gjson.Parse(string(bodyText)).Get(fmt.Sprintf("data.room.%v.nickname", i)).String()
		ddanmu = append(ddanmu, text)
		//fmt.Println(text)

	}
	concatenated := strings.Join(ddanmu, " ") //弹幕列表转化string
	//fmt.Println("-------")
	//fmt.Println("-------")
	//fmt.Println("-------")
	for ii := 0; ii < 10; ii++ {
		str := ddanmu[ii]                         //逐个取出
		count := strings.Count(concatenated, str) //计数重复数量
		if count >= number1 {
			fmt.Println("发现重复弹幕:", str, "---数量:", count)
			fasong(csrf, str, concatenated)

		}

	}

}

//fmt.Println(ddanmu)
//time.Sleep(time.Second * 5) //几秒一次？  底边调10也行 反正没弹幕   弹幕很快的就调0吧

//i := 0

func fasong(csrf string, str string, concatenated string) {
	//con := strings.Join(quchong, " ") //去重转化string
	if switchh == 1 {
		//fmt.Println("去重长度", len(quchong)) //测试

		if quchong != nil {
			for c := 0; c < len(quchong); c++ {
				strr := quchong[c]                          //逐个取出
				countt := strings.Count(concatenated, strr) //计数是否发送过该弹幕

				if countt == 0 {
				} else {
					fmt.Println(strr, "<---已经发送过该弹幕  跳过")
					return
				}
			}
		}
	}
	for o := 0; o < len(str); o++ {
		k := string(str[o])

		if k == "2" || k == "3" || k == "4" || k == "5" || k == "6" || k == "7" || k == "8" || k == "9" {
			fmt.Println("检测到数字 跳过")
			return
		}
	}

	fas := str //弹幕

	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf("------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"bubble\"\r\n\r\n0\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"msg\"\r\n\r\n%v\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"color\"\r\n\r\n16777215\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"mode\"\r\n\r\n1\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"fontsize\"\r\n\r\n25\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"rnd\"\r\n\r\n%v\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"roomid\"\r\n\r\n%v\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"csrf\"\r\n\r\n%v\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf\r\nContent-Disposition: form-data; name=\"csrf_token\"\r\n\r\n%v\r\n------WebKitFormBoundaryWQthq5CRuHi2eUGf--\r\n", fas, time.Now().Unix(), room, csrf, csrf))
	req, err := http.NewRequest("POST", "https://api.live.bilibili.com/msg/send", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.live.bilibili.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "multipart/form-data; boundary=----WebKitFormBoundaryWQthq5CRuHi2eUGf")
	req.Header.Set("cookie", fmt.Sprintf("%v", cookie1))
	req.Header.Set("origin", "https://live.bilibili.com")
	req.Header.Set("referer", fmt.Sprintf("https://live.bilibili.com/%v?live_form=73001", room))
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	code := gjson.Parse(string(bodyText)).Get("code").Int()
	if code == 0 {
		fmt.Println("-------")
		fmt.Println("-------")
		fmt.Println(fas, "---发送成功")
		if fas == "?" || fas == "？" || fas == "呃呃" || fas == "好好好" || fas == "???" || fas == "？？？" {

		} else {
			quchong = append(quchong, fas)
		}

		fmt.Printf("%v秒后开启下一次跟风", time2)
		fmt.Println("-------")
		fmt.Println("-------")
		fmt.Println("-------")
		time.Sleep(time.Duration(time2) * time.Second)

	} else {
		fmt.Printf("%s\n", bodyText)
		fmt.Println("发生了错误  已停止程序")
		time.Sleep(time.Second * 100000)
	}
}
