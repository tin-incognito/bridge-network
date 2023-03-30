package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RegisterKeygenKeyPrefix is the prefix to retrieve all RegisterKeygen
	RegisterKeygenKeyPrefix = "RegisterKeygen/value/"
)

// RegisterKeygenKey returns the store key to retrieve a RegisterKeygen from the index fields
func RegisterKeygenKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
