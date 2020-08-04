package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/tendermint/classic/abci/types"
	cmn "github.com/tendermint/classic/libs/common"
	"github.com/tendermint/go-amino-x"
)

func main() {

	conn, err := cmn.Connect("unix://test.sock")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Make a bunch of requests
	counter := 0
	for i := 0; ; i++ {
		req := types.RequestEcho{Message: "foobar"}
		_, err := makeRequest(conn, req)
		if err != nil {
			log.Fatal(err.Error())
		}
		counter++
		if counter%1000 == 0 {
			fmt.Println(counter)
		}
	}
}

func makeRequest(conn net.Conn, req types.Request) (types.Response, error) {
	var bufWriter = bufio.NewWriter(conn)

	// Write desired request
	err := amino.MarshalLengthPrefixedWriter(&req, bufWriter)
	if err != nil {
		return nil, err
	}
	err = amino.MarshalLengthPrefixedWriter(&types.RequestFlush{}, bufWriter)
	if err != nil {
		return nil, err
	}
	err = bufWriter.Flush()
	if err != nil {
		return nil, err
	}

	// Read desired response
	var res types.Response
	err = amino.UnmarshalLengthPrefixedReader(conn, &res)
	if err != nil {
		return nil, err
	}
	var resFlush types.Response
	err = amino.UnmarshalLengthPrefixedReader(conn, &res)
	if err != nil {
		return nil, err
	}
	if _, ok := resFlush.(types.ResponseFlush); !ok {
		return nil, fmt.Errorf("Expected flush response but got something else: %v", reflect.TypeOf(resFlush))
	}

	return res, nil
}
