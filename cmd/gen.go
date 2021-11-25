/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/dk-sirius/db-builder/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// forkCmd represents the knife command
var forkCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate db table",
	Long:  `fork -n a -t b  -f define file.`,
	Run: func(cmd *cobra.Command, args []string) {
		Exec(cmd, db.Generate)
	},
}

func Exec(cmd *cobra.Command, handle func(dbName, tableName, path string)) {
	handle(
		flagValue(cmd.Flag("name"), true),
		flagValue(cmd.Flag("table"), true),
		flagValue(cmd.Flag("file"), false),
	)
}

func flagValue(flag *pflag.Flag, must bool) string {
	if flag != nil && flag.Value.String() != "" {
		return flag.Value.String()
	} else {
		if must {
			panic("not found db name")
		}
	}
	return ""
}

func init() {
	rootCmd.AddCommand(forkCmd)
	forkCmd.Flags().StringP("name", "n", "", "database name")
	forkCmd.Flags().StringP("table", "t", "", "database table name")
	forkCmd.Flags().StringP("file", "f", "", "the abstract implement file path of table")
}
