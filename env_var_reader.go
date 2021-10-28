package main

type EnvVarReader struct {
	proxy SystemProxy
}

func (evr *EnvVarReader) Read(key string) (*Credential, error) {
	v, err := evr.proxy.ReadEnvVar(key)
	if err == nil {
		return &Credential{
			Name:  key,
			Value: v,
		}, nil
	}

	return nil, err
}
