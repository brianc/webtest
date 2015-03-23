package webtest

import (
	"./test"
	"testing"
)

func Test_Easy(t *testing.T) {
	FromHandler(test.MartiniServer())
}
