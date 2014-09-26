package server

import (
	"crypto/tls"
	"io/ioutil"
	"ngrok/server/assets"
"fmt"
)

func LoadTLSConfig(crtPath string, keyPath string) (tlsConfig *tls.Config, err error) {

fmt.Println("[DEBUG] crtPath=", crtPath, "keyPath=", keyPath)

	fileOrAsset := func(path string, default_path string) ([]byte, error) {
		loadFn := ioutil.ReadFile
		if path == "" {
			loadFn = assets.Asset
			path = default_path
		}

fmt.Println("[HACK] 1")
		return loadFn(path)
	}

	var (
		crt  []byte
		key  []byte
		cert tls.Certificate
	)

	if crt, err = fileOrAsset(crtPath, "assets/server/tls/snakeoil.crt"); err != nil {
fmt.Println("[HACK] 2")
		return
	}

	if key, err = fileOrAsset(keyPath, "assets/server/tls/snakeoil.key"); err != nil {
fmt.Println("[HACK] 3")
		return
	}

	if cert, err = tls.X509KeyPair(crt, key); err != nil {
fmt.Println("[HACK] i4")
		return
	}

	tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

fmt.Println("[HACK] 5")
	return
}
