build:
	powershell -Command "$$env:GOOS = 'linux'; $$env:GOARCH = 'amd64'; go build -o bootstrap"
	powershell -Command "Compress-Archive -Path bootstrap -DestinationPath function.zip -Force"