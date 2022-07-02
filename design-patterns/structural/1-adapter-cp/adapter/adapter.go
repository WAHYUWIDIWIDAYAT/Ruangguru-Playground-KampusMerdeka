package adapter

import "fmt"

type AudioPlayer interface {
	Play(m Mp3) string
}

type Mp3 struct {
	Data []byte
}

func (m Mp3) PlayAudio() string {
	fmt.Println("playing mp3 with data" + string(m.Data))
	return string(m.Data)
}

type Kaset struct {
	PitaMusik string
}

type Walkman struct{}

func (w Walkman) Play(c Kaset) string {
	fmt.Println(c.PitaMusik)
	return c.PitaMusik
}

type Mp3ToKasetAdapter struct {
	Adaptee Walkman
}

func (a Mp3ToKasetAdapter) Play(m Mp3) string {

	k := Kaset{}
	k.PitaMusik = string(m.Data)
	return a.Adaptee.Play(k)

}