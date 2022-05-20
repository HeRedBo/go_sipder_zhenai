package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseMemberListProfile(t *testing.T) {

	bytes , err := ioutil.ReadFile("/Users/hehongbo/www/GO/go_spider/SingleTask/zhenai/parser/merber_list.json")
	if err != nil {
		fmt.Println("读取 json 文件报错",err)
		return
	}
	resp := ParseMemberListProfile(bytes,"")
	fmt.Println(resp)

}
