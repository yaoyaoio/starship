package grafana

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

type AuthToken struct {
	client http.Client
}

func NewAuthToken() *AuthToken {
	return &AuthToken{
		client: http.Client{},
	}
}

func (g *AuthToken) GetToken(apiKey string) (*AuthKeyResult, error) {
	data := make(map[string]interface{})
	data["name"] = apiKey
	data["role"] = "Admin"

	bytesData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", CreateAuthKeysURI, bytes.NewReader(bytesData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(BasicAuth)))
	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err == nil {
		var out AuthKeyResult
		err = json.Unmarshal(body, &out)
		if err != nil {
			return nil, err
		}
		return &out, nil
	}
	return nil, err
}

type AuthKeyResult struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
