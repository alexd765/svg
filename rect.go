package svg

// Rect is a possible shape element in a svg.
type Rect struct {
	X           float64 `xml:"x,attr"`
	Y           float64 `xml:"y,attr"`
	Width       float64 `xml:"width,attr"`
	Height      float64 `xml:"height,attr"`
	Stroke      string  `xml:"stroke,attr"`
	StrokeWidth string  `xml:"stroke-width,attr"`
	Fill        string  `xml:"fill,attr"`
}

func (r Rect) private() {}
