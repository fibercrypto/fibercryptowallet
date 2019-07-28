package wallet

type Address struct {
	address string
}

func (addr Address) IsBip32() bool {
	return true
}

func (addr Address) String() string {
	return addr.address
}
