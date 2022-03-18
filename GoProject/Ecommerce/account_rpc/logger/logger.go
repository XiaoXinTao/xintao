package logger

import (
	"io"
	"log"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	flag       = log.Ldate | log.Ltime | log.Lshortfile
	preDebug   = "[DEBUG]"
	preInfo    = "[INFO]"
	preWarning = "[WARNING]"
	preError   = "[ERROR]"
)

var (
	logFile        io.Writer
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	defaultLogFile = "/var/log/account_rpc2.log"
)

func InitKlog() {
	logFile, err := os.OpenFile(defaultLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("create log file err %+v", err)
	}
	klog.SetOutput(logFile)
}
