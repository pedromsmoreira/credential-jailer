package main

type SystemProxy interface {
	ReadEnvVar(key string) (string, error)
}
