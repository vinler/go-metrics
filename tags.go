package metrics

import (
	"bytes"
	"sort"
	"strings"
)

func EncodeNameWithTags(name string, tags map[string]string) string {

	var b bytes.Buffer
	b.WriteString(name)

	if len(tags) > 0 {

		var keys []string
		for k := range tags {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, v := range keys {
			b.WriteString("|")
			b.WriteString(v)
			b.WriteString("|")
			b.WriteString(tags[v])
		}
	}

	return b.String()
}

func DecodeNameWithTags(full string) (string, map[string]string) {

	name := full
	var tags map[string]string = nil

	splits := strings.Split(full, "|")
	splitsLen := len(splits)

	if splitsLen > 1 {
		name = splits[0]
		tags = make(map[string]string)
		for i := 1; i < splitsLen; i += 2 {
			tags[splits[i]] = splits[i+1]
		}
	}

	return name, tags
}
