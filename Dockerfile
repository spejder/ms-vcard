FROM docker.io/golang:1.22.5-alpine AS build-env

WORKDIR /build/github.com/spejder/ms-vcard

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add git=~2

COPY *.go go.mod go.sum /build/github.com/spejder/ms-vcard/
COPY internal/ms/ /build/github.com/spejder/ms-vcard/internal/ms/

RUN go build -tags docker

FROM scratch

EXPOSE 80

ENV PATH=/

COPY --from=build-env /build/github.com/spejder/ms-vcard/ms-vcard /ms-vcard
COPY --from=build-env /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/ms-vcard"]
