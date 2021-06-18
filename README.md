dp-search-api
================
Digital Publishing Search API

A Go application microservice to provide query functionality on the ONS Website

### Getting started

* Run `make debug`

### Dependencies

Clone and set up the following project following the README instructions:
- [dp-compose](https://github.com/ONSdigital/dp-compose)

No further dependencies other than those defined in `go.mod`

### Configuration

An overview of the configuration options available, either as a table of
environment variables, or with a link to a configuration guide.

| Environment variable      | Default                 | Description
| ------------------------- | ----------------------- | ------------------
| BIND_ADDR                 | :23900                  | The host and port to bind to
| ELASTIC_URL	            | http://localhost:9200 | Http url of the ElasticSearch server
| GRACEFUL_SHUTDOWN_TIMEOUT | 5s                      | The graceful shutdown timeout in seconds (`time.Duration` format)

## Releasing
To package up the API uses `make package`

## Deploying
Export the following variables;
* export `DATA_CENTER` to the nomad datacenter to use.
* export `S3_TAR_FILE` to a S3 location on where a release file can be found.
* export `ELASTIC_SEARCH_URL` to elastic search url.

Then run `make nomad` this shall create a nomad plan within the root directory
called `dp-search-api.nomad`

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2016-2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
