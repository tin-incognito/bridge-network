package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KeygenBlockKeyPrefix is the prefix to retrieve all KeygenBlock
	KeygenBlockKeyPrefix = "KeygenBlock/value/"
)

// KeygenBlockKey returns the store key to retrieve a KeygenBlock from the index fields
func KeygenBlockKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
