package cmd

import (
	"fmt"
	"os"
	"seo-helper/pkg/zola"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var contentFolder string
var filePath string

var zolaCmd = &cobra.Command{
	Use:   "zola",
	Short: "Add description and keywords to Zola content",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load()
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), err)
			os.Exit(1)
		}

		if filePath != "" {
			if err := zola.ProcessFile(filePath); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			}
		} else {
			if err := zola.ProcessFolder(contentFolder); err != nil {
				fmt.Fprintln(cmd.ErrOrStderr(), err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(zolaCmd)
	zolaCmd.Flags().StringVarP(&contentFolder, "content-folder", "c", "content", "Content folder")
	zolaCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "file path")
	zolaCmd.MarkFlagsOneRequired("content-folder", "file-path")
}
