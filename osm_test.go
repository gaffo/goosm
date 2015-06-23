package hunt

import (
	"testing"
)

// func TestGeohashFromLL(t *testing.T) {
// 	equals(t, "u4pruydqqvj", LL{57.64911, 10.40744}.GeoHash())
// }

// func TestLLFromGeohash(t *testing.T) {
// 	equals(t, LL{57.64911, 10.40744}, GeoHash("u4pruydqqvj").LL())
// }

func TestParseChangeEmpty(t *testing.T) {
	osm := ParseChange("blilbirp")
	if osm != nil {
		fail(t, "not nil")
	}
}

func TestParseChange(t *testing.T) {
	osm := ParseChange("test/osm_change.xml")
	if osm == nil {
		fail(t, "nil")
	}

	equals(t, "0.6", osm.Version)
	equals(t, "JOSM", osm.Generator)

	equals(t, 1, len(osm.Creates))
}

func TestParseOregon(t *testing.T) {
	osm := ParseOsm("test/or.osm")
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
	osm := ParseOsm("test/wa.osm")
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
	osm := ParseOsm("blibbity")
	if osm != nil {
		t.FailNow()
	}
}
