package main

import (
    "fmt"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/anypb"
    "github.com/arasevic/savetheworldwithgo/12_protocolbuffers/pb/example_05/user"
)

func main() {
    info := anypb.Any{Value: []byte(`John rules`), TypeUrl: "urltype"}
    userA := user.User{UserId: "John", Email: "john@gmail.com", Info: []*anypb.Any{&info}}
    anypb.a
    fmt.Println("To encode:", userA.String())

    encoded, err := proto.Marshal(&userA)
    if err != nil {
        panic(err)
    }
    recovered := user.User{}
    err = proto.Unmarshal(encoded, &recovered)
    fmt.Println("Recovered:", recovered.String())
}


