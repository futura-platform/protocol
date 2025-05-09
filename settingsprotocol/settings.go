package settingsprotocol

import (
	"context"
)

type UserSettings struct {
	WebhookUrl              string `json:"webhookUrl"`
	QueueWebhookUrl         string `json:"queueWebhookUrl"`
	CapMonsterApiKey        string `json:"capMonsterApiKey"`
	TwoCapApiKey            string `json:"twoCapApiKey"`
	EZCaptchaApiKey         string `json:"ezCaptchaApiKey"`
	EnigmaSolverApiKey      string `json:"enigmaSolverApiKey"`
	CapSolverApiKey         string `json:"capSolverApiKey"`
	AntiCaptchaApiKey       string `json:"antiCaptchaApiKey"`
	FreeCaptchaBypassApiKey string
	TextVerifiedApiKey      string
	SMSPoolApiKey           string
	PVADealsApiKey          string
	MajorPhonesLogin        string
	DaisySMSApiKey          string
}

type Provider interface {
	GetSettings(ctx context.Context) (UserSettings, error)
}
