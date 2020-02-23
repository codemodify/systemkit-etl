package pipeline

type sequentialPipeline struct {
	units []Unit
}

// NewSequentialPipeline -
func NewSequentialPipeline(units []Unit) Pipeline {
	return &sequentialPipeline{
		units: units,
	}
}

func (thisRef sequentialPipeline) Execute(rawData []byte) ([]byte, []error) {
	for _, unit := range thisRef.units {
		// E
		data, err := unit.Extractor.Extract(rawData)
		if err != nil {
			return nil, []error{err}
		}

		// T
		for _, t := range unit.Transformers {
			data, err = t.Transform(data)
			if err != nil {
				return nil, []error{err}
			}
		}

		// L
		rawData, err = unit.Loader.Load(data)
		if err != nil {
			return nil, []error{err}
		}
	}

	return rawData, nil
}
