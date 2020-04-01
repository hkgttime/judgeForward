package slflog

import "log"

func FatalErr(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Info(body string)  {
	log.Printf("msg: %s\n", body)
}

func Debug(msg interface{})  {
	log.Printf("msg: %s\n", msg)
}
