package main

import (
	"github.com/kyokomi/emoji"
)

func main() {
	PrintS()
}

func PrintS() string {
	//  fmt.Println("Hello 🗺️!")
	// want := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})

	return emoji.Sprintf("Hello :world_map:!")
}
