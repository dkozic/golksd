package lk

import "github.com/dkozic/golksd/ber"

type TagResidence struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGR_0D = TagResidence{code: "0D", tagCharset: ber.HEX}
	TAGR_20 = TagResidence{code: "20"}
	TAGR_21 = TagResidence{code: "21"}
	TAGR_22 = TagResidence{code: "22"}
	TAGR_23 = TagResidence{code: "23"}
	TAGR_24 = TagResidence{code: "24"}
	TAGR_2A = TagResidence{code: "2A"}
)
