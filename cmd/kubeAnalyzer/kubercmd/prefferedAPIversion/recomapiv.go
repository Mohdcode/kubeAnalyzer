/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd"
	"github.com/mohdcode/kubeAnalyzer/internal/analyzer"
	"github.com/spf13/cobra"
)

// recomapivCmd represents the recomapiv command
var recomapivCmd = &cobra.Command{
	Use:   "recomapiv",
	Short: "Check prefered Kubernetes API versions",
	Long: `This command analyzes the health of Kubernetes API versions for resources in a given namespace.
It checks the preferred API version of resources and prints the result.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, _ := cmd.Flags().GetString("config")

		err := analyzer.GetPreferedAPIversion(configPath)
		if err != nil {
			log.Print(err)
		}
	},
}

func init() {
	kubercmd.CheckCmd.AddCommand(recomapivCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recomapivCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recomapivCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
