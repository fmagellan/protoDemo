package main

import (
    "fmt"
    "io/ioutil"
    "github.com/golang/protobuf/proto"
    "github.com/fmagellan/protoDemo/demo1"
)

func main() {
    pWriteMessage := &demo1.Person{
        FirstName: "Ferdinand",
        Id: 1519,
        LastName: "Magellan",
    }

    writeBytes, err := proto.Marshal(pWriteMessage)
    if err != nil {
        fmt.Println("Error during marshalling:", err.Error())
    }

    err = ioutil.WriteFile("person1.dat", writeBytes, 0644)
    if err != nil {
        fmt.Println("Error during writing to the file:", err.Error())
    }

    var readBytes []byte
    readBytes, err = ioutil.ReadFile("person1.dat")
    if err != nil {
        fmt.Println("Error during reading from the file:", err.Error())
    }

    var readMessage *demo1.Person = new(demo1.Person)
    err = proto.Unmarshal(readBytes, readMessage)
    if err != nil {
        fmt.Println("Error during Unmarshalling:", err.Error())
    }

    fmt.Println("After unmarshalling:")
    fmt.Println("First Name :", readMessage.FirstName)
    fmt.Println("Last Name  :", readMessage.LastName)
    fmt.Println("Id         :", readMessage.Id)
}
