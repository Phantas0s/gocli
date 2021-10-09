package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrg/xdg"
)

var key string

const envPrefix = "gocket"

func initConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(filepath.Join(xdg.ConfigHome, "gocket"))
	v.AddConfigPath(".")
	v.SetConfigName("config")

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	v.ReadInConfig()

	return v
}

func Execute() {
	rootCmd := rootCmd(initConfig())
	rootCmd.AddCommand(otherCmd())
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Some key")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd(v *viper.Viper) *cobra.Command {
	return &cobra.Command{
		Use:   "gocli",
		Short: "gocli",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bindFlagToConfig(cmd, v)
		},
	}
}

func bindFlagToConfig(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func prompt(message string) bool {
	os.Stdout.WriteString(message + " (y/n)")
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if strings.Trim(string(i), "\n") == "y" {
		return true
	} else {
		os.Stdout.WriteString("Aborted.")
		return false
	}
}
