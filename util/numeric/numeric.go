package numeric

import (
	"fmt"
	"math"
)

// 변환 하는 함수들을 모아서 Convert라는 객체로 맵핑해서 모두 사용하게 하면 어떨까? StringConv, NumberConv 와 같이
func IntToUint8(i int) (uint8, error) {
	if i >= 0 && i <= math.MaxUint8 {
		return uint8(i), nil
	}

	return 0, fmt.Errorf("$%d cannot convert to uint8 type", i)
}
