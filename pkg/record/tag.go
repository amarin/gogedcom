package record

type Tag string

const (
	// Concatenation defines constant tag name for continue current logical line with current physical line value.
	Concatenation Tag = "CONC"
	// Continuation defines constant tag name for continue current logical line with new line and current physical line.
	Continuation Tag = "CONT"
	// Header tag indicate a header record to provide basic information about a GEDCOM file.
	Header Tag = "HEAD"
	// Trailer tag indicates a trailer record to signal the end of a GEDCOM file.
	Trailer Tag = "TRLR"
)
