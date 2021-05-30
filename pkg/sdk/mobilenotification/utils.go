package mobilenotification

func GetSDK(config Config) (SDK, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}
	var sdk SDK
	if config.ServiceToUse == ConnectOneService {
		sdk = ConnectOneAccessor{
			IPPort:   config.IPPort,
			UserName: config.UserName,
			Password: config.Password,
		}
	} else if config.ServiceToUse == YourSMSService {
		sdk = YourMessageAccessor{
			IPPort:   config.IPPort,
			UserName: config.UserName,
			AuthKey:  config.AuthKey,
		}
	} else if config.ServiceToUse == WhatsAppService {
		sdk = WhatsAppAccessor{
			IPPort:   config.IPPort,
			UserName: config.UserName,
			Password: config.Password,
		}
	} else if config.ServiceToUse == JapiService {
		sdk = JapiAccessor{
			IPPort:   config.IPPort,
			UserName: config.UserName,
			Password: config.Password,
			AuthKey:  config.AuthKey,
		}
	} else {
		sdk = GupShupAccessor{
			IPPort:   config.IPPort,
			UserName: config.UserName,
			Password: config.Password,
		}
	}
	return sdk, nil
}

func GetRoundRobinInstance(sdks []Config) (RoundRobinOTPService, error) {

	services := make([]SDK, 0)
	for _, eachConfig := range sdks {
		sdk, err := GetSDK(eachConfig)
		if err != nil {
			return nil, err
		}
		services = append(services, sdk)
	}
	return RoundRobinImpl{
		services:          services,
		RoundRobinCounter: getRoundRobinCounter(len(services)),
	}, nil
}
