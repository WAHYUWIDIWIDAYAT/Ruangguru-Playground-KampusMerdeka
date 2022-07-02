package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Dari contoh yang diberikan, cobalah untuk mengimplementasikan sebuah http client sederhana.
// Buatlah sebuah http client dan lakukan request GET ke API https://www.metaweather.com/api/.
// Buatlah request get ke endpoint /api/location/(woeid)/(date)/ dengan nilai woeid 1047378.
// Untuk date, gunakan format YYYY/MM/dd

func main() {
	resp, err := http.Get("https://www.metaweather.com/api/location/1047378/2022/05/09/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

}
