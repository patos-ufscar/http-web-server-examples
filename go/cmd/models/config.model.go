package models

import "regexp"

type HandlerConfig struct {
	Port			uint16
	HostsRegs		[]regexp.Regexp
	Locations		[]LocationConfig
}

type LocationConfig struct {
	Path			string
	ReturnType		string
	Return			ReturnConfig
}

type ReturnConfig struct {
	Code			uint16
	Headers			map[string]string
	Body			[]byte
}
