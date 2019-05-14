package portainer

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

// AuthContext is the configuration and authentication cache
type AuthContext struct {
	tokens *viper.Viper
}

// NewAuthContext returns the portainer context loading authentication tokens from a file
func NewAuthContext() *AuthContext {
	c := &AuthContext{}
	if err := c.loadConfigFile(); err != nil {
		fmt.Println(err)
	}
	return c
}

func (c *AuthContext) loadConfigFile() error {
	home, _ := homedir.Dir()
	c.tokens = viper.New()
	c.tokens.AddConfigPath(filepath.Join(home, ".portctl"))
	c.tokens.SetConfigName("tokens")
	c.tokens.SetConfigType("yaml")
	return c.tokens.ReadInConfig()
}

// GetToken return the current token associated with an host
func (c *AuthContext) GetToken(host string) (runtime.ClientAuthInfoWriter, error) {
	authConfig := c.tokens.GetStringMapString(host)
	if authConfig != nil {
		if authConfig["mode"] == "basic" {
			bytes, err := base64.StdEncoding.DecodeString(authConfig["token"])
			if err != nil {
				return nil, err
			}
			creds := strings.Split(string(bytes), ":")
			if len(creds) != 2 {
				return nil, errors.New("unable to parse the credentials")
			}
			return httptransport.BasicAuth(creds[0], creds[1]), nil

		}
		return httptransport.BearerToken(authConfig["token"]), nil
	}
	return nil, nil
}

// SetToken sets the current token associated with an host
func (c *AuthContext) SetToken(host, mode, token string) {
	tokenMap := make(map[string]string)
	tokenMap["mode"] = mode
	tokenMap["token"] = token
	c.tokens.Set(host, tokenMap)
}

// SaveContext writes all the current tokens in a file
func (c *AuthContext) SaveContext() error {
	home, _ := homedir.Dir()
	os.MkdirAll(filepath.Join(home, ".portctl"), os.ModePerm)
	return c.tokens.WriteConfigAs(filepath.Join(home, ".portctl", "tokens.yaml"))
}
