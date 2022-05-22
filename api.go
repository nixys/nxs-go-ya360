package ya360

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func (ya *Ya360) get(uri url.URL, resp interface{}) (int, error) {

	u := ya.s.URL + uri.String()

	// Create request
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return 0, fmt.Errorf("can't create new request: %v", err)
	}

	// Set Authorization header
	req.Header.Add("Authorization", fmt.Sprintf("OAuth %s", ya.s.OAuth))

	// Make request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	dJ := json.NewDecoder(res.Body)

	e := errorRx{}
	if res.StatusCode != http.StatusOK {
		if err := dJ.Decode(&e); err != nil {
			return res.StatusCode, err
		}
		return res.StatusCode, fmt.Errorf("wrong status code: %s", e.Message)
	}

	if resp != nil {
		// Decode response
		if err := dJ.Decode(&resp); err != nil {
			return res.StatusCode, fmt.Errorf("can't decode response body: %v", err)
		}
	}

	return res.StatusCode, nil
}

func (ya *Ya360) alter(method string, uri url.URL, req interface{}, resp interface{}) (int, error) {

	var rdr io.Reader

	u := ya.s.URL + uri.String()

	if req != nil {
		s, err := json.Marshal(req)
		if err != nil {
			return 0, err
		}
		rdr = strings.NewReader(string(s))
	} else {
		rdr = nil
	}

	r, err := http.NewRequest(method, u, rdr)
	if err != nil {
		return 0, err
	}

	// Set Authorization header
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("OAuth %s", ya.s.OAuth))

	// Make request
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	dJ := json.NewDecoder(res.Body)

	e := errorRx{}
	if res.StatusCode != http.StatusOK {
		if err := dJ.Decode(&e); err != nil {
			return res.StatusCode, err
		}
		return res.StatusCode, fmt.Errorf("wrong status code: %s", e.Message)
	}

	if resp != nil {
		// Decode response
		if err := dJ.Decode(&resp); err != nil {
			return res.StatusCode, fmt.Errorf("can't decode response body: %v", err)
		}
	}

	return res.StatusCode, nil
}
