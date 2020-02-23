package tests

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/codemodify/systemkit-etl/extlod"
	"github.com/codemodify/systemkit-etl/pipeline"
)

func Test_Sequential01(t *testing.T) {
	rawData, err := ioutil.ReadFile("./data-files/sample.csv")
	if err != nil {
		t.Fatal(err)
	}

	p := pipeline.NewSequentialPipeline([]pipeline.Unit{
		pipeline.Unit{
			Extractor: extlod.NewCSVExtractor(false),
			Loader:    extlod.NewCSVLoader(false),
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
	ioutil.WriteFile("./data-files/sample-ouput.csv", outputRawData, 0644)
}
