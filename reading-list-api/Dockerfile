FROM golang:1.19-alpine as build

WORKDIR /src
COPY go.sum /src/
COPY go.mod /src/
COPY main.go /src/
RUN go mod download && go build -o app


FROM alpine:3.12

RUN mkdir -p /app
COPY --from=build /src/app /app/app
EXPOSE 8080

CMD ["./app/app"]
