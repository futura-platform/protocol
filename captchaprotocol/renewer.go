package captchaprotocol

import (
	"context"
	"fmt"
	"log"
	"time"
)

func GetRenewer[T CaptchaParams](
	ctx context.Context,
	t interface {
		Solver

		GetErrorDelay() time.Duration
		FatalError(err error)
		BLog() *log.Logger
	},
	maxConsecutiveErrors int,
	params T,
) *CaptchaTokenRenewer {
	var solver func() (string, time.Duration, error)
	switch p := any(params).(type) {
	case RecaptchaV2Params:
		solver = func() (string, time.Duration, error) { return t.SolveRecaptchaV2(ctx, p, true) }
	case RecaptchaV3Params:
		solver = func() (string, time.Duration, error) { return t.SolveRecaptchaV3(ctx, p, true) }
	case HcaptchaParams:
		solver = func() (string, time.Duration, error) { return t.SolveHcaptcha(ctx, p, true) }
	case TurnstileParams:
		solver = func() (string, time.Duration, error) { return t.SolveTurnstile(ctx, p, true) }
	case ImageToTextParams:
		solver = func() (string, time.Duration, error) { return t.SolveImageToText(ctx, p, true) }
	default:
		panic("Unknown captcha type: " + fmt.Sprintf("%T", params))
	}
	return NewCaptchaTokenRenewer(
		ctx,
		t.BLog(),
		t.GetErrorDelay(),
		solver,
		maxConsecutiveErrors,
		t.FatalError,
	)
}

func NewCaptchaTokenRenewer(ctx context.Context, blog *log.Logger, errorDelay time.Duration, solver func() (string, time.Duration, error), maxConsecutiveErrors int, onError func(error)) *CaptchaTokenRenewer {
	renewer := &CaptchaTokenRenewer{
		onNextToken: make(chan func(*captchaToken) bool),
	}

	go func() {
		consecutiveErrors := 0
		var latestToken *captchaToken
		var latestValidFor time.Duration
		for {
			if latestToken == nil {
				token, validfor, err := solver()
				if err != nil {
					if blog != nil {
						blog.Println("error renewing captcha:", err)
					}

					consecutiveErrors++
					if consecutiveErrors >= maxConsecutiveErrors {
						onError(fmt.Errorf("too many consecutive captcha solve errors (%d) latest error: %w", maxConsecutiveErrors, err))
						return
					}

					time.Sleep(errorDelay)
					continue
				}

				latestToken = &captchaToken{token: token, createdAt: time.Now()}
				latestValidFor = validfor
			}
			select {
			case <-ctx.Done():
				return
			// give the api 30 seconds of buffer time to gen
			case <-time.After(latestValidFor - time.Second*30):
			case callback := <-renewer.onNextToken:
				if callback(latestToken) {
					latestToken = nil
				}
			}
		}
	}()

	return renewer
}

type CaptchaTokenRenewer struct {
	onNextToken chan func(*captchaToken) (didConsume bool)
	latestToken *captchaToken
}

type captchaToken struct {
	token     string
	createdAt time.Time
}

func (r *CaptchaTokenRenewer) GetToken(ctx context.Context) (string, error) {
	t := r.latestToken
	if r.latestToken == nil {
		c := make(chan *captchaToken)
		defer close(c)
		go func() {
			r.onNextToken <- func(ct *captchaToken) bool {
				ctxActive := ctx.Err() == nil
				if ctxActive {
					c <- ct
				}
				return ctxActive
			}
		}()

		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case t = <-c:
		}
	}
	r.latestToken = nil

	return t.token, nil
}
