package main;

import "fmt";

var(
	n7 = 123
)

func main(){
	// 第一种: 变量的使用方式 指定变量的类型,并且赋值
	var num1 int = 18
	fmt.Println(num1)
	// 第二种: 指定变量类型,但不进行赋值,使用默认值
	var num2 int
	fmt.Println(num2)
	// 第三种: 如果没有写变量类型,那么根据 后面的值进行推断判断变量的类型
	var num3 = 30
	fmt.Println(num3) 
	// 第四种: 省略var,注意 := 不能写成 =
	sex := "男"
	fmt.Println(sex)

	fmt.Println("==========================")
	var n1,n2,n3 int
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4,name,n5 = 10,"Jack",7.8
	fmt.Println(n4)
	fmt.Println(name)
	fmt.Println(n5)

	fmt.Println(n7)


}