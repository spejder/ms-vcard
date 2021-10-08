FROM golang:1.17.2-alpine AS build-env

WORKDIR /build/github.com/spejder/ms-vcard

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add git=~2

COPY *.go go.mod go.sum /build/github.com/spejder/ms-vcard/
COPY internal/odoo/ /build/github.com/spejder/ms-vcard/internal/odoo/

RUN go build

FROM scratch

EXPOSE 80

ENV PATH=/

COPY --from=build-env /build/github.com/spejder/ms-vcard/ms-vcard /ms-vcard
COPY --from=build-env /etc/ssl/certs/ /etc/ssl/certs/

ENTRYPOINT ["/ms-vcard"]
