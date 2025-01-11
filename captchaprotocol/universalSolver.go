package captchaprotocol

import (
	"context"
	"time"
)

type CaptchaParams interface {
	RecaptchaV2Params | RecaptchaV3Params | HcaptchaParams | TurnstileParams | ImageToTextParams | DatadomeParams
}

type Solver interface {
	SolveRecaptchaV2(ctx context.Context, params RecaptchaV2Params, silent bool) (string, time.Duration, error)
	SolveRecaptchaV3(ctx context.Context, params RecaptchaV3Params, silent bool) (string, time.Duration, error)
	SolveHcaptcha(ctx context.Context, params HcaptchaParams, silent bool) (string, time.Duration, error)
	SolveTurnstile(ctx context.Context, params TurnstileParams, silent bool) (string, time.Duration, error)
	SolveImageToText(ctx context.Context, params ImageToTextParams, silent bool) (string, time.Duration, error)
	SolveDatadome(ctx context.Context, params DatadomeParams, silent bool) (string, time.Duration, error)
}
