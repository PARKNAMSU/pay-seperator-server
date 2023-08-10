package common_constant

import common_model "github.com/PARKNAMSU/pay-seperator/model/common"

var OSInformations = map[string]common_model.OSInformation{
	"WINDOWS": {
		Name:   "Windows",
		Number: 1,
		Bit:    2,
	},
	"WINDOWS PHONE": {
		Name:   "Windows Phone",
		Number: 2,
		Bit:    4,
	},
	"ANDROID": {
		Name:   "Android",
		Number: 3,
		Bit:    8,
	},
	"MACOS": {
		Name:   "macOS",
		Number: 4,
		Bit:    16,
	},
	"IOS": {
		Name:   "iOS",
		Number: 5,
		Bit:    32,
	},
	"LINUX": {
		Name:   "Linux",
		Number: 6,
		Bit:    64,
	},
	"CHROMEOS": {
		Name:   "ChromeOS",
		Number: 7,
		Bit:    128,
	},
}

var DefautOS = OSInformations["WINDOWS"]

var DefaultHostname = "test.com"
