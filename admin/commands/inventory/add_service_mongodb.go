// Copyright (C) 2023 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inventory

import (
	"github.com/percona/pmm/admin/commands"
	"github.com/percona/pmm/api/inventorypb/json/client"
	"github.com/percona/pmm/api/inventorypb/json/client/services"
)

var addServiceMongoDBResultT = commands.ParseTemplate(`
MongoDB Service added.
Service ID     : {{ .Service.ServiceID }}
Service name   : {{ .Service.ServiceName }}
Node ID        : {{ .Service.NodeID }}
{{ if .Service.Socket -}}
Socket         : {{ .Service.Socket }}
{{- else -}}
Address        : {{ .Service.Address }}
Port           : {{ .Service.Port }}
{{- end }}
Environment    : {{ .Service.Environment }}
Cluster name   : {{ .Service.Cluster }}
Replication set: {{ .Service.ReplicationSet }}
Custom labels  : {{ .Service.CustomLabels }}
`)

type addServiceMongoDBResult struct {
	Service *services.AddMongoDBServiceOKBodyMongodb `json:"mongodb"`
}

func (res *addServiceMongoDBResult) Result() {}

func (res *addServiceMongoDBResult) String() string {
	return commands.RenderTemplate(addServiceMongoDBResultT, res)
}

// AddServiceMongoDBCommand is used by Kong for CLI flags and commands.
type AddServiceMongoDBCommand struct {
	ServiceName    string            `arg:"" optional:"" name:"name" help:"Service name"`
	NodeID         string            `arg:"" optional:"" help:"Node ID"`
	Address        string            `arg:"" optional:"" help:"Address"`
	Port           int64             `arg:"" optional:"" help:"Port"`
	Socket         string            `help:"Path to socket"`
	Environment    string            `help:"Environment name"`
	Cluster        string            `help:"Cluster name"`
	ReplicationSet string            `help:"Replication set name"`
	CustomLabels   map[string]string `mapsep:"," help:"Custom user-assigned labels"`
}

func (cmd *AddServiceMongoDBCommand) RunCmd() (commands.Result, error) {
	customLabels := commands.ParseCustomLabels(cmd.CustomLabels)
	params := &services.AddMongoDBServiceParams{
		Body: services.AddMongoDBServiceBody{
			ServiceName:    cmd.ServiceName,
			NodeID:         cmd.NodeID,
			Address:        cmd.Address,
			Port:           cmd.Port,
			Socket:         cmd.Socket,
			Environment:    cmd.Environment,
			Cluster:        cmd.Cluster,
			ReplicationSet: cmd.ReplicationSet,
			CustomLabels:   customLabels,
		},
		Context: commands.Ctx,
	}

	resp, err := client.Default.Services.AddMongoDBService(params)
	if err != nil {
		return nil, err
	}
	return &addServiceMongoDBResult{
		Service: resp.Payload.Mongodb,
	}, nil
}
