package captchaprotocol

import (
	"context"
	"time"
)

type CaptchaParams interface {
	RecaptchaV2Params | RecaptchaV3Params | HcaptchaParams | TurnstileParams | ImageToTextParams | DatadomeParams | AWSWAFTokenParams | GeetestParamsV3 | GeetestParamsV4
}

type Solver interface {
	SolveRecaptchaV2(ctx context.Context, params RecaptchaV2Params, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveRecaptchaV3(ctx context.Context, params RecaptchaV3Params, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveHcaptcha(ctx context.Context, params HcaptchaParams, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveTurnstile(ctx context.Context, params TurnstileParams, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveImageToText(ctx context.Context, params ImageToTextParams, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveDatadome(ctx context.Context, params DatadomeParams, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveAWSWAFToken(ctx context.Context, params AWSWAFTokenParams, silent bool) (solution string, delay time.Duration, reportFailure func() error, err error)
	SolveGeetestV3(ctx context.Context, params GeetestParamsV3, silent bool) (solution GeetestSolutionV3, delay time.Duration, reportFailure func() error, err error)
	SolveGeetestV4(ctx context.Context, params GeetestParamsV4, silent bool) (solution GeetestSolutionV4, delay time.Duration, reportFailure func() error, err error)
}
