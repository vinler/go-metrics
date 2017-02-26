package metrics

import (
	"bytes"
	"net/url"
	"sort"
	"strings"
)

const (
	tagEncodingDelimiter = "|"
)

func EncodeNameWithTags(name string, tags map[string]string) string {

	var b bytes.Buffer
	b.WriteString(url.QueryEscape(name))

	if len(tags) > 0 {

		var keys []string
		for k := range tags {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, v := range keys {
			b.WriteString(tagEncodingDelimiter)
			b.WriteString(url.QueryEscape(v))
			b.WriteString(tagEncodingDelimiter)
			b.WriteString(url.QueryEscape(tags[v]))
		}
	}

	return b.String()
}

func DecodeNameWithTags(full string) (string, map[string]string) {

	name := full
	var tags map[string]string = nil

	splits := strings.Split(full, tagEncodingDelimiter)
	splitsLen := len(splits)

	if splitsLen > 1 {
		name = splits[0]
		tags = make(map[string]string)
		for i := 1; i < splitsLen; i += 2 {

			k := splits[i]
			unescapedK, err := url.QueryUnescape(k)
			if err != nil {
				unescapedK = k
			}

			v := splits[i+1]
			unescapedV, err := url.QueryUnescape(v)
			if err != nil {
				unescapedK = v
			}

			tags[unescapedK] = unescapedV
		}
	}

	unescaped, err := url.QueryUnescape(name)
	if err != nil {
		unescaped = name
	}

	return unescaped, tags
}
