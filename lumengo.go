package lumengo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const DefaultUserAgent = "Uselumen.co Go Client/" + Version

const (
	BaseUrl string = "https://api.uselumen.co"
)

type (
	Identifier string
)

type (
	// Lumengo is the main struct of the Lumengo package.
	Lumengo struct {
		apiKey    string
		Client    *http.Client
		UserAgent string
	}

	IdentifyParams struct {
		FirstName  string                 `json:"first_name"` // optional
		LastName   string                 `json:"last_name"`  // optional
		Email      string                 `json:"email"`      // compulsory
		DeviceId   string                 `json:"device_id"`  // optional
		Identifier string                 `json:"identifier"` // optional. One would be generated if none is provided.
		Attributes map[string]interface{} `json:"attributes"` // optional
	}

	TrackParams struct {
		Identifier string                 `json:"identifier"` // compulsory
		Event      string                 `json:"event"`      // compulsory
		Properties map[string]interface{} `json:"properties"` // optional
	}
)

// CustomerIOError is returned by any method that fails at the API level
type LumengoError struct {
	status  int    //
	message string //    string
	body    []byte
}

type ApiError struct {
	Message string `json:"message"`
}

func (e *LumengoError) Error() string {
	return fmt.Sprintf("%v: %v %v", e.status, e.message, string(e.body))
}

func NewLumengo(apiKey string) *Lumengo {
	return &Lumengo{
		apiKey:    apiKey,
		UserAgent: DefaultUserAgent,
	}
}

// Identify with context
func (l *Lumengo) IdentifyCtx(ctx context.Context, identifier Identifier, params IdentifyParams) error {

	if identifier == "" {
		return errors.New("Identifier is required")
	}

	if params.Email == "" {
		return errors.New("Email is required")
	}

	params.SetIdentifier(identifier)
	fmt.Println("params: ", params)

	uri := fmt.Sprintf("%v/customer/identify", BaseUrl)
	return l.request(context.Background(), "POST", uri, params)
}

func (l *Lumengo) Identify(identifier Identifier, params IdentifyParams) error {
	return l.IdentifyCtx(context.Background(), identifier, params)
}

func (l *Lumengo) TrackCtx(ctx context.Context, identifier Identifier, eventName string, params TrackParams) error {

	if identifier == "" {
		return errors.New("Identifier is required")
	}

	uri := fmt.Sprintf("%v/event/track", BaseUrl)
	return l.request(ctx, "POST", uri, params)

}

func (l *Lumengo) Track(identifier Identifier, eventName string, params TrackParams) error {
	return l.TrackCtx(context.Background(), identifier, eventName, params)
}

// Lumengo Methods
func (l *Lumengo) SetApiKey(apiKey string) {
	l.apiKey = apiKey
}

func (l *Lumengo) GetApiKey() string {
	return l.apiKey
}

// Client
func (l *Lumengo) request(ctx context.Context, method, url string, body interface{}) error {
	var req *http.Request

	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(method, url, bytes.NewBuffer(b))
		if err != nil {
			return err
		}
		req = req.WithContext(ctx)

		req.Header.Add("User-Agent", l.UserAgent)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Content-Length", strconv.Itoa(len(b)))
	} else {
		var err error
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	}

	req.Header.Add("api_key", l.GetApiKey())

	resp, err := l.Client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {

		errorBody := &ApiError{}
		b := json.Unmarshal(responseBody, errorBody)

		if b != nil {
			return &LumengoError{
				status:  resp.StatusCode,
				message: "Something went wrong. Please contact support",
				body:    responseBody,
			}
		}

		return &LumengoError{status: resp.StatusCode, message: errorBody.Message, body: responseBody}
	}

	return nil
}

// Identifier

func (i Identifier) ToString() string {
	return string(i)
}
func (i *IdentifyParams) SetIdentifier(identifier Identifier) {
	i.Identifier = strings.TrimSpace(identifier.ToString())
}

// TrackParams
func (i *TrackParams) SetIdentifier(identifier Identifier) {
	i.Identifier = strings.TrimSpace(identifier.ToString())
}
