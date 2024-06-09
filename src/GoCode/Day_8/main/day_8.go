package main

import (
	"fmt"
)

// 首字母大写该函数可以被本包文件和其他文件使用 类似 Public
// 首字母小写只能被本包文件使用,其他包文件不能使用 Private
// 形参列表: 可以是无参数 也可以多参数
// 如果返回值不想接受,可以使用_进行忽略

func addNum(n1 int,n2 int) (int){
	sum := 0
	sum += n1
	sum += n2
	return sum
}

// agrs...int 可以传入多个任务int参数
func test (args...int){
	// 函数内容处理可变参数的时候,可变参数当作切片处理
	for i :=0 ; i< len(args); i++ {
		fmt.Println(args[i])
	}
}

func main(){
	num1 := 1
	num2 := 2
	fmt.Printf("Sum : %d",addNum(num1,num2))

	test(12,222,334,1441)
}