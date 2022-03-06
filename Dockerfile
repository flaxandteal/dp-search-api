FROM golang:1.17.2 as build

ADD . /dp-search-api
WORKDIR /dp-search-api

RUN go get -d

RUN make build && mv build/$(go env GOOS)-$(go env GOARCH)/* ./build

FROM golang:1.17.2

COPY --from=build /dp-search-api/build/dp-search-api /dp-search-api

RUN mkdir /app

WORKDIR /app

ADD templates /app/templates

EXPOSE 8080

CMD [ "/dp-search-api" ]
