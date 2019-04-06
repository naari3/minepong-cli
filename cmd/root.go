package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/naari3/mc-poorcount/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd for all
var RootCmd = &cobra.Command{
	Use:   "mc-poorcount [flags]",
	Short: "A CLI for count number of people present on Minecraft server tool",
	Example: `
mc-poorcount --host mc1 --port 25575
mc-poorcount --port 25575
RCON_PORT=25575 mc-poorcount
`,
	Long: `
mc-poorcount is a CLI for count number of people present on Minecraft server tool.
`,
	Run: func(cmd *cobra.Command, args []string) {

		hostPort := net.JoinHostPort(viper.GetString("host"), strconv.Itoa(viper.GetInt("port")))
		password := viper.GetString("password")

		cli.Execute(hostPort, password)
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().String("host", "localhost", "RCON server's hostname")
	RootCmd.PersistentFlags().String("password", "", "RCON server's password")
	RootCmd.PersistentFlags().Int("port", 25575, "Server's RCON port")
	err := viper.BindPFlags(RootCmd.PersistentFlags())
	if err != nil {
		log.Fatal(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// This will allow for env vars like RCON_PORT
	viper.SetEnvPrefix("rcon")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
