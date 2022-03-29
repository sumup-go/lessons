package main

import (
	"log"

	"github.com/spf13/cobra"
)

// CLI
// DONE order pizza
// TODO check the prices of different pizzas

var cmd = &cobra.Command{
	Use: "pizza",
}

func main() {
	cmd.AddCommand(orderCmd)
	cmd.AddCommand(newPriceCmd())

	if err := cmd.Execute(); err != nil {
		log.Fatalf("executing cmd: %v", err)
	}
}

//go:generate enumer -trimprefix pizza -json -type pizza $GOFILE
type pizza int

const (
	pizzaVege pizza = iota
	pizzaPepperoni
	pizzaChicken
	pizzaHawaiian
)

func (p pizza) Cost() int {
	switch p {
	case pizzaVege:
		return 5
	case pizzaPepperoni:
		return 8
	case pizzaChicken:
		return 7
	case pizzaHawaiian:
		return 9
	}

	return 0
}
