package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/buger/jsonparser"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	Number int
	Name   string
	Time   string
}

var m []Info

var str_off string

//var num_off int64

func main() {
	//ii := 114103501
	ii, _ := strconv.Atoi(fmt.Sprintf("%v", ReadLine(2)))
	ms, _ := strconv.Atoi(fmt.Sprintf("%v", ReadLine(5)))

	now_, name, zb_gid := id_iddd(ii)
	now, _ := strconv.ParseInt(now_, 10, 64)
	fmt.Printf("name:%v   开售时间:%v\n", name, now)

	if now == 0 {
		fmt.Println("似乎config第二行的装扮ID填错了")
		time.Sleep(time.Second * 1000)
	}
	biii()

	end := ReadLine(4)
	if end == "2" || end == "3" {
		fmt.Println("初始值为2或3  不被允许  修改成1即可...可通过修改config来改动程序")
		time.Sleep(time.Second * 200000)

	}

	fmt.Printf("间隔: %vms\n", ms)

	//tm := time.Unix(now, 0)

	nowwww(now)

	rank(ii, ms)

	var uniqueNumbers = make(map[int]string)
	for _, info := range m {
		if t, ok := uniqueNumbers[info.Number]; !ok || t > info.Time {
			uniqueNumbers[info.Number] = info.Time
		}
	}

	var result []Info
	for number, time := range uniqueNumbers {
		for _, info := range m {
			if info.Number == number && info.Time == time {
				result = append(result, info)
			}
		}
	}

	// 对结果按照 Number 排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Number < result[j].Number
	})

	// 输出结果

	noww := time.Now()

	// 提取年、月、日、时、分、秒信息
	year := noww.Year()
	month := int(noww.Month())
	day := noww.Day()
	hour := noww.Hour()
	minute := noww.Minute()
	second := noww.Second()

	// 输出结果
	eee := fmt.Sprintf("%04d年%02d月%02d日 %02d.%02d.%02d", year, month, day, hour, minute, second)

	log.SetFlags(0)

	// 打开 output.txt 文件，并设置为可写模式
	txtl := fmt.Sprintf("%v---%v.txt", name, eee)
	f, err := os.Create(txtl)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 设置输出为文件
	log.SetOutput(f)

	// 进行日志输出
	//log.Print("Hello, World!")
	//log.Print("Hello, rld!")
	//log.Print("Hello, ld!")

	var name_k []string
	for _, info := range result {

		index := strings.Count(strings.Join(name_k, ""), fmt.Sprintf("number:%v,name:%v", info.Number, info.Name)) // 统计 "foo" 在字符串中的出现次数
		name_k = append(name_k, fmt.Sprintf("number:%v,name:%v", info.Number, info.Name))

		if index != 0 {
			fmt.Println("重复 跳过")
			continue
		}
		//var uid_k []string

		//fmt.Println(info.Name, "宽度:", kuan)
		uid := uuid(info.Name)

		if uid == 0 {
			continue
		}

		//kuan := 20 - utf8.RuneCountInString(info.Name)                 //宽度
		//kua := 20 - utf8.RuneCountInString(strconv.FormatInt(uid, 10)) //宽度
		//
		//for k := 0; k < kuan; k++ { //名字宽度调节
		//	name_k = append(name_k, "")
		//}
		//
		//for kk := 0; kk < kua; kk++ { //UID宽度调节
		//	uid_k = append(uid_k, "")
		//}
		//fmt.Printf(" 编号:%v 昵称：%v %vUID:%v %v抓取时间:%v\n", info.Number, info.Name, strings.Join(name_k, " "), uid, strings.Join(uid_k, " "), info.Time)
		//fmt.Printf(" 编号:%v 昵称：%v %vUID:%v %v抓取时间:%v\n", info.Number, info.Name, strings.Join(name_k, " "), uid, strings.Join(uid_k, " "), info.Time)

		//fmt.Printf("编号: %d%20s 昵称：%s%30s UID:%v%20s 抓取时间:%s%20s\n", info.Number, "", info.Name, "", uid, "", info.Time, "")
		//log.Printf("编号: %d%20s 昵称：%s%30s UID:%v%20s 抓取时间:%s%20s\n", info.Number, "", info.Name, "", uid, "", info.Time, "")
		//fmt.Printf("%-30s%-40s%-40s%-50s\n", fmt.Sprintf("编号: %v", info.Number), fmt.Sprintf("昵称: %v", info.Name), fmt.Sprintf("UID: %v", uid), fmt.Sprintf("抓取时间: %v", info.Time))
		//log.Printf("%-30s%-40s%-40s%-50s\n", fmt.Sprintf("编号: %v", info.Number), fmt.Sprintf("昵称: %v", info.Name), fmt.Sprintf("UID: %v", uid), fmt.Sprintf("抓取时间: %v", info.Time))

		fmt.Printf("0,%v,%v,https://space.bilibili.com/%v\n", zb_gid, info.Number, uid)
		log.Printf("0,%v,%v,https://space.bilibili.com/%v\n", zb_gid, info.Number, uid)

	}

	fmt.Printf("txt抓取结束,在当前目录下可查看录制文件(%v---%v.txt)\n", name, eee)
	time.Sleep(time.Second * 1)
	fmt.Printf("\n")
	fmt.Printf("将txt转换成scv中...\n")

	file, err := os.Open(fmt.Sprintf("%v", txtl))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建CSV文件
	csvFile, err := os.Create(fmt.Sprintf("%v---%v.csv", name, eee))
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// 创建CSV写入器
	writer := csv.NewWriter(csvFile)

	// 读取txt文件的内容并逐行处理
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 通过指定的分隔符将每行内容分隔成字段
		line := scanner.Text()

		// 检测并删除引号
		if strings.Contains(line, `"`) {
			line = strings.ReplaceAll(line, `"`, "")
		}

		fields := strings.Split(line, ",")

		// 将字段写入CSV文件
		err := writer.Write(fields)
		if err != nil {
			panic(err)
		}
	}

	// 检查文件读取和写入过程中是否有错误
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	// 刷新写入缓冲并检查错误
	writer.Flush()
	if err := writer.Error(); err != nil {
		panic(err)
	}
	fmt.Printf("转换完成,txt和scv文件已生成在当前工作目录下(%v---%v.scv)", name, eee)

	var p int
	fmt.Printf("")
	fmt.Scan(&p)

}

