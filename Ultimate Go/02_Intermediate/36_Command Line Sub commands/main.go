package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subcommand1 := flag.NewFlagSet("firstsub", flag.ExitOnError)  // go run main.go firstsub --help (to know more info)
	subcommand2 := flag.NewFlagSet("secondsub", flag.ExitOnError) // go run main.go secondsub --help

	firstflag := subcommand1.Bool("Processing", false, "Command processing status")
	secondflag := subcommand1.Int("bytes", 1024, "Byte length of the result")
	// go run main.go firstsub -Processing=true -bytes=2034 (Subcommand1 example)

	flagc2 := subcommand2.String("language", "Go", "Enter your Language")
	// go run main.go secondsub -language=python (Subcommand2 example)

	if len(os.Args) < 2 {
		fmt.Println("This program requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstsub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("SubCommand1 : ")
		fmt.Println("processing : ", *firstflag)
		fmt.Println("bytes : ", *secondflag)
	case "secondsub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2 : ")
		fmt.Println("Language : ", *flagc2)
	default:
		fmt.Println("No subcommand entered")
	}
}
