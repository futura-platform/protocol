package captchaprotocol

import (
	"context"
	"time"
)

type CaptchaParams interface {
	RecaptchaV2Params | RecaptchaV3Params | HcaptchaParams | TurnstileParams | ImageToTextParams | DatadomeParams | AWSWAFTokenParams | GeetestParamsV3 | GeetestParamsV4
}

type Solver interface {
	SolveRecaptchaV2(ctx context.Context, params RecaptchaV2Params, silent bool) (string, time.Duration, error)
	SolveRecaptchaV3(ctx context.Context, params RecaptchaV3Params, silent bool) (string, time.Duration, error)
	SolveHcaptcha(ctx context.Context, params HcaptchaParams, silent bool) (string, time.Duration, error)
	SolveTurnstile(ctx context.Context, params TurnstileParams, silent bool) (string, time.Duration, error)
	SolveImageToText(ctx context.Context, params ImageToTextParams, silent bool) (string, time.Duration, error)
	SolveDatadome(ctx context.Context, params DatadomeParams, silent bool) (string, time.Duration, error)
	SolveAWSWAFToken(ctx context.Context, params AWSWAFTokenParams, silent bool) (string, time.Duration, error)
	SolveGeetestV3(ctx context.Context, params GeetestParamsV3, silent bool) (GeetestSolutionV3, time.Duration, error)
	SolveGeetestV4(ctx context.Context, params GeetestParamsV4, silent bool) (GeetestSolutionV4, time.Duration, error)
}
