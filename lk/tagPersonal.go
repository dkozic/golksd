package lk

import "github.com/dkozic/golksd/ber"

type TagPersonal struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGP_0D   = TagPersonal{code: "0D", tagCharset: ber.HEX}
	TAGP_16   = TagPersonal{code: "16"}
	TAGP_17   = TagPersonal{code: "17", tagCharset: ber.UTF8}
	TAGP_18   = TagPersonal{code: "18"}
	TAGP_19   = TagPersonal{code: "19"}
	TAGP_1A   = TagPersonal{code: "1A"}
	TAGP_1B   = TagPersonal{code: "1B"}
	TAGP_1D   = TagPersonal{code: "1D"}
	TAGP_1E   = TagPersonal{code: "1E"}
	TAGP_1F06 = TagPersonal{code: "1F06"}
)
