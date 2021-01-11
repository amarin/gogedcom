package record

import (
	"fmt"

	"github.com/amarin/gogedcom/pkg/errors"
)

type Level uint

// Validate checks internal data is valid, returns error otherwise.
func (level Level) Validate() error {
	if level > 99 {
		return fmt.Errorf("%w: level: must %v <= 99", errors.ErrValidation, level)
	}

	return nil
}
