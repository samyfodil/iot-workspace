# Use an official development container base image
FROM mcr.microsoft.com/devcontainers/base:ubuntu-22.04

# Install necessary dependencies
RUN apt-get update && apt-get install -y \
    curl \
    wget \
    git \
    build-essential \
    python3 \
    python3-pip \
    python-is-python3 \
    && apt-get clean

ENV GO_VERSION=1.22.0
RUN curl -fsSL https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin

ENV NODE_VERSION=22.11.0
ENV NVM_DIR=/usr/local/nvm
RUN mkdir -p /usr/local/nvm \
    && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash \
    && export NVM_DIR="/usr/local/nvm" \
    && [ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh" \
    && nvm install ${NODE_VERSION} \
    && nvm alias default ${NODE_VERSION} \
    && nvm use default \
    && ln -s "$NVM_DIR/versions/node/v18.18.0/bin/node" /usr/bin/node \
    && ln -s "$NVM_DIR/versions/node/v18.18.0/bin/npm" /usr/bin/npm \
    && ln -s "$NVM_DIR/versions/node/v18.18.0/bin/npx" /usr/bin/npx
ENV PATH=$NVM_DIR/versions/node/v${NODE_VERSION}/bin:$PATH

RUN pip install --no-cache-dir virtualenv pylint black

RUN apt-get autoremove -y && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /workspace
