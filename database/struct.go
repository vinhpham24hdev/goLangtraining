package database

// Envelope struct
type Envelope struct {
	XMLName xml.Name `xml:"Envelope" json:"-"`
	Cube    CubeMain `xml:"Cube" json:"cube"`
	Xmlns   string   `xml:"xmlns,attr"`
}

// CubeMain struct
type CubeMain struct {
	XMLName xml.Name `xml:"Cube" json:"-"`
	Cube    []Cube   `xml:"Cube" json:"cube"`
}

// Cube struct
type Cube struct {
	XMLName xml.Name   `xml:"Cube" json:"-"`
	Cube    []CubeItem `xml:"Cube" json:"cube"`
	Time    string     `xml:"time,attr" json:"time"`
}

// CubeItem struct
type CubeItem struct {
	XMLName  xml.Name `xml:"Cube" json:"-"`
	Currency string   `xml:"currency,attr" json:"currency"`
	Rate     string   `xml:"rate,attr" json:"rate"`
}
