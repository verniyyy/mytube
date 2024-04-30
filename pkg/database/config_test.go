package database

import "testing"

func TestConnectionString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		config Config
		want   string
	}{
		{
			name: "full config",
			config: Config{
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "password",
				Name:     "database",
				SSLMode:  "disable",
			},
			want: "postgres://postgres:password@localhost:5432/database?sslmode=disable",
		},
		{
			name: "no port",
			config: Config{
				Host:     "localhost",
				User:     "postgres",
				Password: "password",
				Name:     "database",
				SSLMode:  "disable",
			},
			want: "postgres://postgres:password@localhost/database?sslmode=disable",
		},
	}

	for _, cs := range cases {
		cs := cs
		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()

			got := cs.config.ConnectionURL()
			if got != cs.want {
				t.Errorf("want: %s, got: %s", cs.want, got)
			}
		})
	}
}
