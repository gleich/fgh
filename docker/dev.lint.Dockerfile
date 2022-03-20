FROM golangci/golangci-lint:v1.38

# Meta data
LABEL maintainer="email@mattglei.ch"
LABEL description="📁 Automate the organization of your cloned GitHub repositories"

# Copying over files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing hadolint:
WORKDIR /usr/bin
RUN curl -sL -o hadolint "https://github.com/hadolint/hadolint/releases/download/v2.7.0/hadolint-$(uname -s)-$(uname -m)" \
    && chmod 700 hadolint

# Installing go 1.18
RUN go install "golang.org/dl/go1.18@latest" \
    && go1.18 download \
    && mv "$(which go1.18)" "$(which go)"

# Installing goreleaser
WORKDIR /
RUN git clone https://github.com/goreleaser/goreleaser
WORKDIR /goreleaser
RUN go get ./... \
    && go build -o goreleaser . \
    && mv goreleaser /usr/bin

# Installing make
RUN apt-get update && apt-get install make=4.2.1-1.2 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

CMD ["make", "local-lint"]
