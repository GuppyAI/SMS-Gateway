package configuration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad_MessagingAllowList_NotSet(t *testing.T) {
	if err := Load(); err != nil {
		t.Fatal(err)
	}

	allowlist := GetConfig().Strings("messaging.allowlist")
	assert.Equal(t, 0, len(allowlist))
}

func TestLoad_MessagingAllowList_Empty(t *testing.T) {
	t.Setenv("GATEWAY_MESSAGING_ALLOWLIST", "")

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	allowlist := GetConfig().Strings("messaging.allowlist")
	assert.Equal(t, 0, len(allowlist))
}

func TestLoad_MessagingAllowList_OneEntry(t *testing.T) {
	t.Setenv("GATEWAY_MESSAGING_ALLOWLIST", "test://testing")

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	allowlist := GetConfig().Strings("messaging.allowlist")
	assert.Equal(t, 1, len(allowlist))
	assert.Contains(t, allowlist, "test://testing")
}

func TestLoad_MessagingAllowList_MultipleEntries(t *testing.T) {
	t.Setenv("GATEWAY_MESSAGING_ALLOWLIST", "test://one,test://two,test://three")

	if err := Load(); err != nil {
		t.Fatal(err)
	}

	allowlist := GetConfig().Strings("messaging.allowlist")
	assert.Equal(t, 3, len(allowlist))
	assert.Contains(t, allowlist, "test://one")
	assert.Contains(t, allowlist, "test://two")
	assert.Contains(t, allowlist, "test://three")
}
