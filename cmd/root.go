// Copyright © 2017 Kris Nova <kris@nivenly.com>
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
	"github.com/kris-nova/tt/interpreter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "tt",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := RunTT(o)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tt.yaml)")
	if len(os.Args) < 2 {
		RootCmd.Help()
		os.Exit(1)
	}
	if strings.Contains(os.Args[1], "--") {
		RootCmd.Help()
		os.Exit(1)
	}
	o.EntryFile = os.Args[1]
}

var o = &ttOptions{
	EntryFile: "",
}

type ttOptions struct {
	EntryFile string
}

func RunTT(options *ttOptions) error {
	err := interpreter.ParseFile(options.EntryFile)
	if err != nil {
		return err
	}
	return nil
}
