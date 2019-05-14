package portainer

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var configTest = []byte(`
bearer:
  mode: token
  token: someCrypticToken
basic:
  mode: basic
  token: YWRtaW46cGFzc3dvcmQ=
`)

func initTestsDir() {

	root, _ := ioutil.TempDir("", "")

	cleanup := true
	defer func() {
		if cleanup {
			os.Chdir("..")
			os.RemoveAll(root)
		}
	}()

	os.Chdir(root)
	os.MkdirAll(".portctl", os.ModePerm)
	ioutil.WriteFile("tokens.yaml", configTest, os.ModePerm)
	os.Setenv("HOME", root)

}

func TestNewAuthContext(t *testing.T) {

	context := NewAuthContext()
	if context == nil {
		t.Errorf("context creation should not fail")
	}
}

func TestGetBasicToken(t *testing.T) {
	context := NewAuthContext()
	auth, err := context.GetToken("basic")
	assert.NoError(t, err)
	assert.NotNil(t, auth)
}
