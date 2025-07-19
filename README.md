# URL-shorter

## How it works
- POST
  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{"src":"_link_"}' http://localhost:8080/urls 
  ```
  Example
 
  _request_
  ```json
   {
    "src": "https://www.youtube.com/watch?v=DdI598gKkKw&list=RDDdI598gKkKw&start_radio=1"
   }
  ```
 
  _response_
  ```json
   {
     "src": "https://www.youtube.com/watch?v=DdI598gKkKw\u0026list=RDDdI598gKkKw\u0026start_radio=1",
      "res": "c72b1d8"
    }
```


- GET
  ```bash
  curl -X GET http://localhost:8080/_short_
  ```

 Example

 _request_

 ```bash
 curl -X GET http://localhost:8080/c72b1d8
 ```

 _response_

```html
<a href="https://www.youtube.com/watch?v=DdI598gKkKw&amp;list=RDDdI598gKkKw&amp;start_radio=1">Found</a>.
```

## Used technologies
- Golang 
- MySql
