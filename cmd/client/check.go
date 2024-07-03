package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tibotix/golanguagetool/pkg/golanguagetool"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cobra.CheckErr("No input")
	}

	client, err := GetLanguageToolClient()
	cobra.CheckErr(err)

	var level golanguagetool.CheckLevel = golanguagetool.CheckLevelDefault
	if viper.GetBool("check.picky") {
		level = golanguagetool.CheckLevelPicky
	}
	checkOptions := golanguagetool.CheckOptions{
		Language:           viper.GetString("check.language"),
		Dicts:              viper.GetStringSlice("check.dicts"),
		MotherTongue:       viper.GetString("check.mother-tongue"),
		PreferredVariants:  viper.GetStringSlice("check.preferred-variants"),
		EnabledRules:       viper.GetIntSlice("check.enabled-rules"),
		DisabledRules:      viper.GetIntSlice("check.disabled-rules"),
		EnabledCategories:  viper.GetIntSlice("check.enabled-rules"),
		DisabledCategories: viper.GetIntSlice("check.disabled-categories"),
		EnabledOnly:        viper.GetBool("check.enabled-only"),
		Level:              level,
	}

	printOptions := PrintOptions{
		ExplainRule:        viper.GetBool("check.explain-rule"),
		ShowRules:          viper.GetBool("check.rules"),
		ShowRuleCategories: viper.GetBool("check.rule-categories"),
	}

	for _, file := range args {
		data, err := client.OpenFile(file)
		cobra.CheckErr(err)

		fileType := golanguagetool.DetermineFileType(file)
		backgroundColor.Printf("Detected file type '%s'", fileType.String())
		if input_type := viper.GetString("check.input-type"); input_type != "auto" {
			fileType = golanguagetool.GetFileTypeFromString(input_type)
			backgroundColor.Printf(", but using '%s' because of setting", fileType.String())
		}
		backgroundColor.Println(".")

		text := golanguagetool.Text{
			Contents: data,
			FileType: fileType,
		}
		results, err := client.CheckText(text, &checkOptions)
		cobra.CheckErr(err)
		PrintCheckResults(results, &printOptions)
	}

}

type PrintOptions struct {
	// showLineNumbers    bool
	ExplainRule        bool
	ShowRules          bool
	ShowRuleCategories bool
}

const (
	surrSelf rune = 0x10000
	maxRune  rune = '\U0010FFFF' // Maximum valid Unicode code point.
)

func codeUnitToByteIndex(s string, n int) int {
	pos := 0
	for i := 0; i < n; i++ {
		r, w := utf8.DecodeRuneInString(s[pos:])
		// if r is a surrogate pair, increase index additionally
		if surrSelf <= r && r <= maxRune {
			i++
		}
		pos += w
	}
	return pos
}

