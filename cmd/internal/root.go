package internal // import "go.fraixed.es/onepagestaticsite/cmd/internal"

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"go.fraixed.es/onepagestaticsite"
)

var cfgFile string // nolint: gochecknoglobals

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{ // nolint: gochecknoglobals
	Use:   "one-page-static-site",
	Short: "Render the template using the content of the Markdown file",
	Long: `Render the Go template file injecting the rendered content of the
Mardown file where the {{.Content}} variable has been placed and write the
output to the indicated file.
NOTE all the input files are considering coming from a trusted source, if
that isn't your case, then DO NOT USE this tool.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if err := onepagestaticsite.Render(
			viper.GetString("template"), viper.GetString("content"), viper.GetString("output"),
		); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// nolint: gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	// Configuration file flag
	rootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "", "config file (default is $HOME/.one-page-static-site.yaml)",
	)

	// Tool flags
	rootCmd.Flags().StringP(
		"template", "t", "src/layout.html",
		`File path of the Go HTML template to use where to render the content.
Can be absolute or relative to the directory where the command is executed.
The path separator is used accordingly the used OS.
`,
	)
	rootCmd.Flags().StringP(
		"content", "c", "content/post.md",
		"File path of the Markdown file which contains the content to render into the layout",
	)
	rootCmd.Flags().StringP(
		"output", "o", "build/index.html",
		"File path where the rendered HTML will be written",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".one-page-static-site" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".one-page-static-site")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// read from flags and or its defaults when no config file
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		fmt.Printf("Unexpected error by viper.BindPFlags: %+v\n", err)
	}
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
