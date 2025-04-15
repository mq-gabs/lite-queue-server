package utils

func FlattenBytes(bytes [][]byte) []byte {
	var res []byte

	for _, b := range bytes {
		res = append(res, b...)
	}

	return res
}
