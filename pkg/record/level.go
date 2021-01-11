package record

import (
	"fmt"

	"github.com/amarin/gogedcom/pkg/errors"
)

// MaxRecordLevel defines maximum allowed record level set to 99 according to GEDCOM 5.5.5 spec.
const MaxRecordLevel = 99

// Level type defines GEDCOM record level.
type Level uint

// Validate checks internal data is valid, returns error otherwise.
func (level Level) Validate() error {
	if level > MaxRecordLevel {
		return fmt.Errorf("%w: level: must %v <= 99", errors.ErrValidation, level)
	}

	return nil
}
