package helpers

import (
	"errors"
	"net"
	"net-cat/internal/myerrors"
	"strconv"
)

func CheckPort(port string) (string, error) {
	n, err := strconv.Atoi(port)
	if err != nil {
		return "", err
	}
	if n < 0 && n > 65353 {
		return "", errors.New(myerrors.IncorectPort)
	}
	return port, nil
}

func IsAscii(name string) bool {
	for _, r := range name {
		if r > 127 {
			return false
		}
	}
	return true
}

func RepeatedName(name string, Listeners map[string]*net.Conn) bool {
	for person, _ := range Listeners {
		if person == name {
			return false
		}
	}
	return true
}
