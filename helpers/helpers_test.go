package helpers
import (
	"testing"
)

func TestStashPath(t *testing.T) {
	f := Flags("foo")
	var a []string
	f.Parse(a)
	if (StashPath() != "cubbyhole/lachash") {
		t.Error("defaults are not set properly")
	}
	a = []string{"-path", "foo"}
	f.Parse(a)
	stash_path = "foo"
	if (StashPath() != "foo") {
		t.Error("changes are not picked up properly")
	}
}

func TestUUIDMangling(t *testing.T) {
	if EncodeUUID("18e9f502-8ebb-c1fd-b7e2-ff91785ccebd") != "GOn1Ao67wf234v+ReFzOvQ" {
		t.Error("Unable to encode UUID")
	}
	if DecodeUUID("GOn1Ao67wf234v+ReFzOvQ") != "18e9f502-8ebb-c1fd-b7e2-ff91785ccebd" {
		t.Error("Unable to decode UUID")
	}
}
