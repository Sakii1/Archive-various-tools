package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Info2 struct {
	Number int
	Name   string
	Time   string
}

func main() {

	content, err := ioutil.ReadFile("five_txt.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := string(content)
	//data := `[{2545 我TM是个笨B 2023-07-16 22:23:50.976} {2544 雲和海的彼端 2023-07-16 22:23:50.976} {2543 49丨51 2023-07-16 22:23:50.976} {2542 星灵执行官 2023-07-16 22:23:50.976} {2541 浅陌ぐ 2023-07-16 22:23:50.976} {2540 每天几瓶水 2023-07-16 22:23:50.976} {2539 36D狐狐超甜 2023-07-16 22:23:50.976} {2538 星约骑士 2023-07-16 22:23:50.976} {2537 JOJO的鸡蛋饼 2023-07-16 22:23:50.976} {2536 沉没捕鱼 2023-07-16 22:23:50.976}]"`

	// 使用正则表达式提取数字、名称和时间
	re := regexp.MustCompile(`\{(\d+)\s(.*?)\s(\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}.\d{3})\}`)
	matches := re.FindAllStringSubmatch(data, -1)

	// 存储解析后的数据
	var mm []Info2
	for _, match := range matches {
		number, _ := strconv.Atoi(match[1])
		name := match[2]
		time := match[3]

		mm = append(mm, Info2{
			Number: number,
			Name:   name,
			Time:   time,
		})
	}

	// 处理数据
	var uniqueNumbers = make(map[int]string)
	for _, info2 := range mm {
		if t, ok := uniqueNumbers[info2.Number]; !ok || t > info2.Time {
			uniqueNumbers[info2.Number] = info2.Time
		}
	}

	var resultt []Info2
	for number, time := range uniqueNumbers {
		for _, info2 := range mm {
			if info2.Number == number && info2.Time == time {
				resultt = append(resultt, info2)
			}
		}
	}

	// 对结果按照 Number 排序
	sort.Slice(resultt, func(i, j int) bool {
		return resultt[i].Number < resultt[j].Number
	})

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
	txtl := fmt.Sprintf("z_five_txt的排序---%v.txt", eee)
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
	fmt.Println(" ")
	log.Println(" ")
	fmt.Println("                                             补充结果                                                               ")
	log.Println("                                             补充结果                                                               ")
	fmt.Println(" ")
	log.Println(" ")
	var name_k []string
	for _, info2 := range resultt {

		index := strings.Count(strings.Join(name_k, ""), fmt.Sprintf("number:%v,name:%v", info2.Number, info2.Name)) // 统计 "foo" 在字符串中的出现次数
		name_k = append(name_k, fmt.Sprintf("number:%v,name:%v", info2.Number, info2.Name))

		if index != 0 {
			fmt.Println("重复 跳过")
			continue
		}
		//var uid_k []string

		//fmt.Println(info.Name, "宽度:", kuan)
		uid := uuid2(info2.Name)
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
		fmt.Printf("%-30s%-40s%-40s%-50s\n", fmt.Sprintf("编号: %v", info2.Number), fmt.Sprintf("昵称: %v", info2.Name), fmt.Sprintf("UID: %v", uid), fmt.Sprintf("抓取时间: %v", info2.Time))
		log.Printf("%-30s%-40s%-40s%-50s\n", fmt.Sprintf("编号: %v", info2.Number), fmt.Sprintf("昵称: %v", info2.Name), fmt.Sprintf("UID: %v", uid), fmt.Sprintf("抓取时间: %v", info2.Time))

	}

	var p int
	fmt.Printf("")
	fmt.Printf("")
	fmt.Printf("")
	fmt.Printf("five_txt排序结束,在当前目录下可查看重新排序的文件(five_txt + 时间戳)")

	fmt.Scan(&p)
}

func uuid2(name string) int64 {
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
