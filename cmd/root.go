package cmd

import (
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "getsize",
	Short: "List the size of a local directory.",
	Long:  `This command will display the size of a directory with several different options.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	//flags
	var Verbose bool
	var Debug bool
	var Highlight int
	var Path string

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	//setting flags using viper
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))

	rootCmd.PersistentFlags().IntVarP(&Highlight, "highlight", "", 500, "Highlight files/directories over this threshold, in MB")
	viper.BindPFlag("highlight", rootCmd.PersistentFlags().Lookup("highlight"))

	rootCmd.PersistentFlags().StringVarP(&Path, "path", "p", "Define the path to scan.", "")
	rootCmd.MarkPersistentFlagRequired("path")
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
