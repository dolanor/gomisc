package main

import (
	"errors"
)

type Alias struct {
	Registry string
	Key      string
	Value    AliasValue
}

type Kind string

const (
	KindUnknown Kind = ""
	KindString       = "string"
	KindArray        = "array"
)

type AliasValue interface {
	Value() interface{}
}

func NewAliasValue(kind Kind, value interface{}) (AliasValue, error) {
	switch kind {
	case KindString:
		s, ok := value.(string)
		if !ok {
			return nil, errors.New("bad alias value string kind")
		}
		return AliasString{value: s}, nil
	case KindArray:
		s, ok := value.([]string)
		if !ok {
			return nil, errors.New("bad alias value array kind")
		}
		return AliasArray{value: s}, nil
	default:
		return nil, errors.New("unknown alias value kind")
	}
}

type AliasString struct {
	value string
}

func (as AliasString) Value() interface{} {
	return as.value
}

type AliasArray struct {
	value []string
}

func (as AliasArray) Value() interface{} {
	return as.value
}
