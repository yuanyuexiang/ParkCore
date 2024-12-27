package utils

import (
	"encoding/json"
	"log"
)

func init() {
	/*
		logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		if err != nil {
			panic(err)
		}
		// 组合一下即可，os.Stdout代表标准输出流
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	*/
}

// Println Println
func Println(v interface{}) {
	o, err := json.MarshalIndent(v, "", "    ")
	if err == nil {
		log.Println((string(o)))
	} else {
		log.Println(err)
	}
}

// PrintlnWithTag PrintlnWithTag
func PrintlnWithTag(tag string, v interface{}) {
	o, err := json.MarshalIndent(v, "", "    ")
	if err == nil {
		log.Println(tag + ": " + string(o))
	} else {
		log.Println(tag + ": " + err.Error())
	}
}
