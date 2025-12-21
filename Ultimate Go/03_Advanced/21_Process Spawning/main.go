package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Basic command :
	// cmd := exec.Command("cmd", "/C", "echo", "Hello World")
	// output, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(output))

	// Finding the words in command
	// cmd := exec.Command("findstr", "foo") // find the particular string

	// // Set Input for the command
	// cmd.Stdin = strings.NewReader("foo\nbar\nbaz\n")

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(output))

	// Sleeping the program
	// cmd := exec.Command("cmd", "/C", "timeout", "/T", "5", "/NOBREAK") // sleep command

	// // Start the command here:
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// // Waiting
	// err = cmd.Wait()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Process is complete")

	// Kill command
	// cmd := exec.Command("cmd", "/C", "timeout", "/T", "60", "/NOBREAK") // sleep command

	// // Start the command here:
	// err := cmd.Start()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// time.Sleep(2 * time.Second)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Process is Killed")

	// Working with environment variables
	// cmd := exec.Command("cmd", "/C", "echo", "%SHELL%")
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Error : ", err)
	// 	return
	// }
	// fmt.Println(string(output))

	// IO package with command
	// pr, pw := io.Pipe()
	// cmd := exec.Command("findstr", "foo")
	// cmd.Stdin = pr

	// go func() {
	// 	defer pw.Close()
	// 	pw.Write([]byte("food is bad\nbar\nbaz"))
	// }()

	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(output))

	cmd := exec.Command("cmd", "/C", "dir")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Output", string(output))
}
