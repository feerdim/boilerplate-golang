package notification

import (
	"context"
	"os"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/feerdim/boilerplate-golang/log"
	"github.com/mailgun/errors"
)

type Notification struct {
	App     *firebase.App
	Timeout time.Duration
}

func NewNotification() (notif *Notification, err error) {
	timeoutStr := os.Getenv("FIREBASE_CLOUD_MESSAGING_TIMEOUT")

	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse duration on timeout firebase cloud messaging env : %s", timeoutStr)
	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error create firebase app")
	}

	notif = &Notification{
		App:     app,
		Timeout: timeout,
	}

	return
}

func (notif *Notification) PushNotification(title, body, priority string, tokens []string) (err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), notif.Timeout*time.Second)
	defer cancelFunc()

	fcmClient, err := notif.App.Messaging(ctx)
	if err != nil {
		log.WithContext(ctx).Error(err, "error create fcm client")
		return
	}

	response, err := fcmClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
			// ImageURL: "",
		},
		Tokens: tokens,
		Android: &messaging.AndroidConfig{
			Priority: priority, // normal or high
		},
	})
	if err != nil {
		log.WithContext(ctx).Error(err, "error push notification")
		return
	}

	if response.FailureCount > 0 {
		log.WithContext(ctx).Error(err, "error count", "count", response.FailureCount)
		return
	}

	return
}
