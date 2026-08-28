package khttp

import (
	"net/http"
	"time"
)

const (
	readTimeout    = 15 * time.Second
	writeTimeout   = 15 * time.Second
	idleTimeout    = 60 * time.Second
	maxHeaderBytes = 1 << 20
)

func NewServer(address string, handler Handler) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        handler,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		IdleTimeout:    idleTimeout,
		MaxHeaderBytes: maxHeaderBytes,
		TLSConfig:      tlsConfig(),
	}
}
