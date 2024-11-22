package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Add a new todo specify title")
	flag.IntVar(&cf.Del, "del", -1, "Add a new todo specify title")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Add a new todo specify title")
	flag.BoolVar(&cf.List, "list", false, "Add a new todo specify title")

	flag.Parse()
	return &cf
}

func (cf *cmdFlags) Execute(tds *Todos) {
	switch {
	case cf.List:
		tds.print()
	case cf.Add != "":
		tds.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid edit format. Use -edit <id>:<new_title>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid id")
			os.Exit(1)
		}
		tds.edit(id, parts[1])
	case cf.Del != -1:
		tds.delete(cf.Del)
	case cf.Toggle != -1:
		tds.toggle(cf.Toggle)
	default:
		fmt.Println("Error: Invalid command")
	}
}
