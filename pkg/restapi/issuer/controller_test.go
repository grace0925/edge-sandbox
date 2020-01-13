/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package issuer

import (
	"testing"

	"golang.org/x/oauth2"

	"github.com/stretchr/testify/require"

	"github.com/trustbloc/edge-sandbox/pkg/restapi/issuer/operation"
)

func TestController_New(t *testing.T) {
	controller, err := New(&operation.Config{OAuth2Config: &oauth2.Config{}})
	require.NoError(t, err)
	require.NotNil(t, controller)
}

func TestController_GetOperations(t *testing.T) {
	controller, err := New(&operation.Config{OAuth2Config: &oauth2.Config{}})
	require.NoError(t, err)
	require.NotNil(t, controller)

	ops := controller.GetOperations()
	require.Equal(t, 2, len(ops))
}