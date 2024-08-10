# Stage 1: Build
FROM rust:1.70 as builder

# Set the working directory inside the container
WORKDIR /usr/src/ethernet

# Copy Cargo.toml and Cargo.lock into the container
COPY Cargo.toml Cargo.lock ./
# Copy the source code into the container
COPY src ./src
# Build the project
RUN cargo build --release

# Stage 2: Final image
FROM rockylinux:8

# Set the working directory inside the container
WORKDIR /usr/local/bin
# Copy the built binary from the builder stage
COPY --from=builder /usr/src/ethernet/target/release/ethernet .

# Make the binary executable
RUN chmod +x ./ethernet

# Command to run the application
CMD ["./ethernet"]
