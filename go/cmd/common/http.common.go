package common

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/patos-ufscar/http-web-server-example-go/models"
)

func ParseHeader(header string) (*http.Request, error) {

	return nil, nil
}

func ParseHttpRequestFrame(requestBytes []byte) (*models.HttpRequestFrame, error) {
	var httpReqFrame models.HttpRequestFrame

	lines := strings.Split(string(requestBytes), "\r\n")
	// request := [][]string{}
	for i, v := range lines {
		if i == 0 {
			words := strings.Split(v, " ")
			httpReqFrame.Method = words[0]
			httpReqFrame.RequestURI = words[1]
			httpReqFrame.HTTPVersion = words[2]
		} else {
			words := strings.SplitN(v, " ", 1)

			parseRequestLine(&httpReqFrame, words)
		}
		// request = append(request, words)
	}

	return &httpReqFrame, nil
}

func parseRequestLine(frame *models.HttpRequestFrame, words []string) error {

	switch words[0] {
	case "Content-Type":
		frame.RequestHeaders.ContentType = words[1]
	case "Content-Length":
		val, err := strconv.ParseUint(words[1], 10, 64)
		if err != nil {
			return errors.New("could not convert to uint64")
		}
		frame.RequestHeaders.ContentLength = val
	case "Content-Encoding":
		frame.RequestHeaders.ContentEncoding = words[1]
	case "Content-Language":
		frame.RequestHeaders.ContentLanguage = words[1]

	default:
		return errors.New("could not find match")
	}

	return nil
}