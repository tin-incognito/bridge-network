package types

func (m *Blame) IsEmpty() bool {
	if m.FailReason == "" {
		return true
	}
	return false
}
