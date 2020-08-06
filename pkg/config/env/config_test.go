package env

import (
	"testing"
)

func TestNameNotice(t *testing.T) {
	v1 := EnvConfig.ProjectEnv
	println("EnvConfig:", v1)
}
