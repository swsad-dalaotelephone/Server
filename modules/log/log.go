package log

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

//log file
func init() {
	//open log file
	dname := "storage/logs"
	os.MkdirAll(dname, os.ModeDir|os.ModePerm)
	errFile, err := os.OpenFile(dname+"/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("cannot open log file", err)
	}

	infoFile, err := os.OpenFile(dname+"/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("cannot open log file", err)
	}
	// set log format
	InfoLog = log.New(infoFile, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(errFile, "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
