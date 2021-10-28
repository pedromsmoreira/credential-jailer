package main

type VaultReader struct {
}

func (vr *VaultReader) Read(key string) (*Credential, error) {
	// https://learn.hashicorp.com/collections/vault/getting-started
	return nil, nil
}
