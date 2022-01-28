package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func Main() {
	fmt.Println("********** Packege logging -> logging.go **********")

	LoggingSettings("test.log")

	log.Println("logging")
	log.Printf("%T %v \n", "test", "test")

	// 出力して終了
	// log.Fatalf("%T %v \n", "test", "test")

	// Fatallnを使うと終了する
	// _, err := os.Open("kdndmkla")
	// if err != nil {
	// 	log.Fatalln("Exit", err)
	// }

	// log.Fatalln("Error!!")

}
