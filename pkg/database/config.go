package database

import "net/url"

// Config holds the configuration to connect to the database.
type Config struct {
	// Host is the address of the database.
	Host string

	// Port is the port of the database.
	Port string

	// User is the user name to connect to the database.
	User string

	// Password is the password to connect to the database.
	Password string

	// Name is the name of the database you want to connect to.
	Name string

	// SSLMode is the SSL mode of the database.
	// For more information, check https://www.postgresql.org/docs/current/libpq-ssl.html
	SSLMode string
}

// ConnectionURL returns a connection string to connect to the database.
// The connection string is in the format of `postgres://user:password@host:port/database?sslmode=mode`
func (c *Config) ConnectionURL() string {

	host := c.Host
	if c.Port != "" {
		host = host + ":" + c.Port
	}

	u := url.URL{
		Scheme: "postgres",
		Host:   host,
		Path:   c.Name,
	}

	if c.Password != "" || c.User != "" {
		u.User = url.UserPassword(c.User, c.Password)
	}

	q := u.Query()
	if c.SSLMode != "" {
		q.Set("sslmode", c.SSLMode)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
