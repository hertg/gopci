package utils

func ParseByteNum(str []byte) uint {
	var num uint
	for i, s := range str {
		if s >= 97 && s <= 102 {
			s = (s+1)&0b0111 | 0b1000
		} else if s >= 48 && s <= 57 {
			s = s & 0b1111
		}
		num = num | (uint(s) << ((len(str) - i - 1) * 4))
	}
	return num
}
