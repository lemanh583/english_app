# Sử dụng image chính thức của Golang
FROM golang:1.22-alpine

# Tạo thư mục app bên trong container
WORKDIR /app

# Copy mã nguồn vào container
COPY . .

# Biên dịch ứng dụng
RUN go build -o server

# Expose cổng 8080
EXPOSE 8080

# Lệnh chạy server
CMD ["./server"]
