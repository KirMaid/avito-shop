#FROM golang:1.22
#
#WORKDIR ${GOPATH}/avito-shop/
#COPY . ${GOPATH}/avito-shop/
#
#RUN go build -o /build ./cmd/app \
#    && go clean -cache -modcache
#
#EXPOSE 8080
#
#CMD ["/build"]


# Step 1: Modules caching
FROM golang:1.24 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.24 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

# Step 3: Final
FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /app/migrations /migrations
COPY --from=builder /bin/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app"]