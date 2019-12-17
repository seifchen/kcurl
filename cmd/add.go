/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"log"

	"github.com/seifchen/kcurl/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var env string
var url string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add name url env to local file",
	Long:  `add name url env to local file then you can use name to request`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	for _, arg := range args {
		item := todo.Item{Name: arg}
		item.SetEnv(env)
		item.SetUrl(url)

		items = append(items, item)
	}
	err = todo.SaveItem(viper.GetString("datafile"), items)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

var test string

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&env, "env", "e", "dev", "env:dev,online")
	addCmd.Flags().StringVarP(&url, "url", "u", "", "url")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
