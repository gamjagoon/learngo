package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
  Use: "do",
  Short : "dos a task to your task list.",
  Run : func(cmd *cobra.Command, args []string) {
    fmt.Println("do called")
  },
}

func init() {
  RootCmd.AddCommand(doCmd)
}