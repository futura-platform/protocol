package captchaprotocol

type ImageToTextParams struct {
	B64Body string `json:"body"`
	Module  string `json:"module,omitempty"`
	Case    bool   `json:"case,omitempty"`
}
