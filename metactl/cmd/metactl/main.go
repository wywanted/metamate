package main

import (
	"fmt"
	"github.com/metamatex/metamatemono/metactl/cmd/metactl/v0"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "metactl",
	Short: "",
	Long:  "",
}

func init() {
	v0.AddV0(rootCmd, true)
	v0.AddV0(rootCmd, false)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
