package internal

import (
	"errors"
	"net/url"
)

func MakeUrl(base, path string) (string, error) {
	if len(base) == 0 {
		return "", errors.New("URL Length Invalid")
	}
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	u.Path = u.Path + path
	return u.String(), nil
}
