FROM ubuntu:20.10

RUN apt-get update
RUN apt-get install -y libvips-dev wget tar

RUN wget -c https://golang.org/dl/go1.16.7.linux-amd64.tar.gz -O - | tar -xz -C /usr/local
ENV PATH="/usr/local/go/bin:${PATH}"
RUN go version

WORKDIR /app