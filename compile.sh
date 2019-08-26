echo "🔥 Compiling to binary..."
cd src/cmd/server
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main
cp main ../../../
rm main