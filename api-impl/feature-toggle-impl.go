package feature_toggle_impl

import (
	"fmt"

	api "github.com/linusbohwalli/feature-toggle-service/api"
	"github.com/linusbohwalli/feature-toggle-service/featuretree"
	"github.com/linusbohwalli/feature-toggle-service/storage"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type FeatureToggleServiceServer struct {
	fs   storage.FeatureToggleStore
	tree *featuretree.ToggleRuleTree
}

func (s *FeatureToggleServiceServer) GetFeaturesForProperties(ctx context.Context, req *api.GetFeaturesByPropertiesRequest) (*api.GetFeaturesByPropertiesResponse, error) {
	if s.tree == nil {
		return nil, errors.New("Feature toggle service not initialized.")
	}
	fmt.Printf("getfeat: %v\n", req)
	return &api.GetFeaturesByPropertiesResponse{Features: s.tree.FindFeatures(req.Properties)}, nil
}

func (s *FeatureToggleServiceServer) CreateToggleRule(ctx context.Context, req *api.CreateToggleRuleRequest) (*api.CreateToggleRuleResponse, error) {
	fmt.Printf("CreateToggleRule: %v\n", req.ToggleRule)

	if req.ToggleRule == nil {
		return nil, errors.New("No Toggle Rule received")
	}

	if req.ToggleRule.Properties == nil {
		return nil, errors.New("No properties received for Toggle Rule")
	}

	feature, err := s.fs.ReadFeatureByName(req.ToggleRule.Name)
	if feature == nil {
		return nil, errors.New("Unknown feature")
	}

	propAsSlice := []string{}
	for k, v := range req.ToggleRule.Properties {
		propAsSlice = append(propAsSlice, k, v)
	}


	ruleId, err := s.fs.CreateToggleRule(*storage.NewToggleRule((*feature).Id, req.ToggleRule.Enabled, propAsSlice...))
	if err != nil {
		return nil, err
	}
	response := new(api.CreateToggleRuleResponse)
	response.Id = *ruleId

	return response, nil
}

func (s *FeatureToggleServiceServer) ReadToggleRule(ctx context.Context, req *api.ReadToggleRuleRequest) (*api.ReadToggleRuleResponse, error) {
	fmt.Printf("ReadToggleRule: id=%s\n", req.Id)

	storeToggleRule, err := s.fs.ReadToggleRule(req.Id)
	if err != nil {
		return nil, err
	}

	if storeToggleRule == nil {
		return nil, errors.New("Unable to find toggle rule")
	}

	toggleRule := storeToggleRuleToResponseToggleRule(storeToggleRule)
	response := new(api.ReadToggleRuleResponse)
	response.ToggleRule = toggleRule

	return response, nil
}

func storeToggleRuleToResponseToggleRule(store *storage.ToggleRule) *api.ToggleRule {
	return &api.ToggleRule{
		Id:         store.Id,
		Enabled:    store.Enabled,
		Properties: store.Properties,
	}
}

func (s *FeatureToggleServiceServer) DeleteToggleRule(ctx context.Context, req *api.DeleteToggleRuleRequest) (*api.DeleteToggleRuleResponse, error) {
	fmt.Printf("DeleteToggleRule: id=%s\n", req.Id)

	if _, err := s.fs.DeleteToggleRule(req.Id); err != nil {
		return nil, err
	}

	return new(api.DeleteToggleRuleResponse), nil
}

func (s *FeatureToggleServiceServer) SearchToggleRule(ctx context.Context, req *api.SearchToggleRuleRequest) (*api.SearchToggleRuleResponse, error) {
	fmt.Printf("SearchToggleRule: %s\n", req)

	//var Filter map[string]string
	//Filter = make(map[string]string)
	//Filter[] =

	/*_, err := s.fs.SearchToggleRule(&req.Name, Filter)
	if err != nil {
		return nil, errors.New( "failed to search for rules")
	}*/

	response := new(api.SearchToggleRuleResponse)
	//response.ToggleRules = []*api.ToggleRule{ToApiToggleRules(rules)}

	return response, nil
}

/*
func ToApiToggleRules(rules *[]storage.ToggleRule) []api.ToggleRule{
	toggleRules := make([]api.ToggleRule, len(rules))
	for i := 0; i < len(rules); i++ {
		toggleRules[i] = ToApiToggleRule( rules[0])
	}
}
func ToApiToggleRule(rule storage.ToggleRule) api.ToggleRule {
	return api.ToggleRule{Id:rule.Id,Name:rule.FeatureId}
}
*/

func (s *FeatureToggleServiceServer) CreateFeature(ctx context.Context, req *api.CreateFeatureRequest) (*api.CreateFeatureResponse, error) {
	fmt.Printf("CreateFeature: %v\n", req.Feature)
	fmt.Printf("CreateFeature: id=%s\n", req.Feature.Name)

	featureId, err := s.fs.CreateFeature(*storage.NewFeature(req.Feature.Name, req.Feature.Enabled, req.Feature.Description))
	if err != nil {
		return nil, err
	}

	response := new(api.CreateFeatureResponse)
	response.Id = *featureId

	return response, nil
}

func (s *FeatureToggleServiceServer) ReadFeature(ctx context.Context, req *api.ReadFeatureRequest) (*api.ReadFeatureResponse, error) {
	fmt.Printf("ReadFeature: id=%s\n", req.Id)

	storeFeature, err := s.fs.ReadFeature(req.Id)
	if err != nil {
		return nil, err
	}

	if storeFeature == nil {
		return nil, errors.New("No such feature exists")
	}

	feature := storeFeatureToResponseFeature(*storeFeature)

	response := new(api.ReadFeatureResponse)
	response.Feature = feature

	return response, nil
}

func (s *FeatureToggleServiceServer) DeleteFeature(ctx context.Context, req *api.DeleteFeatureRequest) (*api.DeleteFeatureResponse, error) {
	fmt.Printf("DeleteFeature: id=%s\n", req.Id)

	_, err := s.fs.DeleteFeature(req.Id)
	if err != nil {
		return nil, err
	}

	return new(api.DeleteFeatureResponse), nil
}

func storeFeatureToResponseFeature(storeFeature storage.Feature) *api.Feature {
	return &api.Feature{
		Name:        storeFeature.Name,
		Id:          storeFeature.Id,
		Description: storeFeature.Description,
		Enabled:     storeFeature.Enabled}
}

func storeFeaturesToResponseFeatures(storeFeatures *[]storage.Feature) []*api.Feature {
	var responseFeatures []*api.Feature
	for _, f := range *storeFeatures {
		responseFeature := storeFeatureToResponseFeature(f)
		responseFeatures = append(responseFeatures, responseFeature)
	}
	return responseFeatures
}

func (s *FeatureToggleServiceServer) SearchFeature(ctx context.Context, req *api.SearchFeatureRequest) (*api.SearchFeatureResponse, error) {
	fmt.Printf("SearchFeature name=: %v\n", req.Name)

	features, err := s.fs.SearchFeature(req.Name)
	if err != nil {
		return nil, err
	}

	responseFeatures := storeFeaturesToResponseFeatures(features)

	response := api.SearchFeatureResponse{Features: responseFeatures}

	return &response, nil
}

func (s *FeatureToggleServiceServer) CreateProperty(ctx context.Context, req *api.CreatePropertyRequest) (*api.CreatePropertyResponse, error) {
	fmt.Printf("CreateProperty: %v\n", req.Property)

	propertyId, err := s.fs.CreateProperty(*storage.NewProperty(req.Property.Name, req.Property.Description))
	if err != nil {
		return nil, err
	}

	response := new(api.CreatePropertyResponse)
	response.Name = *propertyId

	return response, nil
}

func (s *FeatureToggleServiceServer) ReadProperty(ctx context.Context, req *api.ReadPropertyRequest) (*api.ReadPropertyResponse, error) {
	fmt.Printf("ReadProperty: id=%s\n", req.Name)

	property, err := s.fs.ReadProperty(req.Name)
	if err != nil {
		return nil, err
	}

	if property == nil {
		return nil, errors.New("No such property exists")
	}

	return &api.ReadPropertyResponse{&api.Property{Name: property.Name, Description: property.Description}}, nil
}

func (s *FeatureToggleServiceServer) DeleteProperty(ctx context.Context, req *api.DeletePropertyRequest) (*api.DeletePropertyResponse, error) {

	_, err := s.fs.DeleteProperty(req.Name)
	if err != nil {
		return nil, err
	}

	fmt.Printf("DeleteProperty: id=%s\n", req.Name)
	return new(api.DeletePropertyResponse), nil
}

func (s *FeatureToggleServiceServer) SearchProperty(ctx context.Context, req *api.SearchPropertyRequest) (*api.SearchPropertyResponse, error) {
	fmt.Printf("SearchProperty: %s\n", req.Name)

	props, err := s.fs.SearchProperty(req.Name)
	if err != nil {
		return nil, err
	}

	var resp []*api.Property
	for _, v := range *props {
		resp = append(resp, &api.Property{Name: v.Name, Description: v.Description})
	}

	response := new(api.SearchPropertyResponse)
	response.Properties = resp

	return response, nil
}

func newFeatureToggleServiceServer() *FeatureToggleServiceServer {
	s := new(FeatureToggleServiceServer)
	s.fs = storage.NewFeatureToggleStoreImpl()
	s.fs.Open()

	toggleRules, err := s.fs.GetEnabledToggleRules()
	if toggleRules != nil {
		propertyNames, err := s.fs.ReadAllPropertyNames()
		if err != nil {
			fmt.Printf("Failed to read all properties, %v\n", err)
			panic(err.Error())
		}
		tree := featuretree.NewFeatureTree(*propertyNames)

		for _, rule := range *toggleRules {
			err := tree.AddFeature(rule)
			if err != nil {
				fmt.Printf("%s\n", err.Error())
			} else {
				fmt.Printf("Added rule: %v\n", rule)
			}
		}
		s.tree = tree
		fmt.Print(s.tree.String())
	} else {
		fmt.Printf("Failed to init feature toggle service, %v\n", err)
		panic(err.Error())
	}
	return s
}

func RegisterFeatureToggleService(s *grpc.Server) {
	api.RegisterFeatureToggleServiceServer(s, newFeatureToggleServiceServer())
}
