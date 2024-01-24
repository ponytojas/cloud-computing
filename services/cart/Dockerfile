FROM golang:1.21.6-alpine3.19 AS build
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main ./cmd/main.go

FROM scratch
COPY --from=build /app/main /main/
CMD ["/main/main"]
