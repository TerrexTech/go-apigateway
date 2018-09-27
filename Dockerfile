# ===> Build Image
FROM golang:1.11.0-alpine3.8 AS builder
LABEL maintainer="Jaskaranbir Dhillon"

ARG SOURCE_REPO

ENV DEP_VERSION=0.5.0 \
    CGO_ENABLED=0 \
    GOOS=linux

# Download and install dep and git
ADD https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN apk add --update git

WORKDIR $GOPATH/src/github.com/TerrexTech/${SOURCE_REPO}

# Copy the code from the host and compile it
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only -v
COPY . ./

RUN go build -v -a -installsuffix nocgo -o /app ./main

# ===> Run Image
FROM scratch
LABEL maintainer="Jaskaranbir Dhillon"

COPY --from=builder /app ./
ENTRYPOINT ["./app"]
