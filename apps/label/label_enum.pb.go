// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package label

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseOPERATORFromString Parse OPERATOR from string
func ParseOPERATORFromString(str string) (OPERATOR, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := OPERATOR_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown OPERATOR: %s", str)
	}

	return OPERATOR(v), nil
}

// Equal type compare
func (t OPERATOR) Equal(target OPERATOR) bool {
	return t == target
}

// IsIn todo
func (t OPERATOR) IsIn(targets ...OPERATOR) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t OPERATOR) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *OPERATOR) UnmarshalJSON(b []byte) error {
	ins, err := ParseOPERATORFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
