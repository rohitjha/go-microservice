FROM golang:alpine
WORKDIR /build
COPY . .
RUN unset GOPATH && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go && \
    chmod 111 main

FROM scratch
WORKDIR /app
COPY --from=0 /build/main /app/main
COPY --from=0 /build/app.env /app/app.env
COPY --from=0 /build/certs /app/certs
ADD passwd.minimal /etc/passwd
USER nobody
CMD ["/app/main"]
