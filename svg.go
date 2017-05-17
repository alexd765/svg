package svg

import "encoding/xml"

// SVG contains a slice of shape elements in order.
type SVG struct {
	Elements []Element
}

// Element is a shape element of a svg. Circle is the only one supported yet.
type Element interface {
	private()
}

// UnmarshalXML fullfills the xml unmarshaler interface
func (s *SVG) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			switch tok.Name.Local {
			case "circle":
				var c Circle
				if err := d.DecodeElement(&c, &tok); err != nil {
					return err
				}
				s.Elements = append(s.Elements, c)
			}

		case xml.EndElement:
			if tok.Name.Local == start.Name.Local {
				return nil
			}
		}
	}
}
