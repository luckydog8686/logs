package logs
import (
	log "github.com/sirupsen/logrus"
	"os"
)

var Trace=log.Trace
var Debug = log.Debug
var Info = log.Info
var Warn = log.Warn
var Error = log.Error
var Fatal = log.Fatal


func init() {
	if logLevel,err :=log.ParseLevel(os.Getenv("LOG_LEVEL"));err !=nil{
		log.SetLevel(log.InfoLevel)
	}else{
		log.SetLevel(logLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:true,
		FullTimestamp:true,
		TimestampFormat:"2006-01-02T15:04:05.000000000Z07:00",
		ForceQuote:true,
	})
	log.SetReportCaller(true)
}