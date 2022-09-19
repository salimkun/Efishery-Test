module github.com/salimkun/Efishery-Test/Fetch

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/salimkun/Efishery-Test/Auth v0.0.0-20220918122306-ee60ec1763c6
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/salimkun/Efishery-Test/Auth v0.0.0-20220918122306-ee60ec1763c6 => ../Auth
