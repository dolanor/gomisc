package main

import (
	"bytes"
	"encoding/json"
	"log"
)

type Valuer interface {
	Value() interface{}
}

type Alias struct {
	ID    int    `json:"id"`
	Value Valuer `json:"value"`
}

type AliasValue struct {
	Kind     Kind            `json:"type"`
	RawValue json.RawMessage `json:"value"`
	Value    Valuer          `json:"-"`
}

func (a *AliasValue) UnmarshalJSON(b []byte) error {
	log.Println("1")
	type tmp AliasValue
	var av tmp
	err := json.Unmarshal(b, &av)
	if err != nil {
		return err
	}
	log.Println("2", av, av.Kind)
	switch av.Kind {
	case KindString:
		var s StringValue
		err = json.Unmarshal(av.RawValue, &s)
		*a = AliasValue{
			Kind:  av.Kind,
			Value: &s,
		}
	case KindArray:
		var s ArrayValue
		err = json.Unmarshal(av.RawValue, &s)
		*a = AliasValue{
			Kind:  av.Kind,
			Value: &s,
		}
	}
	return err
}
func (a *AliasValue) MarshalJSON() ([]byte, error) {
	var err error
	switch a.Kind {
	case KindString:
		var m json.RawMessage
		m, err = json.Marshal(a.Value)
		a.RawValue = m
	case KindArray:
		var m json.RawMessage
		m, err = json.Marshal(a.Value)
		a.RawValue = m
	}
	if err != nil {
		return nil, err
	}

	return err
}

type Kind string

const (
	KindString = "string"
	KindArray  = "array"
)

type StringValue string

func (sv StringValue) Value() interface{} {
	return sv
}

type ArrayValue []string

func (av ArrayValue) Value() interface{} {
	return av
}

func main() {
	str := `{"type": "array", "value": ["lol", "mdr"]}`
	log.Println(str)
	var av AliasValue
	err := json.NewDecoder(bytes.NewBufferString(str)).Decode(&av)
	log.Println(err)
	log.Println("av", av)

	err = json.NewEncoder(bytes.NewBufferString(str)).Encode(av)
	log.Println(err)
	log.Println(str)

	av = AliasValue{
		Kind:  KindArray,
		Value: ArrayValue([]string{"haha", "hoho"}),
	}
	err = json.NewEncoder(bytes.NewBufferString(str)).Encode(av)
	log.Println(err)
	log.Println(str)
}
