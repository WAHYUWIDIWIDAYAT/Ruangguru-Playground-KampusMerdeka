package main

import (
	"github.com/ruang-guru/playground/backend/design-patterns/creational/1-factory-pattern/factoryAwal"
)

// Ini adalah file client, kita akan menggunakan package yang sudah kita buat sebelumnya
// Coba lah bermain main dengan kode ini dan gunakan contentCreator dengan type NetflixKorea
func main() {
	var contentCreator factoryAwal.ContentCreator
	var contentCreator1 factoryAwal.ContentCreator
	var contentCreator2 factoryAwal.ContentCreator


	contentCreator2 = &factoryAwal.NetflixKorea{}
	contentCreator1 = &factoryAwal.Disney{}
	contentCreator = &factoryAwal.Mappa{}


	content := contentCreator1.Produce()
	content1 := contentCreator.Produce()
	content2 := contentCreator2.Produce()
	content.Play()
	content1.Play()
	content2.Play()
}
