package main;

import (
	"fmt"
)

func main(){
	var count int = 20;
	if count < 30 {
		fmt.Println("库存不足")
	}else{
		fmt.Println("库存充足")
	}
}