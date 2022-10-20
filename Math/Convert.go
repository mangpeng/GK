package Math

import "strconv"

func DecimalStringToDecimal(decimal string) (int64, error) {
	return strconv.ParseInt(decimal, 10, 64)
}

func BinaryStringToDecimal(binary string) (int64, error) {
	return strconv.ParseInt(binary, 2, 64)
}

func OctalStringToDecimal(octal string) (int64, error) {
	return strconv.ParseInt(octal, 8, 64)
}

func HexStringToDecimal(hex string) (int64, error) {
	return strconv.ParseInt(hex, 16, 64)
}
