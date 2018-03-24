package sd

import "github.com/dkozic/golksd/ber"

type TagC struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGC_78 = TagC{code: "78", tagCharset: ber.HEX}
	TAGC_4F = TagC{code: "4F", tagCharset: ber.HEX}
	TAGC_72 = TagC{code: "72", tagCharset: ber.HEX}
	TAGC_80 = TagC{code: "80", tagCharset: ber.HEX}
	TAGC_C1 = TagC{code: "C1"}
	TAGC_C2 = TagC{code: "C2"}
	TAGC_C3 = TagC{code: "C3"}
	TAGC_C4 = TagC{code: "C4"}
	TAGC_C5 = TagC{code: "C5"}
)
