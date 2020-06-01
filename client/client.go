package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nicolaferraro/boxnet/api"
	"io/ioutil"
	"net/http"
	"strings"
)

type BoxnetClient struct {
	server string
}

func New() *BoxnetClient {
	return NewWithServer(api.DefaultServer)
}

func NewWithServer(server string) *BoxnetClient {
	return &BoxnetClient{
		server: server,
	}
}

func (c *BoxnetClient) Create(img api.Image) (string, error) {
	data, err := json.Marshal(img)
	if err != nil {
		return "", err
	}
	res, err := http.Post(fmt.Sprintf("%s/images", c.server), "application/json", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(res.Status, "2") {
		return "", fmt.Errorf("error returned from server: %s", res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
