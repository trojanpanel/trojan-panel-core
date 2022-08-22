package bo

import "errors"

// TypeMessage 自定义类型并自定义序列化方式
type TypeMessage []byte

func (m TypeMessage) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return []byte("null"), nil
	}
	return m, nil
}

func (m *TypeMessage) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
