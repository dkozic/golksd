package lk

import "github.com/dkozic/golksd/ber"

type TagDocument struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGD_0D = TagDocument{code: "0D"}
	TAGD_0A = TagDocument{code: "0A"}
	TAGD_0B = TagDocument{code: "0B", tagCharset: ber.HEX}
	TAGD_0C = TagDocument{code: "0C"}
	TAGD_0E = TagDocument{code: "0E"}
	TAGD_0F = TagDocument{code: "0F"}
	TAGD_09 = TagDocument{code: "09"}
	TAGD_10 = TagDocument{code: "10", tagCharset: ber.HEX}
	TAGD_11 = TagDocument{code: "11", tagCharset: ber.HEX}
)
