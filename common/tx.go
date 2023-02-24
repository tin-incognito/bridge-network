package common

type (
	// TxID is a string that can uniquely represent a transaction on different
	// block chain
	TxID string
	// TxIDs is a slice of TxID
	TxIDs []TxID
)

// BlankTxID represent blank
var BlankTxID = TxID("0000000000000000000000000000000000000000000000000000000000000000")