func PrintCheckResults(results *golanguagetool.CheckResults, options *PrintOptions) {
	backgroundColor.Printf("%s detected.\n", *results.Language.DetectedLanguage.Name)
	if *results.Language.Code != *results.Language.DetectedLanguage.Code {
		backgroundColor.Printf("checking as %s text because of setting.\n", *results.Language.Name)
	}
	fmt.Println()

	ruleExplanations := make(map[string]string)

	for _, match := range results.Matches {
		contextText := *match.Context.Text
		offset := codeUnitToByteIndex(contextText, int(*match.Context.Offset))
		length := codeUnitToByteIndex(contextText[offset:], int(*match.Context.Length))
		endpos := offset + length
		padding := runewidth.StringWidth(contextText[:offset])
		fmt.Printf("%+v,%+v,%+v,%+v,%+v\n", contextText, offset, length, *match.Context.Offset, *match.Context.Length)

		fmt.Printf("%s\n", *match.Message)
		fmt.Printf("  %s %s%s%s\n", cross, backgroundText(contextText[:offset]), redText(contextText[offset:endpos]), backgroundText(contextText[endpos:]))
		fmt.Printf("    %s%s\n", strings.Repeat(" ", padding), redText(strings.Repeat("^", int(*match.Context.Length))))
		for _, replacement := range match.Replacements[:min(len(match.Replacements), 5)] {
			fmt.Printf("  %s %s%s%s\n", tick, backgroundText(contextText[:offset]), greenText(replacement.Value), backgroundText(contextText[endpos:]))
		}
		if match.Rule != nil {
			if options.ShowRules {
				fmt.Printf("  %s%s\n", backgroundText(*match.Rule.ID, ": "), *match.Rule.Description)
			}
			if options.ShowRuleCategories {
				fmt.Printf("  %s%s\n", backgroundText(match.Rule.Category.ID, ": "), match.Rule.Category.Name)
			}
			if options.ExplainRule && len(match.Rule.Urls) > 0 {
				ruleExplanations[*match.Rule.Description] = match.Rule.Urls[0].Value
			}
		}
		fmt.Println()
	}

	if options.ExplainRule {
		for desc, val := range ruleExplanations {
			fmt.Printf("%s: %s\n", desc, val)
		}
		fmt.Println()
	}

	backgroundColor.Printf("Text checked by %s %s\n", *results.Software.Name, *results.Software.Version)
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringP("language", "l", "auto", "A language code like `en-US`, `de-DE`, etc. or `auto`.")
	viper.BindPFlag("check.language", checkCmd.Flags().Lookup("language"))
	checkCmd.Flags().StringSliceP("dicts", "d", []string{}, "List of dictionaries to include words from; uses special default dictionary if this is unset.")
	viper.BindPFlag("check.dicts", checkCmd.Flags().Lookup("dicts"))
	checkCmd.Flags().StringP("mother-tongue", "m", "", "A language code of the user's native language, enabling false friends checks for some language pairs.")
	viper.BindPFlag("check.mother-tongue", checkCmd.Flags().Lookup("mother-tongue"))
	checkCmd.Flags().IntSliceP("preferred-variants", "p", []int{}, "List of preferred language variants. The language detector used with language=auto can detect e.g. English, but it cannot decide whether British English or American English is used. Thus this parameter can be used to specify the preferred variants like en-GB and de-AT. Only available with language=auto.")
	viper.BindPFlag("check.preferred-variants", checkCmd.Flags().Lookup("preferred-variants"))
	checkCmd.Flags().IntSlice("enabled-rules", []int{}, "List of IDs of rules to be enabled.")
	viper.BindPFlag("check.enabled-rules", checkCmd.Flags().Lookup("enabled-rules"))
	checkCmd.Flags().IntSlice("disabled-rules", []int{}, "List of IDs of rules to be disabled.")
	viper.BindPFlag("check.disabled-rules", checkCmd.Flags().Lookup("disabled-rules"))
	checkCmd.Flags().IntSlice("enabled-categories", []int{}, "List of IDs of categories to be enabled.")
	viper.BindPFlag("check.enabled-categories", checkCmd.Flags().Lookup("enabled-categories"))
	checkCmd.Flags().IntSlice("disabled-categories", []int{}, "List of IDs of categories to be disabled.")
	viper.BindPFlag("check.disabled-categories", checkCmd.Flags().Lookup("disabled-categories"))
	checkCmd.Flags().IntSlice("enabled-only", []int{}, "Enable only the rules and categories whose IDs are specified with --enabled-rules or --enabled-categories.")
	viper.BindPFlag("check.enabled-only", checkCmd.Flags().Lookup("enabled-only"))
	checkCmd.Flags().Bool("picky", false, "If enabled, additional rules will be activated.Enable only the rules and categories whose IDs are specified with --enabled-rules or --enabled-categories.")
	viper.BindPFlag("check.picky", checkCmd.Flags().Lookup("picky"))
	checkCmd.Flags().BoolP("explain-rule", "u", false, "Print URLs with more information about rules.")
	viper.BindPFlag("check.explain-rule", checkCmd.Flags().Lookup("explain-rule"))
	checkCmd.Flags().BoolP("rules", "r", false, "Show the matching rules.")
	viper.BindPFlag("check.rules", checkCmd.Flags().Lookup("rules"))
	checkCmd.Flags().Bool("rule-categories", false, "Show the categories of matching rules.")
	viper.BindPFlag("check.rule-categories", checkCmd.Flags().Lookup("rule-categories"))

	allowedFileTypes := append(golanguagetool.SupportedFileTypes, "auto")
	checkCmd.Flags().VarP(newEnum(allowedFileTypes, "auto"), "input-type", "t", fmt.Sprintf("Input text type. Use `auto` to detect it based on file extension. Available are {%s}", strings.Join(allowedFileTypes, ",")))
	viper.BindPFlag("check.input-type", checkCmd.Flags().Lookup("input-type"))
}
