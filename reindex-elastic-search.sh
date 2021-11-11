#!/bin/bash -e

while getopts ":s:i:e:z:c:h" arg
do
  case "$arg" in
    s)
      search=$OPTARG         # 'titles'
    ;;
    i)
      index=$OPTARG          # 'titles-mapping'
    ;;
    e)
      es_url=$OPTARG         # 'http://localhost:9200'
    ;;
    z)
      zebedee_url=$OPTARG      # 'http://localhost:8080'
    ;;
    c)
      create_mapping=$OPTARG # 'true'
    ;;
    h)
      echo ""
      echo "Elastic Loadinator Script"
      echo ""
      echo " * Empties elastic search index of data"
      echo " * Creates new mapping to elastic search index"
      echo " * Calls the relevant bindex to load in data to elastic search index"
      echo ""
      echo "Options are:"
      echo ""
      echo "      OPTION    ENV-VAR        DESCRIPTION                              EXAMPLE ('' does NOT indicate default value)"
      echo "        -s      search         The search type you intend to write to.  titles"
      echo "        -e      es_url         The elastic search url.                  http://localhost:9200"
      echo "        -i      index          The name of the elastic search index.    test-titles"
      echo "        -z      zebedee_url    The zebedee url.                         http://localhost:8082"
      echo "        -c      create_mapping Boolean flag to create mapping or not.   true or false (defaults to false)"
      echo ""
      exit 0
    ;;
    \?)
      echo "ERROR: Unknown option $OPTARG"
    ;;
  esac
done

search=${search:?ERROR: var not set [-s search]}
index=${index:?ERROR: var not set [-i index]}
es_url=${es_url:?ERROR: var not set [-e es_url]}
zebedee_url=${zebedee_url:?ERROR: var not set [-m zebedee_url]}

full_es_url="$es_url/"$index
echo "            search: $search"
echo "elastic search url: $full_es_url"
echo "      zebedee url: $zebedee_url"

# Check load type
if [ $search = "titles" ]
then
    bindex="./cmd/reindex/reindex"
    scheme="search-index-settings.json"
else
    echo "Incorrect search - use titles"
    echo "Use -h for further options"
    exit 1
fi

echo "++++++++++++++++++++++++++++++++++++++++++++++++++++++"
echo "STEP 1: Delete existing index if -c flag set to true"
if [ $create_mapping = "true" ]
then
    echo "DELETING INDEX $full_es_url"
    delete_index="curl -XDELETE $full_es_url"
    echo $delete_index
    delete_index_response=`$delete_index`
    echo $delete_index_response
else
    echo "NOT DELETING INDEX"
fi

echo "++++++++++++++++++++++++++++++++++++++++++++++++++++++"
echo "STEP 2: Create index with new mapping if -c flag set to true"
if [ $create_mapping = "true" ]
then
    echo "CREATING INDEX WITH NEW MAPPING $full_es_url"
    curl -XPUT -H "Content-Type: application/json" $full_es_url -d@./elasticsearch/$scheme
else
    echo "NOT CREATING INDEX WITH NEW MAPPING"
fi

echo "bindex: $bindex"

echo "++++++++++++++++++++++++++++++++++++++++++++++++++++++"
echo "STEP 3: Start $type load"
upload="$bindex -zebedee-url=$zebedee_url -es-dest-url=$es_url -es-dest-index=$index"
echo $upload
exec $upload


