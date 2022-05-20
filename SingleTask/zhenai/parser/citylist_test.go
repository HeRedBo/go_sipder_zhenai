package parser

import (
	"GoSpider/SingleTask/fetcher"
	"GoSpider/SingleTask/zhenai/model"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contens , err := fetcher.Fetch("http://www.zhenai.com/zhenghun/akesu")
	//contens , err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}



	//filePath := "citylist_test_data.html"

	//var InstallRe2 = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	var InstallRe = `window.__INITIAL_STATE__=(.*?)</script>`

	matches := regexp.MustCompile(InstallRe).FindAllSubmatch(contens,-1)
	fmt.Println(len(matches))
	var json_str string
	for _,m := range matches {
		json_str = string(m[1])
	}
	json_str2  := strings.Replace(json_str,";(function(){var s;(s=document.currentScript||document.scripts[document.scripts.length-1]).parentNode.removeChild(s);}());","",1)
	fmt.Println(json_str2)
	data := model.CityDetailList{}
	if err := json.Unmarshal([]byte(json_str2),&data); err == nil {
		fmt.Println(data.MemberListData)

	} else {
		fmt.Println("err",err)
	}





	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Println("文件打开失败", err)
	//}
	////及时关闭file句柄
	//defer file.Close()
	////写入文件时，使用带缓存的 *Writer
	//write := bufio.NewWriter(file)
	//write.WriteString(string(contens))
	////Flush将缓存的文件真正写入到文件中
	//write.Flush()


}
