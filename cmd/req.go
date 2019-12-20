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
	"os"
	"strings"

	"github.com/seifchen/kcurl/req"

	"github.com/seifchen/kcurl/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// reqCmd represents the req command
var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "request url use name",
	Long: `request url use name and path、params、header and so on, 
	If you use post the default content-type is json.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("At least one args")
			os.Exit(1)
		}
		switch strings.ToLower(contentType) {
		case "json":
			headers = append(headers, "Content-type:application/json")
		case "form":
			headers = append(headers, "Content-type:application/x-www-form-urlencoded")
		default:
			log.Printf("--type Not support Content-Type:%s\nUse \"kcurl req --help\" for more information", contentType)
			os.Exit(1)
		}
		option = strings.ToUpper(option)
		if option != "GET" && option != "POST" && option != "OPTIONS" {
			log.Printf("--option Not support:%s\nUse \"kcurl req --help\" for more information", option)
			os.Exit(1)
		}

	},
	Run: reqRun,
}

func reqRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	reqArgs := strings.Join(parameters, "&")
	for _, name := range args {
		for _, item := range items {
			if env != "" && item.Env != env {
				continue
			}
			if name == item.Name {
				err := req.DoReq(item.Url, option, path, headers, reqArgs, body)
				if err != nil {
					log.Printf("req:%v error:%s", item, err.Error())
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
var contentType string

func init() {
	rootCmd.AddCommand(reqCmd)
	reqCmd.Flags().StringVarP(&option, "option", "o", "GET", "option:GET,POST,OPTIONS")
	reqCmd.Flags().StringVarP(&path, "path", "p", "/", "path:get path")
	reqCmd.Flags().StringSliceVarP(&headers, "header", "", nil, "headers:req head")
	reqCmd.Flags().StringSliceVarP(&parameters, "params", "", nil, "parameters")
	reqCmd.Flags().StringVarP(&body, "body", "", "", "request body")
	reqCmd.Flags().StringVarP(&contentType, "type", "", "json", "content-type:json,form;default:json")
}
