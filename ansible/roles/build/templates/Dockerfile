#Docker
#docker build -t validator:latest .

FROM golang:alpine AS build-env

# Install minimum necessary dependencies,
ENV PACKAGES pass curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /build

# Add source files
COPY . .

# install mbcorecr, remove packages
RUN make

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /build/mbcorecrd /usr/bin/mbcorecrd
EXPOSE 26656 26657 1317 9090
