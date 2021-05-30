package mobilenotification

import (
	"sync"
)

const (
	ConnectOneService = "connect one service to send msg"
	GupShupService    = "gup shup service to send message"
	YourSMSService    = "your sms service to send message"
	WhatsAppService   = "WhatsAppService to send message"
	YourSMSPath       = "/api/pushsms"
	YourSMSSender     = "BHRTPE"
	JapiService       = "japi service"
)

type roundRobinCounter struct {
	total    int
	lastUsed int32
	m        *sync.RWMutex
}

func (this roundRobinCounter) getCounter() int {
	this.m.RLock()
	if this.total == int(this.lastUsed) {
		this.lastUsed = 0
		this.m.RUnlock()
		return 0
	}
	this.lastUsed = this.lastUsed + 1
	this.m.RUnlock()
	return int(this.lastUsed)
}

func getRoundRobinCounter(serviceCount int) roundRobinCounter {
	return roundRobinCounter{
		total: serviceCount,
	}
}
