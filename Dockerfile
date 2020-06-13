FROM golang:latest As build
RUN mkdir -p /go/src/github.com/mattdavis0351/day3api/ && go get github.com/gorilla/mux
COPY . /go/src/github.com/mattdavis0351/day3api/
WORKDIR /go/src/github.com/mattdavis0351/day3api/cmd/
# EXPOSE 8080
RUN CGO_ENABLED=0 GOOS=linux go build -o app . 
# WORKDIR /go/src/github.com/mattdavis0351/day3api/
# CMD [ "./day3api" ]

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/mattdavis0351/day3api/cmd/app .
COPY --from=build /go/src/github.com/mattdavis0351/day3api/assets/ assets/
CMD ["./app"]  