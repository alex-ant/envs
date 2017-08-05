# envs

Flags to environment variables for Go projects

### Usage example

```go
package main

import (
	"flag"
	"log"

	"github.com/LoudRun/envs"
)

var (
	APIPort = flag.Int("api-port", 30303, "HTTP API port number")

	PgHost    = flag.String("pg-host", "127.0.0.1", "PostgreSQL host")
	PgPort    = flag.Int("pg-port", 5432, "PostgreSQL port")
	PgUser    = flag.String("pg-user", "postgres", "PostgreSQL user")
	PgPass    = flag.String("pg-pass", "postgres", "PostgreSQL password")
	PgDB      = flag.String("pg-db", "events", "PostgreSQL database name")
	PgTimeout = flag.Int("pg-timeout", 30, "PostgreSQL connection timeout in seconds")
)

func main() {
	// Parse flags.
	flag.Parse()

	// Determine and read environment variables.
	flagsErr := envs.GetAllFlags()
	if flagsErr != nil {
		log.Fatal(flagsErr)
	}

	// At this point the application has overwritten its flags' values with the
	// values of the following environment variables if those have been provided:
	// API_PORT, PG_DB, PG_HOST, PG_PASS, PG_PORT, PG_TIMEOUT, PG_USER.
}
```
