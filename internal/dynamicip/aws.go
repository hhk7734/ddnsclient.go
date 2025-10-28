package dynamicip

import (
	"errors"
	"io"
	"net"
	"net/http"
	"regexp"
)

var _ IPer = (*AWSIPer)(nil)

type AWSIPer struct{}

func NewAWSIPer() *AWSIPer {
	return &AWSIPer{}
}

func (a AWSIPer) IP() (net.IP, error) {
	res, err := http.Get("https://checkip.amazonaws.com/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ip := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`).Find(data)
	if ip == nil {
		return nil, errors.New("no IP address found in response")
	}

	return net.ParseIP(string(ip)), nil
}
