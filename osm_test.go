package goosm_test

import (
	"testing"
	"bytes"
	"runtime"
	"fmt"
	"path/filepath"
	"reflect"
	"github.com/gaffo/goosm"
)

func fail(tb testing.TB, reason string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%d: %s\n\n", filepath.Base(file), line, reason)
	tb.FailNow()
}

func contains(tb testing.TB, s []string, e string) {
	for _, a := range s {
		if a == e {
			return
		}
	}
	fmt.Printf("Expected to contain %s\n", e)
	tb.FailNow()
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func TestParseChangeEmpty(t *testing.T) {
	osm := goosm.ParseChange("blilbirp")
	if osm != nil {
		fail(t, "not nil")
	}
}

func TestParseChange(t *testing.T) {
	osm := goosm.ParseChange("test/osm_change.xml")
	if osm == nil {
		fail(t, "nil")
	}

	equals(t, "0.6", osm.Version)
	equals(t, "JOSM", osm.Generator)

	equals(t, 1, len(osm.Creates))
}

func TestParseOregon(t *testing.T) {
	osm := goosm.ParseOsm("test/or.osm")
	if osm == nil {
		t.FailNow()
	}

	equals(t, "0.6", osm.Version)
	equals(t, true, osm.Upload)
	equals(t, "JOSM", osm.Generator)

	equals(t, 198, len(osm.Nodes))
	node := osm.Nodes[0]
	equals(t, "-19857", node.Id)
	equals(t, true, node.Visible)
	equals(t, 45.99517536936, node.Lat)
	equals(t, -116.91913230528, node.Lon)

	equals(t, 1, len(osm.Ways))
	way := osm.Ways[0]
	equals(t, "-19859", way.Id)
	equals(t, true, way.Visible)

	equals(t, 199, len(way.Nds))
	nd := way.Nds[0]
	equals(t, "-19555", nd.Ref)

	equals(t, 6, len(way.Tags))
	tag := way.Tags[0]
	equals(t, "DRAWSEQ", tag.Key)
	equals(t, "12", tag.Value)
}

func TestParseWashington(t *testing.T) {
	osm := goosm.ParseOsm("test/wa.osm")
	if osm == nil {
		t.FailNow()
	}

	equals(t, "0.6", osm.Version)
	equals(t, true, osm.Upload)
	equals(t, "JOSM", osm.Generator)

	equals(t, 273, len(osm.Nodes))
	node := osm.Nodes[0]
	equals(t, "-19221", node.Id)
	equals(t, true, node.Visible)
	equals(t, 48.22521637144, node.Lat)
	equals(t, -122.40201531038, node.Lon)

	equals(t, 3, len(osm.Ways))
	way := osm.Ways[0]
	equals(t, "-19229", way.Id)
	equals(t, true, way.Visible)

	equals(t, 244, len(way.Nds))
	nd := way.Nds[0]
	equals(t, "-19221", nd.Ref)

	equals(t, 1, len(osm.Relations))
	relation := osm.Relations[0]
	equals(t, "-19265", relation.Id)
	equals(t, true, relation.Visible)

	equals(t, 3, len(relation.Members))
	member := relation.Members[0]
	equals(t, "way", member.Type)
	equals(t, "-19229", member.Ref)
	equals(t, "outer", member.Role)

	equals(t, 7, len(relation.Tags))
	tag := relation.Tags[0]
	equals(t, "DRAWSEQ", tag.Key)
	equals(t, "2", tag.Value)
}

func TestParseNonexistent(t *testing.T) {
	osm := goosm.ParseOsm("blibbity")
	if osm != nil {
		t.FailNow()
	}
}

func TestEncode(t *testing.T) {
	var b bytes.Buffer

	osm := goosm.NewOsm()
	osm.Write(&b)
	equals(t, `<osm version="0.6" upload="true" generator="github.com/gaffo/goosm"></osm>`, b.String())
}

func TestEncodeWithNode(t *testing.T) {
	var b bytes.Buffer
	osm := goosm.NewOsm()
	osm.Nodes = append(osm.Nodes, goosm.Node{Id: "-1", Visible: true, Lat: 30.1, Lon: -122.33})
	osm.Write(&b)
	fmt.Println(b.String())
	equals(t, `<osm version="0.6" upload="true" generator="github.com/gaffo/goosm"><node id="-1" visible="true" lat="30.1" lon="-122.33"></node></osm>`, b.String())
}