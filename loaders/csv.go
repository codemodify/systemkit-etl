package loaders

import (
	"encoding/csv"
	"strings"

	etlContracts "github.com/codemodify/systemkit-etl/contracts"
)

// CSVLoaderLineHandler -
type CSVLoaderLineHandler func(etlContracts.DATAPayload) (bool, []string)

// csvLoader -
type csvLoader struct {
	isFirstLineHeader bool
	lineHandler       CSVLoaderLineHandler
}

// NewCSVLoader -
func NewCSVLoader(isFirstLineHeader bool) etlContracts.Loader {
	return NewCSVLoaderWithDelegate(isFirstLineHeader, func(line etlContracts.DATAPayload) (bool, []string) {
		if v, ok := line[""].([]string); ok {
			return true, v
		}

		return false, []string{}
	})
}

// NewCSVLoaderWithDelegate -
func NewCSVLoaderWithDelegate(isFirstLineHeader bool, lineHandler CSVLoaderLineHandler) etlContracts.Loader {
	return &csvLoader{
		isFirstLineHeader: isFirstLineHeader,
		lineHandler:       lineHandler,
	}
}

// Load - `github.com/codemodify/systemkit-etl/contracts/Loader` interface
func (thisRef csvLoader) Load(data etlContracts.DATA) ([]byte, error) {
	if thisRef.lineHandler != nil {
		result := [][]string{}

		for _, value := range data {
			isValid, line := thisRef.lineHandler(value)
			if isValid {
				result = append(result, line)
			}
		}

		sb := strings.Builder{}
		csv.NewWriter(&sb).WriteAll(result)

		return []byte(sb.String()), nil
	}

	return []byte{}, nil
}
