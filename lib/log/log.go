package log

import (
	"log"
	"log/slog"
	"os"
)

// var logfile = "logs/log"
// var errorLog *os.File
var Logger *slog.Logger
var Log *log.Logger

//var err error

func init() {
	//errorLog, err = os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	fmt.Printf("error opening file: %v", err)
	//	os.Exit(1)
	//}

	Log = log.New(os.Stdout, "CC:", log.Ldate|log.Ltime|log.Lshortfile)

	h0 := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: new(slog.LevelVar)})

	Logger = slog.New(h0)

}
