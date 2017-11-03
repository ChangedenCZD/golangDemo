package base64Demo

import (
	"encoding/base64"
	"fmt"
)

var p = fmt.Println
/**base64编码/解码*/
func Run() {
	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}
