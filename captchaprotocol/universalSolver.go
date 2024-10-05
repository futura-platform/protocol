package captchaprotocol

import (
	"time"
)

type SolverService string

const (
	TwoCaptcha SolverService = "2captcha"
	CapMonster SolverService = "capmonster"
	CapSolver  SolverService = "capsolver"
)

type CaptchaParams interface {
	RecaptchaV2Params | RecaptchaV3Params | HcaptchaParams | TurnstileParams | ImageToTextParams
}

type Solver interface {
	SetSolverService(solverService SolverService)

	SolveRecaptchaV2(params RecaptchaV2Params, silent bool) (string, time.Duration, error)
	SolveRecaptchaV3(params RecaptchaV3Params, silent bool) (string, time.Duration, error)
	SolveHcaptcha(params HcaptchaParams, silent bool) (string, time.Duration, error)
	SolveTurnstile(params TurnstileParams, silent bool) (string, time.Duration, error)
	SolveImageToText(params ImageToTextParams, silent bool) (string, time.Duration, error)
}
