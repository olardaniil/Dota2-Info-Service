package entity

type Line struct {
	Name     string `json:"name,omitempty"`
	Presence string `json:"presence,omitempty"`
	WinRate  string `json:"win_rate,omitempty"`
}
