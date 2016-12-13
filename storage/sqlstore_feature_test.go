package storage

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var defaultDb = DataBase{USER: "postgres", PASSWORD: "ftftft", NAME: "postgres", HOST: "localhost", PORT: "5432"}
var defaultConfig = Config{defaultDb}

func TestFeatureToggleStoreImpl_CreateFeature(t *testing.T) {
	var fs FeatureToggleStore = NewFeatureToggleStoreImpl()
	fs.SetConfig(defaultConfig)

	err := fs.Open()
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %v", err))
	}
	defer fs.Close()

	feature := NewFeature(randomSufix("Feature-"), true, "f description")
	featureId, err := fs.CreateFeature(*feature)

	require.NotNil(t, featureId, "Should get featureId, %v", err)
}

func TestFeatureToggleStoreImpl_ReadFeature(t *testing.T) {
	var fs FeatureToggleStore = NewFeatureToggleStoreImpl()
	fs.SetConfig(defaultConfig)

	err := fs.Open()
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %v", err))
	}
	defer fs.Close()

	feature := NewFeature(randomSufix("Feature-"), true, "f description")
	featureId, err := fs.CreateFeature(*feature)

	require.NotNil(t, featureId, "Should get featureId, %v", err)

	f, err := fs.ReadFeature(*featureId)
	require.NotNil(t, *f, "Should get feature from featureId %s, err", featureId, err)
	assert.Equal(t, feature.Name, f.Name, "Should get feature name, %v", err)
	assert.Equal(t, feature.Description, f.Description, "Should get feature description, %v", err)
	assert.Equal(t, feature.Enabled, f.Enabled, "Should get feature enabled, %v", err)
}

func TestFeatureToggleStoreImpl_ReadFeatureByName(t *testing.T) {
	var fs FeatureToggleStore = NewFeatureToggleStoreImpl()
	fs.SetConfig(defaultConfig)

	err := fs.Open()
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %v", err))
	}
	defer fs.Close()

	featureName := randomSufix("Feature-")
	feature := NewFeature(featureName, true, "f description")
	featureId, err := fs.CreateFeature(*feature)

	require.NotNil(t, featureId, "Should get featureId, %v", err)

	f, err := fs.ReadFeatureByName(featureName)
	require.NotNil(t, *f, "Should get feature from featureId %s, err", featureId, err)
	assert.Equal(t, feature.Name, f.Name, "Should get feature name, %v", err)
	assert.Equal(t, feature.Description, f.Description, "Should get feature description, %v", err)
	assert.Equal(t, feature.Enabled, f.Enabled, "Should get feature enabled, %v", err)
}

func TestFeatureToggleStoreImpl_ReadFeatureByName__unknown(t *testing.T) {
	var fs FeatureToggleStore = NewFeatureToggleStoreImpl()
	fs.SetConfig(defaultConfig)

	err := fs.Open()
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %v", err))
	}
	defer fs.Close()

	f, err := fs.ReadFeatureByName("unknown feature")
	assert.Nil(t, f, "Should not get a feature, %v", f)
	assert.Nil(t, err, "Should not get an error, %v", err)
}

func TestFeatureToggleStoreImpl_DeleteFeature(t *testing.T) {
	var fs FeatureToggleStore = NewFeatureToggleStoreImpl()
	fs.SetConfig(defaultConfig)

	err := fs.Open()
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %v", err))
	}
	defer fs.Close()

	feature := NewFeature(randomSufix("Feature-"), true, "f description")
	featureId, err := fs.CreateFeature(*feature)

	require.NotNil(t, featureId, "Should get featureId, %v", err)

	res, err := fs.DeleteFeature(*featureId)

	require.True(t, *res, "Should get true from delete operation for featureId %s, %v", featureId, err)
}

func TestFeatureToggleStoreImpl_SearchFeature(t *testing.T) {
	var fs FeatureToggleStore = NewFeatureToggleStoreImpl()
	fs.SetConfig(defaultConfig)

	err := fs.Open()
	if err != nil {
		panic(fmt.Sprintf("Failed to open database, %v", err))
	}
	defer fs.Close()

	featureName := randomSufix("Feature-")
	feature := NewFeature(featureName, true, "Fake description")
	_, err = fs.CreateFeature(*feature)

	if err != nil {
		t.Fail()
	}

	res, err := fs.SearchFeature(featureName)

	require.Equal(t, 1, len(*res), "Should receive exactly ONE feature back, got %v, %v", *res, err)
	require.Equal(t, featureName, (*res)[0].Name, "Names should match for feature with name %s, %v", featureName, err)
}