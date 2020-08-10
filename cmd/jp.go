package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gopkg.in/spf13/cobra.v0"
)

var (
	key       string
	showValue bool
)

var jpCmd = &cobra.Command{
	Use:   "jp action [flags]",
	Short: "JsonPath finds paths to a key in nested JSON structures.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if key == "" {
			return fmt.Errorf("key must not be empty")
		}

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
			return err
		}

		return nil
	},
}

func init() {
	flags := jpCmd.PersistentFlags()
	flags.AddGoFlagSet(flag.CommandLine)
	flags.StringVarP(&key, "key", "k", "", "Key to search for")
	flags.BoolVarP(&showValue, "show-value", "v", false, "If enabled, will show the value for found path")
}

func printPaths(v interface{}, key, path string) {
	switch v := v.(type) {
	case map[string]interface{}:
		for mk, mv := range v {
			p := path + "." + mk
			if mk == key {
				if showValue {
					fmt.Printf("%v [%v]\n", p, mv)
				} else {
					fmt.Println(p)
				}
			}
			printPaths(mv, key, p)
		}
	case []interface{}:
		for i, sv := range v {
			printPaths(sv, key, fmt.Sprintf("%s[%d]", path, i))
		}
	}
}

func fatalln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	if err := jpCmd.Execute(); err != nil {
		fatalln(err)
	}
}
