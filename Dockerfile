#Docker
#docker build -t mbcore .

FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /build

# Add source files
COPY . .

# install mbcore, remove packages
RUN make

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /build/mbcorecrd /usr/bin/mbcorecrd
COPY --from=build-env /build/ops .
EXPOSE 26656 26657 1317 9090

# Run mbcore by default
CMD ["/bin/sh", "/init.sh"]
