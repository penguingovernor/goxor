// GoXor is a command line tool that encrypts and decrypts data.
//
// For usagae run `goxor --help`
//
// Encryption Specifics:
// Unless specified otherwise, encrypting a file yields a file with a go xor encrypted file extension (.gxef)
// and a key (.gxef.key) with the original file name preceding the extension.
// It is crucial the the key file is not renamed as `goxor` will look for this file in decryption.
//
// Without the key file the encrypted data is utterly useless, unless of coure you're a cryptanalyst.
package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ogier/pflag"
)

// Flag Related Code

var (
	inputFile   string
	outputFile  string
	decryptFlag bool
	cleanFlag   bool
)

func init() {
	pflag.BoolVarP(&decryptFlag, "decrypt", "d", false, "Toggle decrypt mode")
	pflag.BoolVarP(&cleanFlag, "clean", "c", false, "Remove decrypted data and key among completion")
	pflag.StringVarP(&inputFile, "input", "i", "", "Input file")
	pflag.StringVarP(&outputFile, "output", "o", "", "Output file")
}

func helpMsg() string {
	return "Try 'goxor --help' for more information"
}

func verifyFlags() {
	pflag.Parse()
	if !decryptFlag && inputFile == "" && outputFile == "" {
		log.Fatalf("\nNo output file while trying to read from stdin\n%s\n", helpMsg())
	}
	if decryptFlag && inputFile == "" {
		log.Fatalf("\nNo input file while trying to to decrypt\n%s\n", helpMsg())
	}
}

// End Flag Related Code

// File Related Code
func getFileBytes(fileName string) []byte {
	var (
		fileBytes []byte
		err       error
	)

	if fileName == "" {
		fileBytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		fileBytes, err = ioutil.ReadFile(fileName)
	}

	if err != nil {
		log.Fatal(err)
	}

	return fileBytes

}

func writeEncryptedData(data []byte) {
	var fileName string
	if outputFile == "" {
		fileName = inputFile + ".gxef"
	} else {
		fileName = outputFile
	}
	ioutil.WriteFile(fileName, data, 0666)
}

func writeKey(key []byte) {
	var fileName string
	if outputFile == "" {
		fileName = inputFile + ".gxef.key"
	} else {
		fileName = outputFile + ".key"
	}
	ioutil.WriteFile(fileName, key, 0666)
}

func writeDecryptedData(data []byte) {
	if outputFile == "" {
		fmt.Print(string(data))
		return
	}
	ioutil.WriteFile(outputFile, data, 0666)
}

func removeEvidence() {
	fmt.Println("Cleaning up files...")
	err := os.Remove(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(inputFile + ".key")
	if err != nil {
		log.Fatal(err)
	}
}

// End File Related Code

// Encryption Code

func encrytFile() {
	fileBytes := getFileBytes(inputFile)
	fileKey := generateFileKey(len(fileBytes))
	encrytedBytes := encrypt(fileBytes, fileKey)
	writeEncryptedData(encrytedBytes)
	writeKey(fileKey)
}

func generateFileKey(n int) []byte {
	fileKey := make([]byte, n)
	rand.Read(fileKey)
	return fileKey
}

func encrypt(bytes, key []byte) []byte {
	encryptedBytes := make([]byte, len(key))

	for i := range encryptedBytes {
		encryptedBytes[i] = bytes[i] ^ key[i]
	}

	return encryptedBytes

}

// End Encryption Code

// Decryption Code

func decryptFile() {
	encryptedBytes := getFileBytes(inputFile)
	fileKey := getFileBytes(inputFile + ".key")
	decryptedData := decrypt(encryptedBytes, fileKey)
	writeDecryptedData(decryptedData)
	if cleanFlag {
		removeEvidence()
	}
}

func decrypt(encryptedData, key []byte) []byte {
	return encrypt(encryptedData, key)
}

// End Decryption Code

func main() {
	verifyFlags()

	if !decryptFlag {
		encrytFile()
	} else {
		decryptFile()
	}
}
