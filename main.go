package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	// Setup logging
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("err: fail to create log file: %s", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	var screenId int = 0
	if len(os.Args) >= 2 {
		id, _ := strconv.ParseInt(os.Args[1], 10, 32)
		screenId = int(id)
	}
	count := screenshot.NumActiveDisplays()
	if screenId >= count {
		log.Fatalf("err: invalid screen id: %d", screenId)
	}

	filePath := ""
	if len(os.Args) >= 3 {
		filePath = os.Args[2]
	} else {
		now := time.Now()
		datetimeText := now.Format("20060102_150405")
		filePath = fmt.Sprintf("%s%03d.png", datetimeText, now.Nanosecond()/1000000)
	}

	bounds := screenshot.GetDisplayBounds(screenId)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Fatalf("err: fail to capture screen: %d", screenId)
	}
	log.Printf("info: capture screen: %d, size: %dx%d, file: %s", screenId, bounds.Dx(), bounds.Dy(), filePath)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("err: fail to create file: %s\n", filePath)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("err: fail to encode image: %s\n", filePath)
	}
}
