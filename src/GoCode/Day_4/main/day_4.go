package main;

// import "fmt";
// import "unsafe";

import (
	"fmt"
	"unsafe"
)

func main(){

	var num int8 = 123
	// Printf函数的作用就是: 格式化,将num的类型填充到%T的位置
	fmt.Printf("num的数据类型: %T\n",num)

	fmt.Println(unsafe.Sizeof(num))
}