package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func printPaths(v interface{}, key, path string) {
	switch v := v.(type) {
	case map[string]interface{}:
		for mk, mv := range v {
			p := path + "." + mk
			if mk == key {
				fmt.Printf("%v [%v]\n", p, mv)
			}
			printPaths(mv, key, p)
		}
	case []interface{}:
		for i, sv := range v {
			printPaths(sv, key, fmt.Sprintf("%s[%d]", path, i))
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `usage:
  %s key
where key is the key to search for in JSON structures passed to standard input.
`, os.Args[0])
}

func fatalln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(1)
	}

	key := flag.Arg(0)
	scanner := bufio.NewScanner(os.Stdin)
	bytes := make([]byte, 0)

	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) == 0 {
			continue
		}
		bytes = append(bytes, b...)
	}

	var v interface{}
	if err := json.Unmarshal(bytes, &v); err != nil {
		fmt.Fprintf(os.Stderr, "JSON parsing error: %s\n", err)
	}
	printPaths(v, key, "")

	if err := scanner.Err(); err != nil {
		fatalln(err)
	}
}
