# URL-shorter
> This is a link shortener written in golang, using the standard net/http package and mySQL

## How it works
- POST
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"src":"_link_"}' http://localhost:8080/urls 
  ```
  
  > _request_
  ```json
   {
    "src": "https://www.youtube.com/watch?v=DdI598gKkKw&list=RDDdI598gKkKw&start_radio=1"
   }
  ```
 
  > _response_
  ```json
   {
     "src": "https://www.youtube.com/watch?v=DdI598gKkKw\u0026list=RDDdI598gKkKw\u0026start_radio=1",
      "res": "c72b1d8"
    }
  ```

  ***

- GET
   ```bash
   curl -X GET http://localhost:8080/_short_
   ```

   > _request_

   ```bash
   curl -X GET http://localhost:8080/c72b1d8
   ```

   > _response_

  ```html
  <a href="https://www.youtube.com/watch?v=DdI598gKkKw&amp;list=RDDdI598gKkKw&amp;start_radio=1">Found</a>.
  ```

***

## Shortener function
  ```golang
  const shortLength = 7

func ShortenUrl(src string) string {
	hash := md5.Sum([]byte(src))
	hexHash := hex.EncodeToString(hash[:])

	res := hexHash[:shortLength]
	return res
}
  ```
  
***

## Database
| short_url | long_url                                                                     |
|-----------|------------------------------------------------------------------------------|
| bcf62df   | https://github.com/IvanDrf/url-shorter/main                                  |
| c72b1d8   | https://www.youtube.com/watch?v=DdI598gKkKw&list=RDDdI598gKkKw&start_radio=1 |
 

***

## Used technologies
- Golang 
- MySql
