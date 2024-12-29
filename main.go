package main

import (
	clienthttp "ali.go/go-tools/client_http"
	logger "ali.go/go-tools/logger"
	"ali.go/go-tools/osdetector"
)

func main() {
	logger.BeginWriteLogDetails()
	// Initialiser le logger depuis le package logger
	//logDetails := logger.InitLogger("client_http.log", "D:\\03.TechMania\\04.Dev\\ali.go\\")
	logDetails := logger.InitLogger("goGetMyPubIP.log", osdetector.GetLogPath())

	//url := "http://192.168.118.110:9003"
	//clienthttp.ReqHTTP(url, logDetails)
	url := "https://ifconfig.co"
	clienthttp.ReqHTTP(url, logDetails)
	logger.EndWriteLogDetails()
}
