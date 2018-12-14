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
	"log"
	"os"
	"runtime"

	"github.com/penguingovernor/goxor/internal/constants"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goxor",
	Short: "goxor is a cli encryption and decryption tool",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		version, err := cmd.Flags().GetBool("version")
		if err != nil {
			log.Fatalf("internal error: could not get flag %q: %v\n", "version", err)
		}
		if version {
			fmt.Printf("Goxor version %s %s/%s\n", constants.Version, runtime.GOOS, runtime.GOARCH)
		}
	},
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print version infromation and quit")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
