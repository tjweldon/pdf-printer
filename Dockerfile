# syntax=docker/dockerfile:1
FROM openlabs/docker-wkhtmltopdf

# Copy go package
WORKDIR /
COPY './go1.18.linux-amd64.tar.gz' './golang.tar.gz'

# Install git
RUN apt-get update
RUN apt-get install -y git

# Install go and remove archive
RUN tar -C '/usr/local' -xzf 'golang.tar.gz' && \
    rm -f 'golag.tar.gz'

## Setup go env
RUN mkdir /app
RUN ln -s /usr/local/go/bin/go /bin/go

# copy source files to app dir
WORKDIR /app/src/tjweldon/pdf-printer
COPY . .
RUN go build && cp pdf-printer /


ENTRYPOINT /pdf-printer
