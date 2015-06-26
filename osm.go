package goosm

import (
	"encoding/xml"
	"io/ioutil"
	"io"
)

type OsmChange struct {
	Creates   []Create `xml:"create"`
	Version   string   `xml:"version,attr"`
	Generator string   `xml:"generator,attr"`
}

type Create struct {
	Nodes []Node `xml:"node"`
	Ways  []Way  `xml:"way"`
}

type Osm struct {
	Nodes     []Node     `xml:"node"`
	Ways      []Way      `xml:"way"`
	Relations []Relation `xml:"relation"`
	Version   string     `xml:"version,attr"`
	Upload    bool       `xml:"upload,attr"`
	Generator string     `xml:"generator,attr"`
	XMLName struct{}	`xml:"osm"`
}

func NewOsm() *Osm {
	return &Osm{
		Version: "0.6",
		Generator: "github.com/gaffo/goosm",
		Upload: true,
		Nodes: make([]Node, 0, 128),
		Ways: make([]Way, 0, 128),
		Relations: make([]Relation, 0, 128),
	}
}

func (osm *Osm) Write(writer io.Writer) {
	e := xml.NewEncoder(writer)
	e.Encode(osm)
	e.Flush()
}

type Node struct {
	Id      string  `xml:"id,attr"`
	Visible bool    `xml:"visible,attr"`
	Lat     float64 `xml:"lat,attr"`
	Lon     float64 `xml:"lon,attr"`
}

type Way struct {
	Tags    []Tag  `xml:"tag"`
	Nds     []Nd   `xml:"nd"`
	Id      string `xml:"id,attr"`
	Visible bool   `xml:"visible,attr"`
}

type Nd struct {
	Ref string `xml:"ref,attr"`
}

type Tag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

type Relation struct {
	Tags    []Tag    `xml:"tag"`
	Members []Member `xml:"member"`
	Id      string   `xml:"id,attr"`
	Visible bool     `xml:"visible,attr"`
}

type Member struct {
	Type string `xml:"type,attr"`
	Ref  string `xml:"ref,attr"`
	Role string `xml:"role,attr"`
}

func ParseOsm(path string) *Osm {
	var osm Osm
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	if err = xml.Unmarshal(data, &osm); err != nil {
		return nil
	}
	return &osm
}

func ParseChange(path string) *OsmChange {
	var osm OsmChange
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	if err = xml.Unmarshal(data, &osm); err != nil {
		return nil
	}
	return &osm
}