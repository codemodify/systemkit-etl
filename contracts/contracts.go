package contracts

// DATA -
type DATA map[string]interface{}

// // DATAContainer -
// type DATAContainer map[uint64]DATA

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
