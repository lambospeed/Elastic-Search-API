# Golang integration with Elasticsearch 5.0
Rest Service API based in Golang which communicates with Elasticsearch 5.0.  
The communication between the Golang application and Elasticsearch was created using the Elastic, a Elasticsearch client for Go.   
Elastic can be found at https://github.com/olivere/elastic

# Must haves:  
1. Elasticsearch 5.0
	* Running at "http://localhost:9200"  
	* An index named "series"  
	* An type named "serie" inside the index "series"
2. Golang
	* Last version installed

# Rest API Information
1. Welcome index.
	* Type:  
			GET
	* URL:  
			http://localhost:8080
	* Result:  
			Welcome to the Rest API for Elastic Search!

2. Search for all series in the index.
	* Type:  
			GET 
	* URL:  
			http://localhost:8080/all_series
	* Result:  
			Searching for all series:  
			Results:  
			[{Game of Thrones 9.5} {Breaking Bad 9.5} {Test Serie 2 9} {Dexter 8.8} {Sons of Anarchy 8.6}  
			{The Walking Dead 8.5}  { }]

3. Search for a specific serie in the index using the name as filter.
	* Type:  
			GET 
	* URL:  
			http://localhost:8080/series_byname/{serieName}	
	* Parameter:  
			/Breaking Bad
	* Result:  
			Searching for series with name: Breaking Bad  
			Results:  
			[{Breaking Bad 9.5}]

4. Search for a specific serie in the index using the imdb rating as filter.
	* Type:  
			GET 
	* URL:  
			http://localhost:8080/series_byrate/{rate}	
	* Parameter:  
			/9.5
	* Result:  
			Searching for series with rate: 9.5  
			Results:  
			[{Breaking Bad 9.5} {Game of Thrones 9.5}]

5. Create a serie in the Series index.
	* Type:  
			POST
	* URL:  
			http://localhost:8080/create_serie
	* Parameter:  
			{  
				"name": "Breaking Bad",  
				"imdb_rating": "9.5"  
			}
	* Result:  
			Creating serie in the index: {Breaking Bad 9.5}  
			Results:  
			{Breaking Bad 9.5}

		
