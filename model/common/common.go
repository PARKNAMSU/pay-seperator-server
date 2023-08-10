package common_model

type OSInformation struct {
	Name   string
	Number int
	Bit    int64
}

type ApiCommonParams struct {
	UserId         int
	ClientIp       string
	Hostname       string
	UserAgent      string
	ClientOS       OSInformation
	ClientPlatform string
	ClientDevice   string
}
