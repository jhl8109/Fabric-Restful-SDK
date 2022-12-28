package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mb "github.com/hyperledger/fabric-protos-go/msp"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/status"
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	contextImpl "github.com/hyperledger/fabric-sdk-go/pkg/context"
	lcpackager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/lifecycle"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/policydsl"
	"net/http"
	"os"
	"restfulsdk/sdkInit"
	"strings"
)

const (
	cc_name    = "simplecc"
	cc_version = "1.0.0"
)

var (
	GOPATH      = os.Getenv("GOPATH")
	NetworkOrgs = []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: fmt.Sprintf("%s/src/restfulsdk/fixtures/channel-artifacts/Org1MSPanchors.tx", GOPATH),
		},
	}
	NetworkInfo = sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    fmt.Sprintf("%s/src/restfulsdk/fixtures/channel-artifacts/channel.tx", GOPATH),
		Orgs:             NetworkOrgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    fmt.Sprintf("%s/src/restfulsdk/chaincode/", GOPATH),
		ChaincodeVersion: cc_version,
	}
	packageID = ""
	label     = ""
	ccPkg     = []byte("string")
	sdk       = new(fabsdk.FabricSDK)
	App       = sdkInit.Application{}
)

type channelInfo struct {
	ChannelName string `json:"channel_name`
}
type packageInfo struct {
	ChaincodeId      string `json:"chaincode_id"`
	ChaincodeVersion string `json:"chaincode_version"`
}
type installInfo struct {
	ChaincodeId      string `json:"chaincode_id"`
	ChaincodeVersion string `json:"chaincode_version"`
}
type approveInfo struct {
	ChaincodeID      string `json:"chaincode_id"`
	ChaincodeVersion string `json:"chaincode_version"`
	Sequence         int64  `json:"sequence"`
	InitRequired     bool   `json:"init_required"`
}
type commitInfo struct {
	ChaincodeID      string `json:"chaincode_id"`
	ChaincodeVersion string `json:"chaincode_version"`
	Sequence         int64  `json:"sequence"`
	InitRequired     bool   `json:"init_required"`
}
type initInfo struct {
	ChannelName  string `json:"channel_name"`
	ChaincodeID  string `json:"chaincode_id"`
	InitRequired bool   `json:"init_required"`
}

