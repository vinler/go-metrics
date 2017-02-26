package metrics

import (
	"testing"
)

func TestTagsEcnoding(t *testing.T) {

	tests := []struct {
		Name string
		Tags map[string]string
	}{
		{Name: "a", Tags: map[string]string{}},
		{Name: "field", Tags: map[string]string{}},
		{Name: "a-longer-one", Tags: map[string]string{}},
		{Name: "tagged", Tags: map[string]string{"a": "a", "b": "b", "c": "c"}},
		{Name: "delimiter|escape", Tags: map[string]string{"ok": "true"}},
		{Name: "delimiter|escape", Tags: map[string]string{}},
		{Name: "xxx|a.large.and.complex.one.with.escaping|xxx", Tags: map[string]string{"a||": "123", "b%%": "[|]"}},
	}

	for _, test := range tests {
		nameEncoded := EncodeNameWithTags(test.Name, test.Tags)
		nameDecoded, tagsDecoded := DecodeNameWithTags(nameEncoded)

		if test.Name != nameDecoded {
			t.Fatalf("encoded and decoded names do not match (%s/%s)", test.Name, nameDecoded)
		}

		if len(test.Tags) != len(tagsDecoded) {
			t.Fatalf("encoded decoded tags count do not match (%d/%d)", len(test.Tags), len(tagsDecoded))
		}

		for k, v := range test.Tags {
			decV, ok := tagsDecoded[k]
			if !ok {
				t.Fatalf("could not find key %s in decoded tags", k)
			}
			if decV != v {
				t.Fatalf("encoded decoded tags values do not match for key %s (%s/%s)", k, v, decV)
			}
		}

	}

}
