package main

import (
	"fmt"
	"onlinejudgeForward/mq"
)

func main() {
	mq.RabbitConsume()
	fmt.Println("hello world")
	/*channel.StartCatch()*/
}
