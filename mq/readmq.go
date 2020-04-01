package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"onlinejudgeForward/slflog"
	"time"
)

var solution []RunCode

func init() {
	solution = make([]RunCode, 20)
}

func GetSolution() []RunCode {
	return solution
}

type RunCode struct {
	Eid string
	Pid string
	Uid string
	Language string
	Data string
}

func (RunCode *RunCode) addRunCode(body string) {
	err := json.Unmarshal([]byte(body), RunCode)
	slflog.FatalErr(err, "Unmarshal err")
	solution = append(solution, *RunCode)
}

func RabbitConsume() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	slflog.FatalErr(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	channel, err := conn.Channel()
	slflog.FatalErr(err, "Failed to open a channel")
	defer channel.Close()
	channel.ExchangeDeclare("exec_dir_c", "direct", true, false, false, false, nil)
	channel.QueueDeclare("exec_queue_c", true, false, false, false, nil)
	channel.QueueBind("exec_queue_c", "exec.run.c", "exec_dir_c", false, nil)
	msgs, err := channel.Consume("exec_queue_c", "rabbit", false, false, false, false, nil)
	slflog.FatalErr(err, "get msg err")
	ch := make(chan int)
	defer close(ch)
	go func() {
		for d := range msgs {
			msg := d.Body
			slflog.Info(string(msg))
			code := new(RunCode)
			code.addRunCode(string(msg))
			time.Sleep(2 *  time.Second)
			d.Ack(false)
		}
		ch<-10
	}()
	<-ch
}
