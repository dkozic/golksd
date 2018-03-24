package sd

import (
	"fmt"
	"log"

	"github.com/dkozic/golksd/ber"
)

func bytesToHex(bytes []byte) string {
	return fmt.Sprintf("%X", bytes)
}

func logUnknownTag(tlv ber.TLV, bytes []byte) {
	log.Printf("Unknown tag: %#v, bytes: %#v", tlv, bytes)
}
