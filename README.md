# Rest api with openapi3.0 swagger + grpc client

spring boot with grpc client example



swagger ui

> http://localhost:8080/swagger-ui/index.html


test controller

> localhost:8080/api/v1/test


# env

grpc server addr

> ${TEST_ADDR} : static://0.0.0.0:50010

### if use dns name contains "_", please check this issue

it still making some problem

https://github.com/grpc-ecosystem/grpc-spring/issues/871

# Response

> curl --location 'http://localhost:8080/api/v1/test' \
--header 'Accept-Language: en-US'

> {
"status": 200,
"message": "processing completed",
"ts": "20240115120543",
"response": {
"Hello": "result {\n  result: \"hi\"\n}\n"
}
}


# Build with Docker

> Docker build -t {Tag} . 

or 

> docker-compose build