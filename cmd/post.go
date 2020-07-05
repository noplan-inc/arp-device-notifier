/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/noplan-inc/arp-device-notifier/lib"
	"github.com/spf13/cobra"
	"log"
	"net/url"
)

var (
	o = &lib.PostOptions{}
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "It is a tool to get a list of devices by Arp protocol and request it to the server.",
	Long: `It is a tool to get a list of devices by Arp protocol and request it to the server. For example:

## Normal
arp-device-notifier -e http://example.com

## Verbose
arp-device-notifier -e http://example.com -v

## Post Interval(default is 10s)
arp-device-notifier -e http://example.com -i 3
`,
	Run: func(cmd *cobra.Command, args []string) {
		validateOptions()
		lib.Run(o)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().StringVarP(&o.Endpoint, "endpoint", "e", "", "A endpoint to send request")
	postCmd.Flags().StringVarP(&o.Authorization, "authorization", "a", "", "A request header to be used for authorization")
	postCmd.Flags().Int64VarP(&o.PostInterval, "post-interval", "i", 10, "Update interval in seconds")
	postCmd.Flags().BoolVarP(&o.Verbose, "verbose", "v", false, "Print debugging messages about its progress")

	// Required Options
	postCmd.MarkFlagRequired("endpoint")
}

func validateOptions() {
	if _, err := url.ParseRequestURI(o.Endpoint); err != nil {
		log.Fatalf("%s is invalid URL", o.Endpoint)
	}
}
