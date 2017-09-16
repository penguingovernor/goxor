# goxor 
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
go test -v github.com/penguingovernor/goxor
```
### Documentation 
After installing, you can use `go doc` to get documenation:
```shell
go doc github.com/penguingovernor/goxor
```

### Running goxorcrypt 
`goxor` can be run in two modes of operation: encrypt and decrypt mode.

Example:
```shell
# Encrypting a file 
goxor --input="exampleFile" # 'goxor -i exampleFile' works too!

# Decrypting a file 
goxor --decrypt --input="exampleFile.gxef" # Again, 'goxor -d -i exampleFile.gxef' works as well!
```

The `goxor` command can encrypt and decrypt files. It supports the following flags:

* `--decrypt` or `-d`: Toggle decryption mode. Omitting this flag causes `goxor` to be encryption mode.
* `--input=` or `-i`: The file to encrypted or decrypted. In encryption mode, omitting this flag causes `goxor` to read from stdin
* `--output=` or `-o`: The desired file name to output the encrypted or decrypted data to. Omitting this flag in encryption mode causes `goxor` to output to fileName.gxef. Omitting this flag in decryption mode causes `goxor` to output to stdout.
* `--clean` or `-c`: Remove encrypted data and key file upon completion of decryption.