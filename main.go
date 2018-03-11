package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Created by Prabhat Sharma <hi.prabhat@gmail.com>")
	fmt.Println("Moves mouse every 30 seconds by 1, 1 pixel")
	fmt.Println("Use at your own risk :-)")
	var x, y int

	for {
		x, y = robotgo.GetMousePos()
		fmt.Println("mouse coordinates: x=", x, ", y=", y)
		robotgo.MoveMouse(x+1, y+1)
		time.Sleep(30 * time.Second)
	}
}
