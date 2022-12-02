FROM golang:alpine as build
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates
WORKDIR /builder
COPY . /builder
RUN go mod download
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /minimalback server.go

FROM scratch as final
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /minimalback /minimalback
EXPOSE 6666
ENTRYPOINT ["/minimalback"]