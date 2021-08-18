FROM ubuntu:20.04
RUN apt-get update 
RUN apt-get install -y wget gcc
RUN apt-get install -y libjpeg-dev libpng-dev
RUN wget https://dl.google.com/go/go1.16.4.linux-amd64.tar.gz
RUN tar -xvf go1.16.4.linux-amd64.tar.gz   
RUN mv go /usr/local  
RUN export GOROOT=/usr/local/go 
RUN export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 
RUN export GOPATH=/app/

RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge.sh
RUN bash ./install_wasmedge.sh /usr/local
RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge_tensorflow_deps.sh
RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge_tensorflow.sh
RUN bash ./install_wasmedge_tensorflow_deps.sh /usr/local
RUN bash ./install_wasmedge_tensorflow.sh /usr/local
RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge_image.sh
RUN bash ./install_wasmedge_image.sh /usr/local


WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN /usr/local/go/bin/go mod download
COPY *.go ./
RUN mkdir lib
COPY lib ./lib/
RUN /usr/local/go/bin/go build --tags "tensorflow image" -o image-api-go
EXPOSE 8086

CMD ["./image-api-go"]
