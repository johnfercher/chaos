package main

import (
	"fmt"
	"github.com/johnfercher/chaos/internal/services"
	"github.com/johnfercher/chaos/internal/template/chaos"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "deco",
	Short: "Deco is a decorator generator for go code",
	Run:   Command,
}

func Command(cmd *cobra.Command, args []string) {
	_type, _ := cmd.Flags().GetString("type")

	input, err := cmd.Flags().GetString("input")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if input == "" {
		fmt.Fprintln(os.Stderr, "error: input is empty")
		os.Exit(1)
	}

	_interface, err := cmd.Flags().GetString("interface")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Generate %s decorator for %s in %s\n", _type, _interface, input)

	file := services.NewFile()
	interpreter := services.NewInterfaceInterpreter()
	decorator := services.NewDecoratorGenerator("Chaos", chaos.Decorator, chaos.Method)
	orchestrator := services.NewOrchestrator(file, interpreter, decorator)

	err = orchestrator.Generate(input, _interface)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.PersistentFlags().String("type", "chaos", "The decorator type generation")
	rootCmd.PersistentFlags().String("input", "", "Input file")
	rootCmd.PersistentFlags().String("interface", "", "Interface to generate decorator")
	//rootCmd.SetArgs([]string{"--type=chaos", "--input=docs/examples/interfaces.go", "--interface=SingleParameterWithTwoReturns"})
	rootCmd.SetArgs([]string{"--type=chaos", "--input=docs/examples/interfaces.go"})

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
