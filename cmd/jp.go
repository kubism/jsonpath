package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/spf13/cobra.v0"
)

var (
	key        string
	showValue  bool
	showType   bool
	ignoreCase bool
)

var jpCmd = &cobra.Command{
	Use:   "jp [flags]",
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
	flags.BoolVarP(&showType, "show-type", "t", false, "If enabled, will show the type of the value for found path")
	flags.BoolVarP(&ignoreCase, "ignore-case", "i", false, "If enabled, will ignores case when matching key")
}

func printPaths(v interface{}, key, path string) {
	switch v := v.(type) {
	case map[string]interface{}:
		for mk, mv := range v {
			p := path + "." + mk
			if (ignoreCase && strings.EqualFold(mk, key)) || mk == key {
				fmt.Print(p)
				if showValue {
					fmt.Printf(":%v", mv)
				}
				if showType {
					fmt.Printf(":%T", mv)
				}
				fmt.Print("\n")
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
