/*
Copyright 2019-2020 vChain, Inc.

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

package immuclient

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/codenotary/immudb/cmd/immuclient/immuclienttest"
	"github.com/codenotary/immudb/pkg/server"
	"github.com/codenotary/immudb/pkg/server/servertest"
	"github.com/spf13/cobra"
)

func TestLogin(t *testing.T) {
	options := server.Options{}.WithAuth(true).WithInMemoryStore(true)
	bs := servertest.NewBufconnServer(options)
	bs.Start()

	cmdl := commandline{
		immucl: newClient(&immuclienttest.PasswordReader{
			Pass: []string{"immudb"},
		}, bs.Dialer, nil),
	}
	cmd := cobra.Command{}
	cmdl.login(&cmd)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"login", "immudb"})
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	msg, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(msg), "Successfully logged in") {
		t.Fatal(err)
	}

	cmdl.logout(&cmd)
	cmd.SetOut(b)
	cmd.SetArgs([]string{"logout"})
	err = cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	msg, err = ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}