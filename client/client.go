package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/container-tools/boxit/api"
	"io/ioutil"
	"net/http"
	"strings"
)

type BoxitClient struct {
	server string
}

func New() *BoxitClient {
	return NewWithServer(api.DefaultServer)
}

func NewWithServer(server string) *BoxitClient {
	return &BoxitClient{
		server: server,
	}
}

func (c *BoxitClient) Create(img api.ImageRequest) (api.ImageResult, error) {
	data, err := json.Marshal(img)
	if err != nil {
		return api.ImageResult{}, err
	}
	res, err := http.Post(fmt.Sprintf("%s/images", c.server), "application/json", bytes.NewReader(data))
	if err != nil {
		return api.ImageResult{}, err
	}
	if !strings.HasPrefix(res.Status, "2") {
		return api.ImageResult{}, fmt.Errorf("error returned from server: %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return api.ImageResult{}, err
	}
	var result api.ImageResult
	if err := json.Unmarshal(body, &result); err != nil {
		return api.ImageResult{}, err
	}
	return result, nil
}
