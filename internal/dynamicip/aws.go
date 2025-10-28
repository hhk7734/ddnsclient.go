package dynamicip

import (
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

	r := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	ip := r.Find(data)

	return net.ParseIP(string(ip)), nil
}
