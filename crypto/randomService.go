package crypto

import (
	"crypto/rand"
	"math/big"
)

type RandomService struct{}

func (rs *RandomService) GenerateString(n int) (string, error) {
	buffer := make([]byte, n)
	table := []byte{
		// numbers
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		// ascii uppercase
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		// ascii lowercase
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	for i := 0; i < n; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(table))))
		if err != nil {
			return "", err
		}

		buffer[i] = table[index.Uint64()]
	}

	return string(buffer), nil
}
