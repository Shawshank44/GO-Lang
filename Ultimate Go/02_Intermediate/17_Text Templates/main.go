package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	// temp := template.New("Example")

	// temp, err := template.New("Example").Parse("Welcome, {{.name}}! How are you doing")

	// if err != nil {
	// 	panic(err)
	// }

	// // Defining the data :
	// data := map[string]any{
	// 	"name": "John",
	// }

	// err = temp.Execute(os.Stdout, data)

	// if err != nil {
	// 	panic(err)
	// }

	// // Other ways to create templates :
	// templ := template.Must(template.New("Example").Parse("Welcome, {{.name}}! How are you doing")) // instead of handeling errors

	// data := map[string]any{
	// 	"name": "John",
	// }

	// templ.Execute(os.Stdout, data)

	// console application :
	// getting input name
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define named templates for types of
	temps := map[string]string{
		"welcome":      "Welcome, {{.name}}! we're glad you joined",
		"notification": "{{.nm}}, you have a new Notification {{.ntf}}.",
		"error":        "OOPS! An error occured : {{.em}}",
	}

	// Parse and store templates :
	parsedTemplates := make(map[string]*template.Template)
	for name, temp := range temps {
		parsedTemplates[name] = template.Must(template.New(name).Parse(temp))
	}

	for {
		// Show Menu
		fmt.Println("\n Menu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get error")
		fmt.Println("4. exit")
		fmt.Println("choose an option")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var templ *template.Template

		switch choice {
		case "1":
			templ = parsedTemplates["welcome"]
			data = map[string]any{
				"name": name,
			}
		case "2":
			fmt.Println("Enter your notification message : ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			templ = parsedTemplates["notification"]
			data = map[string]any{
				"nm":  name,
				"ntf": notification,
			}
		case "3":
			fmt.Println("Enter you error message : ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			templ = parsedTemplates["error"]
			data = map[string]any{
				"nm": name,
				"em": errorMessage,
			}

		case "4":
			fmt.Println("exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option")
			continue
		}

		err := templ.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("error while executing the template")
		}
	}
}
