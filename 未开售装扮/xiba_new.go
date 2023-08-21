package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var listt = make([]interface{}, 0)

type Item struct {
	Name            string
	Zbid            int
	Raw             string
	Quan            string
	Sheng           int64
	Sale_time_begin string
	Bilitime        string
	Saletime        int64
}

func main() {

	list()

	var items []Item
	for i := 0; i < len(listt); i += 7 {
		name := listt[i].(string)
		zbid := listt[i+1].(int64)
		raw := listt[i+2].(string)
		quan := listt[i+3].(string)
		sheng := listt[i+4].(int64)
		//shengValue := listt[i+4]
		//var sheng int64
		//switch shengValue.(type) {
		//case int:
		//	sheng = shengValue.(int64)
		//case float64:
		//	sheng = int64(shengValue.(float64))
		//}
		sale_time_begin := listt[i+5].(string)
		bilitime := listt[i+6].(string)
		saletime, _ := strconv.ParseInt(sale_time_begin, 10, 64)
		item := Item{name, int(zbid), raw, quan, sheng, sale_time_begin, bilitime, saletime}
		items = append(items, item)
	}
	// ...

	sort.Slice(items, func(i, j int) bool {
		return items[i].Saletime < items[j].Saletime
	})

	log.SetFlags(0)

	// 打开 output.txt 文件，并设置为可写模式
	txtl := fmt.Sprintf("new.txt")
	f, err := os.Create(txtl)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 设置输出为文件
	log.SetOutput(f)
	log.Println("            【当前未开售装扮】")
	log.Println("  ")
	for _, item := range items {
		rraw, _ := strconv.Atoi(item.Raw)

		ccpf, xg, stock := cpf(item.Zbid)
		var minNum int64

		//fmt.Println(stock)

		if stock >= item.Sheng {
			minNum = stock
		} else {
			minNum = item.Sheng
		}

		fmt.Printf("  名称: %s\n", item.Name)
		fmt.Printf("  售价: %d\n", rraw/100)
		fmt.Printf("  数量: %v/%v \n", minNum, item.Quan)
		fmt.Printf("  限购：%v\n", xg)
		fmt.Printf("  装扮id: %d\n", item.Zbid)
		fmt.Printf("  出品方: %v\n", ccpf)

		log.Printf("  名称: %s\n", item.Name)
		log.Printf("  售价: %d\n", rraw/100)
		log.Printf("  数量: %v/%v \n", minNum, item.Quan)
		log.Printf("  限购：%v\n", xg)
		log.Printf("  装扮id: %d\n", item.Zbid)
		log.Printf("  出品方: %v\n", ccpf)

		//fmt.Printf("  直达链接: https://www.bilibili.com/h5/mall/suit/detail?id=%v\n", item.Zbid)
		yuyue(item.Zbid)
		fmt.Printf("  %s\n", item.Bilitime)
		log.Printf("  %s\n", item.Bilitime)

		fmt.Println("")
		log.Println("")
	}

	var p int
	fmt.Printf("")
	fmt.Scan(&p)

}

func list() {
	for ii := 0; ii < 10; ii++ {
		clientt := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com/x/garb/v2/mall/partition/item/list?group_id=0&part_id=6&pn=%v&ps=50&sort_type=2", ii), nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("authority", "api.bilibili.com")
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

		respp, _ := clientt.Do(req)

		defer respp.Body.Close()
		boodyText, err := io.ReadAll(respp.Body)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 50; i++ {
			name, _ := jsonparser.GetString(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "name")
			zbid, _ := jsonparser.GetInt(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "item_id")
			raw, _ := jsonparser.GetString(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "properties", "sale_bp_forever_raw")
			quan, _ := jsonparser.GetString(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "properties", "sale_quantity")
			sheng, _ := jsonparser.GetInt(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "sale_surplus")
			saletime, _ := jsonparser.GetInt(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "sale_left_time")
			sale_time_begin, _ := jsonparser.GetString(boodyText, "data", "list", fmt.Sprintf("[%v]", i), "properties", "sale_time_begin")

			//qua, i := jsonparser.GetString(boodyText, "data", "list", fmt.Sprintf("[%v]", 1))
			//fmt.Println(qua, i)
			//fmt.Println(raw)

			if saletime > 0 {

				//fmt.Println("名称：", name)
				//
				//fmt.Println("装扮id：", zbid)
				//
				//fmt.Println("全部数量：", quan)
				//fmt.Println("剩余数量：", sheng)
				//
				//fmt.Println("开售倒计时：", saletime)

				now, _ := strconv.ParseInt(sale_time_begin, 10, 64)

				t := time.Unix(now, 0) // 转换为time.Time

				year := t.Year()     // 年
				month := t.Month()   // 月
				day := t.Day()       // 日
				hour := t.Hour()     // 时
				minute := t.Minute() // 分
				second := t.Second() // 秒

				bilitime := fmt.Sprintf("开售时间:  %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
				//fmt.Printf("开售时间:  %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
				//fmt.Println(" ")
				//fmt.Println(" ")

				listt = append(listt, name, zbid, raw, quan, sheng, sale_time_begin, bilitime) // int
			}
		}

	}
}
func yuyue(itid int) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com/x/garb/user/reserve/state?item_id=%v&part=suit", itid), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.57")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	yy, _ := jsonparser.GetInt(bodyText, "data", "reserve_count")
	fmt.Printf("  预约数: %v\n", yy)
	log.Printf("  预约数: %v\n", yy)

}

func cpf(iidd int) (s, ss string, sss int64) { //出品方
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com/x/garb/mall/item/suit/v2?item_id=%v&part=suit", iidd), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.95 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not-A.Brand";v="24", "Chromium";v="14"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
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
	sb, _ := jsonparser.GetString([]byte(string(bodyText)), "data", "fan_user", "nickname")
	xiangou, _ := jsonparser.GetString(bodyText, "data", "item", "properties", "sale_buy_num_limit") //限购检测

	var limit int64
	if strings.Contains(string(bodyText), "item_stock_surplus") {
		lim, _ := jsonparser.GetString(bodyText, "data", "item", "properties", "item_stock_surplus")
		limit, _ = strconv.ParseInt(lim, 10, 64)

	} else {
		limit = 99999
	} //实际售卖量

	return sb, xiangou, limit

}
