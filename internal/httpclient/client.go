package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/weibaohui/atomgit-cli/internal/config"
)

type HttpOptions struct {
	Method  string
	Body    interface{}
	Headers map[string]string
}

func Request(path string, options HttpOptions) (*http.Response, error) {
	baseURL := config.GetBaseURL()

	apiPath := path
	if !strings.HasPrefix(path, "/api/") {
		if strings.HasPrefix(path, "/") {
			apiPath = "/api/v5" + path
		} else {
			apiPath = "/api/v5/" + path
		}
	}

	url := strings.TrimSuffix(baseURL, "/") + apiPath

	headers := map[string]string{
		"Accept":     "application/json",
		"User-Agent": "atomgit-cli/0.1.0",
	}

	token := config.GetToken()
	if token != "" {
		headers["Authorization"] = "Bearer " + token
	}

	for k, v := range options.Headers {
		headers[k] = v
	}

	var body io.Reader
	if options.Body != nil {
		headers["Content-Type"] = "application/json"
		jsonData, err := json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		body = bytes.NewBuffer(jsonData)
	}

	method := options.Method
	if method == "" {
		method = "GET"
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return http.DefaultClient.Do(req)
}

func Get(path string) (interface{}, error) {
	resp, err := Request(path, HttpOptions{Method: "GET"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s\n%s", resp.StatusCode, resp.Status, string(respBody))
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		respBody, _ := io.ReadAll(resp.Body)
		return string(respBody), nil
	}
	return result, nil
}

func Post(path string, body interface{}) (interface{}, error) {
	resp, err := Request(path, HttpOptions{Method: "POST", Body: body})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s\n%s", resp.StatusCode, resp.Status, string(respBody))
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		respBody, _ := io.ReadAll(resp.Body)
		return string(respBody), nil
	}
	return result, nil
}

func Delete(path string) error {
	resp, err := Request(path, HttpOptions{Method: "DELETE"})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("HTTP %d: %s\n%s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

func Patch(path string, body interface{}) (interface{}, error) {
	resp, err := Request(path, HttpOptions{Method: "PATCH", Body: body})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s\n%s", resp.StatusCode, resp.Status, string(respBody))
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		respBody, _ := io.ReadAll(resp.Body)
		return string(respBody), nil
	}
	return result, nil
}

func Put(path string, body interface{}) (interface{}, error) {
	resp, err := Request(path, HttpOptions{Method: "PUT", Body: body})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s\n%s", resp.StatusCode, resp.Status, string(respBody))
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		respBody, _ := io.ReadAll(resp.Body)
		return string(respBody), nil
	}
	return result, nil
}
