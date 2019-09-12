package main

import (
	"fmt"
	"os"
)

func main() {
	op := &Options{}
	app := build(op)

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}

	fmt.Println("env: " + op.Env)
	fmt.Println("cp: " + op.Classpath)
}
