# goxor #
[![Go Report Card](https://goreportcard.com/badge/github.com/PenguinGovernor/goxor)](https://goreportcard.com/report/github.com/PenguinGovernor/goxor) [![Build Status](https://travis-ci.com/PenguinGovernor/goxor.svg?branch=master)](https://travis-ci.com/PenguinGovernor/goxor)

GoXor is an encryption method written in the [Go programming language](https://golang.org/), and is based on a [xor cipher](https://en.wikipedia.org/wiki/XOR_cipher) in conjunction with a [one-time pad](https://en.wikipedia.org/wiki/One-time_pad). 

### Download and Install 
1. Install and configure Go for you operating system. [See Here](https://golang.org/doc/install)
2. Once Go is installed and configured, run this command to install the `goxor` tool 

```shell
go get github.com/penguingovernor/goxor
``` 
### Testing 
To ensure that installation went smoothly run the `go test` tool 
```shell
go test -v github.com/penguingovernor/goxor/xor
```
### Documentation 
After installing, you can use `goxor help` to get documentation:
```shell
goxor help
```

### Running goxor
`goxor` has two sub-commands: encrypt and decrypt.

Example:
```shell
# Encrypts input hello.txt using a one time pad as the key, 'goxor' as the signature, outputs out.xor and out.xor.key
goxor encrypt -i hello.txt
# Decrypts input out.xor using out.xor.key as the key, this will output to stdout
goxor decrypt -i out.xor -k out.xor.key
```

The `goxor encrypt` command can encrypt files. It supports the following flags:

* `--input=` or `-i`: The file to be encrypted. If the file cannot be found then the input is treated as a string. If omitted `goxor encrypt` will read from stdin

* `--output=` or `-o`: The desired file name to output the encrypted data to. Omitting this flag mode causes `goxor encrypt` to output to out.xor. If the string 'stdout' is passed to `goxor encrypt --output` then the encrypted data is written to stdout

* `--key_out` or `-K`: The desired file name to output the key to. Omitting this flag causes `goxor encrypt` to output to out.xor.key. If the string 'stdout' is passed to `goxor encrypt --key_out` then the key is written to stdout

* `--signature` or `-s`: The file or string to use as the encryption signature. If omitted the string `goxor` is used as the signature. If stdin is passed then the input from stdin is used as the signature.

* `--key` or `-k`: The file or string to use as the encryption key. If omitted a one time pad is used as the key. If stdin is passed then the input from stdin is used as the key.

The `goxor decrypt` command can decrypt files. It supports the following flags:

* `--input=` or `-i`: The file to be decrypted. 

* `--output=` or `-o`: The desired file name to output the decrypted data to. Omitting this flag mode causes `goxor decrypt` to output to stdout.

* `--key` or `-k`: The file to use as the decryption key. 
