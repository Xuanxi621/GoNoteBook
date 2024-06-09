package main;

import "fmt";

func test01(num int){
	fmt.Println(num)
}

tyo

func test02(num1 int,num2 int, testFunc func(int)){
	fmt.Println("==== TEST 02 ========")
}

func main(){
	a := test01
	fmt.Printf("a的数据类型是: %T,test01函数的数据类型: %T\n",a,test01)

	a(10)

	test02(10,3,test01)
	test02(11,3,a)

	// 自定义数据类型: （相当于Python as ）别名 但是在Go编译时认为myInt和int不是一种数据类型
	type myInt int
	var num1 myInt = 30
	fmt.Println("num1:",num1)
}