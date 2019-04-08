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
	Use:   "minepong-cli [flags]",
	Short: "A CLI for get information about Minecraft server",
	Example: `
minepong-cli --host mc1 --port 25575
minepong-cli --port 25575
MC_PORT=25575 minepong-cli
`,
	Long: `
minepong-cli is a CLI for get information about Minecraft server metadata.
`,
	Run: func(cmd *cobra.Command, args []string) {
		hostPort := net.JoinHostPort(viper.GetString("host"), strconv.Itoa(viper.GetInt("port")))
		pretty := viper.GetBool("pretty")

		cli.Execute(hostPort, pretty)
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

	RootCmd.PersistentFlags().String("host", "localhost", "server's hostname")
	RootCmd.PersistentFlags().Int("port", 25565, "Server's port")
	RootCmd.PersistentFlags().BoolP("pretty", "p", false, "Use pretty printing")
	err := viper.BindPFlags(RootCmd.PersistentFlags())
	if err != nil {
		log.Fatal(err)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// This will allow for env vars like MC_PORT
	viper.SetEnvPrefix("mc")
	viper.AutomaticEnv() // read in environment variables that match
}
