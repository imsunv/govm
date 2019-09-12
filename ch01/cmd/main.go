package main

import (
	"fmt"
	"os"
	//prompt "github.com/c-bata/go-prompt"
)

func main() {
	op := &Options{}
	cliApp := build(op)

	err := cliApp.Run(os.Args)
	if err != nil {
		panic(err)
	}

	fmt.Println("env: " + op.Env)
	fmt.Println("cp: " + op.Classpath)

	//cliApp := build(app)
	//suggests := make([]prompt.Suggest, 0)
	//for _, cmd := range cliApp.Commands {
	//	suggests = append(suggests, prompt.Suggest{
	//		Text:        cmd.Name,
	//		Description: cmd.Usage,
	//	})
	//}
	//
	//p := prompt.New(func(cmd string) {
	//	if cmd == "exit" {
	//		os.Exit(0)
	//	}
	//	err := cliApp.Run(strings.Split("console "+cmd, " "))
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}, func(d prompt.Document) []prompt.Suggest {
	//	w := d.GetWordBeforeCursor()
	//	if w == "" {
	//		return []prompt.Suggest{}
	//	}
	//	return prompt.FilterHasPrefix(suggests, w, true)
	//}, prompt.OptionPrefix("govm> "), prompt.OptionTitle("govm command line"))
	//p.Run()
}
