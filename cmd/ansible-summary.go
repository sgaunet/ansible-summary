package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/ansible-summary/pkg/ansiblesummary"
)

func main() {
	var inputFile string
	flag.StringVar(&inputFile, "input", "", "input file")
	flag.Parse()

	if inputFile == "" {
		fmt.Println("input file is mandatory")
		os.Exit(1)
	}

	summary, err := ansiblesummary.NewAnsibleSummaryFromFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	summary.PrintNameOfTaksNotOK()
	fmt.Println("************************************")
	summary.PrintHTMLStats()
}
