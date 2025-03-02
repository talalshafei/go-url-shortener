package shortener

import (
	"crypto/sha256"
	"log"
	"math/big"
	"strconv"

	"github.com/itchyny/base58-go"
)

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha2560f(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(strconv.FormatUint(generatedNumber, 10)))

	return finalString[:8]
}

func sha2560f(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Panic(err.Error())
	}

	return string(encoded)
}
