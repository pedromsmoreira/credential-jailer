package main

type MockSystemProxy struct {
	mValue string
	mError error
}

func (mr *MockSystemProxy) ReadEnvVar(key string) (string, error) {
	return mr.mValue, mr.mError
}
