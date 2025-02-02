package config

type NullableValue struct {
	Value uint16
}

func NewNullable(val uint16) *NullableValue {
	return &NullableValue{Value: val}
}
