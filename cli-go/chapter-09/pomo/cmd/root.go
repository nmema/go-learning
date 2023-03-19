/*
Copyright © 2023 crocs-rosas
Copyrights apply to this source code.
Check LICENSE for details.
*/
package cmd

import (
	"io"
	"os"
	"time"

	"cli-go/chapter-09/pomo/app"
	"cli-go/chapter-09/pomo/pomodoro"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "Interactive Pomodoro Timer",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := getRepo()
		if err != nil {
			return err
		}

		config := pomodoro.NewConfig(
			repo,
			viper.GetDuration("pomo"),
			viper.GetDuration("short"),
			viper.GetDuration("long"),
		)

		return rootAction(os.Stdout, config)
	},
}

func rootAction(out io.Writer, config *pomodoro.IntervalConfig) error {
	a, err := app.New(config)
	if err != nil {
		return err
	}

	return a.Run()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pomo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().DurationP("pomo", "p", 25*time.Minute,
		"Pomodoro duration")
	rootCmd.Flags().DurationP("short", "s", 5*time.Minute,
		"Short break duration")
	rootCmd.Flags().DurationP("long", "l", 15*time.Minute,
		"Long break duration")

	viper.BindPFlag("pomo", rootCmd.Flags().Lookup("pomo"))
	viper.BindPFlag("short", rootCmd.Flags().Lookup("short"))
	viper.BindPFlag("long", rootCmd.Flags().Lookup("long"))
}