//for {
//	a := ReadLine(1)
//	fmt.Println(a)
//	time.Sleep(time.Second * 2)
//}

func rank(ii int, ms int) {

	//infoo := Info{
	//	Number: int(10000),
	//	Name:   "test",
	//	Time:   time.Now().Format("2006-01-02 15:04:05.000"),
	//}
	//
	//fmt.Println(infoo)
	//
	////fmt.Println(name, ":", numberr, "---", tioem)
	//m = append(m, infoo)

	for {

		end := ReadLine(4)
		if end == "2" {
			fmt.Println("已暂停...")
			time.Sleep(time.Second * 2)
			continue
		}

		if end == "3" {
			fmt.Println("结束抓取---生成log中...")
			break
		}

		time.Sleep(time.Duration(ms) * time.Millisecond) //自定义延迟

		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com/x/garb/rank/fan/recent?item_id=%v&spm_id_from=333.1035.rich-text.link.click", ii), nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("authority", "api.bilibili.com")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		req.Header.Set("cache-control", "max-age=0")
		req.Header.Set("cookie", fmt.Sprintf("%v", ReadLine(3)))
		req.Header.Set("sec-ch-ua", `"Chromium";v="110", "Not A(Brand";v="24", "Microsoft Edge";v="110"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
		req.Header.Set("sec-fetch-dest", "document")
		req.Header.Set("sec-fetch-mode", "navigate")
		req.Header.Set("sec-fetch-site", "none")
		req.Header.Set("sec-fetch-user", "?1")
		req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.41")
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
		if len(bodyText) < 70 {
			fmt.Println("空")
			continue
		}
		count := strings.Count(string(bodyText), "nickname")
		tioem := time.Now().Format("2006-01-02 15:04:05.000")

		if str_off == string(bodyText) {
			fmt.Println("榜单无变动 跳过存储处理    <-")
			for i := 0; i < count; i++ {
				name, _ := jsonparser.GetString(bodyText, "data", "rank", fmt.Sprintf("[%v]", i), "nickname")
				numberr, _ := jsonparser.GetInt(bodyText, "data", "rank", fmt.Sprintf("[%v]", i), "number")

				info := Info{
					Number: int(numberr),
					Name:   name,
					Time:   tioem,
				}

				fmt.Println(info)

				//fmt.Println(name, ":", numberr, "---", tioem)
				//m = append(m, info)
				//m = append(m, strconv.FormatInt(numberr, 10)+"**"+name+tioem)
				//fmt.Println(numberr)
			}

		} else {
			fmt.Println("榜单有变动 进行存储    <-")
			str_off = string(bodyText)
			for i := 0; i < count; i++ {
				name, _ := jsonparser.GetString(bodyText, "data", "rank", fmt.Sprintf("[%v]", i), "nickname")
				numberr, _ := jsonparser.GetInt(bodyText, "data", "rank", fmt.Sprintf("[%v]", i), "number")

				info := Info{
					Number: int(numberr),
					Name:   name,
					Time:   tioem,
				}

				fmt.Println(info)

				//fmt.Println(name, ":", numberr, "---", tioem)
				m = append(m, info)
				//m = append(m, strconv.FormatInt(numberr, 10)+"**"+name+tioem)
				//fmt.Println(numberr)
				if i == count-1 {
					log.SetFlags(0)

					// 打开 output.txt 文件，并设置为可写模式
					txtl := fmt.Sprintf("five_txt.txt")
					f, err := os.Create(txtl)
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()

					// 设置输出为文件
					log.SetOutput(f)
					log.Println(m)
					log.SetOutput(os.Stdout)
				}

			}

		}

	}
	//fmt.Println(m)

}

func ReadLine(lineNumber int) string {
	file, _ := os.Open("config.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	defer file.Close()
	return ""
} //读文本

func nowwww(now int64) { //服务器时间戳
	for {
		client := &http.Client{}
		var data = strings.NewReader(`
`)
		req, err := http.NewRequest("GET", fmt.Sprintf("http://api.bilibili.com/x/report/click/now"), data)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer func(Body io.ReadCloser) {
			er := Body.Close()
			if er != nil {

			}
		}(resp.Body)
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		miao, err := jsonparser.GetInt(bodyText, "data", "now")

		if miao >= now-1 {

			break
		} else {
			fmt.Printf("off:%v\n", miao)

			continue
		}

	}

} //服务器时间戳

func id_iddd(ii int) (sale_time_begin string, name string, gid int64) {
	client := &http.Client{}
	var data = strings.NewReader(`
`)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com/x/garb/mall/item/suit/v2?item_id=%d&part=suit", ii), data) //必须修改id
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		er := Body.Close()
		if er != nil {

		}
	}(resp.Body)
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	nnname, err := jsonparser.GetString(bodyText, "data", "item", "name") //该装扮名称

	sale_time_begin, err = jsonparser.GetString(bodyText, "data", "item", "properties", "sale_time_begin") //该装扮开售时间戳

	gid, _ = jsonparser.GetInt(bodyText, "data", "suit_items", "space_bg", "[0]", "item_id") //该装扮开售时间戳

	//fmt.Printf("%v\n", name)
	//log.Printf("%v\n", name)
	return sale_time_begin, nnname, gid
}

func uuid(name string) int64 {
	//fmt.Println("传入的昵称： ", name)
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.vc.bilibili.com/dynamic_mix/v1/dynamic_mix/name_to_uid?names=%v", url.QueryEscape(name)), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "l=v")
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
	//fmt.Printf("%s\n", bodyText)
	uid, _ := jsonparser.GetInt(bodyText, "data", "uid_list", "[0]", "uid")

	return uid
}

func biii() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/nav", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.bilibili.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36")
	req.Header.Set("origin", "https://account.bilibili.com")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("referer", "https://account.bilibili.com/account/home")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cookie", fmt.Sprintf("%v", ReadLine(3)))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)

	}

	code, err := jsonparser.GetInt(bodyText, "code")
	if code == 0 {

		fmt.Printf("cookie验证通过\n")
		fmt.Printf("-\n")

		//fmt.Printf("%s\n", bodyText)
	} else {
		var t int
		fmt.Printf("cookie未通过,请重新抓取\n")
		fmt.Scan(&t)
	}

} //验证cookie是否有效   以及检测b币
