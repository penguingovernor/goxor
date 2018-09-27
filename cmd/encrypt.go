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
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/penguingovernor/goxor/protocol"

	"github.com/penguingovernor/goxor/xor"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt encrypts data using a key and a signature with a xor cipher",
	Long: `Examples:

# Encrypts input hello.txt using a one time pad as the key, 'goxor' as the signature, outputs out.xor and out.xor.key
goxor encrypt -i hello.txt

# Encrypts input hello.txt using a one time pad as the key, 'goxor' as the signature, outputs data.xor and key.xor.key
goxor encrypt -i hello.txt -o data -K key

# Encrypts input hello.txt using key 'test' and signature 'sig', outputs out.xor and out.xor.key
goxor encrypt -i hello -k test -s sig`,
	Run: func(cmd *cobra.Command, args []string) {
		// Grab the flags we defined earlier
		inFlag, err := cmd.Flags().GetString("input")
		if err != nil {
			log.Println(err)
		}
		keyFlag, err := cmd.Flags().GetString("key")
		if err != nil {
			log.Println(err)
		}
		outputFlag, err := cmd.Flags().GetString("output")
		if err != nil {
			log.Println(err)
		}
		keyOutFlag, err := cmd.Flags().GetString("key_out")
		if err != nil {
			log.Println(err)
		}
		sigFlag, err := cmd.Flags().GetString("signature")
		if err != nil {
			log.Println(err)
		}
		// Encrypt the files
		encrypt(inFlag, keyFlag, sigFlag, outputFlag, keyOutFlag)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	const (
		inputMsg string = `the file to encrypt
if the file cannot be found, then the input is treated as a string and that is encrypted
if the flag is omitted, then stdin will be used as the source for encryption`

		outMsg string = `the file to output the encrypted data to
if omitted, then the file will be out.xor
if the input is "stdout", then the data will be placed to stdout`

		outKeyMsg string = `the file to output the encrypted data to
if omitted, then the file will be out.xor.key
if the input is "stdout", then the data will be placed to stdout`

		keyMsg string = `the file to use as the key to the enryption process
if the file cannot be found, then the input is treated as a string and that is used as the key
if the flag is omitted, then a one time pad will be used as the key
if the input is "stdin", then stdin will be used as the key`

		sigMsg string = `the file to use as the signature to the enryption process
if the file cannot be found, then the input is treated as a string and that is used as the signature
if the flag is omitted, then "goxor" will be used as the signature
if the input is "stdin", then stdin will be used as the signautre`
	)
	encryptCmd.Flags().StringP("input", "i", "", inputMsg)
	encryptCmd.Flags().StringP("key", "k", "", keyMsg)
	encryptCmd.Flags().StringP("output", "o", "", outMsg)
	encryptCmd.Flags().StringP("key_out", "K", "", outMsg)
	encryptCmd.Flags().StringP("signature", "s", "", sigMsg)

}

func encrypt(inputFlag, keyFlag, signatureFlag, outputFlag, outputKeyFlag string) {

	// Get the bytes from user input
	inputBytes := getInput(inputFlag)
	keyBytes := getKey(keyFlag, len(inputBytes))
	signatureBytes := getSignature(signatureFlag)

	// Generate the appropriate data
	data := xor.GenerateData(inputBytes, signatureBytes)
	key := xor.GenerateKey(keyBytes, signatureBytes)

	// Encrypt the data
	eData, err := xor.Encrypt(data, key)
	if err != nil {
		log.Fatalf("error while encryptng file: %v", err)
	}

	// Say we're done and output the files
	fmt.Println("Done encrypting file")
	writeData(eData, outputFlag)
	writeKey(key, outputKeyFlag)

}

