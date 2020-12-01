package main

import (
	"context"
	"fmt"
	"log"

	"Go-000/week02/01/service"
)

func main() {
	data, err := service.DoSomething(context.Background())
	if err != nil {
		fmt.Printf("catch err: %+v", err) // 输出详细栈信息
		return
	}

	log.Printf("data: %v", data)
}
