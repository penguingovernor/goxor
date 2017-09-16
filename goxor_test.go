package main

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestXor is a unit test to ensure that
// encryption and decryption went smoothly
func TestXor(t *testing.T) {
	sampleText := []byte("Sample")
	err := ioutil.WriteFile("sample.file", sampleText, 0666)
	if err != nil {
		t.Error(err)
	}

	fileData := getFileBytes("sample.file")
	fileKey := generateFileKey(len(fileData))

	encryptedData := encrypt(fileData, fileKey)

	decryptedBytes := decrypt(encryptedData, fileKey)

	if string(decryptedBytes) != string(sampleText) {
		t.Errorf("Decryption error: got %s wanted %s\n", string(decryptedBytes), string(sampleText))
	}

	err = os.Remove("sample.file")
	if err != nil {
		t.Error(err)
	}
}
