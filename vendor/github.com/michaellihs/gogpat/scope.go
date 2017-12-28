package main

import "net/url"

// Scope define the scope of the personal access token a minimum of api must be set
type Scope struct {
	API          bool
	ReadUser     bool
	Sudo         bool
	ReadRegistry bool
}

const gitlabPersonalAccessTokenScope = "personal_access_token[scopes][]"
const apiScope = "api"
const readUserScope = "read_user"
const sudoScope = "sudo"
const readRegistryScope = "read_registry"

// NewScope exports the permission for gitlab that is currently requested for the token
func NewScope(values url.Values, s Scope) url.Values {
	if s.API == true {
		key, value := addAPIScope()
		values.Add(key, value)
	}

	if s.ReadUser == true {
		key, value := addReadUserScope()
		values.Add(key, value)
	}

	if s.Sudo == true {
		key, value := addSudoScope()
		values.Add(key, value)
	}

	if s.ReadRegistry == true {
		key, value := addReadRegistryScope()
		values.Add(key, value)
	}
	return values
}

func addAPIScope() (key string, value string) {
	return gitlabPersonalAccessTokenScope, apiScope
}

func addReadUserScope() (key string, value string) {
	return gitlabPersonalAccessTokenScope, readUserScope
}

func addSudoScope() (key string, value string) {
	return gitlabPersonalAccessTokenScope, sudoScope
}

func addReadRegistryScope() (key string, value string) {
	return gitlabPersonalAccessTokenScope, readRegistryScope
}
