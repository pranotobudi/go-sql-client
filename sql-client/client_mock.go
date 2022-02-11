package sqlclient

type ClientMock struct {
}

func (c *ClientMock) Query(query string, args ...interface{}) (*SqlRows, error) {
	return nil, nil
}
