package testutil

import (
	"net/http"
	"testing"

	"github.com/joshiaj7/vessel-management/module/core/entity"
	"github.com/stretchr/testify/assert"
)

var (
	ErrorUnexpected = entity.NewError("Unexpected error", http.StatusInternalServerError)
)

// TODO: Deprecate Later.
func AssertError(t *testing.T, actual error, expected interface{}) bool {
	t.Helper()

	return AssertErrorExAc(t, expected, actual)
}

// revive:disable:cognitive-complexity,cyclomatic

func AssertErrorExAc(t *testing.T, expected interface{}, actual error) bool {
	t.Helper()

	switch value := expected.(type) {
	case error:
		return assert.ErrorIs(t, actual, value)
	case string:
		return assert.Equal(t, value, actual.Error())
	default:
		return assert.Nil(t, actual)
	}
}
