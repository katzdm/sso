package mocks

import "github.com/buzzfeed/sso/internal/pkg/groups"

// MockCache is a mock of MemberSetCache that can be used for testing purposes
type MockCache struct {
	GetMembersFunc func(string) (groups.MemberSet, bool)
	Exists         bool
	Updated        bool
	UpdateError    error
	Refreshed      bool
}

// Get returns the members and if the set exists
func (mc *MockCache) Get(group string) (groups.MemberSet, bool) {
	return mc.GetMembersFunc(group)
}

// Update updates the cache
func (mc *MockCache) Update(string) (bool, error) {
	return mc.Updated, mc.UpdateError
}

// RefreshLoop returns a boolean of if the refresh loop is refreshed
func (mc *MockCache) RefreshLoop(string) bool {
	return mc.Refreshed
}

// Stop satisfies the MemberSetCache interface
func (mc *MockCache) Stop() {
	return
}
