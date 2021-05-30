package payout

func GetSDK(config Config) (SDK, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}
	sdk := Accessor{
		IPPort:          config.IPPort,
		Version:         config.Version,
		clientvalidator: config.ClientValidator,
	}

	return sdk, nil
}
