package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/brecht-vde/prompter/engine"
)

func main() {
	template := ""
	data := make(map[string]interface{})

	args := os.Args[1:]

	for i := 0; i < len(args); i += 2 {
		cmd := args[i]
		val := args[i+1]

		switch cmd {
		case "-t":
			template = strings.TrimSuffix(strings.TrimPrefix(val, "\""), "\"")
		case "-j":
			joinArgs := strings.Split(val, "=")
			joinVals := strings.Split(joinArgs[1], ",")

			for i := 0; i < len(joinVals); i++ {
				joinVals[i] = strings.TrimSuffix(strings.TrimPrefix(joinVals[i], "\""), "\"")
			}

			data[joinArgs[0]] = joinVals
		case "-v":
			varArgs := strings.Split(val, "=")
			data[varArgs[0]] = strings.TrimSuffix(strings.TrimPrefix(varArgs[1], "\""), "\"")
		}
	}

	engine := engine.NewEngine()
	result, err := engine.Render(template, data)

	if err != nil {
		fmt.Printf("could not process template: \n%s", err.Error())
	}

	fmt.Print(result)
}
