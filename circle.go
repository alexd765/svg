package svg

// Circle is a possible shape element in a SVG.
type Circle struct {
	CX          float64 `xml:"cx,attr"`
	CY          float64 `xml:"cy,attr"`
	R           float64 `xml:"r,attr"`
	Stroke      string  `xml:"stroke,attr"`
	StrokeWidth string  `xml:"stroke-width,attr"`
	Fill        string  `xml:"fill,attr"`
}

func (c Circle) private() {}
