# ./dockerfile
# 使用官方的 Golang 映像作為基礎映像
FROM golang:1.22.4

# 設置工作目錄
WORKDIR /app

# 將當前目錄中的所有文件複製到工作目錄中
COPY . .

# 下載所需套件
RUN go mod download

# 執行編譯好的二進制文件
CMD ["go", "run", "main.go"]

EXPOSE 8080
