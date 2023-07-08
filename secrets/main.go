package main

import (
	"encoding/json"
	"fmt"
)

type SecretStringValue string

func (s SecretStringValue) String() string {
	return "*****"
}

func (s SecretStringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

type Provider struct {
	OrgURL   string
	APIToken string
}

type ProviderSecret struct {
	OrgURL   string
	APIToken SecretStringValue
}

func main() {
	oops := Provider{OrgURL: "commonfate.io", APIToken: "secret"}
	oopsSecret := ProviderSecret{OrgURL: "commonfate.io", APIToken: "secret"}
	fmt.Printf("%s\n", oops.APIToken)
	fmt.Printf("%s", oopsSecret.APIToken)
}
