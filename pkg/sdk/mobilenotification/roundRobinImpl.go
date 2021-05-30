package mobilenotification

import "fmt"

type RoundRobinImpl struct {
	services          []SDK
	RoundRobinCounter roundRobinCounter
}

func (this RoundRobinImpl) SendMessage(number string, message string) (*SendOTPRes, error) {

	serviceToUse := this.RoundRobinCounter.getCounter()

	if serviceToUse > len(this.services) {
		return nil, fmt.Errorf("wrong service used")
	}
	return this.services[serviceToUse].SendMessage(number, message)
}
