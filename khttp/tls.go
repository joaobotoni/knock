package khttp

import "crypto/tls"

const minTLSVersion = tls.VersionTLS12

func tlsCurves() []tls.CurveID {
	return []tls.CurveID{
		tls.X25519,
		tls.CurveP256,
	}
}

func tlsCiphers() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}
}

func tlsConfig() *tls.Config {
	return &tls.Config{
		MinVersion:       minTLSVersion,
		CurvePreferences: tlsCurves(),
		CipherSuites:     tlsCiphers(),
	}
}
