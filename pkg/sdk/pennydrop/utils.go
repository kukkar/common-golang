package pennydrop

func GetSDK(config Config) (Sdk, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}

	sdk := PennyDropAccessor{
		IPPort:          config.IPPort,
		Version:         config.Version,
		SuperKey:        config.SuperKey,
		MerchantDBToUse: config.MerchantDBToUse,
	}

	return sdk, nil
}
