package extractors

import (
	"encoding/csv"
	"strings"

	etlContracts "github.com/codemodify/systemkit-etl/contracts"
)

// CSVExtractorLineHandler -
type CSVExtractorLineHandler func(line []string, lineIndex int, isFirstLineHeader bool) (bool, etlContracts.DATAPayload)

// csvExtractor -
type csvExtractor struct {
	isFirstLineHeader bool
	lineHandler       CSVExtractorLineHandler
}

// NewCSVExtractor -
func NewCSVExtractor(isFirstLineHeader bool) etlContracts.Extractor {
	return NewCSVExtractorWithDelegate(isFirstLineHeader, func(line []string, lineIndex int, isFirstLineHeader bool) (bool, etlContracts.DATAPayload) {
		if len(line) > 0 {
			return true, etlContracts.DATAPayload{
				"": line,
			}
		}
		return false, etlContracts.DATAPayload{}
	})
}

// NewCSVExtractorWithDelegate -
func NewCSVExtractorWithDelegate(isFirstLineHeader bool, lineHandler CSVExtractorLineHandler) etlContracts.Extractor {
	return &csvExtractor{
		isFirstLineHeader: isFirstLineHeader,
		lineHandler:       lineHandler,
	}
}

// Extract - `github.com/codemodify/systemkit-etl/contracts/Extractor` interface
func (thisRef csvExtractor) Extract(data []byte) (etlContracts.DATA, error) {
	lines, err := csv.NewReader(strings.NewReader(string(data))).ReadAll()
	if err != nil {
		return nil, err
	}

	if thisRef.lineHandler != nil {
		result := etlContracts.DATA{}

		for lineIndex, line := range lines {
			isValid, obj := thisRef.lineHandler(line, lineIndex, thisRef.isFirstLineHeader)
			if isValid {
				result = append(result, obj)
			}
		}

		return result, nil
	}

	return etlContracts.DATA{}, nil
}
