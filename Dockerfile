FROM golang:1.11.1 as builder
WORKDIR /
COPY .    /
# Build
RUN make deps && make build

FROM debian:9.5-slim
WORKDIR /root/
COPY --from=builder /bin/ .
EXPOSE 9999
ENTRYPOINT ["skipper -address :9999 -inline-routes 'r: * -> setQuery("lang", "pt") -> "http://10.15.0.80"'"]