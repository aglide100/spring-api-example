# Rest api with openapi3.0 swagger

> http://localhost:8080/swagger-ui/index.html


test controller

> localhost:8080/api/v1/test


# Response

> curl --location 'http://localhost:8080/api/v1/test' \
--header 'Accept-Language: en-US'

> {
"status": 200,
"message": "processing completed",
"ts": "20240114135908",
"response": {
"Hello": "world"
}
}


# Build with Docker

> Docker build -t {Tag} . 

or 

> docker-compose build