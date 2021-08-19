FROM rust:1.54.0

RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge.sh
RUN bash ./install_wasmedge.sh /usr/local
RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge_tensorflow_deps.sh
RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge_tensorflow.sh
RUN bash ./install_wasmedge_tensorflow_deps.sh /usr/local
RUN bash ./install_wasmedge_tensorflow.sh /usr/local
RUN wget https://github.com/second-state/WasmEdge-go/releases/download/v0.8.1/install_wasmedge_image.sh
RUN bash ./install_wasmedge_image.sh /usr/local

WORKDIR /app
COPY Cargo.* ./
RUN mkdir -p ./src
COPY src ./src
RUN cargo install --path .
RUN cargo build 
RUN mkdir -p ./lib 
COPY lib ./lib

CMD ["./target/debug/image-api-rs"]

