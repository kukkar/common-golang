package handshake

func GetSDK(config SDKConfig) (SDK, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}

	sdk := Accessor{
		IPPort:  config.IPPort,
		Version: config.Version,
	}

	return sdk, nil
}
