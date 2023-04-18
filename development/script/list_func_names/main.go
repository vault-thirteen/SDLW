package main

import (
	"fmt"
	"log"
	"os"

	c "github.com/vault-thirteen/SDLW/development/common"
	"github.com/vault-thirteen/SDLW/development/parser"
)

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Command line arguments.
	if len(os.Args) != 2 {
		log.Fatal("os args")
	}
	filePath := os.Args[1]
	log.Println(fmt.Sprintf("Parsing a file: %s.", filePath))

	var funcData []*c.FuncData
	var err error
	funcData, err = parser.ParseCFunctions(filePath)
	mustBeNoError(err)
	fmt.Println(funcData) //TODO
}
