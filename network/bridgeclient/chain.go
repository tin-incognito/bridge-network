package bridgeclient

type BridgeNetworkClient struct {
	/*httpClient *retryablehttp.Client*/
}

func NewBridgeNetworkClient() *BridgeNetworkClient {
	/*httpClient := retryablehttp.NewClient()*/
	return &BridgeNetworkClient{
		//httpClient: httpClient,
	}
}
