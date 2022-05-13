package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contens , err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contens , err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}


	fmt.Printf("%s \n", contens)

	//filePath := "citylist_test_data.html"



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
