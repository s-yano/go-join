package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fp := os.Stdin
	var err error

	sep := flag.String("s", ",", "セパレータの指定")
	ignore_blank := flag.Bool("b", true, "空白行の無視")
	last_newline := flag.Bool("n", true, "最後に改行を挿入する")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		// ファイルオープン
		fp, err = os.Open(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	s := ""

	for scanner.Scan() {
		text := scanner.Text()
		if *ignore_blank && len(text) == 0 {
			continue
		}
		fmt.Print(s, text)
		s = *sep
	}
	if *last_newline {
		fmt.Println()
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
