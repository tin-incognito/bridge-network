package types

import (
	"bridge/common"
	"crypto/sha256"
	"encoding/hex"
	fmt "fmt"
	"sort"
	"strconv"
	"strings"
)

type KeygenType int

type Keygen struct {
	Id      string
	Members []string
	Type    KeygenType
}

// NewKeygen create a new instance of Keygen
func NewKeygen(height int64, members []string, keygenType KeygenType) (Keygen, error) {
	// sort the members
	sort.SliceStable(members, func(i, j int) bool {
		return members[i] < members[j]
	})
	id, err := getKeygenID(height, members, keygenType)
	if err != nil {
		return Keygen{}, fmt.Errorf("fail to create new keygen: %w", err)
	}
	return Keygen{
		Id:      string(id),
		Members: members,
		Type:    keygenType,
	}, nil
}

// getKeygenID will create ID based on the pub keys
func getKeygenID(height int64, members []string, keygenType KeygenType) (common.TxID, error) {
	sb := strings.Builder{}
	sb.WriteString(strconv.FormatInt(height, 10))
	sb.WriteString(strconv.Itoa(int(keygenType)))
	for _, m := range members {
		sb.WriteString(m)
	}
	h := sha256.New()
	_, err := h.Write([]byte(sb.String()))
	if err != nil {
		return "", fmt.Errorf("fail to write to hash: %w", err)
	}

	return common.TxID(hex.EncodeToString(h.Sum(nil))), nil
}
