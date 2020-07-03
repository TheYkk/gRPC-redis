# Grpc to Redis
Client python 3 ile yazilmistir
Server ise golang ile
## Protobuf
Ilk once protoyu client ve server icin kaynak koda cevirmek lazim bunun icin golang ve python icin  ayri ayri araclar var.

Ilk once python icin gereken araclar `grpcio` ve `grpcio-tools`, kurmak icin ise 
```
pip3 install -r requirements.txt
```

Golang icin gereken araclar `protobuf-compiler` ve `protoc-gen-go`, kurmak icin ise
```
sudo apt install protobuf-compiler
```

```
export GO111MODULE=on
go get github.com/golang/protobuf/protoc-gen-go
export PATH="$PATH:$(go env GOPATH)/bin"
```

Gerekli toolari kurduktan sonra protolari uretebiliriz
```
./proto_gen.sh
```

## Docker compose
Uygulama tamamen docker uyumludur protolari urettikten sonra tek yapmak gereken docker compose ile ayaga kaldirmak
```
docker-compose up --build
```