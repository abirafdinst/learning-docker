# Menggunakan gambar golang sebagai base image
FROM golang:1.20-alpine

# Set working directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependensi
RUN go mod init my-docker
RUN go mod tidy

# Build aplikasi
RUN go build -o my-docker

# Jalankan aplikasi
CMD ["./my-docker"]

# Expose port 8080
EXPOSE 8080