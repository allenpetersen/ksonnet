// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package actions

import (
	"bytes"
	"testing"

	"github.com/ksonnet/ksonnet/pkg/app"
	amocks "github.com/ksonnet/ksonnet/pkg/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestPrototypePreview(t *testing.T) {
	withApp(t, func(appMock *amocks.App) {
		libaries := app.LibraryRefSpecs{}

		appMock.On("Libraries").Return(libaries, nil)

		args := []string{
			"--name", "myDeployment",
			"--image", "nginx",
			"--port", "80",
		}

		in := map[string]interface{}{
			OptionApp:       appMock,
			OptionQuery:     "single-port-deployment",
			OptionArguments: args,
		}

		a, err := NewPrototypePreview(in)
		require.NoError(t, err)

		var buf bytes.Buffer
		a.out = &buf

		err = a.Run()
		require.NoError(t, err)

		assertOutput(t, "prototype/preview/output.txt", buf.String())
	})
}
