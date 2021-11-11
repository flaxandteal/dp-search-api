### Script to reindex ONS index

* The script `reindex-elastic-search` sets up an index with the correct settings and mappings required for
search using the elasticsearch/search-index-settings.json file. 

* It then calls the relevant go scripts (titlebindex) which will copy and transforms data from zebedee to ElasticSearch.


### Getting started

* Run `make build-reindex`

* Also to clear reindex binary - `make clean-reindex`


## Examples 

```bash
./reindex-elastic-search.sh -s titles -i titles-mapping -e http://localhost:9200 -m localhost:8080 -c true
```
* Anything to be executed should be executed from the project root — ie this directory.

* The script `reindex-elastic-search` sets up an index with the correct settings and mappings required for
search using the elasticsearch/search-index-settings.json file. It then calls the relevant go scripts (cmd/reindex/reindex) which will copy and transforms data from zebedee to ElasticSearch.

* `reindex-elastic-search` will ask for several parameters, to view these use the help parameter `-h`
