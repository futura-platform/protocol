package settingsprotocol

import (
	"context"
	"fmt"
)

type UserSettings struct {
	WebhookUrl         string `json:"webhookUrl"`
	QueueWebhookUrl    string `json:"queueWebhookUrl"`
	CapMonsterApiKey   string `json:"capMonsterApiKey"`
	TwoCapApiKey       string `json:"twoCapApiKey"`
	EZCaptchaApiKey    string `json:"ezCaptchaApiKey"`
	CapSolverApiKey    string `json:"capSolverApiKey"`
	TextVerifiedApiKey string
	SMSPoolApiKey      string
	PVADealsApiKey     string
	MajorPhonesLogin   string
	DaisySMSApiKey     string
}

type Provider interface {
	GetSettings(ctx context.Context) (UserSettings, error)
}

var AmbientProvider Provider

func GetSettings(ctx context.Context) (UserSettings, error) {
	if AmbientProvider == nil {
		return UserSettings{}, fmt.Errorf("ambient provider not set")
	}

	return AmbientProvider.GetSettings(ctx)
}
