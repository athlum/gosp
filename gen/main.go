package main

import (
	"github.com/athlum/gosp/log"
	"github.com/athlum/pkg/utils"
	"github.com/spf13/cobra"
	"os"
)

var dir string

var rootCmd = &cobra.Command{
	Use:   "gosp-gen",
	Short: "gosp orm implement generator",
	Long:  "gosp orm implement generator",
	Run: func(cmd *cobra.Command, args []string) {
		if ok := utils.RunningHelp(cmd, args); ok {
			return
		}
		if dir == "" {
			log.Fatal("Please give a valid path, run `gosp-gen help` for more.")
		}
		if err := parseDir(dir); err != nil {
			log.Fatal("Parse dir failed: %s", err.Error())
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "D", "./", "The path of your model package")
	rootCmd.Flags().Parse(os.Args)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
