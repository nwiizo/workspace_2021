package domain

type PiggyBankState uint

const (
	Healthy PiggyBankState = iota + 1
	Broken
)
