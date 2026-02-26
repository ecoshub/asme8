package config

type NullableValue struct {
	Value uint16
}

func (nv *NullableValue) Get() uint16 {
	if nv == nil {
		return 0
	}
	return nv.Value
}

func NewNullable(val uint16) *NullableValue {
	return &NullableValue{Value: val}
}
