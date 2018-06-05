# envs

Flags to environment variables for Go projects

### Installation

`go get github.com/alex-ant/envs`

### About the project

`envs` allows Go projects to read environment variables by setting the corresponding flags. Consider the execution example below:

```
go run main.go -api-port 8080
```

This program expects the `api-port` flag of the `int` type.

Importing the `envs` library, the mentioned program will automatically determine the `API_PORT` environment variable which makes it convenient to run it in an environment like Docker Compose.

The library reads all the flags expected by the binary replacing dashes (`-`) with underscores (`_`), capitalizing the letters and looking for the respective environment variables whose values, if set, are set to flags' values.

### Usage example

```go
package main

import (
	"flag"
	"log"

	"github.com/alex-ant/envs"
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

### Printing all expected flags and environment variables

After importing the `envs` library, a program can be run with `-envs` flag which will force it to print the table of expected flags, the corresponding environment variables, default values and exit at once.

Example or running the program described in "Usage example" section:

```
[alex.ant@localhost tmp]$ PG_DB=my-db go run env-example.go -envs
     FLAG    | ENVIRONMENT VAR | DEFAULT VALUE | CURRENT VALUE |          DESCRIPTION            
+------------+-----------------+---------------+---------------+--------------------------------+
  api-port   | API_PORT        | 30303         | 30303         | HTTP API port number            
  pg-db      | PG_DB           | events        | my-db         | PostgreSQL database name        
  pg-host    | PG_HOST         | 127.0.0.1     | 127.0.0.1     | PostgreSQL host                 
  pg-pass    | PG_PASS         | postgres      | postgres      | PostgreSQL password             
  pg-port    | PG_PORT         | 5432          | 5432          | PostgreSQL port                 
  pg-timeout | PG_TIMEOUT      | 30            | 30            | PostgreSQL connection timeout   
             |                 |               |               | in seconds                      
  pg-user    | PG_USER         | postgres      | postgres      | PostgreSQL user                 
```

This action will also produce the `envs.md` file with the same data formatted as markdown which can be added to project's README in the following form:

```
|Flag|Env. variable|Default value|Description|
|:----|:----|:---|:---|
|api-port|API_PORT|30303|HTTP API port number|
|pg-db|PG_DB|events|PostgreSQL database name|
|pg-host|PG_HOST|127.0.0.1|PostgreSQL host|
|pg-pass|PG_PASS|postgres|PostgreSQL password|
|pg-port|PG_PORT|5432|PostgreSQL port|
|pg-timeout|PG_TIMEOUT|30|PostgreSQL connection timeout in seconds|
|pg-user|PG_USER|postgres|PostgreSQL user|
```
