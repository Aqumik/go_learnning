package main

var x,y int
var (	//声明全局变量
		//全局变量允许声明但不使用
	a int
	b bool


)

var c,d int = 1, 2
var e,f = 123, "hello"


//var aaa, bb int
//var cc string




func main()  {
	//这种不带声明格式的只能在函数体中出现
	// g,h :=123 , "hello"
	//函数哪变量声明了以后需要使用，否则会报错
	g, h := 123, "hello"
	//多变量可以在同一行进行赋值
	aa, bb, cc := 5, 7, "abc"
	println(x,y,a,b,c,d,e,f,g,h,aa,bb,cc)
}
