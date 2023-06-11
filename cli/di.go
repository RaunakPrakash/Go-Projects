package cli

import "github.com/spf13/cobra"

type DIContainer struct {
	cli func() *cobra.Command
}

func NewDIContainer(use, description string) *DIContainer {
	dic := &DIContainer{}
	dic.cli = newCLIProvider(use, description)
	return dic
}

// CLI provides new cobra root command
func (dic *DIContainer) CLI() *cobra.Command {
	return dic.cli()
}
