package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type pricer interface {
	Cost() int
}

func newPriceCmd() *cobra.Command {
	c := &cobra.Command{
		Use:       "price pizza",
		Short:     "Check the price of a pizza",
		Run:       runPriceCmd,
		ValidArgs: pizzaStrings(),
		Args:      cobra.OnlyValidArgs,
	}

	c.Flags().StringP("currency", "c", "EUR", "set the currency of your purchase")

	return c
}

func runPriceCmd(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("price requires at least 1 argument")
		os.Exit(1)
	}

	var price int
	for _, a := range args {
		p, err := pizzaString(a)
		cobra.CheckErr(err)

		price += p.Cost()
	}

	cur, err := cmd.Flags().GetString("currency")
	cobra.CheckErr(err)

	fmt.Printf("The price would be: %s %d\n", cur, price)
}
