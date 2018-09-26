// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"io/ioutil"
	"log"
	"os"

	"github.com/penguingovernor/goxor/xor"

	"github.com/penguingovernor/goxor/protocol"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		decrypt(intputFlag, keyFlag, outFlag)

	},
}

func decrypt(in, key, out string) {
	input := dGetInput(in)
	keyData := dGetKey(key)
	data, err := xor.Decrypt(input, keyData)
	if err != nil {
		log.Fatal(err)
	}
	dWriteOutput(data, out)
}

func dWriteOutput(data *protocol.Data, out string) {

	if out == "" {
		if _, err := os.Stdout.Write(data.PayLoad); err != nil {
			log.Fatalf("could not write data: %v", err)
		}
		return
	}

	file, err := os.Create(out + ".xor")
	if err != nil {
		log.Fatalf("Could not create file %s.xor: %v", out, err)
	}

	if _, err := file.Write(data.PayLoad); err != nil {
		log.Fatalf("could not write to file %s.xor: %v", out, err)
	}

	if err := file.Close(); err != nil {
		log.Fatalf("Could not close file %s.xor: %v", out, err)
	}

}

func dGetInput(in string) *protocol.Data {
	if in == "" {
		log.Fatal("flag input is required")
	}

	fileBytes, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatalf("could not read %s: %v\n", in, err)
	}

	data, err := xor.LoadData(fileBytes)
	if err != nil {
		log.Fatalf("could not load file %s: %v\n", in, err)
	}

	return data
}

func dGetKey(in string) *protocol.Key {
	if in == "" {
		log.Fatal("flag key is required")
	}

	fileBytes, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatalf("could not read %s: %v", in, err)
	}

	data, err := xor.LoadKey(fileBytes)
	if err != nil {
		log.Fatalf("could not load file %s: %v", in, err)
	}

	return data
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
if ommitted, then the file will be out.xor
if the input is "stdout", then the data will be placed to stdout`
	)
	decryptCmd.Flags().StringP("input", "i", "", inputMsg)
	decryptCmd.Flags().StringP("output", "o", "", outMsg)
	decryptCmd.Flags().StringP("key", "k", "", keyMsg)
}
