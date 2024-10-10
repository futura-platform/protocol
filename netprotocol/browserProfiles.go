package netprotocol

import (
	"fmt"
)

type BrowserProfile struct { // https://browserleaks.com/client-hints
	UserAgent string

	Brands          string
	FullVersionList string

	Platform        string
	PlatformVersion string
	Architecture    string
	Model           string
	Mobile          bool
	Bitness         string
}

func (b BrowserProfile) GetHeaderDefaults() map[string]string {
	h := map[string]string{
		"user-agent": b.UserAgent,

		"sec-ch-ua":                   b.Brands,
		"sec-ch-ua-full-version-list": b.FullVersionList,

		"sec-ch-ua-platform":         fmt.Sprintf(`"%s"`, b.Platform),
		"sec-ch-ua-platform-version": fmt.Sprintf(`"%s"`, b.PlatformVersion),
		"sec-ch-ua-arch":             fmt.Sprintf(`"%s"`, b.Architecture),
		"sec-ch-ua-model":            fmt.Sprintf(`"%s"`, b.Model),
		"sec-ch-ua-bitness":          fmt.Sprintf(`"%s"`, b.Bitness),
	}
	if b.Mobile {
		h["sec-ch-ua-mobile"] = "?1"
	} else {
		h["sec-ch-ua-mobile"] = "?0"
	}

	return h
}

var (
	Chrome126_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",

		Brands:          `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`,
		FullVersionList: `"Not)A;Brand";v="99.0.0.0", "Google Chrome";v="126.0.6478.178", "Chromium";v="126.0.6478.178"`,

		Platform:        `macOS`,
		PlatformVersion: "126.0.6478.178",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome127_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",

		Brands:          `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`,
		FullVersionList: `"Not)A;Brand";v="99.0.0.0", "Google Chrome";v="127.0.6533.72", "Chromium";v="127.0.6533.72"`,

		Platform:        `macOS`,
		PlatformVersion: "127.0.6533.72",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome128_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",

		Brands:          `"Chromium";v="128", "Not;A=Brand";v="24", "Google Chrome";v="128"`,
		FullVersionList: `"Chromium";v="128.0.6613.137", "Not;A=Brand";v="24.0.0.0", "Google Chrome";v="128.0.6613.137"`,

		Platform:        `macOS`,
		PlatformVersion: "128.0.6613.137",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome129_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`,
		FullVersionList: `"Google Chrome";v="129.0.6668.90", "Not=A?Brand";v="8.0.0.0", "Chromium";v="129.0.6668.90"`,

		Platform:        `macOS`,
		PlatformVersion: "129.0.6668.90",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	ChromeLatest_mac = Chrome129_mac

	Chrome127_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36",

		Brands:          `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`,
		FullVersionList: `"Not)A;Brand";v="99.0.0.0", "Google Chrome";v="127.0.6533.89", "Chromium";v="127.0.6533.89"`,

		Platform:        `Windows`,
		PlatformVersion: "127.0.6533.89",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}

	Chrome127_android = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36",

		Brands:          `"Not)A;Brand";v="99", "Google Chrome";v="127", "Chromium";v="127"`,
		FullVersionList: `"Not)A;Brand";v="99.0.0.0", "Google Chrome";v="127.0.6533.84", "Chromium";v="127.0.6533.84"`,

		Platform:        `Android`,
		PlatformVersion: "127.0.6533.84",
		Architecture:    "",
		Model:           "sdk_gphone64_arm64",
		Mobile:          true,
		Bitness:         "",
	}
)
