// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/ory/x/cmdx"
	"github.com/ory/x/flagx"
	"github.com/spf13/cobra"
	"time"
)

// quickstartCommand represents the quickstart command
var quickstartCommand = &cobra.Command{
	Use:   "quickstart <endpoint>",
	Short: "Get started with zero configuration and protect a local API or web service",
	Long: `This command follows the quickstart guide on the Ory cloud platform. Assuming that you are running
some type of HTTP(s) web service locally (e.g. NodeJS, Java, PHP,) - be it an API or something that serves HTML, images,
or some other type of data over HTTP, this command will add authentication and authorization to that service without
any additional work.

Given a web service that is running on http://localhost:1234, you would run the quickstart command as follows:

$ ory quickstart -k <your-api-key> http://localhost:1234

All requests to http://localhost:1234/ now require authentication and authorization when accessed via:

	http://localhost:3313

Requests to http://localhost:3313/ will be forwarded to your service (http://localhost:1234/) when authorized:

* Accessing http://localhost:3313/foo/bar will proxy the request to http://localhost:1234/foo/bar when a valid user
  session exists.
* Accessing http://localhost:3313/foo/bar directly in the browser will redirect the browser to the login screen
  when no valid user session exists.
* Accessing http://localhost:3313/foo/bar via an API request (e.g. AJAX) will respond with 401 Unauthorized and a JSON
  containing information how to continue.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdx.ExactArgs(cmd, args, 1)
		port := flagx.MustGetInt(cmd, "port")
		host := flagx.MustGetString(cmd, "host")
		if host == "" {
			host = "localhost"
		}

		fmt.Printf(`Web service located at %s is now protected with authentication and authorization at:

	http://%s:%d/
`, args[0], host, port)

		for {
			time.Sleep(time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(quickstartCommand)

	quickstartCommand.PersistentFlags().StringP("key", "k", "", "The API key generated in the Ory Cloud Console")
	quickstartCommand.PersistentFlags().IntP("port", "p", 4000, "The port Ory should listen on")
	quickstartCommand.PersistentFlags().String("host", "", "The host Ory should listen on")
}
