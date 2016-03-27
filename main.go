package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parsed()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	conf, err := ReadConfig("~/.we/settings.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if args[0] == "pull" {
		Pull(conf)
	} else if args[0] == "push" {
		Push(conf)
	} else if args[0] == "run" {
		Run(conf)
	} else if args[0] == "clean" {
		Clean(conf)
	} else {
		return
	}
}
