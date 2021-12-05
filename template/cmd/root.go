package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrg/xdg"
)

const appName = "{{.Name}}"

func initConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(filepath.Join(xdg.ConfigHome, appName))
	v.AddConfigPath(".")
	v.SetConfigName("config")

	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	v.ReadInConfig()

	return v
}

func Execute() {
	rootCmd := rootCmd(initConfig())
	rootCmd.AddCommand(exampleCmd())
	rootCmd.AddCommand(versionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd(v *viper.Viper) *cobra.Command {
	return &cobra.Command{
		Use:   appName,
		Short: appName,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bindFlagToConfig(cmd, v)
		},
	}
}

func bindFlagToConfig(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", appName, envVarSuffix))
		}
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
