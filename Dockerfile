FROM golang
WORKDIR /app
COPY . .
RUN unset GOPATH && \
    go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM scratch
COPY --from=0 /app/main /main
COPY --from=0 /app/certs /certs
CMD ["/main"]
