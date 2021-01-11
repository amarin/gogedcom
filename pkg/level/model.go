package level

import (
	"fmt"

	"github.com/amarin/gogedcom/pkg/errors"
)

type RecordLevel uint

// Validate checks internal data is valid, returns error otherwise.
func (level RecordLevel) Validate() error {
	if level > 99 {
		return fmt.Errorf("%w: record level: must %v <= 99", errors.ErrValidation, level)
	}

	return nil
}
