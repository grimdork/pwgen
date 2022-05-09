// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/grimdork/climate/arg"
)

func main() {
	opt := arg.New("pwgen")
	opt.SetDefaultHelp(true)
	opt.SetOption("", "l", "length", "Length of passwords.", 16, false, arg.VarInt, nil)
	opt.SetOption("", "c", "count", "Number of passwords to generate.", 140, false, arg.VarInt, nil)
	opt.SetOption("", "n", "nonce", "Generate a binary string more suitable for cryptography. Sets count to 1. Should be piped to file.", false, false, arg.VarBool, nil)
	opt.SetOption("", "w", "words", "Generate word-based passwords.", false, false, arg.VarBool, nil)
	opt.SetOption("", "W", "wordcount", "Number of words in word-based passwords.", 6, false, arg.VarInt, nil)
	opt.SetOption("", "C", "completions", "Show script for Bash completions. Source it to enable.", false, false, arg.VarBool, nil)
	opt.Parse(os.Args)
	if opt.GetBool("help") {
		opt.PrintHelp()
		return
	}

	if opt.GetBool("completions") {
		comp, err := opt.Completions("pwgen")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		fmt.Printf("%s\n", comp)
		return
	}

	if opt.GetBool("nonce") && !opt.GetBool("words") {
		fmt.Printf("%s", RandNonce(opt.GetInt("length")))
		return
	}

	if opt.GetBool("words") {
		wc := opt.GetInt("wordcount")
		for i := 0; i < opt.GetInt("count"); i++ {
			pr("\t%s", RandWords(wc))
		}
		return
	}

	// Issues with TinyGo; temporarily removed.
	// w, _, err := terminal.GetSize(int(os.Stdout.Fd()))
	// if err != nil {
	// 	pr("Couldn't get terminal width: %s", err.Error())
	// 	os.Exit(2)
	// }
	w := 120
	length := opt.GetInt("length")
	maxcount := opt.GetInt("count")
	maxw := w / (length + 2)
	count := 0
	for i := 0; i < maxcount; i++ {
		fmt.Printf("%s  ", RandString(length))
		count++
		if count == maxw {
			println("")
			count = 0
		}
	}

	// Breather
	println("")
}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}
