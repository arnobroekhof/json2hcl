package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
	"bytes"
	"io/ioutil"
	"os"
)

var (
	inputFile  string
	outputFile string
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert files",
	Run: func(cmd *cobra.Command, args []string) {
		convert(inputFile, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.PersistentFlags().StringVarP(&inputFile, "input", "i", "", "json formatted input file")
	convertCmd.MarkPersistentFlagRequired("input")
	convertCmd.PersistentFlags().StringVarP(&inputFile, "output", "o", "", "output file if none provided the input file will be overwritten")
}

func readJSONFile(inputFile string) (string, error) {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func json2hcl(inputJSON string) ([]byte, error) {

	ast, err := jsonParser.Parse([]byte(inputJSON))
	if err != nil {
		return nil, fmt.Errorf("unable to parse JSON: %s", err)
	}

	buffer := new(bytes.Buffer)

	err = printer.Fprint(buffer, ast)
	if err != nil {
		return nil, fmt.Errorf("unable to construct HCL file from JSON: %s", err)
	}

	return buffer.Bytes(), nil
}

func convert(inputFile string, outputFile string) error {
	jsonFile, readErr := readJSONFile(inputFile)
	if readErr != nil {
		return readErr
	}

	hclParsed, parseErr := json2hcl(jsonFile)
	if parseErr != nil {
		return parseErr
	}

	if outputFile == "" {
		outputFile = inputFile
	}

	writeErr := ioutil.WriteFile(outputFile, hclParsed, os.FileMode(0644))
	if writeErr != nil {
		return writeErr
	}

	return nil

}
