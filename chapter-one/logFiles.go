package main


import (
	"fmt"
	"log"
	"log/syslog"
	"path/filepath"
	"os"
)


func main(){
	progName := filepath.Base(os.Args[0])
	fmt.Println(progName)
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, progName)
	if err != nil{
		log.Fatal(err.Error())
	} else {
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")

	sysLog, err = syslog.New(syslog.LOG_MAIL, "this is not it")
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
	}
}