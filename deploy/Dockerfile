ARG MISSKEY_GO_VERSION=master

####################################################################################################
## Builder
####################################################################################################
FROM golang:1-alpine AS builder

ARG MISSKEY_GO_VERSION

ENV GO111MODULE=on \
    UPSTREAM=http://localhost:3000

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    git \
    make \
    postgresql-client \
    postgresql \
    tar

WORKDIR /misskey

RUN git clone --recursive --branch ${MISSKEY_GO_VERSION} https://random.chars.jp/git/misskey /tmp/misskey \ 
    && cp -r /tmp/misskey/. /misskey

# Build misskey
RUN make boil && make && make assets-download && make

####################################################################################################
## Final image
####################################################################################################
FROM alpine:3.15

ARG MISSKEY_GO_VERSION

ENV UPSTREAM=http://localhost:3000

RUN apk add --no-cache \
    ca-certificates \
    tini \
    shadow \
    postgresql-client \
    tzdata

WORKDIR /misskey

COPY --from=builder /misskey /misskey

# Add an unprivileged user and set directory permissions
RUN adduser --disabled-password --gecos "" --no-create-home misskey \
    && chown -R misskey:misskey /misskey

# Make persistent data directory
RUN mkdir -p /misskey/ \
    && chown -R misskey:misskey /misskey/

ENTRYPOINT ["/sbin/tini", "--"]

USER misskey

CMD ["./misskey"]

EXPOSE 3000

STOPSIGNAL SIGTERM

HEALTHCHECK \
    --start-period=15s \
    --interval=1m \
    --timeout=5s \
    CMD wget --spider --q http://localhost:3000/api/server-info || exit 1

# Image metadata
LABEL org.opencontainers.image.version=${MISSKEY_GO_VERSION}
LABEL org.opencontainers.image.title=Misskey
LABEL org.opencontainers.image.description="Alternative backend for Misskey written in Go."
LABEL org.opencontainers.image.url=https://random.chars.jp/git/misskey
LABEL org.opencontainers.image.licenses=GPL-3.0-or-later
LABEL org.opencontainers.image.source="https://random.chars.jp/git/misskey"

