FROM golang:1.18

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="📁 Automate the organization of your cloned GitHub repositories"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["make", "local-test"]
