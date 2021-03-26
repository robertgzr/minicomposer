# syntax=docker/dockerfile:1.2-labs
FROM docker:dind

RUN apk update && apk add git go make
RUN --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go/pkg \
    set -ex; \
    git clone https://github.com/compose-spec/compose-ref; \
    cd compose-ref || false; \
    make build; \
    install bin/compose-ref /usr/local/bin/;

COPY ./entry.sh /usr/local/bin/
WORKDIR /src
ENTRYPOINT [ "entry.sh" ]
