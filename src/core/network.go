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
	GetBlock() uint64
	IsTrusted() bool
	GetLastSeenIn() int64
	GetLastSeenOut() int64
}

