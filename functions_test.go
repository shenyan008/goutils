package goutils

import "testing"

func TestNamedFormatString(t *testing.T) {
	format := "key1: %{key1}s, key2: %{key2}s"
	m := map[string]interface{}{
		"key1": "val1",
		"key2": "val2",
	}
	expected := "key1: val1, key2: val2"
	actural := NamedFormatString(format, m)
	//t.Log(actural)
	if expected != actural {
		t.Error("Test Failed")
	}
}
