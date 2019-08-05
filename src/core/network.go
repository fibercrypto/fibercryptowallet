package core


type NetworkIterator interface {
	Value() INetwork
	Next() bool
	HasNext() bool
}

type NetworkSet interface {
	ListNetworks() NetworkIterator
}

type INetwork interface {
	GetIp() string
	GetPort() uint16
	GetSource() string
	GetBlock() uint64
	IsTrusted() bool
	GetAgent() string
}

