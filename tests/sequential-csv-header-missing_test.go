package tests

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/codemodify/systemkit-etl/extractors"
	"github.com/codemodify/systemkit-etl/loaders"
	"github.com/codemodify/systemkit-etl/pipeline"
)

func Test_Sequential_CSV_Header(t *testing.T) {
	rawData, err := ioutil.ReadFile("./data-files/sample-header-mising.csv")
	if err != nil {
		t.Fatal(err)
	}

	p := pipeline.NewSequentialPipeline([]pipeline.Unit{
		pipeline.Unit{
			Extractor: extractors.NewCSVExtractor(false),
			Loader:    loaders.NewCSVLoader(false),
		},
	})

	outputRawData, errs := p.Execute(rawData)
	if errs != nil {
		sb := strings.Builder{}
		for _, err := range errs {
			sb.WriteString(err.Error())
		}

		t.Fatal(sb.String())
	}

	fmt.Println(string(outputRawData))
	ioutil.WriteFile("./data-files/sample-header-mising-ouput.csv", outputRawData, 0644)
}
