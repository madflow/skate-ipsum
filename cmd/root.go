package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/madflow/skate-ipsum/generator"
)

type SkateOpt struct {
	LeadingText string
	Paragraphs  int
	Sentences   int
	DoLead      bool
}

var skateOpt SkateOpt

var rootCmd = &cobra.Command{
	Use:   "skate-ipsum",
	Short: "A gnarlier ipsum generator",
	Run:   ipsumRun,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&skateOpt.Paragraphs, "paragraphs", "p", 10, "number of paragraphs")
}

func ipsumRun(cmd *cobra.Command, args []string) {
	if skateOpt.Paragraphs > 1080 {
		panic("too many paragraphs")
	}

	outputText := generator.IpsumText(skateOpt.Paragraphs)

	fmt.Println(outputText)
}
