// Copyright © 2018 JORGE HENRIQUEZ <JOAHENRI@UCSC.EDU>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/penguingovernor/goxor/internal/cmdutil"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt decrypts data. The inputs should come from goxor encrypt",
	Long: `Examples:

# Decrypts input out.xor using out.xor.key as the key, this will output to stdout
$ goxor decrypt -i out.xor -k out.xor.key

# Decrypts inpuyt out.xor using out.xor.key as the key, this will output to test.xor
$ goxor decrypt -i out.xor -k out.xor.key -o test`,
	Run: func(cmd *cobra.Command, args []string) {
		intputFlag, err := cmd.Flags().GetString("input")
		if err != nil {
			log.Println(err)
		}
		keyFlag, err := cmd.Flags().GetString("key")
		if err != nil {
			log.Println(err)
		}
		outFlag, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Println(err)
		}

		if err := cmdutil.Decrypt(intputFlag, keyFlag, outFlag); err != nil {
			log.Fatalf("decryption err: %v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	const (
		inputMsg string = `the file to decrypt`
		keyMsg   string = `the file to use as the key to the decryption process`
		outMsg   string = `the file to output the decrypted data to
if omitted, then the file will be out.xor
if the input is "stdout", then the data will be placed to stdout`
	)
	decryptCmd.Flags().StringP("input", "i", "", inputMsg)
	decryptCmd.Flags().StringP("output", "o", "", outMsg)
	decryptCmd.Flags().StringP("key", "k", "", keyMsg)
}
