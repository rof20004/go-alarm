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

	var days int
	var hours int
	var minutes int
	var alarmTime string

	// Check length of arguments passed to process
	if len(os.Args) != 3 {
		log.Fatalln("Arguments mismatch, needs two arguments [time->HH:MM] [days-foward->N], type 0 to days-foward for today")
	}

	// Get hours and minutes, separated by :
	alarmTime = os.Args[1]

	// Get days
	days, _ = strconv.Atoi(os.Args[2])

	// Check hours formatt
	if len(alarmTime) != 5 {
		log.Fatalln("Invalid time. Correct formatt: 13:01")
	}

	if days <= 0 {
		log.Fatalln("Invalid days. Days need to be greater than ZERO")
	}

	// Split hours into two elements array
	s := strings.Split(alarmTime, ":")

	// Convert hours to int
	hours, _ = strconv.Atoi(s[0])

	// Convert minutes to int
	minutes, _ = strconv.Atoi(s[1])

	// Get current time
	timeNow := time.Now()

	// Set new time variable with current time
	alarmDate := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), hours, minutes, 0, 0, time.Local)

	// Add days foward to alarmDate
	alarmDate = alarmDate.AddDate(0, 0, days)

	// Check if alarm time is after current time
	if alarmDate.Before(timeNow) {
		log.Fatalln("Alarm time error, needs to be after current time")
	}

	fmt.Println("Alarm setted to:", alarmDate)

	for {

		// While current time is after alarm continue checking inside loop
		// If not, break loop and plays alarm recursively
		if time.Now().After(alarmDate) {
			break
		}

	}

	// Verify if Operation System is windows or linux
	switch runtime.GOOS {
	case "windows":
		playOnWindows("alarm.mp3")
	case "linux":
		playOnLinux("alarm.mp3")
	default:
		fmt.Println("O.S not supported")
	}

}

// Play sound using mpg123 package
func playOnLinux(fileName string) {

	// Check if alarm file sound exists
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Check if mpg123 package can be used
	cmd := exec.Command("mpg123", fileName)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	// Call recursively
	playOnLinux(fileName)

}

// Play sound using cmdmp3 package
func playOnWindows(fileName string) {

	// Check if alarm file sound exists
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Check if mpg123 package can be used
	cmd := exec.Command("cmdmp3", fileName)
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}

	// Call recursively
	playOnWindows(fileName)

}
