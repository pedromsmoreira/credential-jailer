package main

type Reader interface {
	Read(key string) (*Credential, error)
}
