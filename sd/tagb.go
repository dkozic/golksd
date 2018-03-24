package sd

import "github.com/dkozic/golksd/ber"

type TagB struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGB_78    = TagB{code: "78", tagCharset: ber.HEX}
	TAGB_4F    = TagB{code: "4F", tagCharset: ber.HEX}
	TAGB_72    = TagB{code: "72", tagCharset: ber.HEX}
	TAGB_80    = TagB{code: "80", tagCharset: ber.HEX}
	TAGB_A1    = TagB{code: "A1", tagCharset: ber.HEX}
	TAGB_A7    = TagB{code: "A7", tagCharset: ber.HEX}
	TAGB_A7_83 = TagB{code: "83"}
	TAGB_A7_84 = TagB{code: "84"}
	TAGB_A7_85 = TagB{code: "85"}
	TAGB_A8    = TagB{code: "A8", tagCharset: ber.HEX}
	TAGB_A8_83 = TagB{code: "83"}
	TAGB_A8_84 = TagB{code: "84"}
	TAGB_A8_85 = TagB{code: "85"}
	TAGB_A9    = TagB{code: "A9", tagCharset: ber.HEX}
	TAGB_A9_83 = TagB{code: "83"}
	TAGB_A9_84 = TagB{code: "84"}
	TAGB_A9_85 = TagB{code: "85"}
	TAGB_A4    = TagB{code: "A4", tagCharset: ber.HEX}
	TAGB_96    = TagB{code: "96"}
	TAGB_97    = TagB{code: "97"}
	TAGB_98    = TagB{code: "98"}
	TAGB_99    = TagB{code: "99"}
	TAGB_9A    = TagB{code: "9A"}
	TAGB_AD    = TagB{code: "AD", tagCharset: ber.HEX}
	TAGB_9F1F  = TagB{code: "9F1F"}
	TAGB_9F20  = TagB{code: "9F20"}
	TAGB_9F21  = TagB{code: "9F21"}
	TAGB_9F22  = TagB{code: "9F22"}
	TAGB_9F23  = TagB{code: "9F23"}
	TAGB_AE    = TagB{code: "AE", tagCharset: ber.HEX}
	TAGB_9B    = TagB{code: "9B"}
	TAGB_9C    = TagB{code: "9C"}
	TAGB_A5    = TagB{code: "A5", tagCharset: ber.HEX}
	TAGB_9D    = TagB{code: "9D"}
	TAGB_9E    = TagB{code: "9E"}
	TAGB_9F24  = TagB{code: "9F24"}
	TAGB_9F25  = TagB{code: "9F25"}
	TAGB_AF    = TagB{code: "AF", tagCharset: ber.HEX}
	TAGB_9F26  = TagB{code: "9F26"}
	TAGB_9F27  = TagB{code: "9F27"}
	TAGB_9F28  = TagB{code: "9F28"}
	TAGB_B0    = TagB{code: "B0", tagCharset: ber.HEX}
	TAGB_9F29  = TagB{code: "9F29"}
	TAGB_9F2A  = TagB{code: "9F2A"}
	TAGB_9F2B  = TagB{code: "9F2B"}
	TAGB_9F2C  = TagB{code: "9F2C"}
	TAGB_9F2D  = TagB{code: "9F2D"}
	TAGB_9F2E  = TagB{code: "9F2E"}
	TAGB_9F2F  = TagB{code: "9F2F"}
	TAGB_9F30  = TagB{code: "9F30"}
	TAGB_9F31  = TagB{code: "9F31"}
	TAGB_9F32  = TagB{code: "9F32"}
)
