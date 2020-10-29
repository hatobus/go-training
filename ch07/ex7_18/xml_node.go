package ex7_18

import (
	"encoding/xml"
	"fmt"
)

type Node interface {
	String() string
}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (c CharData) String() string {
	return string(c)
}

func (e *Element) String() string {
	var attr string
	for _, a := range e.Attr {
		attr += fmt.Sprintf(" %v=\"%v\"", a.Name.Local, a.Value)
	}

	var children string
	for _, child := range e.Children {
		children += child.String()
	}

	return fmt.Sprintf("<%v%v>%v</%v>", e.Type.Local, attr, children, e.Type.Local)
}

func Parse(d *xml.Decoder) (Node, error) {
	var stk []*Element

	for {
		tok, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			element := Element{
				Type:     tok.Name,
				Attr:     tok.Attr,
				Children: []Node{},
			}
			if len(stk) > 0 {
				stk[len(stk)-1].Children = append(stk[len(stk)-1].Children, &element)
			}
			stk = append(stk, &element)
		case xml.EndElement:
			if len(stk) == 0 {
				return nil, fmt.Errorf("unexpected tag closing")
			} else if len(stk) == 1 {
				// find out end tag
				return stk[0], nil
			}
			stk = stk[:len(stk)-1]
		case xml.CharData:
			if len(stk) > 0 {
				// ignore has not parent attributes
				stk[len(stk)-1].Children = append(stk[len(stk)-1].Children, CharData(tok))
			}
		}
	}
}
