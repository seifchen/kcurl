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
	"log"
	"strings"

	"github.com/seifchen/kcurl/req"

	"github.com/seifchen/kcurl/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reqCmd represents the req command
var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: reqRun,
}

func reqRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	if isJson {
		headers = append(headers, "Content-type:application/json")
	} else {
		headers = append(headers, "Content-type:application/x-www-form-urlencoded")
	}

	reqArgs := strings.Join(parameters, "&")
	for _, name := range args {
		for _, item := range items {
			if name == item.Name && env == item.Env {
				err := req.DoReq(item.Url, option, path, headers, reqArgs, body)
				if err != nil {
					log.Printf("req:%v error:%e", item, err)
				}
			}
		}
	}
}

var option string
var headers []string
var path string
var parameters []string
var body string
var isJson bool

func init() {
	rootCmd.AddCommand(reqCmd)
	reqCmd.Flags().StringVarP(&env, "env", "e", "dev", "env:dev,online")
	reqCmd.Flags().StringVarP(&option, "option", "o", "", "option:GET,POST,OPTIONS")
	reqCmd.Flags().StringVarP(&path, "path", "p", "/", "path:get path")
	reqCmd.Flags().StringSliceVarP(&headers, "headers", "", nil, "headers:req head")
	reqCmd.Flags().StringSliceVarP(&parameters, "params", "", nil, "parameters")
	reqCmd.Flags().StringVarP(&body, "body", "b", "", "request body")
	reqCmd.Flags().BoolVarP(&isJson, "json", "j", true, "json:is json or form")
}
