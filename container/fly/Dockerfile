FROM ubuntu:20.04 AS builder
WORKDIR /tmp

SHELL ["/bin/bash", "-o", "pipefail", "-c"]


RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        locales \
        tzdata \
        openssl \
        netbase \
        apt-utils \
        apt-transport-https \
        libreadline8 \
        ca-certificates \
        curl \
    && rm -rf /var/lib/apt/lists/*

RUN locale-gen en_US.UTF-8 \
    && update-locale LANG=en_US.UTF-8 \
    && echo "Etc/UTC" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

ENV FLY_VERSION=6.7.5
RUN curl https://github.com/concourse/concourse/releases/download/v${FLY_VERSION}/fly-${FLY_VERSION}-linux-amd64.tgz | \
    tar zxf - -C /usr/local/bin --strip-components 1
    
