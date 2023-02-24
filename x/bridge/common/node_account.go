package common

type NodeAddress []byte
type NodeStatus int
type NodeAccounts []NodeAccount
type SignerMembership string

type NodeAccount struct {
	NodeAddress      NodeAddress
	Status           NodeStatus
	PubKeySet        PubKeySet
	SignerMembership []SignerMembership
}
