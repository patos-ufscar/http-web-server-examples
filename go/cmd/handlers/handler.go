package handlers

type Handler interface {
	ValidHost(host string)							bool
}
