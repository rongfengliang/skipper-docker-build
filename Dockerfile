FROM golang:1.11.1 as builder
WORKDIR /
COPY .    /
# Build
RUN make build

# FROM debian:9.5-slim
# add luarocks for lua module install
FROM abaez/luarocks:lua5.1
WORKDIR /root
COPY --from=builder /bin/ .
ENV PATH=$PATH:/root
EXPOSE 9999
EXPOSE 9911
EXPOSE 9090