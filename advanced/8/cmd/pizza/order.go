package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var orderCmd = &cobra.Command{
	Use:       "order pizza ...",
	Short:     "Order pizza of your choice",
	Run:       runOrderCmd,
	ValidArgs: pizzaStrings(),
	Args:      cobra.OnlyValidArgs,
}

func runOrderCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Ordering pizzas:", args)
}
