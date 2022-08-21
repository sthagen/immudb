/*
Copyright 2022 Codenotary Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package tokenservice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewInmemoryTokenService(t *testing.T) {
	ts := NewInmemoryTokenService()
	err := ts.SetToken("db1", "")
	require.Equal(t, ErrEmptyTokenProvided, err)
	err = ts.SetToken("db", "tk")
	require.NoError(t, err)
	present, err := ts.IsTokenPresent()
	require.True(t, present)
	require.NoError(t, err)
	db, err := ts.GetDatabase()
	require.Equal(t, db, "db")
	require.NoError(t, err)
	tk, err := ts.GetToken()
	require.NoError(t, err)
	require.Equal(t, tk, "tk")
	err = ts.DeleteToken()
	require.NoError(t, err)
}
