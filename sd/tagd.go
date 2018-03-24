package sd

import "github.com/dkozic/golksd/ber"

type TagD struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGD_78 = TagD{code: "78", tagCharset: ber.HEX}
	TAGD_4F = TagD{code: "4F", tagCharset: ber.HEX}
	TAGD_72 = TagD{code: "72", tagCharset: ber.HEX}
	TAGD_80 = TagD{code: "80", tagCharset: ber.HEX}
	TAGD_C9 = TagD{code: "C9"}
)
