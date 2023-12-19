package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Flag Definitions
var (
	helpFlag      bool
	shortHelpFlag bool
	userFlag      string
	dbPathFlag    string
	currDBFlag    string
	startDBFlag   string
	endDBFlag     string
	SQLFlag       string
)

func configFlags() {
	flag.BoolVar(&helpFlag, "help", false, "Show help message")
	flag.BoolVar(&shortHelpFlag, "h", false, "Show help message (short)")
	flag.StringVar(&userFlag, "name", "Guest", "Specify a name")
	flag.StringVar(&dbPathFlag, "set-db", "", "Set the Current Database Path")
	flag.StringVar(&currDBFlag, "pwd-db", "", "Print the Current Database Path")

	flag.StringVar(&startDBFlag, "start-db", "", "Stop Current Database")
	flag.StringVar(&endDBFlag, "kill-db", "", "End Current Database")

	flag.StringVar(&SQLFlag, "sql-db", "", "Import Current Database to MySQL Local Instance")
}

func main() {
	configFlags()
	flag.Parse()

	// Check for the help flags
	if helpFlag || shortHelpFlag {
		HelpMenu()
		os.Exit(0)
	}

	// Add your CLI logic here
	fmt.Printf("Hello, %s!\n", userFlag)

	// You can add more functionality based on other flags or arguments
	if userFlag != "" {
		fmt.Printf("You provided a value for the long flag: %s\n", userFlag)
	}
}

func HelpMenu() {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	fmt.Printf("  %s, %s: %s\n", red("-help"), red("-h"), green("Show this help message"))
	fmt.Printf("  %s: %s\n", red("-name"), green("Specify a name (default: Guest)"))
	fmt.Printf("  %s: %s\n", red("--long-flag"), green("A flag with a long name"))
	// Add more colored options as needed
}
