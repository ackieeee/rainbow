package main

import "github.com/ackieeee/rainbow/rainbow"

func main() {
	texts := []rainbow.Text{
		{
			Param: "Hello, ",
			Color: "#ff0000",
		},
		{
			Param: "World",
			Color: "#00ff00",
		},
	}
	textList := rainbow.New(texts...)
	textList.Print()
}
