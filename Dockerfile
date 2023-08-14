FROM golang:1.16 as build

WORKDIR /tmp/blowfish/
ADD . /tmp/blowfish/
RUN go build -o /tmp/bin/blowfish

FROM scratch
COPY --from=build /tmp/bin/blowfish /
CMD ["/blowfish"]
