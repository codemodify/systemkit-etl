package contracts

// DATAPayload -
type DATAPayload map[string]interface{}

// DATA -
type DATA []DATAPayload

// Extractor -
type Extractor interface {
	Extract(rawData []byte) (DATA, error)
}

// Transformer -
type Transformer interface {
	Transform(data DATA) (DATA, error)
}

// Loader -
type Loader interface {
	Load(data DATA) ([]byte, error)
}
