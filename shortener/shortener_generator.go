package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func hash(input string) []byte {
	algorithm := sha256.New();
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil);
}

func encode(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1);
	}

	return string(encoded)
}

func GenerateShortLink(url string, userId string) string {
	hash := hash(url + userId)
	id := new(big.Int).SetBytes(hash).Uint64()
	encoded := encode([]byte(fmt.Sprintf("%d", id)))
	return encoded[:8]
}