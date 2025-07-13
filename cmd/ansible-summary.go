// Package main provides a CLI tool to summarize Ansible task execution states.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sgaunet/ansible-summary/pkg/ansiblesummary"
)

var version = "dev"

const (
	exitCodeSuccess           = 0
	exitCodeError             = 1
	exitCodeChangesOrFailures = 2
)

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
		os.Exit(exitCodeSuccess)
	}

	if inputFile == "" {
		fmt.Println("input file is mandatory")
		os.Exit(exitCodeError)
	}

	output := ansiblesummary.NewOutput()
	summary, err := ansiblesummary.NewAnsibleSummaryFromFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(exitCodeError)
	}

	if jsonFlag {
		err = output.WriteStatsJSON(summary)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(exitCodeError)
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
		os.Exit(exitCodeError)
	}
	exitWithSummaryStatus(summary)
}

func exitWithSummaryStatus(a *ansiblesummary.AnsibleSummary) {
	if a.HasChangedOrFailed() {
		os.Exit(exitCodeChangesOrFailures)
	}
	os.Exit(exitCodeSuccess)
}
