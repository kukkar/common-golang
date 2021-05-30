package emailnotificaiton

import "fmt"

func GetSDK(config Config) (SDK, error) {
	err := config.validateConfig()
	if err != nil {
		return nil, err
	}
	if config.ServiceToUse == SendGrid {
		sdk := SendGridAccessor{
			APIKey: config.APIKey,
		}
		return sdk, nil
	}
	return nil, fmt.Errorf("Not a valid service")
}
