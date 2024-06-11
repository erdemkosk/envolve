package command

import "github.com/spf13/cobra"

type ICommand interface {
	Execute(cmd *cobra.Command, args []string)
}
