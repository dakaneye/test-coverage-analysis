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
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/dakaneye/test-coverage-analysis"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-coverage-analysis",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {

			// Open our jsonFile
			jsonFile, err := os.Open("/Users/samdacanay/dev/repos/anchore-engine/.tox/unit-test-report.log")
			// if we os.Open returns an error then handle it
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Successfully Opened unit test report log")
			// defer the closing of our jsonFile so that we can parse it later on
			defer jsonFile.Close()

			// read our opened jsonFile as a byte array.
			byteValue, _ := ioutil.ReadAll(jsonFile)

			// we initialize our Users array
			var testResults Results

			// we unmarshal our byteArray which contains our
			// jsonFile's content into 'users' which we defined above
			json.Unmarshal(byteValue, &users)

			// we iterate through every user within our users array and
			// print out the user Type, their name, and their facebook url
			// as just an example
			for i := 0; i < len(users.Users); i++ {
				fmt.Println("User Type: " + users.Users[i].Type)
				fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
				fmt.Println("User Name: " + users.Users[i].Name)
				fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
			}
		},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-coverage-analysis.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".test-coverage-analysis" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".test-coverage-analysis")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
