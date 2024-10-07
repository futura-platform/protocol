package captchaprotocol

type Provider interface {
	CreateSolver(serviceType string) (Solver, error)
}
