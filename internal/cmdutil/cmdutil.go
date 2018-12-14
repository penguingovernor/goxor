package cmdutil

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/penguingovernor/goxor/api/protocol"
	"github.com/penguingovernor/goxor/pkg/xor"
)

// Decrypt decrypts the data based on the interpretations of in,key, and out
func Decrypt(in, key, out string) error {
	input, err := GetInputFromFile(in)
	if err != nil {
		return err
	}

	keyData, err := GetKeyFromFile(key)
	if err != nil {
		return err
	}

	data, err := xor.Decrypt(input, keyData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done decrypting file")
	WriteDecryptedData(data, out)

	return nil
}

// Encrypt encrypts the data based on the interpretations of in,key, and out
func Encrypt(inputFlag, keyFlag, signatureFlag, outputFlag, outputKeyFlag string) error {

	// Get the bytes from user input
	inputBytes := GetInputFromUser(inputFlag)
	keyBytes := GetKeyFromUser(keyFlag, len(inputBytes))
	signatureBytes := GetSigFromUser(signatureFlag)

	// Generate the appropriate data
	data := xor.GenerateData(inputBytes, signatureBytes)
	key := xor.GenerateKey(keyBytes, signatureBytes)

	// Encrypt the data
	eData, err := xor.Encrypt(data, key)
	if err != nil {
		return fmt.Errorf("error while encrypting file: %v", err)
	}

	// Say we're done and output the files
	fmt.Println("Done encrypting file")
	if err := WriteDataToFile(eData, outputFlag); err != nil {
		return err
	}

	if err := WriteKeyToFile(key, outputKeyFlag); err != nil {
		return err
	}

	return nil
}

// WriteDecryptedData writes decrypted data to the interpretation of out
func WriteDecryptedData(data *protocol.Data, out string) error {
	if out == "" {
		fmt.Println("---BEGIN DECRYPTED DATA---")
		if _, err := os.Stdout.Write(data.PayLoad); err != nil {
			log.Fatalf("could not write data: %v", err)
		}
		fmt.Println("---END DECRYPTED DATA---")
		return nil
	}

	file, err := os.Create(out + ".xor")
	if err != nil {
		return fmt.Errorf("Could not create file %s.xor: %v", out, err)
	}

	if _, err := file.Write(data.PayLoad); err != nil {
		return fmt.Errorf("could not write to file %s.xor: %v", out, err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("Could not close file %s.xor: %v", out, err)
	}

	fmt.Println("Data written to:", out, ".xor")
	return nil
}

//WriteDataToFile writes encrypted data to the interpretation of dest
func WriteDataToFile(data *protocol.Data, dest string) error {
	outname := dest

	if strings.ToLower(dest) == "stdout" {
		fmt.Println("---BEGIN ENCRYPTED DATA---")
		xor.WriteData(os.Stdout, data)
		fmt.Println("---END ENCRYPTED DATA---")
		return nil
	}

	if dest == "" {
		outname = "out"
	}

	file, err := os.Create(outname + ".xor")
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", outname+".xor", err)
	}

	if err := xor.WriteData(file, data); err != nil {
		return fmt.Errorf("could not write key %s: %v", outname+".xor", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", outname+".xor.key", err)
	}

	fmt.Println("Data written to:", outname+".xor")
	return nil
}

// WriteKeyToFile writes the key to the interpretation of dest
func WriteKeyToFile(data *protocol.Key, dest string) error {
	outname := dest

	if strings.ToLower(dest) == "stdout" {
		fmt.Println("---BEGIN KEY DATA---")
		xor.WriteKey(os.Stdout, data)
		fmt.Println("---END KEY DATA---")
		return nil
	}

	if dest == "" {
		outname = "out"
	}

	file, err := os.Create(outname + ".xor.key")
	if err != nil {
		return fmt.Errorf("Could not create file %s: %v", outname+".xor.key", err)
	}

	if err := xor.WriteKey(file, data); err != nil {
		return fmt.Errorf("could not write key %s: %v", outname+".xor.key", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", outname+".xor.key", err)
	}

	fmt.Println("Key written to:", outname+".xor.key")
	return nil
}

// GetInputFromUser gets data based on the interpretation of input
func GetInputFromUser(input string) []byte {
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

// GetKeyFromUser gets the key based on the interpretation of input
// parameter length is only used if the input is empty signifying a OTP is to be used
func GetKeyFromUser(input string, length int) []byte {

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

// GetSigFromUser gets the signautre based on the interpretation of input
func GetSigFromUser(input string) []byte {
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

// GetInputFromFile gets data from the file in
func GetInputFromFile(in string) (*protocol.Data, error) {
	if in == "" {
		return nil, fmt.Errorf("missing input file")
	}

	fileBytes, err := ioutil.ReadFile(in)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %v", in, err)
	}

	data, err := xor.LoadData(fileBytes)
	if err != nil {
		return nil, fmt.Errorf("could not load file %s: %v", in, err)
	}

	fmt.Println("Using file:", in, "as input")

	return data, nil
}

// GetKeyFromFile gets the key from the file in
func GetKeyFromFile(in string) (*protocol.Key, error) {
	if in == "" {
		return nil, fmt.Errorf("flag key is required")
	}

	fileBytes, err := ioutil.ReadFile(in)
	if err != nil {
		return nil, fmt.Errorf("could not read %s: %v", in, err)
	}

	data, err := xor.LoadKey(fileBytes)
	if err != nil {
		return nil, fmt.Errorf("could not load file %s: %v", in, err)
	}

	fmt.Println("Using file:", in, "as key")

	return data, nil
}
