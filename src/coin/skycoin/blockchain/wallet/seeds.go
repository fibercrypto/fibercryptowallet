package wallet


type SeedService{}


func (seedService *SeedService) GenerateMnemonic(entropy int) (string, error) {
	c := util.NewClient()
	seed, err := c.NewSeed(entropy)
	if err != nil {
		return nil, err
	}

	return seed, nil
}

func (seedService *SeedService) VerifyMnemonic(seed string) (bool, error) {
	c := util.NewClient()
	ok, err := c.VerifySeed(seed)
	if err != nil {
		return nil, err
	}
	return ok, nil
}