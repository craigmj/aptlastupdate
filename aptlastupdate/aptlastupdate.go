package main

import (
	"fmt"
	"os"
	"time"

	"github.com/craigmj/aptlastupdate"
)

func main() {
	var err error
	d := time.Hour
	if 1<len(os.Args) {
		d,err = time.ParseDuration(os.Args[1])
		if nil!=err {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	b, err := aptlastupdate.Within(d)
	if nil!=err {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if b {
		os.Exit(0)
	}
	os.Exit(1)
}
		
	
