# syntax=docker/dockerfile:1
FROM openlabs/docker-wkhtmltopdf
RUN apt-get update
RUN apt-get install -y git

RUN wget 'https://go.dev/dl/go1.18.linux-amd64.tar.gz'
RUN tar -C '/usr/local' -xzf 'go1.18.linux-amd64.tar.gz'
RUN rm -f 'go1.18.linux-amd64.tar.gz'
RUN export PATH="$PATH:/usr/local/go/bin"
RUN export GOPATH="/app"
RUN export PATH="$PATH:$GOPATH/bin"

RUN mkdir -p "/app/src/tjweldon/pdf-printer"
WORKDIR "/app/src/tjweldon/pdf-printer"
COPY . .
RUN ls -alh
RUN /usr/local/go/bin/go build
RUN mv "./data/tick.sh" /tick.sh

ENTRYPOINT /app/src/tjweldon/pdf-printer/pdf-printer
