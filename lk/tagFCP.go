package lk

import "github.com/dkozic/golksd/ber"

type TagFCP struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGF_6F = TagFCP{code: "6F"}
	TAGF_62 = TagFCP{code: "62"}
	TAGF_80 = TagFCP{code: "80"}
	TAGF_81 = TagFCP{code: "81"}
	TAGF_82 = TagFCP{code: "82"}
	TAGF_83 = TagFCP{code: "83"}
	TAGF_84 = TagFCP{code: "84"}
	TAGF_85 = TagFCP{code: "85"}
	TAGF_86 = TagFCP{code: "86"}
	TAGF_87 = TagFCP{code: "87"}
	TAGF_88 = TagFCP{code: "88"}
	TAGF_8A = TagFCP{code: "8A"}
	TAGF_8B = TagFCP{code: "8B"}
	TAGF_8C = TagFCP{code: "8C"}
	TAGF_8D = TagFCP{code: "8D"}
	TAGF_8E = TagFCP{code: "8E"}
	TAGF_A0 = TagFCP{code: "A0"}
	TAGF_A1 = TagFCP{code: "A1"}
	TAGF_A2 = TagFCP{code: "A2"}
	TAGF_A5 = TagFCP{code: "A5"}
	TAGF_AB = TagFCP{code: "AB"}
	TAGF_AC = TagFCP{code: "AC"}
)
