// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package namespace

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseVISIBLEFromString Parse VISIBLE from string
func ParseVISIBLEFromString(str string) (VISIBLE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := VISIBLE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown VISIBLE: %s", str)
	}

	return VISIBLE(v), nil
}

// Equal type compare
func (t VISIBLE) Equal(target VISIBLE) bool {
	return t == target
}

// IsIn todo
func (t VISIBLE) IsIn(targets ...VISIBLE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t VISIBLE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *VISIBLE) UnmarshalJSON(b []byte) error {
	ins, err := ParseVISIBLEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
