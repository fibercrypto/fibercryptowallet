package core


type NetworkIterator interface {
	Value() Network
	Next() bool
	HasNext() bool
}

type NetworkSet interface {
	ListNetwork() NetworkIterator
}

type Network interface {
	GetIp() string
	GetPort() string
	GetSource() string
	GetBlock() string
	IsTrusted() bool
	GetAgent() string
}

