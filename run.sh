rm -rf out
mkdir -p out
go build -o ./out/providers ./providers/...
go build -o ./out/driver ./cmd
./out/driver