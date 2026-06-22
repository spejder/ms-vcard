FROM scratch

EXPOSE 80

RUN echo "nonroot:x:65532:65532::/:/sbin/nologin" > /etc/passwd && \
    echo "nonroot:x:65532:" > /etc/group

ARG TARGETPLATFORM
COPY $TARGETPLATFORM/ms-vcard /ms-vcard

ENTRYPOINT ["/ms-vcard"]
