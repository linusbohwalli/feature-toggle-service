package go_client

import (
	"context"

	api "github.com/linusbohwalli/feature-toggle-service/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const grpcServerAddr = "localhost:9090"

var opts []grpc.DialOption

//Client wraps grpc connection and FeatureToggleServiceClient
type Client struct {
	cc *grpc.ClientConn
	fs api.FeatureToggleServiceClient
}

func (c *Client) Open() error {

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(grpcServerAddr, opts...)
	if err != nil {
		grpclog.Fatalf("Dial failed: %v", err)
		return err
	}

	c.cc = conn
	c.fs = api.NewFeatureToggleServiceClient(conn)
	return nil
}

func (c *Client) Close() error {

	if err := c.cc.Close(); err != nil {
		return err
	}
	return nil
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GetFeaturesForProperties(req *api.GetFeaturesByPropertiesRequest) (*api.GetFeaturesByPropertiesResponse, error) {

	resp, err := c.fs.GetFeaturesForProperties(context.Background(), req)
	if err != nil {
		grpclog.Printf("Get features for properties failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateToggleRule(req *api.CreateToggleRuleRequest) (*api.CreateToggleRuleResponse, error) {

	resp, err := c.fs.CreateToggleRule(context.Background(), req)
	if err != nil {
		grpclog.Printf("Create toggle rule failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) ReadToggleRule(req *api.ReadToggleRuleRequest) (*api.ReadToggleRuleResponse, error) {

	resp, err := c.fs.ReadToggleRule(context.Background(), req)
	if err != nil {
		grpclog.Printf("Read toggle rule failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteToggleRule(req *api.DeleteToggleRuleRequest) (*api.DeleteToggleRuleResponse, error) {

	resp, err := c.fs.DeleteToggleRule(context.Background(), req)
	if err != nil {
		grpclog.Printf("Delete toggle rule failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) SearchToggleRule(req *api.SearchToggleRuleRequest) (*api.SearchToggleRuleResponse, error) {

	resp, err := c.fs.SearchToggleRule(context.Background(), req)
	if err != nil {
		grpclog.Printf("Search toggle rule failed: %v", err)
		return nil, err
	}

	return resp, nil

}

func (c *Client) CreateFeature(req *api.CreateFeatureRequest) (*api.CreateFeatureResponse, error) {

	resp, err := c.fs.CreateFeature(context.Background(), req)
	if err != nil {
		grpclog.Printf("Create feature failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) ReadFeature(req *api.ReadFeatureRequest) (*api.ReadFeatureResponse, error) {

	resp, err := c.fs.ReadFeature(context.Background(), req)
	if err != nil {
		grpclog.Printf("Read feature failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteFeature(req *api.DeleteFeatureRequest) (*api.DeleteFeatureResponse, error) {

	resp, err := c.fs.DeleteFeature(context.Background(), req)
	if err != nil {
		grpclog.Printf("Delete feature failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) SearchFeature(req *api.SearchFeatureRequest) (*api.SearchFeatureResponse, error) {

	resp, err := c.fs.SearchFeature(context.Background(), req)
	if err != nil {
		grpclog.Printf("Search feature failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateProperty(req *api.CreatePropertyRequest) (*api.CreatePropertyResponse, error) {

	resp, err := c.fs.CreateProperty(context.Background(), req)
	if err != nil {
		grpclog.Printf("Create property failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) ReadProperty(req *api.ReadPropertyRequest) (*api.ReadPropertyResponse, error) {

	resp, err := c.fs.ReadProperty(context.Background(), req)
	if err != nil {
		grpclog.Printf("Read property failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteProperty(req *api.DeletePropertyRequest) (*api.DeletePropertyResponse, error) {

	resp, err := c.fs.DeleteProperty(context.Background(), req)
	if err != nil {
		grpclog.Printf("Delete property failed: %v", err)
		return nil, err
	}

	return resp, nil
}

func (c *Client) SearchProperty(req *api.SearchPropertyRequest) (*api.SearchPropertyResponse, error) {

	resp, err := c.fs.SearchProperty(context.Background(), req)
	if err != nil {
		grpclog.Printf("Search property failed: %v", err)
		return nil, err
	}

	return resp, nil
}
