FROM mariadb:10.5.9-focal

copy gogin.exe /app/gogin.exe
RUN go get github.com/joho/godotenv
cmd ["./app/gogin.exe"]