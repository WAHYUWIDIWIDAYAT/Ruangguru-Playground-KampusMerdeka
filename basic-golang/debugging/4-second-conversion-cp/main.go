package main

import "fmt"

func main() {
	/*
		Convert the given second to 00:00:00 hour minute second format

		Example Input/Output
		30 -> 00:00:30
		70 -> 00:01:10
		67812 -> 18:50:12
		678120 -> 188:22:00

	*/
	res := ConvertSecondToTimeString(30)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ConvertSecondToTimeStringCorrect(arr)
	// fmt.Println(resCorrect)
	resCorrect := ConvertSecondToTimeStringCorrect(30)
	fmt.Println(resCorrect)
}

func ConvertSecondToTimeString(second int) string {
	hours := second / 3600
	minutes := second / 60

	timeString := fmt.Sprintf("%d:%d:%d", hours, minutes, second)
	return timeString
}

func ConvertSecondToTimeStringCorrect(second int) string {
	if second < 0 {
		return "00:00:00"
	}
	var hours, minutes, seconds int
	hours = second / 3600
	minutes = (second % 3600) / 60
	seconds = second % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
