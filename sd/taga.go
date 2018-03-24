package sd

import "github.com/dkozic/golksd/ber"

type TagA struct {
	code       string
	tagCharset ber.TagCharset
}

var (
	TAGA_78    = TagA{code: "78", tagCharset: ber.HEX}
	TAGA_4F    = TagA{code: "4F", tagCharset: ber.HEX}
	TAGA_71    = TagA{code: "71", tagCharset: ber.HEX}
	TAGA_80    = TagA{code: "80", tagCharset: ber.HEX}
	TAGA_9F33  = TagA{code: "9F33"}
	TAGA_9F34  = TagA{code: "9F34"}
	TAGA_9F35  = TagA{code: "9F35"}
	TAGA_9F36  = TagA{code: "9F36"}
	TAGA_9F37  = TagA{code: "9F37", tagCharset: ber.HEX}
	TAGA_9F38  = TagA{code: "9F38"}
	TAGA_81    = TagA{code: "81"}
	TAGA_82    = TagA{code: "82"}
	TAGA_A1    = TagA{code: "A1", tagCharset: ber.HEX}
	TAGA_A2    = TagA{code: "A2", tagCharset: ber.HEX}
	TAGA_A2_83 = TagA{code: "83"}
	TAGA_A2_84 = TagA{code: "84"}
	TAGA_A2_85 = TagA{code: "85"}
	TAGA_86    = TagA{code: "86", tagCharset: ber.HEX}
	TAGA_A3    = TagA{code: "A3", tagCharset: ber.HEX}
	TAGA_87    = TagA{code: "87"}
	TAGA_88    = TagA{code: "88"}
	TAGA_89    = TagA{code: "89"}
	TAGA_8A    = TagA{code: "8A"}
	TAGA_A4    = TagA{code: "A4", tagCharset: ber.HEX}
	TAGA_8B    = TagA{code: "8B"}
	TAGA_8C    = TagA{code: "8C"}
	TAGA_8D    = TagA{code: "8D"}
	TAGA_8E    = TagA{code: "8E"}
	TAGA_8F    = TagA{code: "8F"}
	TAGA_A5    = TagA{code: "A5", tagCharset: ber.HEX}
	TAGA_90    = TagA{code: "90"}
	TAGA_91    = TagA{code: "91"}
	TAGA_92    = TagA{code: "92"}
	TAGA_93    = TagA{code: "93"}
	TAGA_A6    = TagA{code: "A6", tagCharset: ber.HEX}
	TAGA_94    = TagA{code: "94"}
	TAGA_95    = TagA{code: "95"}
)
