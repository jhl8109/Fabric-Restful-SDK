package sdkInit

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *Application) Set(args []string) (string, error) {
	var tempArgs [][]byte
	for i := 1; i < len(args); i++ {
		tempArgs = append(tempArgs, []byte(args[i]))
	}

	request := channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1]), []byte(args[2])}}
	response, err := t.SdkEnvInfo.ChClient.Execute(request)
	if err != nil {
		return "", err
	}

	//fmt.Println("============== response:",response)

	return string(response.TransactionID), nil
}
