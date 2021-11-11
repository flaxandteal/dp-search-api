#!/bin/bash -e

./reindex-elastic-search.sh -s titles -i titles-mapping -e http://localhost:11200 -z localhost:8080 -c true

