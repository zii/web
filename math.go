package web

import (
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"math"
	"math/rand"
	"strconv"
)

func RandDigitCode(l int) string {
	var out string
	for i := 0; i < l; i++ {
		a := rand.Intn(10)
		out += strconv.Itoa(a)
	}
	return out
}

func GenerateNonce(size int) []byte {
	b := make([]byte, size)
	_, _ = crand.Read(b)
	return b
}

// 生成的字符串长度是size的2倍
func GenerateStringNonce(size int) string {
	b := make([]byte, size)
	_, _ = crand.Read(b)
	return hex.EncodeToString(b)
}

func AbsInt64(d int64) int64 {
	if d < 0 {
		return -d
	}
	return d
}

func Sha1Digest(data []byte) []byte {
	r := sha1.Sum(data)
	return r[:]
}

func Sha1String(data []byte) string {
	r := sha1.Sum(data)
	return hex.EncodeToString(r[:])
}

func Md5String(data []byte) string {
	r := md5.Sum(data)
	return hex.EncodeToString(r[:])
}

func FileMd5String(f io.Reader) string {
	h := md5.New()
	_, err := io.Copy(h, f)
	Raise(err)
	return hex.EncodeToString(h.Sum(nil))
}

func ShuffleInt64(a []int64) {
	for i := 0; i < len(a); i++ {
		j := rand.Int63n(int64(len(a)))
		a[i], a[j] = a[j], a[i]
	}
}

// 随机采样n条数据
func SampleInt64(src []int64, n int) []int64 {
	var out []int64
	if n <= 0 {
		return out
	}
	sz := len(src)
	if sz == 0 {
		return out
	}
	if n > sz {
		n = sz
	}
	m := rand.Perm(sz)
	for _, i := range m[:n] {
		out = append(out, src[i])
	}
	return out
}

// 随机选取一条
func ChoiceInt64(src []int64) int64 {
	if len(src) == 0 {
		return 0
	}
	i := rand.Intn(len(src))
	return src[i]
}

func ChoiceInts(src []int) int {
	if len(src) == 0 {
		return 0
	}
	i := rand.Intn(len(src))
	return src[i]
}

// 集合: 求差集 A-B
func SubtractInt64s(a, b []int64) []int64 {
	var out []int64
	if len(a) == 0 {
		return out
	}
	if len(b) == 0 {
		copy(out, a)
		return out
	}
	bm := make(map[int64]struct{}, len(a))
	for _, v := range b {
		bm[v] = struct{}{}
	}
	for _, v := range a {
		if _, ok := bm[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

// 保留两位小数(keep 2 decimal places)
func Double2dp(n float64) float64 {
	return math.Trunc(n*1e2+0.5) * 1e-2
}

// 保留两位小数(keep 2 decimal places)
func Float2dp(n float32) float32 {
	return float32(Double2dp(float64(n)))
}
