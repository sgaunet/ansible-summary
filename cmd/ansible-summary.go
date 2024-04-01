package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/ansible-summary/pkg/ansiblesummary"
)

var version string = "dev"

func printVersion() {
	fmt.Println(version)
}

func main() {
	var inputFile string
	var jsonFlag bool
	var versionFlag bool
	flag.StringVar(&inputFile, "input", "", "input file")
	flag.BoolVar(&jsonFlag, "json", false, "output format as JSON")
	flag.BoolVar(&versionFlag, "version", false, "print version")
	flag.Parse()

	if versionFlag {
		printVersion()
		os.Exit(0)
	}

	if inputFile == "" {
		fmt.Println("input file is mandatory")
		os.Exit(1)
	}

	output := ansiblesummary.NewOutput()
	summary, err := ansiblesummary.NewAnsibleSummaryFromFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if jsonFlag {
		err = output.WriteStatsJSON(summary)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		exitWithSummaryStatus(summary)
	}

	summary.PrintNameOfTaksNotOK()
	fmt.Println("************************************")
	errs := output.WriteStatsHTML(summary)
	if len(errs) > 0 {
		for idx := range errs {
			fmt.Fprintln(os.Stderr, errs[idx].Error())
		}
		os.Exit(1)
	}
	exitWithSummaryStatus(summary)
}

func exitWithSummaryStatus(a *ansiblesummary.AnsibleSummary) {
	if a.HasChangedOrFailed() {
		os.Exit(2)
	}
	os.Exit(0)
}
