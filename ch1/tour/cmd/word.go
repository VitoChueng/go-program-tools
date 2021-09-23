package cmd

import "github.com/spf13/cobra"

var wordCmd = &cobra.Command{
	Use:"word",
	Short: "word transfer",
	Long: "multiple lang",
	Run: func(cmd *cobra.Command,args []string){},
}

func init(){}