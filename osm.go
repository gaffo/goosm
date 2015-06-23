package hunt

import (
	"encoding/xml"
	"io/ioutil"
)

// Decimal 	0 	1 	2 	3 	4 	5 	6 	7 	8 	9 	10 	11 	12 	13 	14 	15
// Base 32 	0 	1 	2 	3 	4 	5 	6 	7 	8 	9 	b 	c 	d 	e 	f 	g
//
// Decimal 	16 	17 	18 	19 	20 	21 	22 	23 	24 	25 	26 	27 	28 	29 	30 	31
// Base 32 	h 	j 	k 	m 	n 	p 	q 	r 	s 	t 	u 	v 	w 	x 	y 	z

var geoHashDict = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"b": 10,
	"c": 11,
	"d": 12,
	"e": 13,
	"f": 14,
	"g": 15,
	"h": 16,
	"j": 17,
	"k": 18,
	"m": 19,
	"n": 20,
	"p": 21,
	"q": 22,
	"r": 23,
	"s": 24,
	"t": 25,
	"u": 26,
	"v": 27,
	"w": 28,
	"x": 29,
	"y": 30,
	"z": 31,
}

type LL struct {
	Lat float64
	Lon float64
}

func (ll LL) GeoHash() string {
	return ""
}

type GeoHash string

func (geohash GeoHash) LL() LL {
	return LL{}
}

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