func CreateChannel(c *gin.Context) {
	var requestBody channelInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
	}
	//NetworkInfo.ChannelID = requestBody.ChannelName
	sdk, _ = sdkInit.InitOrgContext("config.yaml", &NetworkInfo)
	if len(NetworkInfo.Orgs) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Org don't exist."})
		return
	}
	signIds := []msp.SigningIdentity{}
	for _, org := range NetworkInfo.Orgs {
		// Get signing identity that is user to sign create channel request
		orgSignId, err := org.OrgMspClient.GetSigningIdentity(org.OrgAdminUser)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		}
		signIds = append(signIds, orgSignId)
	}
	if err := sdkInit.CreateChannel(signIds, &NetworkInfo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Create channel successfully."})
}
func JoinChannel(c *gin.Context) {
	var requestBody channelInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
	}
	//NetworkInfo.ChannelID = requestBody.ChannelName
	for _, org := range NetworkInfo.Orgs {
		if err := org.OrgResMgmt.JoinChannel(NetworkInfo.ChannelID, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com")); err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		}
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Join channel successfully."})
}
func PackageCC(c *gin.Context) {
	var requestBody packageInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}
	NetworkInfo.ChaincodeID = requestBody.ChaincodeId
	NetworkInfo.ChaincodeVersion = requestBody.ChaincodeVersion
	if len(NetworkInfo.Orgs) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Org don't exist."})
		return
	}
	label = NetworkInfo.ChaincodeID + "_" + NetworkInfo.ChaincodeVersion
	desc := &lcpackager.Descriptor{
		Path:  NetworkInfo.ChaincodePath,
		Type:  pb.ChaincodeSpec_GOLANG,
		Label: label,
	}
	ccPkg2, err := lcpackager.NewCCPackage(desc)
	ccPkg = ccPkg2
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
	packageID = lcpackager.ComputePackageID(label, ccPkg2)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "package successfully : " + packageID})
}
func InstallCC(c *gin.Context) {
	var requestBody installInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}
	label = requestBody.ChaincodeId + "_" + requestBody.ChaincodeVersion
	installCCReq := resmgmt.LifecycleInstallCCRequest{
		Label:   label,
		Package: ccPkg,
	}
	packageID = lcpackager.ComputePackageID(label, ccPkg)
	for _, org := range NetworkOrgs {
		orgPeers, err := DiscoverLocalPeers(*org.OrgAdminClientContext, org.OrgPeerNum)
		if err != nil {
			fmt.Errorf("DiscoverLocalPeers error: %v", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "DiscoverLocalPeers error"})
		}
		if flag, _ := checkInstalled(packageID, orgPeers[0], org.OrgResMgmt); flag == false {
			if _, err := org.OrgResMgmt.LifecycleInstallCC(installCCReq, resmgmt.WithTargets(orgPeers...), resmgmt.WithRetry(retry.DefaultResMgmtOpts)); err != nil {

				fmt.Println(err)
				fmt.Errorf("LifecycleInstallCC error: %v", err)
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
			}
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Install chaincode : " + label})
}
func InstalledCC(c *gin.Context) {
	orgPeers, err := DiscoverLocalPeers(*NetworkInfo.Orgs[0].OrgAdminClientContext, 1)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "getInstallCCPackage err"})
		fmt.Errorf("DiscoverLocalPeers error: %v", err)
		return
	}
	_, err = NetworkInfo.Orgs[0].OrgResMgmt.LifecycleGetInstalledCCPackage(packageID, resmgmt.WithTargets([]fab.Peer{orgPeers[0]}...))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "LifecycleGetInstalledCCPackage error"})
		fmt.Errorf("LifecycleGetInstallCC error: %v", err)
		return
	}
	resp1, err := NetworkInfo.Orgs[0].OrgResMgmt.LifecycleQueryInstalledCC(resmgmt.WithTargets([]fab.Peer{orgPeers[0]}...))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "LifecycleQueryInstalledCC error"})
		fmt.Errorf("LifecycleQueryInstalledCC error: %v", err)
		return
	}
	packageID1 := ""
	for _, t := range resp1 {
		if t.PackageID == packageID {
			packageID1 = t.PackageID
		}
	}
	if !strings.EqualFold(packageID, packageID1) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "check package id error"})
		fmt.Errorf("check package id error")
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "installed package : " + resp1[0].Label})
}
func ApproveCC(c *gin.Context) {
	var requestBody approveInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}
	mspIDs := []string{}
	for _, org := range NetworkInfo.Orgs {
		mspIDs = append(mspIDs, org.OrgMspId)
	}
	ccPolicy := policydsl.SignedByNOutOfGivenRole(int32(len(mspIDs)), mb.MSPRole_MEMBER, mspIDs)
	approveCCReq := resmgmt.LifecycleApproveCCRequest{
		Name:              requestBody.ChaincodeID,
		Version:           requestBody.ChaincodeVersion,
		PackageID:         packageID,
		Sequence:          requestBody.Sequence,
		EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",
		SignaturePolicy:   ccPolicy,
		InitRequired:      requestBody.InitRequired,
	}
	for _, org := range NetworkOrgs {
		orgPeers, err := DiscoverLocalPeers(*org.OrgAdminClientContext, org.OrgPeerNum)
		fmt.Printf(">>> chaincode approved by %s peers:\n", org.OrgName)
		for _, p := range orgPeers {
			fmt.Printf("	%s\n", p.URL())
		}

		if err != nil {
			fmt.Errorf("DiscoverLocalPeers error: %v", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "DiscoverLocalPeers error"})
			return
		}
		_, err = org.OrgResMgmt.LifecycleApproveCC(NetworkInfo.ChannelID, approveCCReq, resmgmt.WithTargets(orgPeers...), resmgmt.WithOrdererEndpoint(NetworkInfo.OrdererEndpoint), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
		if err != nil {
			fmt.Errorf("LifecycleApproveCC error: %v", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "LifecycleApproveCC error"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Approve chaincodeDefinition successfully"})
}
func ApprovedCC(c *gin.Context) {
	// Query approve cc
	queryApprovedCCReq := resmgmt.LifecycleQueryApprovedCCRequest{
		Name:     NetworkInfo.ChaincodeID,
		Sequence: 1,
	}

	for _, org := range NetworkOrgs {
		orgPeers, err := DiscoverLocalPeers(*org.OrgAdminClientContext, org.OrgPeerNum)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "DiscoverLocalPeers error"})
			fmt.Errorf("DiscoverLocalPeers error: %v", err)
			return
		}
		// Query approve cc
		for _, p := range orgPeers {
			resp, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
				func() (interface{}, error) {
					resp1, err := org.OrgResMgmt.LifecycleQueryApprovedCC(NetworkInfo.ChannelID, queryApprovedCCReq, resmgmt.WithTargets(p))
					if err != nil {
						return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleQueryApprovedCC returned error: %v", err), nil)
					}
					return resp1, err
				},
			)
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "NewInvoker error"})
				fmt.Errorf("Org %s Peer %s NewInvoker error: %v", org.OrgName, p.URL(), err)
				return
			}
			if resp == nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Got nil invoker"})
				fmt.Errorf("Org %s Peer %s Got nil invoker", org.OrgName, p.URL())
				return
			}
		}
	}
	// Check commit readiness
	mspIds := []string{}
	for _, org := range NetworkInfo.Orgs {
		mspIds = append(mspIds, org.OrgMspId)
	}
	ccPolicy := policydsl.SignedByNOutOfGivenRole(int32(len(mspIds)), mb.MSPRole_MEMBER, mspIds)
	req := resmgmt.LifecycleCheckCCCommitReadinessRequest{
		Name:    NetworkInfo.ChaincodeID,
		Version: NetworkInfo.ChaincodeVersion,
		//PackageID:         packageID,
		EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",
		SignaturePolicy:   ccPolicy,
		Sequence:          1,
		InitRequired:      true,
	}
	for _, org := range NetworkOrgs {
		orgPeers, err := DiscoverLocalPeers(*org.OrgAdminClientContext, org.OrgPeerNum)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "DiscoverLocalPeers error"})
			fmt.Errorf("DiscoverLocalPeers error: %v", err)
			return
		}
		for _, p := range orgPeers {
			resp, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
				func() (interface{}, error) {
					resp1, err := org.OrgResMgmt.LifecycleCheckCCCommitReadiness(NetworkInfo.ChannelID, req, resmgmt.WithTargets(p))
					fmt.Printf("LifecycleCheckCCCommitReadiness cc = %v, = %v\n", NetworkInfo.ChaincodeID, resp1)
					if err != nil {
						return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleCheckCCCommitReadiness returned error: %v", err), nil)
					}
					flag := true
					for _, r := range resp1.Approvals {
						flag = flag && r
					}
					if !flag {
						return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleCheckCCCommitReadiness returned : %v", resp1), nil)
					}
					return resp1, err
				},
			)
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "NewInvoker error"})
				fmt.Errorf("Org %s Peer %s NewInvoker error: %v", org.OrgName, p.URL(), err)
				return
			}
			if resp == nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Got nil invoker"})
				fmt.Errorf("Got nill invoker response")
				return
			}
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Get Approved chaincodeDefinition successfully"})
}
func CommitCC(c *gin.Context) {
	var requestBody commitInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}
	mspIDs := []string{}
	for _, org := range NetworkInfo.Orgs {
		mspIDs = append(mspIDs, org.OrgMspId)
	}
	ccPolicy := policydsl.SignedByNOutOfGivenRole(int32(len(mspIDs)), mb.MSPRole_MEMBER, mspIDs)

	req := resmgmt.LifecycleCommitCCRequest{
		Name:              requestBody.ChaincodeID,
		Version:           requestBody.ChaincodeVersion,
		Sequence:          requestBody.Sequence,
		EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",
		SignaturePolicy:   ccPolicy,
		InitRequired:      requestBody.InitRequired,
	}
	_, err := NetworkInfo.Orgs[0].OrgResMgmt.LifecycleCommitCC(NetworkInfo.ChannelID, req, resmgmt.WithOrdererEndpoint(NetworkInfo.OrdererEndpoint), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "LifecycleCommitCC error"})
		fmt.Errorf("LifecycleCommitCC error: %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "commit chaincode successfully"})
}
func CommittedCC(c *gin.Context) {
	req := resmgmt.LifecycleQueryCommittedCCRequest{
		Name: NetworkInfo.ChaincodeID,
	}
	flag2 := false
	for _, org := range NetworkInfo.Orgs {
		orgPeers, err := DiscoverLocalPeers(*org.OrgAdminClientContext, org.OrgPeerNum)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "DiscoverLocalPeers error"})
			fmt.Errorf("DiscoverLocalPeers error: %v", err)
			return
		}
		for _, p := range orgPeers {
			resp, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
				func() (interface{}, error) {
					resp1, err := org.OrgResMgmt.LifecycleQueryCommittedCC(NetworkInfo.ChannelID, req, resmgmt.WithTargets(p))
					if err != nil {
						return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleQueryCommittedCC returned error: %v", err), nil)
					}
					flag := false
					for _, r := range resp1 {
						if r.Name == NetworkInfo.ChaincodeID && r.Sequence == 1 {
							flag = true
							break
						}
					}
					if (!flag2) {
						c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("LifecycleQueryCommittedCC returned : %v", resp1)})
						flag2 = true
					}
					if !flag {
						return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("LifecycleQueryCommittedCC returned : %v", resp1), nil)
					}
					return resp1, err
				},
			)
			if err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "NewInvoker error"})
				fmt.Errorf("NewInvoker error: %v", err)
				return
			}
			if resp == nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Got nil invoker"})
				fmt.Errorf("Got nil invoker response")
				return
			}
		}
	}
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "Query committed chaincode successfully"})
}
func InitCC(c *gin.Context) {
	var requestBody initInfo
	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request format."})
		return
	}
	App = sdkInit.Application{
		SdkEnvInfo: &NetworkInfo,
	}
	fmt.Println(requestBody)
	fmt.Println(requestBody.ChaincodeID)
	fmt.Println(requestBody.ChannelName)
	fmt.Println(NetworkInfo.ChaincodeID)
	//prepare channel client context using client context
	clientChannelContext := sdk.ChannelContext(requestBody.ChannelName, fabsdk.WithUser(NetworkInfo.Orgs[0].OrgUser), fabsdk.WithOrg(NetworkInfo.Orgs[0].OrgName))
	// Channel client is used to query and execute transactions (Org1 is default org)
	client, err := channel.New(clientChannelContext)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to create new channel client"})
		fmt.Errorf("Failed to create new channel client: %s", err)
		return
	}

	// init
	_, err = client.Execute(channel.Request{ChaincodeID: requestBody.ChaincodeID, Fcn: "init", Args: nil, IsInit: true},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to init"})
		fmt.Errorf("Failed to init: %s", err)
		return
	}
	App = sdkInit.Application{
		SdkEnvInfo: &NetworkInfo,
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Init chaincode successfully"})
}
func TestCC(c *gin.Context) {
	a := []string{"set", "ID1", "123"}
	ret, err := App.Set(a)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to test"})
		fmt.Println(err)
	}
	fmt.Println("<--- add row1　--->：", ret)

	a = []string{"set", "ID2", "456"}
	ret, err = App.Set(a)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to test"})
		fmt.Println(err)
	}
	fmt.Println("<--- add row2　--->：", ret)

	a = []string{"set", "ID3", "789"}
	ret, err = App.Set(a)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to test"})
		fmt.Println(err)
	}
	fmt.Println("<--- add row3　--->：", ret)

	a = []string{"get", "ID3"}
	response, err := App.Get(a)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to test"})
		fmt.Println(err)
	}
	fmt.Println("<--- get row3　--->：", response)

	a = []string{"get", "ID2"}
	response, err = App.Get(a)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to test"})
		fmt.Println(err)
	}
	fmt.Println("<--- get row2　--->：", response)
	a = []string{"get", "ID1"}
	response, err = App.Get(a)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to test"})
		fmt.Println(err)
	}
	fmt.Println("<--- get row1　--->：", response)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Test chaincode successfully"})
}

