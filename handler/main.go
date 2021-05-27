package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	Models "goeth-api/models"
	Modules "goeth-api/modules"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get parameter from url request
	vars := mux.Vars(r)
	module := vars["module"]

	// Get the query parameters from url request
	address := r.URL.Query().Get("address")
	hash := r.URL.Query().Get("hash")
	blockNum := r.URL.Query().Get("blockNum")
	limit := r.URL.Query().Get("limit")

	// Set our response header
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter:
	switch module {
	case "latest-block":
		_block := Modules.GetLatestBlock(*client.Client)
		json.NewEncoder(w).Encode(_block)

	case "get-block-by-number":
		if blockNum == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Please input block number!",
			})
			return
		}
		_block := Modules.GetBlockByNumber(*client.Client, blockNum)
		if _block != nil {
			json.NewEncoder(w).Encode(_block)
			return
		}
		json.NewEncoder(w).Encode(&Models.Error{
			Code:    404,
			Message: "Block Not Found!",
		})

	case "get-recent-block":
		if limit == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Please input block limit!",
			})
			return
		}
		_block := Modules.GetRecentBlock(*client.Client, limit)
		if _block == "ok" {
			return
		}

	case "get-tx":
		if hash == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		txHash := common.HexToHash(hash)
		_tx := Modules.GetTxByHash(*client.Client, txHash)

		if _tx != nil {
			json.NewEncoder(w).Encode(_tx)
			return
		}

		json.NewEncoder(w).Encode(&Models.Error{
			Code:    404,
			Message: "Tx Not Found!",
		})

	case "send-eth":
		decoder := json.NewDecoder(r.Body)
		var t Models.TransferEthRequest

		err := decoder.Decode(&t)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_hash, err := Modules.TransferEth(*client.Client, t.PrivKey, t.To, t.Amount)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&Models.HashResponse{
			Hash: _hash,
		})

	case "get-balance":
		if address == "" {
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}

		balance, err := Modules.GetAddressBalance(*client.Client, address)

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    500,
				Message: "Internal server error",
			})
			return
		}

		json.NewEncoder(w).Encode(&Models.BalanceResponse{
			Address: address,
			Balance: balance,
			Symbol:  "Ether",
			Units:   "Wei",
		})

	}

}
