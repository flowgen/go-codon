package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	codon_generator "go-codon/generator"
)

var (
	generateClients  bool
	generateWorkflow bool
	generateServer   bool
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates service code alongside codon service specification",
	Long: `After you have run init on your project folder you should edit
the specification and configuration files according to your project. Run
this command to finally generate code according to the specification. Your
working directory must be your project directory. Not specifying any
generator flags will generate everything.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !generateClients && !generateWorkflow && !generateServer {
			generateClients = true
			generateWorkflow = true
			generateServer = true
		}
		opts := codon_generator.GenOpts{
			Language:         "golang",
			GenerateClients:  generateClients,
			GenerateWorkflow: generateWorkflow,
			GenerateServer:   generateServer,
		}
		if codon_generator.Generate(opts) {
			log.Println("Generate successful")
		} else {
			log.Println("Generate failed")
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().BoolVarP(&generateClients, "clients", "c", false, "Generate clients")
	generateCmd.Flags().BoolVarP(&generateWorkflow, "workflow", "w", false, "Generate workflow")
	generateCmd.Flags().BoolVarP(&generateServer, "server", "s", false, "Generate server")
}
