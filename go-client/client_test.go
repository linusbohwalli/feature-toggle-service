package go_client_test

import "testing"

func TestOpen(t *testing.T) {
	//create mocks?
}

/*
func TestGetFeaturesForProperties(t *testing.T) {
	c := NewClient()
	c.cc = grpc.Dial
	c.fs = api.NewFeatureToggleServiceClient(cc)

}

func (c *Client) GetFeaturesForProperties(req *api.GetFeaturesByPropertiesRequest) (*api.GetFeaturesByPropertiesResponse, error) {

	resp, err := c.fs.GetFeaturesForProperties(context.Background(), req)
	if err != nil {
		grpclog.Printf("Get features for properties failed: %v", err)
		return nil, err
	}

	return resp, nil
}
*/
