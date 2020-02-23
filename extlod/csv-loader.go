package extlod

import (
	"encoding/csv"
	"strings"

	etlContracts "github.com/codemodify/systemkit-etl/contracts"
)

// CSVLoaderLineHandler -
type CSVLoaderLineHandler func(key string, value interface{}) (bool, []string)

// csvLoader -
type csvLoader struct {
	isFirstLineHeader bool
	lineHandler       CSVLoaderLineHandler
}

// NewCSVLoader -
func NewCSVLoader(isFirstLineHeader bool) etlContracts.Loader {
	return NewCSVLoaderWithDelegate(isFirstLineHeader, func(key string, value interface{}) (bool, []string) {
		if v, ok := value.([]string); ok {
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

		for key, value := range data {
			isValid, line := thisRef.lineHandler(key, value)
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
