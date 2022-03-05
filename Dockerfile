FROM golang:1.17-alpine as chef
WORKDIR /app

RUN apk update && apk add build-base bash git

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
ADD templates /app/templates

EXPOSE 23900

CMD [ "./dp-search-api" ]
