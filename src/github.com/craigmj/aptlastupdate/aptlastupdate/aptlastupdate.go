package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/craigmj/aptlastupdate"
)

func main() {
	flag.Parse()
	d := time.Hour
	if ""!=flag.Arg(0) {
		var err error
		d, err = time.ParseDuration(flag.Arg(0))	
		if nil!=err {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	}	
	b, err := aptlastupdate.Within(d)
	if nil!=err {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	// Command line returns TRUE if last update occurred within the specified duration
	if b {
		os.Exit(0)
	}
	os.Exit(1)
}
