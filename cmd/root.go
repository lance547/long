package cmd

import "github.com/spf13/cobra"

// 根命令设置
var rootCmd = &cobra.Command{}

func Excute() error {
	return rootCmd.Execute()
}
func init() {
	rootCmd.AddCommand(wordCmd) //增加word子命令
}