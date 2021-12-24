module github.com/ONSdigital/dp-search-api

go 1.17

require (
	github.com/ONSdigital/dp-api-clients-go/v2 v2.5.0
	github.com/ONSdigital/dp-elasticsearch/v2 v2.3.0
	github.com/ONSdigital/dp-healthcheck v1.1.3
	github.com/ONSdigital/dp-net v1.2.0
	github.com/ONSdigital/dp-search-data-extractor v0.2.1-0.20211111103442-c98a47313ed5
	github.com/ONSdigital/dp-search-data-importer v0.1.1-0.20211110125918-56c246e502f7
	github.com/ONSdigital/go-ns v0.0.0-20210831102424-ebdecc20fe9e
	github.com/ONSdigital/log.go/v2 v2.0.9
	github.com/elastic/go-elasticsearch/v7 v7.10.0
	github.com/gorilla/mux v1.8.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.7.2
	github.com/tdewolff/minify v2.3.6+incompatible
)

require (
	github.com/ONSdigital/dp-api-clients-go v1.43.0 // indirect
	github.com/aws/aws-sdk-go v1.38.65 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20210202160940-bed99a852dfe // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20211117102719-0474bc63780f // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/justinas/alice v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/smartystreets/assertions v1.2.0 // indirect
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
	github.com/tdewolff/test v1.0.6 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
)

replace (
	github.com/ONSdigital/dp-elasticsearch/v2 => ../libraries/dp-elasticsearch
	github.com/ONSdigital/dp-search-data-extractor => ../dp-search-data-extractor
	github.com/ONSdigital/dp-search-data-importer => ../dp-search-data-importer
)
