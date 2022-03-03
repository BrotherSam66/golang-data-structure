package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

// 函数
func convert(tableName string) {
	filePath := "./models/" + tableName + ".go"
	err := converter.NewTable2Struct().
		SavePath(filePath).
		Dsn("ferry:ferry|A1@tcp(localhost:3417)/ferry?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		Table(tableName).
		Run()
	fmt.Println(err)
}

func main() {
	params := []string{
		"f_mission",
		"f_mission_file",
		"f_mission_user",
		"f_file",
		"f_mission_action_log",
		"f_mission_his",
		"f_mission_file_his",
		"f_mission_user_his",
		"f_dict_group",
		"f_dict_item",
	}
	for _, str := range params {
		convert(str)
	}

}
