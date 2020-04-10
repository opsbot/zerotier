package cli

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/opsbot/cli-go/utils"
	"github.com/opsbot/zerotier/api"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// Version is set during build
	Version string
	// GitHash is set during build
	GitHash string
	// BuildTime is set during build
	BuildTime string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zerotier",
	Short: "CLI client for zerotier REST APIs",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.SetVerbosity(cmd.Flag("verbosity").Value.String())
		api.SetEnvironment(viper.GetStringMap(cmd.Flag("env").Value.String()))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, commit string, date string) {
	Version = version
	GitHash = commit
	BuildTime = date

	AddCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// allow user to load a different config file
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.zerotier.yaml)")

	// allow user to specify environemnt
	// this will allow for a single config file to contain values for different environments
	rootCmd.PersistentFlags().StringP("env", "", "default", "the environment config to load")

	// allow user to specify verbosity
	// verbosity level will dynamically change the log output level
	rootCmd.PersistentFlags().CountP("verbosity", "v", "the verbosity level")

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
			log.Error(err)
			os.Exit(1)
		}

		viper.AddConfigPath(fmt.Sprintf("%v/.config", home))
		viper.SetConfigName(".zerotier")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Trace("Using config file:", viper.ConfigFileUsed())
	}
}

// fmt.Printf("Version    : %s\n", Version)
// fmt.Printf("Git Hash   : %s\n", GitHash)
// fmt.Printf("Build Time : %s\n", BuildTime)
