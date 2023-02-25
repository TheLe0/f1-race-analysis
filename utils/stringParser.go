package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseStringToUint8(input string) uint8 {

	valInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Printf("Could not cast the %s string input into uint8 type due to the following error \n", input)
		panic(err)
	}

	return uint8(valInt)
}

func ParseStringToUint32(input string) uint32 {

	valInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Printf("Could not cast the %s string input into uint8 type due to the following error \n", input)
		panic(err)
	}

	return uint32(valInt)
}

func ParseTimeStringToMilliseconds(localTime string) uint32 {

	stringSplit := strings.Split(localTime, ".")

	milliseconds := ParseStringToUint32(stringSplit[1])

	stringSplit = strings.Split(stringSplit[0], ":")

	seconds := ParseStringToUint32(stringSplit[2]) * 1000
	minutes := ParseStringToUint32(stringSplit[1]) * 60000
	hours := ParseStringToUint32(stringSplit[0]) * 3600000

	totalMilliseconds := milliseconds + seconds + minutes + hours

	return totalMilliseconds
}

func ParseLapTimeStringToMilliseconds(localTime string) uint32 {

	stringSplit := strings.Split(localTime, ".")

	milliseconds := ParseStringToUint32(stringSplit[1])

	stringSplit = strings.Split(stringSplit[0], ":")

	seconds := ParseStringToUint32(stringSplit[1]) * 1000
	minutes := ParseStringToUint32(stringSplit[0]) * 60000

	totalMilliseconds := milliseconds + seconds + minutes

	return totalMilliseconds
}

func ParseMillisecondsToLapTimeString(localTime uint32) string {

	var minuteStr string
	var secondStr string
	var millisecondStr string

	minutes, remainder := localTime/60000, localTime%60000

	seconds, milliseconds := remainder/1000, remainder%1000

	if minutes <= 9 {
		minuteStr = fmt.Sprintf("0%d", minutes)
	} else {
		minuteStr = fmt.Sprintf("%d", minutes)
	}

	if seconds <= 9 {
		secondStr = fmt.Sprintf("0%d", seconds)
	} else {
		secondStr = fmt.Sprintf("%d", seconds)
	}

	if milliseconds <= 9 {
		millisecondStr = fmt.Sprintf("00%d", milliseconds)
	} else if milliseconds <= 99 {
		millisecondStr = fmt.Sprintf("0%d", milliseconds)
	} else {
		millisecondStr = fmt.Sprintf("%d", milliseconds)
	}

	return fmt.Sprintf("%s:%s.%s", minuteStr, secondStr, millisecondStr)
}

func ParseMillisecondsToTimeString(localTime uint32) string {

	hour, remainder := localTime/3600000, localTime%3600000
	minutes, remainder1 := remainder/60000, remainder%60000
	seconds, milliseconds := remainder1/1000, remainder1%1000

	var hourStr string
	var minuteStr string
	var secondStr string
	var millisecondStr string

	if hour <= 9 {
		hourStr = fmt.Sprintf("0%d", hour)
	} else {
		hourStr = fmt.Sprintf("%d", hour)
	}

	if minutes <= 9 {
		minuteStr = fmt.Sprintf("0%d", minutes)
	} else {
		minuteStr = fmt.Sprintf("%d", minutes)
	}

	if seconds <= 9 {
		secondStr = fmt.Sprintf("0%d", seconds)
	} else {
		secondStr = fmt.Sprintf("%d", seconds)
	}

	if milliseconds <= 9 {
		millisecondStr = fmt.Sprintf("00%d", milliseconds)
	} else if milliseconds <= 99 {
		millisecondStr = fmt.Sprintf("0%d", milliseconds)
	} else {
		millisecondStr = fmt.Sprintf("%d", milliseconds)
	}

	return fmt.Sprintf("%s:%s:%s.%s", hourStr, minuteStr, secondStr, millisecondStr)
}

func ParseSpeedString(speedString string) uint32 {

	stringSplit := strings.Split(speedString, ",")

	return ParseStringToUint32(stringSplit[1]) + (ParseStringToUint32(stringSplit[0]) * 1000)
}

func ParseSpeedToString(speed float32) string {

	thousand := speed / 1000

	return fmt.Sprintf("%0.3f", thousand)
}
