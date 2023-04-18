package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	nfe := NewNotFoundError("test_nfe")
	if !IsNotFoundError(nfe) {
		t.Errorf("Expected error to be a NotFoundError")
	}

	bre := NewBadRequestError("test_bre")
	if !IsBadRequestError(bre) {
		t.Errorf("Expected error to be a BadRequestError")
	}

	ere := NewExpiredResourceError("test_ere")
	if !IsExpiredResourceError(ere) {
		t.Errorf("Expected error to be a ExpiredResourceError")
	}

	ise := NewInternalServerError("test_ise")
	if !IsInternalServerError(ise) {
		t.Errorf("Expected error to be a InternalServerError")
	}

	uae := NewUnauthorizedError("test_uae")
	if !IsUnauthorizedError(uae) {
		t.Errorf("Expected error to be a UnauthorizedError")
	}

	rae := NewResourceAlreadyCreatedError("test_rae")
	if !IsResourceAlreadyCreatedError(rae) {
		t.Errorf("Expected error to be a ResourceAlreadyCreatedError")
	}

	ore := NewOutdatedResourceError("test_ore")
	if !IsOutdatedResourceError(ore) {
		t.Errorf("Expected error to be a OutdatedResourceError")
	}

	// Test error messages
	assert.Equal(t, "test_nfe", nfe.Error())
	assert.Equal(t, "test_bre", bre.Error())
	assert.Equal(t, "test_ere", ere.Error())
	assert.Equal(t, "test_ise", ise.Error())
	assert.Equal(t, "test_uae", uae.Error())
	assert.Equal(t, "test_rae", rae.Error())
	assert.Equal(t, "test_ore", ore.Error())
}
