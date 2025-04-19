package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// SHA256 used to hash the initial inputs during the process.
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// Adding Base58 -> binary to text encoding will be used to provide the final output of the process.
// Why Base58 instead of Base64? -> Base58 reduces confusion in character output.
// The characters 0,O, I, l are highly confusing when used in certain fonts

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// Hashing  initialUrl + userId url with sha256
func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
