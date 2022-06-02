package goproxy

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func (proxy *ProxyHttpServer) AuthProxy(w http.ResponseWriter, r *http.Request) (forbidden bool) {
	if r.Method != "CONNECT" && !r.URL.IsAbs() {
		return
	}
	if proxy.Auth == nil || proxy.Auth.Username == "" || proxy.Auth.Password == "" {
		return
	}
	r.Header.Get("Proxy-Authorization")
	authParts := strings.SplitN(r.Header.Get("Proxy-Authorization"), " ", 2)
	if len(authParts) >= 2 {
		authType := authParts[0]
		authData := authParts[1]
		switch authType {
		case "Basic":
			userPassRaw, err := base64.StdEncoding.DecodeString(authData)
			if err == nil {
				userPass := strings.SplitN(string(userPassRaw), ":", 2)
				if len(userPass) == 2 && proxy.Auth.Username == userPass[0] && proxy.Auth.Password == userPass[1] {
					return false
				}
			}
		default:
		}
	}
	if r.Body != nil {
		defer r.Body.Close()
	}
	w.WriteHeader(http.StatusProxyAuthRequired)
	w.Write([]byte("Proxy Authentication Required"))
	return true
}
