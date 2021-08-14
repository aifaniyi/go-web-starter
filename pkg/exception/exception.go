package exception

// Client-side errors
type ClientError struct {
	Err error
}

func (c *ClientError) Error() string {
	return c.Err.Error()
}

// Server-side errors
type ServerError struct {
	Err error
}

func (c *ServerError) Error() string {
	return c.Err.Error()
}