func writeData(data *protocol.Data, dest string) {
	outname := dest

	if strings.ToLower(dest) == "stdout" {
		fmt.Println("---BEGIN ENCRYPTED DATA---")
		xor.WriteData(os.Stdout, data)
		fmt.Println("---END ENCRYPTED DATA---")
		return
	}

	if dest == "" {
		outname = "out"
	}

	file, err := os.Create(outname + ".xor")
	if err != nil {
		log.Fatalf("Could not create file %s: %v\n", outname+".xor", err)
	}

	if err := xor.WriteData(file, data); err != nil {
		log.Fatalf("could not write key %s: %v\n", outname+".xor", err)
	}

	if err := file.Close(); err != nil {
		log.Fatalf("could not close file %s: %v\n", outname+".xor.key", err)
	}

	fmt.Println("Data written to:", outname+".xor")
}

func writeKey(data *protocol.Key, dest string) {
	outname := dest

	if strings.ToLower(dest) == "stdout" {
		fmt.Println("---BEGIN KEY DATA---")
		xor.WriteKey(os.Stdout, data)
		fmt.Println("---END KEY DATA---")
		return
	}

	if dest == "" {
		outname = "out"
	}

	file, err := os.Create(outname + ".xor.key")
	if err != nil {
		log.Fatalf("Could not create file %s: %v\n", outname+".xor.key", err)
	}

	if err := xor.WriteKey(file, data); err != nil {
		log.Fatalf("could not write key %s: %v\n", outname+".xor.key", err)
	}

	if err := file.Close(); err != nil {
		log.Fatalf("could not close file %s: %v\n", outname+".xor.key", err)
	}

	fmt.Println("Key written to:", outname+".xor.key")

}

func getInput(input string) []byte {
	// If omitted
	if input == "" {
		fmt.Println("Reading from stdin for input, press ctrl-d to stop")
		stdinBytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Could not read from stdin: %v", err)
		}
		return stdinBytes
	}
	// Try to open the file
	fileBytes, err := ioutil.ReadFile(input)
	// If we couldn't open the file, treat input as string
	if err != nil {
		fmt.Printf("Using string: \"%s\" as input\n", input)
		return []byte(input)
	}
	// If we could open the file, return the bytes
	fmt.Println("Using file:", input, "as input")
	return fileBytes
}

func getKey(input string, length int) []byte {

	// If omitted, use otp
	if input == "" {
		fmt.Println("Using one time pad as key")
		// Seed random
		rand.Seed(time.Now().Unix())
		// Make the pad
		key := make([]byte, length)
		_, err := rand.Read(key)
		if err != nil {
			log.Fatalf("Could not generate one time pad: %v", err)
		}
		return key
	}

	// If stdin
	if strings.ToLower(input) == "stdin" {
		fmt.Println("Reading from stdin for key, press ctrl-d to stop")
		key, err := ioutil.ReadAll(os.Stdin)
		fmt.Println("")
		if err != nil {
			log.Fatalf("Could not read from stdin: %v", err)
		}
		return key
	}

	// Try to open the file
	fileBytes, err := ioutil.ReadFile(input)
	// If we couldn't open the file, treat input as string
	if err != nil {
		fmt.Printf("Using string: \"%s\" as key\n", input)
		return []byte(input)
	}

	// If we could open the file, return the bytes
	fmt.Println("Using file:", input, "as key")
	return fileBytes
}

func getSignature(input string) []byte {
	// If omitted, use goxor
	if input == "" {
		fmt.Println("Using \"goxor\" as signature")
		return []byte("goxor")
	}

	// If stdin
	if strings.ToLower(input) == "stdin" {
		fmt.Println("Reading from stdin for signature, press ctrl-d to stop")
		key, err := ioutil.ReadAll(os.Stdin)
		fmt.Println("")
		if err != nil {
			log.Fatalf("Could not read from stdin: %v", err)
		}
		return key
	}

	// Try to open the file
	fileBytes, err := ioutil.ReadFile(input)
	// If we couldn't open the file, treat input as string
	if err != nil {
		fmt.Printf("Using string: \"%s\" as signature\n", input)
		return []byte(input)
	}

	// If we could open the file, return the bytes
	fmt.Println("Using file:", input, "as signature")
	return fileBytes

}
