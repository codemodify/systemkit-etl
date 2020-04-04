# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Extract Transform Load
[![](https://img.shields.io/github/v/release/codemodify/systemkit-etl?style=flat-square)](https://github.com/codemodify/systemkit-etl/releases/latest)
![](https://img.shields.io/github/languages/code-size/codemodify/systemkit-etl?style=flat-square)
![](https://img.shields.io/github/last-commit/codemodify/systemkit-etl?style=flat-square)
[![](https://img.shields.io/badge/license-0--license-brightgreen?style=flat-square)](https://github.com/codemodify/TheFreeLicense)

![](https://img.shields.io/github/workflow/status/codemodify/systemkit-etl/qa?style=flat-square)
![](https://img.shields.io/github/issues/codemodify/systemkit-etl?style=flat-square)
[![](https://goreportcard.com/badge/github.com/codemodify/systemkit-etl?style=flat-square)](https://goreportcard.com/report/github.com/codemodify/systemkit-etl)

[![](https://img.shields.io/badge/godoc-reference-brightgreen?style=flat-square)](https://godoc.org/github.com/codemodify/systemkit-etl)
![](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)
![](https://img.shields.io/gitter/room/codemodify/systemkit-etl?style=flat-square)

![](https://img.shields.io/github/contributors/codemodify/systemkit-etl?style=flat-square)
![](https://img.shields.io/github/stars/codemodify/systemkit-etl?style=flat-square)
![](https://img.shields.io/github/watchers/codemodify/systemkit-etl?style=flat-square)
![](https://img.shields.io/github/forks/codemodify/systemkit-etl?style=flat-square)


>### ETL is set of convertors from one format to another
 - #### The one and only high performance and complete ETL framework in Go
 - #### Supported: Linux, Raspberry Pi, FreeBSD, Mac OS, Windows, Solaris

>### Write your ETL plug-ings, sample are under
- #### extlod
- #### transformers

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Install
```go
go get github.com/codemodify/systemkit-etl
```
# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) API

&nbsp;																| &nbsp;
---     															| ---
NewCSVExtractor() |
NewCSVExtractorWithDelegate() |
Extract(`data` []byte) |
NewCSVLoader() |
NewCSVLoaderWithDelegate() |
Load(`data` etlContracts.DATA) |
NewSequentialPipeline(`units` []Unit) |
Execute(`rawData` []byte) |

# ![](https://fonts.gstatic.com/s/i/materialicons/bookmarks/v4/24px.svg) Usage
```go
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/codemodify/systemkit-etl/extractors"
	"github.com/codemodify/systemkit-etl/loaders"
	"github.com/codemodify/systemkit-etl/pipeline"
)

func main(t *testing.T) {
	rawData, err := ioutil.ReadFile("./")
	if err != nil {
		t.Fatal(err)
	}

	p := pipeline.NewSequentialPipeline([]pipeline.Unit{
		pipeline.Unit{
			Extractor: extractors.NewCSVExtractor(true),
			Loader:    loaders.NewCSVLoader(true),
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
	ioutil.WriteFile("./data-files/sample-header-ouput.csv", outputRawData, 0644)
}
```



> #### For examples see `tests`
