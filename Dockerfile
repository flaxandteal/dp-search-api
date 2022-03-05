FROM golang:1.17.2 as chef
WORKDIR /app

# Install dependencies - this is the caching Docker layer
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM chef as builder
COPY . .

RUN make build && mv build/$(go env GOOS)-$(go env GOARCH)/* ./build

FROM alpine
WORKDIR /app
COPY --from=builder /app/build/dp-search-api .

EXPOSE 8080

CMD [ "/dp-search-api" ]
