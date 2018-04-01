package lk

import (
	"fmt"
	"testing"

	"github.com/dkozic/golksd/ber"
)

func TestParseFileLength(t *testing.T) {

	tlv := ber.TLV{Tag: []byte{0x81}, Length: 2, Value: []byte{0x0, 0x8c}}

	fl := parseFileLength(tlv)

	fmt.Printf("fileLength: %d\n", fl)

	if fl != 2 {
		t.Errorf(`fileLength != 2, %d`, fl)
	}

}
