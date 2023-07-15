package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/f91og/fy/src/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
	// model := flag.String("model", "word", "sentence or word")
	// flag.Parse()
	// fmt.Println(os.Args[1])
	// fmt.Println(engine.Translate(os.Args[1], "", ""))
}
