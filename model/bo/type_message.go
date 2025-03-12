package bo

import (
	"fmt"
)

// TypeMessage custom types and custom serialization methods
type TypeMessage []byte

func (m TypeMessage) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *TypeMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return fmt.Errorf("UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
