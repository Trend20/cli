package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// filesCmd represents the files command
var filesCmd = &cobra.Command{
	Use:   "files",
	Short: "Show the largest files in the given path.",
	Long:  `Quickly scan a directory and find large files.`,
	Run: func(cmd *cobra.Command, args []string) {
		for key, value := range viper.GetViper().AllSettings() {
			log.WithFields(log.Fields{
				key: value,
			}).Info("Command Flag")
		}
	},
}

func init() {

	//flags
	var Filecount int
	var Minfilesize int64

	rootCmd.AddCommand(filesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	filesCmd.PersistentFlags().IntVarP(&Filecount, "filecount", "f", 10, "Limit the number of files returned")
	viper.BindPFlag("filecount", filesCmd.PersistentFlags().Lookup("filecount"))

	filesCmd.PersistentFlags().Int64VarP(&Minfilesize, "minfilesize", "", 50, "Minimum size for files in search in MB.")
	viper.BindPFlag("minfilesize", filesCmd.PersistentFlags().Lookup("minfilesize"))
}
