package record

// Terminator defines line terminator type.
type Terminator string

const (
	// CR provides predefined Carriage Return character line terminator.
	CR Terminator = "\r"
	// LF provides predefined Line Feed character line terminator.
	LF Terminator = "\n"
	// CRLF provides predefined combined CR/LF characters line terminator.
	CRLF Terminator = "\r\n"
)
