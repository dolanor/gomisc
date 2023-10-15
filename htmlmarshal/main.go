package main

import (
	"encoding/xml"
	"fmt"
)

type Attrs struct {
	ID      string `xml:"id,attr,omitempty"`
	Name    string `xml:"name,attr,omitempty"`
	Class   Class  `xml:"class,attr,omitempty"`
	Style   string `xml:"style,attr,omitempty"`
	OnClick string `xml:"onclick,attr,omitempty"`
	//Name    string `xml:"name,attr,omitempty"`
	//Name    string `xml:"name,attr,omitempty"`
}

//TODO create some CSS system
type Class string

type HTML struct {
	XMLName xml.Name `xml:"html"`
	Header  Header   `xml:"header"`
}

type Header struct {
	Meta Meta `xml:"meta"`
	Body Body `xml:"body"`
}

type Meta struct {
	Title string `xml:"title"`
}

type Body struct {
	El
}

type El struct {
	Attrs
	Div *Div `xml:"div,omitempty"`
}

type Div struct {
	Attrs
	Text string `xml:"text,omitempty"`
}

func main() {
	h := HTML{
		Header: Header{
			Meta: Meta{
				Title: "Hello",
			},
			Body: Body{
				El: El{
					//Div: &Div{
					//	Text: "LOLO",
					//},
					Attrs: Attrs{
						Class: "button-big",
					},
				},
			},
		},
	}
	b, err := xml.MarshalIndent(h, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
