package main;

import (
	"fmt"
)

func main(){
	age := 18
	fmt.Println(true&&true)
	fmt.Println(false&&true)
	fmt.Println(true||false)
	fmt.Println(!true)
	fmt.Println(&age) // age对应的存储空间地址
	var ptr *int= &age
	fmt.Println("ptr这个指针获取的数值:",*ptr)
}