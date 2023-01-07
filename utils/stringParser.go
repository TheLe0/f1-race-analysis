package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseStringToUint8(input string) uint8 {

	valInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Printf("Could not cast the %s string input into uint8 type due to this %s error \n", input, err)
	}

	return uint8(valInt)
}

func ParseStringToUint32(input string) uint32 {

	valInt, err := strconv.Atoi(input)

	if err != nil {
		fmt.Printf("Could not cast the %s string input into uint8 type due to this %s error \n", input, err)
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

func ParseSpeedString(speedString string) uint32 {

	stringSplit := strings.Split(speedString, ",")

	return ParseStringToUint32(stringSplit[1]) + (ParseStringToUint32(stringSplit[0]) * 1000)
}
