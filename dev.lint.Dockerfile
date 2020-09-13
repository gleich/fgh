FROM golangci/golangci-lint:v1.30.0

# Meta data
LABEL maintainer="project_author_email"
LABEL description="project_description"

# Copying over files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing hadolint:
WORKDIR /usr/bin
RUN curl -sL -o hadolint "https://github.com/hadolint/hadolint/releases/download/v1.17.6/hadolint-$(uname -s)-$(uname -m)"
RUN chmod 700 hadolint

# Installing goreleaser
WORKDIR /
RUN git clone https://github.com/goreleaser/goreleaser
WORKDIR /goreleaser
RUN go get ./...
RUN go build -o goreleaser .
RUN mv goreleaser /usr/bin

# Installing make
RUN apt-get update && apt-get install make=4.2.1-1.2 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

CMD ["make", "local-lint"]
