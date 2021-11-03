package main
//package awesomeProject

import (
		"fmt"
)



//声明变量 必须用空格隔开 ，如 var age int;
//变量部分
//变量首字母不能为数字
func var_test1() {
	var a string = "var test"
	var b, c int = 1, 2
	fmt.Println(a,b,c)

	//当没有指定变量的值，即没有初始化，变量默认为零
	var stin = "TEST1"
	fmt.Println("有值",stin)
	var no_int int
	fmt.Println("没有值的int",no_int)
	var no_bool bool
	fmt.Println("没有值的判断",no_bool)

	//简写方法，不需要再直接 var去定义
	intVal := 1
	stringVal := "test"
	fmt.Println(intVal,stringVal)
}






func main() {
	var_test1()
	//fmt.Println("this is test!")
	//var scode = 123
	//var enddate = "3030-12-12"
	//var url = "Code=%d&date=%s"
	//fmt.Sprintf() 用于格化式字符串
	//var target_url = fmt.Sprintf(url,scode,enddate)
	//fmt.Println(target_url)
}




