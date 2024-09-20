FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /app

COPY . .

RUN go test --cover -v ./...
RUN go build -v -o morsify

FROM alpine:latest

LABEL authors="SpauriRosso"\
      description="Morsify is a simple API translating text to morse code. Writing in go, it has the sole purpose of training in server and API setup."\
      licence="GNU GPL V3.0-or-later"\
      maintainer="See author label"\
      contact="hey@spyrisk.fr"

WORKDIR /app
COPY --from=build /app/morsify /app/morsify
COPY --from=build /app/src /app/src

EXPOSE 5827

CMD ["/app/morsify"]
