package generatepassword

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword(n int) string {
	symbols := "abcdifghjklmnopqrstuvw"
	symbolsRune := []rune(symbols)
	required := make([]rune, 0, n)
	for i := 0; i < n; i++ {
		index := randomInt(len(symbolsRune))
		required = append(required, symbolsRune[index])
	}
	return string(required)
}
func randomInt(n int) int {
	if n <= 0 {
		return 0
	}
	bigN := big.NewInt(int64(n))
	val, err := rand.Int(rand.Reader, bigN)
	if err != nil {
		panic(err)
	}
	return int(val.Int64())
}
