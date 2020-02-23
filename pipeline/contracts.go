package pipeline

import (
	etlContracts "github.com/codemodify/systemkit-etl/contracts"
)

// Pipeline -
type Pipeline interface {
	Execute(rawData []byte) ([]byte, []error)
}

// Unit -
type Unit struct {
	Extractor    etlContracts.Extractor
	Transformers []etlContracts.Transformer
	Loader       etlContracts.Loader
}
