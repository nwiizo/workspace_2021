/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func fizzbuzz(max int) {
	for i := 0; i <= max; i++ {
		fizz := i%3 == 0
		buzz := i%5 == 0
		switch {
		case fizz && buzz:
			fmt.Println("fizzbuzz")
		case fizz && !buzz:
			fmt.Println("fizz")
		case !fizz && buzz:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}

// fizzbuzzCmd represents the fizzbuzz command
var fizzbuzzCmd = &cobra.Command{
	Use:   "fizzbuzz [int]",
	Short: "return Fizzbuzz",
	Long: `return Fizzbuzz

return Fizzbuzz There is no particular reason because it is a suitable sample`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fizzbuzz called")
		var m int
		m, _ = strconv.Atoi(args[0])
		fizzbuzz(m)
	},
}

func init() {
	rootCmd.AddCommand(fizzbuzzCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fizzbuzzCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fizzbuzzCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
