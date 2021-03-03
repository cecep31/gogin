FROM mariadb:10.5.9-focal

copy gogin.exe /app/gogin.exe

cmd ["./app/gogin.exe"]