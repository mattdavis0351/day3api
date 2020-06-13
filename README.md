# Day 3 API

This API contains three distinct routes that serve different types of responses.

`/api/v1/site`
`/api/v1/byte-json`
`/api/v1/struct-life`

## Usage

### Go

You can clone the source code and execute `go build` or `go run` to run this API locally on a server

### Docker

You can either use `docker build` to build the `Dockerfile` in this repo

OR

You can run `docker pull containers.pkg.github.com/mattdavis0351/day3api:v1.0` to download this image. Once downloaded run `docker run -d -p 80:8080 containers.pkg.github.com/mattdavis0351/day3api:v1.0` to make the container available at `localhost/<route>` for example `localhost/api/v1/site`

OR

You can include this as a base image in your own Docker file (which would be pointless... don't do that!)
