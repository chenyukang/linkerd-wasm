FROM rust:1.54.0

WORKDIR /app
COPY . .
RUN cargo install --path .
RUN cargo build 
CMD ["./target/debug/image-api-rs"]

