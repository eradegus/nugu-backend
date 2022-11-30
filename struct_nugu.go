package main

////////////////////////////////////////////////////////////////////////////////
// Nugu Backend Request
////////////////////////////////////////////////////////////////////////////////

type NuguRequest struct {
	Version string  `json:"version"`
	Action  Action  `json:"action"`
	Event   Event   `json:"event"`
	Context Context `json:"context"`
}
type ResultString struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Parameters struct {
	ResultString ResultString `json:"resultString"`
}
type Action struct {
	ActionName string     `json:"actionName"`
	Parameters Parameters `json:"parameters"`
}
type Event struct {
	Type string `json:"type"`
}
type Session struct {
	AccessToken          string `json:"accessToken"`
	Id                   string `json:"id"`
	IsNew                bool   `json:"isNew"`
	isPlayBuilderRequest bool   `json:"isPlayBuilderRequest"`
}
type State struct {
}
type Device struct {
	Type  string `json:"type"`
	State State  `json:"state"`
}
type AudioPlayer struct {
	PlayerActivity       string `json:"playerActivity"`
	Token                string `json:"token"`
	OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
}
type SupportedInterfaces struct {
	AudioPlayer AudioPlayer `json:"AudioPlayer"`
}
type PrivatePlay struct {
}
type Context struct {
	Session             Session             `json:"session"`
	Device              Device              `json:"device"`
	SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces"`
	PrivatePlay         PrivatePlay         `json:"privatePlay"`
}

////////////////////////////////////////////////////////////////////////////////
// Nugu Backend Response
////////////////////////////////////////////////////////////////////////////////

type NuguResponse struct {
	Version    string       `json:"version"`
	ResultCode string       `json:"resultCode"`
	Output     Output       `json:"output"`
	Directives []Directives `json:"directives"`
}
type Output struct {
	ResultGoodmorning   string `json:"resultGoodmorning"`
	ResultSeeya         string `json:"resultSeeya"`
}
type ProgressReport struct {
	ProgressReportDelayInMilliseconds    int64 `json:"progressReportDelayInMilliseconds"`
	ProgressReportIntervalInMilliseconds int64 `json:"progressReportIntervalInMilliseconds"`
}
type Stream struct {
	URL                   string         `json:"url"`
	OffsetInMilliseconds  int64          `json:"offsetInMilliseconds"`
	ProgressReport        ProgressReport `json:"progressReport"`
	Token                 string         `json:"token"`
	ExpectedPreviousToken string         `json:"expectedPreviousToken"`
}
type Metadata struct {
}
type AudioItem struct {
	Stream   Stream   `json:"stream"`
	Metadata Metadata `json:"metadata"`
}
type Directives struct {
	Type      string    `json:"type"`
	AudioItem AudioItem `json:"audioItem"`
}
