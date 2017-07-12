FROM golang:1.8.3 as builder
WORKDIR /go/src/hello-world
COPY main.go ./
RUN go install
RUN ldd /go/bin/hello-world | grep -q "not a dynamic executable"

FROM scratch
COPY --from=builder /go/bin/hello-world /hello-world
CMD ["/hello-world"]
