## Build
FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /gerador-qrcode .

FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=build /gerador-qrcode /gerador-qrcode

EXPOSE 8080

ENTRYPOINT ["/gerador-qrcode"]