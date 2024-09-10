package utils_test

import "github.com/loickreitmann/utils"

func cyclicDataStructure() utils.JSONResponse {
	node1 := &Node{Value: "first"}
	node2 := &Node{Value: "second", NextNode: node1}
	node1.NextNode = node2
	jsonResp := utils.JSONResponse{
		Data: node1,
	}
	return jsonResp
}
