set GOARCH=amd64
set GOOS=linux

go build -o chalkboard

docker build -t mailtruck/chalkboard:latest .
docker push mailtruck/chalkboard:latest