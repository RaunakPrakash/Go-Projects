package cli

import (
	"github.com/spf13/cobra"
	"sync"
)

func newCLIProvider(use string, description string) func() *cobra.Command {
	var mu sync.Mutex
	var c *cobra.Command
	return func() *cobra.Command {
		mu.Lock()
		defer mu.Unlock()
		if c == nil {
			c = newCLI(use, description)
		}
		return c
	}

}

func newCLI(use string, description string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   use,
		Short: description,
	}

	return rootCmd
}
