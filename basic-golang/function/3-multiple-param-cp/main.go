package main

import "fmt"

//Memanggil fungsi goodAfternoon
//dari dalam good afternoon akan dilakukan print "selamat sore name1 dan name2"
func main() {
	goodAfternoon("adi", "anti")
	goodAfternoon("ado", "suci")

}
func goodAfternoon(nama, kota string) {
	fmt.Printf("Halo %s dari %s\n", nama, kota)
}

// TODO: answer here
