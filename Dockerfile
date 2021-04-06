FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/sentinel-official/hub

# Add source files
COPY . .

# install sentinelhub, remove packages
RUN make build

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/sentinel-official/hub/bin/sentinelhubd /usr/bin/sentinelhubd
COPY --from=build-env /go/src/github.com/sentinel-official/hub/bin/sentinelhubcli /usr/bin/sentinelhubcli

EXPOSE 26656 26657 1317 9090

# Run sentinelhubd by default, omit entrypoint to ease using container with sentinelhubcli
CMD ["sentinelhubd"]