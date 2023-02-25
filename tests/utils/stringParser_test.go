package tests

import (
	"testing"

	"github.com/TheLe0/f1-race-analysis/utils"
)

func TestParseStringToUint8(t *testing.T) {

	strNum := "8"
	numberParsed := utils.ParseStringToUint8(strNum)

	if numberParsed != 8 {
		t.Fatalf("Expected the number to be %v but found %v", 8, numberParsed)
	}
}

func TestParseStringToUint8Fail(t *testing.T) {

	strNum := "Loreum ipsul dolor"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	utils.ParseStringToUint8(strNum)
}

func TestParseStringToUint32(t *testing.T) {

	strNum := "8"
	numberParsed := utils.ParseStringToUint32(strNum)

	if numberParsed != 8 {
		t.Fatalf("Expected the number to be %v but found %v", 8, numberParsed)
	}
}

func TestParseStringToUint32Fail(t *testing.T) {

	strNum := "Loreum ipsul dolor"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	utils.ParseStringToUint32(strNum)
}

func TestParseTimeStringToMilliseconds(t *testing.T) {

	dateTime := "23:49:08.277"
	milliseconds := utils.ParseTimeStringToMilliseconds(dateTime)

	if milliseconds != 85748277 {
		t.Fatalf("Expected the milliseconds to be %v but found %v", 85748277, milliseconds)
	}
}

func TestParseLapTimeStringToMilliseconds(t *testing.T) {

	lapTimeStr := "1:02.852"
	lapTime := utils.ParseLapTimeStringToMilliseconds(lapTimeStr)

	if lapTime != 62852 {
		t.Fatalf("Expected the lap time to be %v but found %v", 62852, lapTime)
	}
}

func TestParseMillisecondsToLapTimeString(t *testing.T) {

	lapTime := utils.ParseMillisecondsToLapTimeString(62852)

	if lapTime != "01:02.852" {
		t.Fatalf("Expected the lap time to be %s but found %s", "01:02.852", lapTime)
	}
}

func TestParseMillisecondsToTimeString(t *testing.T) {

	milliseconds := utils.ParseMillisecondsToTimeString(85748277)

	if milliseconds != "23:49:08.277" {
		t.Fatalf("Expected the lap time to be %s but found %s", "23:49:08.277", milliseconds)
	}
}

func TestParseSpeedString(t *testing.T) {

	speed := utils.ParseSpeedString("44,275")

	if speed != 44275 {
		t.Fatalf("Expected the lap time to be %v but found %v", 44275, speed)
	}
}

func TestParseSpeedToString(t *testing.T) {

	speed := utils.ParseSpeedToString(44275)

	if speed != "44.275" {
		t.Fatalf("Expected the lap time to be %s but found %s", "44.275", speed)
	}
}
