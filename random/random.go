package random

import (
	"math"
	"math/rand"
	"time"
)

// Ints returns a random integer array with the specified from, to and size.
func Ints(from, to, size int) []int {
	if to-from < size {
		size = to - from
	}

	var slice []int
	for i := from; i < to; i++ {
		slice = append(slice, i)
	}

	var ret []int
	for i := 0; i < size; i++ {
		idx := rand.Intn(len(slice))
		ret = append(ret, slice[idx])
		slice = append(slice[:idx], slice[idx+1:]...)
	}

	return ret
}

// String returns a random string ['a', 'z'] in the specified length
func String(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Nanosecond)

	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

// Int returns a random integer in range [min, max].
func Int(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Nanosecond)

	return min + rand.Intn(max-min)
}

func Dichotomy(star, count int) int {
	if count == 0 {
		return 0
	}

	res := float64(star) / float64(count)
	pow10_n := math.Pow10(1)
	res2 := math.Trunc(res*pow10_n+0.5) / pow10_n

	return int(res2 * 10)
}

// Int64Rand 随机数
func Int64Rand(size int) string {
	result := make([]byte, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		scope := r.Intn(10)
		if i == 0 && scope == 0 {
			scope = 1
		}

		result[i] = uint8(48 + scope)
	}

	return string(result) //strconv.FormatInt(ret, 10)
}

// Krand 随机字符串
func Krand(size int, rands ...int) []byte {
	kind := 3
	if len(rands) > 0 {
		kind = rands[0]
	}
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

// RandString 随机字符串
func RandString(size int) string {
	return string(Krand(size, 3))
}

// RandNum RandNum
func RandNum(size int) string {
	return string(Krand(size, 0))
}
