package common

import (
	cryRand "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

// 得到一个 16 位由小写字母和数字组成的 id, 如： 00347135859b4561
func Get16RandId() string {
	kinds := [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}
	res := make([]byte, 16)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 16; i++ {
		// random ikind
		ikind := rand.Intn(2)
		scope := kinds[ikind][0]
		base := kinds[ikind][1]
		rand.Intn(2)

		res[i] = uint8(base + rand.Intn(scope))
	}

	return string(res)
}

// 获取查询数据的偏移量和查询的限制条数
func GetOffsetLimit(page, pageSize int32) (int, int, error) {
	if page <= 0 {
		page = 1
	}

	limit := pageSize
	if limit <= 0 {
		limit = 15
	} else if limit > 50 {
		limit = 50
	}

	offset := (page - 1) * limit

	return int(offset), int(limit), nil
}

// 在 startNum-maxNum 之间的随机数
func RandNum(startNum, maxNum int64) int64 {
	maxBigInt := big.NewInt(maxNum)
	tmp, _ := cryRand.Int(cryRand.Reader, maxBigInt)
	tmpInt := tmp.Int64()

	if tmpInt >= startNum {
		return tmpInt
	} else {
		return RandNum(startNum, maxNum)
	}
}
