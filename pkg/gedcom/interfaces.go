package gedcom

// Validator interface requires implementations provide its own internal data validation method.
type Validator interface {
	// Validate checks internal data is valid, returns error otherwise.
	Validate() error
}
