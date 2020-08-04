package example

import (
	"fmt"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"golang.org/x/net/context"

	cmn "github.com/tendermint/classic/libs/common"
	"github.com/tendermint/classic/libs/log"

	abcicli "github.com/tendermint/classic/abci/client"
	"github.com/tendermint/classic/abci/example/errors"
	"github.com/tendermint/classic/abci/example/kvstore"
	abciserver "github.com/tendermint/classic/abci/server"
	"github.com/tendermint/classic/abci/types"
)

func TestKVStore(t *testing.T) {
	fmt.Println("### Testing KVStore")
	testStream(t, kvstore.NewKVStoreApplication())
}

func TestBaseApp(t *testing.T) {
	fmt.Println("### Testing BaseApp")
	testStream(t, types.NewBaseApplication())
}

func TestGRPC(t *testing.T) {
	fmt.Println("### Testing GRPC")
	testGRPCSync(t, types.NewGRPCApplication(types.NewBaseApplication()))
}

func testStream(t *testing.T, app types.Application) {
	numDeliverTxs := 20000

	// Start the listener
	server := abciserver.NewSocketServer("unix://test.sock", app)
	server.SetLogger(log.TestingLogger().With("module", "abci-server"))
	if err := server.Start(); err != nil {
		require.NoError(t, err, "Error starting socket server")
	}
	defer server.Stop()

	// Connect to the socket
	client := abcicli.NewSocketClient("unix://test.sock", false)
	client.SetLogger(log.TestingLogger().With("module", "abci-client"))
	if err := client.Start(); err != nil {
		t.Fatalf("Error starting socket client: %v", err.Error())
	}
	defer client.Stop()

	done := make(chan struct{})
	counter := 0
	client.SetResponseCallback(func(req *types.Request, res *types.Response) {
		// Process response
		switch r := res.Value.(type) {
		case *types.Response_DeliverTx:
			counter++
			if r.DeliverTx.Error != nil {
				t.Error("DeliverTx failed with error", r.DeliverTx.Error)
			}
			if counter > numDeliverTxs {
				t.Fatalf("Too many DeliverTx responses. Got %d, expected %d", counter, numDeliverTxs)
			}
			if counter == numDeliverTxs {
				go func() {
					time.Sleep(time.Second * 1) // Wait for a bit to allow counter overflow
					close(done)
				}()
				return
			}
		case *types.Response_Flush:
			// ignore
		default:
			t.Error("Unexpected response type", reflect.TypeOf(res.Value))
		}
	})

	// Write requests
	for counter := 0; counter < numDeliverTxs; counter++ {
		// Send request
		reqRes := client.DeliverTxAsync(types.RequestDeliverTx{Tx: []byte("test")})
		_ = reqRes
		// check err ?

		// Sometimes send flush messages
		if counter%123 == 0 {
			client.FlushAsync()
			// check err ?
		}
	}

	// Send final flush message
	client.FlushAsync()

	<-done
}
