// GoXor is a command line tool that encrypts and decrypts data.
//
// For usagae run `goxor --help`
//
// Encryption Specifics:
// Unless specified otherwise, encrypting a file yields a file with a go xor encrypted file extension (.gxef)
// and a key (.gxef.key) with the original file name preceding the extension.
// It is crucial that the key file is not renamed as `goxor` will look for this file in decryption.
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

func writeEncryptedData(in, out string, data []byte) {
	fileName := outputFile

	if out == "" {
		fileName = in + ".gxef"
	}

	err := ioutil.WriteFile(fileName, data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func writeKey(in, out string, key []byte) {
	fileName := out + ".key"

	if out == "" {
		fileName = in + ".gxef.key"
	}
	err := ioutil.WriteFile(fileName, key, 0666)

	if err != nil {
		log.Fatal(err)
	}

}

func writeDecryptedData(out string, data []byte) {
	if out == "" {
		fmt.Print(string(data))
		return
	}

	err := ioutil.WriteFile(out, data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func removeEvidence(in string) {
	fmt.Println("Cleaning up files...")
	err := os.Remove(in)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove(in + ".key")
	if err != nil {
		log.Fatal(err)
	}
}

// End File Related Code

// Encryption Code

func encrytFile(in, out string) {
	fileBytes := getFileBytes(in)
	fileKey := generateFileKey(len(fileBytes))
	encrytedBytes := encrypt(fileBytes, fileKey)
	writeEncryptedData(in, out, encrytedBytes)
	writeKey(in, out, fileKey)
}

func generateFileKey(n int) []byte {
	fileKey := make([]byte, n)
	_, err := rand.Read(fileKey)
	if err != nil {
		log.Fatal(err)
	}
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

func decryptFile(in, out string) {
	encryptedBytes := getFileBytes(in)
	fileKey := getFileBytes(in + ".key")
	decryptedData := decrypt(encryptedBytes, fileKey)
	writeDecryptedData(out, decryptedData)
	if cleanFlag {
		removeEvidence(in)
	}
}

func decrypt(encryptedData, key []byte) []byte {
	return encrypt(encryptedData, key)
}

// End Decryption Code

func main() {
	verifyFlags()

	if !decryptFlag {
		encrytFile(inputFile, outputFile)
	} else {
		decryptFile(inputFile, outputFile)
	}
}
