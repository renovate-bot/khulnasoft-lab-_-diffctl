package remote

import (
	"testing"

	"github.com/khulnasoft-lab/driftctl/enumeration"
	"github.com/khulnasoft-lab/driftctl/enumeration/alerter"
	"github.com/khulnasoft-lab/driftctl/enumeration/remote/common"

	"github.com/khulnasoft-lab/driftctl/enumeration/resource"

	"github.com/stretchr/testify/assert"
)

func TestScannerShouldIgnoreType(t *testing.T) {

	// Initialize mocks
	alerter := alerter.NewAlerter()
	fakeEnumerator := &common.MockEnumerator{}
	fakeEnumerator.On("SupportedType").Return(resource.ResourceType("FakeType"))
	fakeEnumerator.AssertNotCalled(t, "Enumerate")

	remoteLibrary := common.NewRemoteLibrary()
	remoteLibrary.AddEnumerator(fakeEnumerator)

	testFilter := &enumeration.MockFilter{}
	testFilter.On("IsTypeIgnored", resource.ResourceType("FakeType")).Return(true)

	s := NewScanner(remoteLibrary, alerter, testFilter)
	_, err := s.Resources()
	assert.Nil(t, err)
	fakeEnumerator.AssertExpectations(t)
}
