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
	Chrome130_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",

		Brands:          `"Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"`,
		FullVersionList: `"Chromium";v="130.0.6723.60", "Google Chrome";v="130.0.6723.60", "Not?A_Brand";v="99.0.0.0"`,

		Platform:        `macOS`,
		PlatformVersion: "130.0.6723.60",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome131_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
		FullVersionList: `"Google Chrome";v="131.0.6778.86", "Chromium";v="131.0.6778.86", "Not_A Brand";v="24.0.0.0"`,

		Platform:        `macOS`,
		PlatformVersion: "131.0.6778.86",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome132_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36",

		Brands:          `"Not A(Brand";v="8", "Chromium";v="132", "Google Chrome";v="132"`,
		FullVersionList: `"Not A(Brand";v="8.0.0.0", "Chromium";v="132.0.6834.160", "Google Chrome";v="132.0.6834.160"`,

		Platform:        `macOS`,
		PlatformVersion: "132.0.6834.160",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome136_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36",

		Brands:          `"Chromium";v="136", "Google Chrome";v="136", "Not.A/Brand";v="99"`,
		FullVersionList: `"Chromium";v="136.0.7103.114", "Google Chrome";v="136.0.7103.114", "Not.A/Brand";v="99.0.0.0"`,

		Platform:        `macOS`,
		PlatformVersion: "136.0.7103.114",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome137_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`,
		FullVersionList: `"Google Chrome";v="137.0.7151.56", "Chromium";v="137.0.7151.56", "Not/A)Brand";v="24.0.0.0"`,

		Platform:        `macOS`,
		PlatformVersion: "137.0.7151.56",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome138_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",

		Brands:          `"Not)A;Brand";v="8", "Chromium";v="138", "Google Chrome";v="138"`,
		FullVersionList: `"Not)A;Brand";v="8.0.0.0", "Chromium";v="138.0.7204.49", "Google Chrome";v="138.0.7204.49"`,

		Platform:        `macOS`,
		PlatformVersion: "138.0.7204.49",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome139_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36",

		Brands:          `"Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"`,
		FullVersionList: `"Not;A=Brand";v="99.0.0.0", "Google Chrome";v="139.0.7258.67", "Chromium";v="139.0.7258.67"`,

		Platform:        `macOS`,
		PlatformVersion: "139.0.7258.67",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome140_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36",

		Brands:          `"Chromium";v="140", "Not=A?Brand";v="24", "Google Chrome";v="140"`,
		FullVersionList: `"Chromium";v="140.0.7339.81", "Not=A?Brand";v="24.0.0.0", "Google Chrome";v="140.0.7339.81"`,

		Platform:        `macOS`,
		PlatformVersion: "140.0.7339.81",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome141_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="141", "Not?A_Brand";v="8", "Chromium";v="141"`,
		FullVersionList: `"Google Chrome";v="141.0.7390.55", "Not?A_Brand";v="8.0.0.0", "Chromium";v="141.0.7390.55"`,

		Platform:        `macOS`,
		PlatformVersion: "141.0.7390.55",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome142_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36",

		Brands:          `"Chromium";v="142", "Google Chrome";v="142", "Not_A Brand";v="99"`,
		FullVersionList: `"Chromium";v="142.0.7444.60", "Google Chrome";v="142.0.7444.60", "Not_A Brand";v="99.0.0.0"`,

		Platform:        `macOS`,
		PlatformVersion: "142.0.7444.60",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome143_mac = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="143", "Chromium";v="143", "Not A(Brand";v="24"`,
		FullVersionList: `"Google Chrome";v="143.0.7499.41", "Chromium";v="143.0.7499.41", "Not A(Brand";v="24.0.0.0"`,

		Platform:        `macOS`,
		PlatformVersion: "143.0.7499.41",
		Architecture:    "arm",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	ChromeLatest_mac = Chrome142_mac

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
	Chrome135_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="135", "Not-A.Brand";v="8", "Chromium";v="135"`,
		FullVersionList: `"Google Chrome";v="135.0.7049.116", "Not-A.Brand";v="8.0.0.0", "Chromium";v="135.0.7049.116"`,

		Platform:        `Windows`,
		PlatformVersion: "135.0.7049.116",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome137_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`,
		FullVersionList: `"Google Chrome";v="137.0.7151.56", "Chromium";v="137.0.7151.56", "Not/A)Brand";v="24.0.0.0"`,

		Platform:        `Windows`,
		PlatformVersion: "137.0.7151.56",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome138_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",

		Brands:          `"Not)A;Brand";v="8", "Chromium";v="138", "Google Chrome";v="138"`,
		FullVersionList: `"Not)A;Brand";v="8.0.0.0", "Chromium";v="138.0.7204.49", "Google Chrome";v="138.0.7204.49"`,

		Platform:        `Windows`,
		PlatformVersion: "138.0.7204.49",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome139_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36",

		Brands:          `"Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"`,
		FullVersionList: `"Not;A=Brand";v="99.0.0.0", "Google Chrome";v="139.0.7258.67", "Chromium";v="139.0.7258.67"`,

		Platform:        `Windows`,
		PlatformVersion: "139.0.7258.67",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome141_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/141.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="141", "Not?A_Brand";v="8", "Chromium";v="141"`,
		FullVersionList: `"Google Chrome";v="141.0.7390.55", "Not?A_Brand";v="8.0.0.0", "Chromium";v="141.0.7390.55"`,

		Platform:        `Windows`,
		PlatformVersion: "141.0.7390.55",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome142_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36",

		Brands:          `"Chromium";v="142", "Google Chrome";v="142", "Not_A Brand";v="99"`,
		FullVersionList: `"Chromium";v="142.0.7444.60", "Google Chrome";v="142.0.7444.60", "Not_A Brand";v="99.0.0.0"`,

		Platform:        `Windows`,
		PlatformVersion: "142.0.7444.60",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	Chrome143_windows = BrowserProfile{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36",

		Brands:          `"Google Chrome";v="143", "Chromium";v="143", "Not A(Brand";v="24"`,
		FullVersionList: `"Google Chrome";v="143.0.7499.41", "Chromium";v="143.0.7499.41", "Not A(Brand";v="24.0.0.0"`,

		Platform:        `Windows`,
		PlatformVersion: "143.0.7499.41",
		Architecture:    "x86",
		Model:           "",
		Mobile:          false,
		Bitness:         "64",
	}
	ChromeLatest_windows = Chrome143_windows

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
