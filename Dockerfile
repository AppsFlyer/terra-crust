ARG GO_VERSION=1.16
FROM golang:${GO_VERSION} AS builder
COPY . /terra-crust
WORKDIR /terra-crust
RUN make compile


FROM golang:${GO_VERSION}
ARG ARCH=amd64
WORKDIR /opt/
COPY --from=builder /terra-crust/bin/terra-crust-linux-${ARCH} ./
RUN mv terra-crust-linux-${ARCH} terra-crust
ENTRYPOINT ["./terra-crust"]  
