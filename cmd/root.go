package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/text"
	"github.com/spf13/cobra"

	"github.com/madflow/skate-ipsum/generator"
)

type SkateOpt struct {
	Paragraphs           int
	PrintWidth           int
	LeadWithDolorSitAmet bool
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
	rootCmd.Flags().IntVarP(&skateOpt.PrintWidth, "width", "w", 120, "print width")
	rootCmd.Flags().BoolVarP(&skateOpt.LeadWithDolorSitAmet, "lead", "l", false, "lead with dolor sit amet")
}

func ipsumRun(cmd *cobra.Command, args []string) {
	leadingText := "Skate ipsum dolor sit amet"

	if skateOpt.Paragraphs > 1080 {
		panic("too many paragraphs")
	}

	outputTexts := generator.IpsumArray(skateOpt.Paragraphs)

	needsLead := skateOpt.LeadWithDolorSitAmet

	fmt.Println()
	for _, outputText := range outputTexts {
		if needsLead {
			outputText = leadingText + ", " + outputText
			needsLead = false
		}
		fmt.Println(text.WrapSoft(outputText, skateOpt.PrintWidth))
		fmt.Println()
	}
}
