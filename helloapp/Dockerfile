FROM golang:1-alpine AS build-env

ARG VERSION
ARG REVISION
ARG BRANCH

ENV REPO_PATH=$GOPATH/src/gitlab.messagebird.com/infrastructure-team/snippets

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux

ADD ./ $REPO_PATH

RUN ( \
        cd $REPO_PATH && \
        go build -mod=vendor && \
        mv ./helloapp / ; \
    )

FROM alpine:3.10

COPY --from=build-env /helloapp /usr/local/bin/helloapp

USER nobody:nogroup
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/helloapp"]

LABEL name=helloapp
LABEL version=$VERSION
LABEL maintainer="MessageBird B.V. Infrastructure Team <infrastructure-networking@messagebird.com>"
