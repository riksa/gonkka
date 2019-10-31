package gocd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRest_History(t *testing.T) {
	assert := assert.New(t)

	rest := NewRest("jenkins2.local:8153", context.TODO())
	response, err := rest.History("ios-build-app", 190)
	assert.NoError(err)

	bytes, err := response.MarshalBinary()
	assert.NoError(err)

	t.Log(string(bytes))
}

func TestRest_Instance(t *testing.T) {
	assert := assert.New(t)

	rest := NewRest("jenkins2.local:8153", context.TODO())
	response, err := rest.Instance("ios-build-app", 190)
	assert.NoError(err)

	bytes, err := response.MarshalBinary()
	assert.NoError(err)

	t.Log(string(bytes))
}

