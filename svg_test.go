package svg

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var testdata = []string{
	`<svg>
		<circle cx="50" cy="40" r="30" stroke="#000000" stroke-width="3" fill="#ff0000" />
	</svg>`,
	`<svg>
		<circle cx="50" cy="40" r="30" stroke="#000000"></circle>
	</svg>`,
}

func TestUnmarshal(t *testing.T) {
	for _, td := range testdata {
		var svg SVG
		if err := xml.Unmarshal([]byte(td), &svg); err != nil {
			t.Errorf("unexpected err: %s", err)
		}
		fmt.Printf("svg: %+v\n", svg)
	}
}
