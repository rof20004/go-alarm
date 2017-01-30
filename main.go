package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {

	alarmTime := os.Args[1]

	if len(alarmTime) != 5 {
		log.Fatalln("Invalid time. Correct format: 13:01")
	}

	s := strings.Split(alarmTime, ":")
	hour, _ := strconv.Atoi(s[0])
	minute, _ := strconv.Atoi(s[1])
	timeNow := time.Now()
	alarmDate := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hour, minute, 0, 0, time.Local)

	if alarmDate.Before(timeNow) {
		log.Fatalln("Alarm time error, needs to be after time now")
	}

	fmt.Println("Alarm setted to:", alarmDate)

	for {

		if time.Now().After(alarmDate) {
			break
		}

	}

	switch runtime.GOOS {
	case "windows":
		playOnWindows("alarm.mp3")
	case "linux":
		playOnLinux("alarm.mp3")
	default:
		fmt.Println("O.S not supported")
	}

}

func playOnLinux(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	cmd := exec.Command("mpg123", fileName)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	playOnLinux(fileName)

}

func playOnWindows(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	cmd := exec.Command("cmdmp3", fileName)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	playOnWindows(fileName)

}
