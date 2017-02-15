package feature_toggle_impl

import (
	api "github.com/linusbohwalli/feature-toggle-service/api"
	"github.com/linusbohwalli/feature-toggle-service/storage"
	"github.com/stretchr/testify/require"
	"testing"
)

//Storage Mock
func storage_ReadFeature(Id string) *storage.Feature {
	return &storage.Feature{Name: "test", Id: Id, Enabled: false, Description: "test desc"}
}

func TestFeatureToggleServiceServer_ReadFeature(t *testing.T) {

	req := &api.ReadFeatureRequest{Id: "d6cfdca1-2051-4cb8-9f77-cf69aa69993a"}

	storeFeature := storage_ReadFeature(req.Id)

	feature := storeFeatureToResponseFeature(*storeFeature)
	require.NotNil(t, feature, "should get feature after convert from storeFeature")
	require.EqualValues(t, feature.Name, storeFeature.Name, "Name in feature and storeFeature should be same")
	require.EqualValues(t, feature.Id, storeFeature.Id, "Id in feature and storeFeature should be same")
	require.EqualValues(t, feature.Enabled, storeFeature.Enabled, "Enabled in feature and storeFeature should be same")
	require.EqualValues(t, feature.Description, storeFeature.Description, "Description in feature and storeFeature should be same")
}

//CreateFeature
//DeleteFeature
//SearchFeature