func checkInstalled(packageID string, peer fab.Peer, client *resmgmt.Client) (bool, error) {
	flag := false
	resp1, err := client.LifecycleQueryInstalledCC(resmgmt.WithTargets(peer))
	if err != nil {
		return flag, fmt.Errorf("LifecycleQueryInstalledCC error: %v", err)
	}
	for _, t := range resp1 {
		if t.PackageID == packageID {
			flag = true
		}
	}
	return flag, nil
}
func DiscoverLocalPeers(ctxProvider contextAPI.ClientProvider, expectedPeers int) ([]fabAPI.Peer, error) {
	ctx, err := contextImpl.NewLocal(ctxProvider)
	if err != nil {
		return nil, fmt.Errorf("error creating local context: %v", err)
	}

	discoveredPeers, err := retry.NewInvoker(retry.New(retry.TestRetryOpts)).Invoke(
		func() (interface{}, error) {
			peers, serviceErr := ctx.LocalDiscoveryService().GetPeers()
			if serviceErr != nil {
				return nil, fmt.Errorf("getting peers for MSP [%s] error: %v", ctx.Identifier().MSPID, serviceErr)
			}
			if len(peers) < expectedPeers {
				return nil, status.New(status.TestStatus, status.GenericTransient.ToInt32(), fmt.Sprintf("Expecting %d peers but got %d", expectedPeers, len(peers)), nil)
			}
			return peers, nil
		},
	)
	if err != nil {
		return nil, err
	}

	return discoveredPeers.([]fabAPI.Peer), nil
}
