package envs

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

const (
	maxRowLength int    = 30
	envsFlag     string = "envs"
)

func wordWrap(str string) (newStr string) {
	for i, v := range str {
		if i > 0 && i%30 == 0 {
			newStr += "\n"
		}
		newStr += string(v)
	}

	return
}

// GetAllFlags defines corresponding environment variables for each flag and
// returns a table representing all the application flags and environment variables.
func GetAllFlags() error {
	// Check if flags are parsed.
	if !flag.Parsed() {
		return errors.New("flags are not parsed")
	}

	// Define the table.
	var tableBuff *bytes.Buffer
	var table *tablewriter.Table

	printHelp := flag.Lookup(envsFlag).Value.String() == "true"

	// Assign table variables if table representation is required.
	if printHelp {
		tableBuff = new(bytes.Buffer)
		table = tablewriter.NewWriter(tableBuff)
	}

	// Loop through the flags and fill up the table.
	var data [][]string
	flag.VisitAll(func(currentFlag *flag.Flag) {
		// Skip reserved "envs" flag.
		if currentFlag.Name == envsFlag {
			return
		}

		// Define the corresponding environment variable.
		envVar := strings.ToUpper(strings.Replace(currentFlag.Name, "-", "_", -1))

		// Overwrite flag value if the environment variable is set.
		envVarValue := os.Getenv(envVar)
		if envVarValue != "" {
			flag.Set(currentFlag.Name, envVarValue)
		}

		// Append to data if table representation is required.
		if printHelp {
			data = append(data, []string{
				currentFlag.Name,
				envVar,
				wordWrap(currentFlag.DefValue),
				wordWrap(currentFlag.Value.String()),
				currentFlag.Usage,
			})
		}
	})

	// Skip further execution if table representation is not required.
	if !printHelp {
		return nil
	}

	// Check whether the default value column is empty.
	defaultValIndex := 2

	defaultEmpty := true
	for _, row := range data {
		if row[defaultValIndex] != "" {
			defaultEmpty = false
			break
		}
	}

	// Delete column if empty for all rows.
	if defaultEmpty {
		for i := 0; i < len(data); i++ {
			data[i] = append(data[i][:defaultValIndex], data[i][defaultValIndex+1:]...)
		}
	}

	// Assemble header slice.
	header := []string{
		"Flag",
		"Environment Var",
		"Default Value",
		"Current Value",
		"Description",
	}

	if defaultEmpty {
		header = append(header[:defaultValIndex], header[defaultValIndex+1:]...)
	}

	table.SetHeader(header)

	// Append bulk data to the table.
	table.AppendBulk(data)

	// Render the table.
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetBorder(false)

	table.Render()

	// Print table data.
	fmt.Println(tableBuff.String())

	return nil
}
