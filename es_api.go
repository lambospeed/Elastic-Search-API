package main

import "encoding/json"
import "fmt"
import elastic "gopkg.in/olivere/elastic.v5"
import "golang.org/x/net/context"

// Serie is a structure used for serializing/deserializing data in Elasticsearch.
type Serie struct {
	Name     	 string                `json:"name"`
	IMDBRating   string                `json:"imdb_rating"`
}

func main() {
    fmt.Printf("Elastic Search API\n")
	
	searchByRating(9.5)
}

// Search a in the series index for a serie with a specific rating, ordering the results by name.
func searchByRating(rate float32) {
	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	
	if err != nil {
		// Handle error
		panic(err)
	}
	
	// Search with a term query
	termQuery := elastic.NewTermQuery("imdb_rating.keyword", rate)
	searchResult, err := client.Search().
		Index("series").   // search in index "series"
		Query(termQuery).   // specify the query
		Sort("name.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute

		
	if err != nil {
		// Handle error
		panic(err)
	}
	
	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)
	
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d series\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Serie (could also be just a map[string]interface{}).
			var serie Serie
			err := json.Unmarshal(*hit.Source, &serie)
			if err != nil {
				// Deserialization failed
			}

			// Print the serie obj
			fmt.Printf("Serie name: %s: rating: %s\n", serie.Name, serie.IMDBRating)
		}
	} else {
		// No hits
		fmt.Print("Found no series\n")
	}
}