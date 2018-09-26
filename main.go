// GoXor is a command line tool that encrypts and decrypts data.
//
// For usagae run `goxor --help`
//
// Encryption Specifics:
// Unless specified otherwise, encrypting a file yields a file with a xor file extension (.xor)
// and a key (.xor.key) with the word out preceding the extension.
//
// Without the key file the encrypted data is utterly useless, unless of coure you're a cryptanalyst.

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

package main

import "github.com/penguingovernor/goxor/cmd"

func main() {
	cmd.Execute()
}
