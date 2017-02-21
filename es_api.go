package main

import "fmt"
import "encoding/json"
import elastic "gopkg.in/olivere/elastic.v5"
import "golang.org/x/net/context"

// Serie is a structure used for serializing/deserializing data in Elasticsearch.
type Serie struct {
	Name     	 string                `json:"name"`
	IMDBRating   string                `json:"imdb_rating"`
}

// Search all series.
func getAllSeries() []Serie{

	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	
	if err != nil {
		// Handle error
		panic(err)
	}
	
	// Search with a term query
	searchResult, err := client.Search().
		Index("series").
		Sort("imdb_rating.keyword", false).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)
	
	if err != nil {
		// Handle error
		panic(err)
	}
	
	// get the number of results and create array
	resultSize := searchResult.Hits.TotalHits
	series := make([]Serie, resultSize, 2*resultSize)
	
	if searchResult.Hits.TotalHits > 0 {
		// Iterate through results
		for i, hit := range searchResult.Hits.Hits {
		
			// Deserialize hit.Source into a Serie
			var serie Serie
			err := json.Unmarshal(*hit.Source, &serie)
			if err != nil {} // Deserialization failed
			series[i] = serie
		}
	} 
	
	return series
}

// Search a in the series index for a serie with a specific name.
func getSeriesByName(serieName string) []Serie{
	fmt.Printf("Looking for %s \n", serieName)

	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	
	if err != nil {
		// Handle error
		panic(err)
	}
	
	// Search with a term query
	termQuery := elastic.NewTermQuery("name.keyword", serieName)
	searchResult, err := client.Search().
		Index("series").
		Query(termQuery).
		Sort("name.keyword", true).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
	
	// get the number of results and create array
	resultSize := searchResult.Hits.TotalHits
	series := make([]Serie, resultSize, 2*resultSize)
	fmt.Printf("Serie name: %s\n", resultSize)

	if searchResult.Hits.TotalHits > 0 {
		// Iterate through results
		for i, hit := range searchResult.Hits.Hits {
			
			// Deserialize hit.Source into a Serie (could also be just a map[string]interface{}).
			var serie Serie
			err := json.Unmarshal(*hit.Source, &serie)
			if err != nil {} // Deserialization failed
			series[i] = serie
			fmt.Printf("Serie name: %s: rating: %s\n", serie.Name, serie.IMDBRating)

		}
	}
	return series
}

// Search a in the series index for a serie with a specific rating, ordering the results by name.
func getSeriesByRate(rate string) []Serie{

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
		Index("series").
		Query(termQuery).
		Sort("name.keyword", true).
		From(0).Size(10).
		Pretty(true).
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}
	
	// get the number of results and create array
	resultSize := searchResult.Hits.TotalHits
	series := make([]Serie, resultSize, 2*resultSize)
	
	if searchResult.Hits.TotalHits > 0 {
		// Iterate through results
		for i, hit := range searchResult.Hits.Hits {
			
			// Deserialize hit.Source into a Serie (could also be just a map[string]interface{}).
			var serie Serie
			err := json.Unmarshal(*hit.Source, &serie)
			if err != nil {} // Deserialization failed
			series[i] = serie
		}
	}
	return series
}

func createSerie(serie Serie) Serie{

	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	
	if err != nil {
		// Handle error
		panic(err)
	}
	
	// Add a document to the index
	createdSerie := Serie{Name: serie.Name, IMDBRating: serie.IMDBRating}
	_, err = client.Index().
		Index("series").
		Type("serie").
		//Id("1").
		BodyJson(createdSerie).
		//Refresh(true).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	
	return createdSerie
}