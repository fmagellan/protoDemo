# protoDemo
This repository will be used to practice protobuf.

1. Install protoc compiler.
I used windows desktop. So, I downloaded windows zip file, from the following location.
https://github.com/protocolbuffers/protobuf/releases/tag/v3.14.0
After unzipping it, I copied the protoc.exe to ~/go/bin/ directory.

2. Install go packages and plugins for protobuf.
Execute the following two comands.
go get github.com/golang/protobuf
go get github.com/golang/protobuf/proto

3. Write proto file.
1. Learn the basics of protobuf syntax and semantics.
2. Write the protofile, with a simple message structure.
3. Add "package demo1". The name of the package should be same as the directory, where the go-structures will be generated.
3. Add option package=github.com/fmagellan/protoDemo/demo1, to avoid some warnings, while generating go-structures.

4. Generate go-structures with protoc compiler.
Go to the directory where proto file is there, and execute the following command.
protoc -I=. --go_out=paths=source_relative:. person1.proto
go_out flag indicates protoc to generate go structures.
paths=source_relative:. indicates to generate the go-structures in the current directory, which I intended to use for go-package.

5. Initialise the go-module for our demo, which will setup the package structure properly, with the following command, in the protoDemo directory.
go mod init github.com/fmagellan/protoDemo

6. Write the sample program. It does the following
1. Fill one protoMessage structure.
2. Marshal/Encode the protoMessage, which will result into []byte.
3. Write it into a file, person1.dat
4. Read the file, person1.dat, into a []byte.
5. Unmarshal/Decode the []byte into protoMessage structure.
6. Print the protoMessage structure and verify your data is correct.

Et voila !
