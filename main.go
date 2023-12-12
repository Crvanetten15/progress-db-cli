package main

import (
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

var options = []string{"Option 1", "Option 2", "Option 3", "Exit"}

func main() {
	fmt.Println("Welcome to the Arrow Key Survey CLI!")

	err := run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func run() error {
	err := keyboard.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Use arrow keys to navigate and press Enter to select an option. Press 'q' to exit.")

	selectedIndex := 0

	for {
		printOptions(selectedIndex)

		_, key, err := keyboard.GetKey() //nolint:unused
		if err != nil {
			return err
		}

		switch key {
		case keyboard.KeyArrowUp:
			selectedIndex = (selectedIndex - 1 + len(options)) % len(options)
		case keyboard.KeyArrowDown:
			selectedIndex = (selectedIndex + 1) % len(options)
		case keyboard.KeyEnter:
			handleSelection(selectedIndex)
		case keyboard.KeyEsc, keyboard.KeyCtrlC:
			fmt.Println("\nExiting the program. Goodbye!")
			os.Exit(0)
		case 'q':
			fmt.Println("\nExiting the program. Goodbye!")
			os.Exit(0)
		}

		// Clear screen
		fmt.Print("\033[H\033[2J")

		time.Sleep(100 * time.Millisecond)
	}
}

func printOptions(selectedIndex int) {
	fmt.Println("Select an option:")
	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("\033[1m%s\033[0m\n", option) // Highlight the selected option
		} else {
			fmt.Println(option)
		}
	}
}

func handleSelection(selectedIndex int) {
	fmt.Printf("\nYou selected: %s\n", options[selectedIndex])

	// Add your actions based on the selected option here

	os.Exit(0)
}
