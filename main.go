package artLogger

import (
	"flag"
	"log"
	"os"
	"path"
)

var (
	// LogError Log error var - used to log in errors.log
	LogError *log.Logger
	// AppLog log application events - logs to app.log
	AppLog *log.Logger
	// QueueLog log queue events - logs to queue.log
	QueueLog *log.Logger
	// AccessLog log requests
	AccessLog *log.Logger
)

func FatalLog(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	if _, err := os.Stat("logs"); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("logs", 0755)
			FatalLog(err)
		}
	}
	flag.Parse()
	var file, err = os.OpenFile(path.Join("logs", "error.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	FatalLog(err)
	LogError = log.New(file, "", log.LstdFlags|log.Lshortfile)
	var AppFile, err1 = os.OpenFile(path.Join("logs", "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	FatalLog(err1)
	AppLog = log.New(AppFile, "", log.LstdFlags|log.Lshortfile)
	var QueueFile, err2 = os.OpenFile(path.Join("logs", "queue.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	FatalLog(err2)
	QueueLog = log.New(QueueFile, "", log.LstdFlags|log.Lshortfile)
	var AccessFile, err3 = os.OpenFile(path.Join("logs", "access.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	FatalLog(err3)
	AccessLog = log.New(AccessFile, "", log.LstdFlags|log.Lshortfile)
}
