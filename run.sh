cd providers || return
go build -o ./../out/providers
cd ..
go build -o ./out/driver ./cmd
./out/driver