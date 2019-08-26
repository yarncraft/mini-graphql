FROM golang:1.13rc1-buster as builder

RUN apt-get update && apt-get install -y \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

# install UPX for executable compression
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local
RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

WORKDIR /root

COPY main main

# compress the binary
RUN strip --strip-unneeded main
RUN upx main



FROM scratch
# ADD ca-certificates.crt /etc/ssl/certs/
COPY --from=builder root/main .
CMD ["/main"]