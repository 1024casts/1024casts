package flash

// https://www.alexedwards.net/blog/simple-flash-messages-in-golang

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/lexkong/log"
)

func HasFlash(r *http.Request, name string) bool {
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return false
		default:
			return false
		}
		log.Warnf("[flash] read cookie err: %v", err)
	}
	if c.Value != "" {
		return true
	}

	return false
}

func SetFlash(w http.ResponseWriter, name string, value []byte) {
	log.Info("begin set flash...")
	expire := time.Now().Add(3 * time.Second)
	c := &http.Cookie{Name: name, Value: encode(value), Path: "/", Expires: expire, MaxAge: 3}
	http.SetCookie(w, c)
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string) ([]byte, error) {
	log.Info("begin get flash...")
	c, err := r.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}
	// delete cookie
	// todo: 删除cookie有问题，暂时通过控制有效期来处理
	dc := &http.Cookie{Name: name, Path: "/", MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return value, nil
}

// -------------------------

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
