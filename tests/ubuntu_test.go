//go:build ubuntu
// +build ubuntu

package tests

import (
	"context"
	"fmt"
	"github.com/morganhein/shoelace/pkg/io"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

////Puts the config file in /usr/var/shoelace/default.toml and defaults are loaded
//func TestLoadConfigFromUsrDefault(t *testing.T) {
//	defaultLocation := "/usr/share/shoelace/default.toml"
//	err := os.Mkdir("/usr/share/shoelace", os.ModeDir)
//	assert.NoError(t, err)
//	_, err = copy("../configs/default.toml", defaultLocation)
//	assert.NoError(t, err)
//	e, err := exists(defaultLocation)
//	assert.NoError(t, err)
//	assert.True(t, e)
//
//	cmd.Execute()
//}
//
////Puts config in $HOME/.config/shoelace/default.toml and defaults are loaded
//func TestLoadDefaultConfigFromHomeConfig(t *testing.T) {
//	homeDir, err := os.UserHomeDir()
//	assert.NoError(t, err)
//	homeConfigLocation := fmt.Sprintf("%v/.config/shoelace/ubuntu.toml", homeDir)
//	err = os.MkdirAll(fmt.Sprintf("%v/.config/shoelace/", homeDir), os.ModeDir)
//	assert.NoError(t, err)
//	_, err = copy("../configs/default.toml", homeConfigLocation)
//	assert.NoError(t, err)
//	e, err := exists(homeConfigLocation)
//	assert.NoError(t, err)
//	assert.True(t, e)
//
//	cmd.Execute()
//}

func TestInstallCommandInstallsPackage(t *testing.T) {
	defaultLocation := "/usr/share/shoelace/default.toml"
	err := os.Mkdir("/usr/share/shoelace", os.ModeDir)
	assert.NoError(t, err)
	_, err = copy("../configs/default.toml", defaultLocation)

	r := io.NewShellRunner()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	res, err := r.Run(ctx, false, "go run main.go install vim")
	assert.NoError(t, err)
	fmt.Println(res)
}
