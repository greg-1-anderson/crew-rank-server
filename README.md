# Crew Rank Server

# Running

```
$ make generate
$ make create-non-production-certs
$ go run server.go
```

# Sample Query

```
query {
  crewRank {
    ladder {
      name
      overall {
        games {
          won
          lost
          total
        }
      }
      imposter {
        games {
          won
          lost
          total
        }
      }
      crewmate {
        games {
          won
          lost
          total
        }
      }
    }
  }
}
```
