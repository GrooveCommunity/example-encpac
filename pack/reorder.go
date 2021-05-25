package pack

import (
	"bytes"
	"strconv"
)

func DoReorderAndPack(token []byte, fileName string, keyFct uint64) {
	var buffer bytes.Buffer

	for _, elem := range token {
		newElem := uint64(elem)
		nc := newElem * keyFct
		buffer.WriteString(strconv.FormatUint(nc, 2) + " ")
	}

	DoPackage(buffer, fileName)
}
