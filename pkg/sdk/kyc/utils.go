package kyc

func GetSDK(config Config) (SDK, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}
	sdk := Accessor{
		IPPort:          config.IPPort,
		Version:         config.Version,
		MerchantDBToUse: config.MerchantDBToUse,
		SuperKey:        config.SuperKey,
		clientvalidator: config.ClientValidator,
	}

	return sdk, nil
}
