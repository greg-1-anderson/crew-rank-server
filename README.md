# Crew Rank Server

## Running

```
$ make generate
$ make create-non-production-certs
$ make run
```

## Sample Query

```
query {
  crewRank {
    ladder {
      name
      overall {
        games {
          ratio
          won
          lost
          total
        }
      }
      imposter {
        games {
          ratio
          won
          lost
          total
        }
      }
      crewmate {
        games {
          ratio
          won
          lost
          total
        }
      }
    }
  }
}
```
