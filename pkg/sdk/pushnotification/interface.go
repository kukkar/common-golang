package pushnotification

type SDK interface {
	SendPushNotification(req RequestNotification) error
}
