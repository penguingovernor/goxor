// Package xor is a library
// that handles encryptions and decryptions
package xor

import (
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/golang/protobuf/proto"

	"github.com/penguingovernor/goxor/api/protocol"
)

// LoadKey loads the key from the given byte slice
func LoadKey(data []byte) (*protocol.Key, error) {
	key := &protocol.Key{}
	err := proto.Unmarshal(data, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// LoadData loads the data from the given byte slice
func LoadData(slice []byte) (*protocol.Data, error) {
	data := &protocol.Data{}
	err := proto.Unmarshal(slice, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GenerateKey generates a properly formatted key
// that contains the provided signature and key.
func GenerateKey(key, signature []byte) *protocol.Key {
	hashedData := sha256.Sum256(signature)
	return &protocol.Key{
		PayLoad:   key,
		Signature: hashedData[:],
	}
}

// GenerateData generates properly formatted data
// based upon the provided signature and data.
func GenerateData(data, signature []byte) *protocol.Data {
	hashedData := sha256.Sum256(signature)
	return &protocol.Data{
		PayLoad:   data,
		Signature: hashedData[:],
	}
}

// Write writes the message to w.
func Write(w io.Writer, message proto.Message) error {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	if _, err := w.Write(bytes); err != nil {
		return err
	}

	return nil
}

// WriteKey writes the provided key to w.
func WriteKey(w io.Writer, key *protocol.Key) error {
	return Write(w, key)
}

// WriteData writes the provided data to w.
func WriteData(w io.Writer, data *protocol.Data) error {
	return Write(w, data)
}

// Encrypt encrypts the given data with the provided
// key. It returns encrypted data that follows the goxor protocol.
func Encrypt(data *protocol.Data, key *protocol.Key) (*protocol.Data, error) {

	if !validate(data.Signature, key.Signature) {
		return nil, fmt.Errorf("data signature doese not match key signature: got %s, wanted %s", data.Signature, key.Signature)
	}

	eData := &protocol.Data{
		PayLoad:   []byte{},
		Signature: key.Signature,
	}

	for i := range data.PayLoad {

		eData.PayLoad = append(eData.PayLoad, data.PayLoad[i]^key.PayLoad[i%len(key.PayLoad)])

	}

	return eData, nil
}

// Decrypt encrypts the given data with the provided
// key. It returns encrypted data that follows the goxor protocol.
func Decrypt(data *protocol.Data, key *protocol.Key) (*protocol.Data, error) {
	return Encrypt(data, key)
}

func validate(x, y []byte) bool {

	if (x == nil) != (y == nil) {
		return false
	}

	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
