package httpc

import (
	"go.mamad.dev/telebot/log"
	"net/http"
	"net/http/httputil"
)

type LoggingTransport struct {
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface
func (lt *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log the request
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Errorf("Error dumping request: %s", err)
	} else {
		log.Debugf("Request: %s\n", string(requestDump))
	}

	// Perform the actual request using the original Transport
	resp, err := lt.Transport.RoundTrip(req)

	// Log the response
	if err != nil {
		log.Warnf("Error making request: %s", err)
	} else {
		responseDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Errorf("Error dumping response: %s", err)
		} else {
			log.Debugf("Response: %s", string(responseDump))
		}
	}

	return resp, err
}
