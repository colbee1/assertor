package assertor

import (
	"fmt"
	"strings"
)

// Assertor exposes Separator used between message.
type Assertor struct {
	Separator string
	msgs      []string
}

// New returns a new assertors with default Separator (;).
func New() *Assertor {
	return &Assertor{
		msgs:      []string{},
		Separator: "; ",
	}
}

// Assert returns <ok>
func (a *Assertor) Assert(ok bool, format string, args ...any) bool {
	if !ok {
		a.msgs = append(a.msgs, fmt.Errorf(format, args...).Error())
	}

	return ok
}

// Validate returns an error if at least on Assert() has failed.
// Error message contains the list of unsatisfied requirements.
func (a *Assertor) Validate() error {
	if len(a.msgs) == 0 {
		return nil // No errors
	}

	return fmt.Errorf("%w: %d unsatisfied requirement(s): %s",
		ErrValidate,
		len(a.msgs),
		strings.Join(a.msgs, a.Separator))

}
