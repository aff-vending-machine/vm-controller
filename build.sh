#download dependencies
go mod vendor

# build amd64
docker build --platform linux/amd64 -f jetts.dockerfile -t vm-controller:0.0.0 .

#build arm64
docker build --platform linux/arm64 --build-arg GOARCH=arm64 -f jetts.dockerfile -t vm-controller:0.0.0-arm64 .

# chown -R 1000:1000 /db

# test
docker run -it vm-controller:0.0.0-arm64