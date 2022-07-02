package main

import (
	"fmt"
	"time"
)

type character struct {
	name            string
	activity        string
	defaultActivity string
}

func (c *character) awake(movementInput, attackInput chan string) {
	for {
		//sama seperti checkpoint sebelumnya
		//tetapi sekarang menambahkan default activity yang menjalankan
		//fmt.Printf("%s %s\n", c.name, c.defaultActivity)
		//time.Sleep(100 * time.Millisecond)
		// TODO: answer here
		select {
		case movement := <-movementInput:
			c.activity = movement
		case attack := <-attackInput:
			c.activity = attack
		default:
			fmt.Printf("%s %s\n", c.name, c.defaultActivity)
			time.Sleep(100 * time.Millisecond)
		}

	}
}
