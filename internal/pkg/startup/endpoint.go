/*******************************************************************************
 * Copyright 2018 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/
package startup

import (
	"fmt"
	"os"
	"time"

	consulclient "sup-sys-mgmt-agent/internal/pkg/consul"
	"sup-sys-mgmt-agent/pkg/clients/types"
)

type Endpoint struct{}

func (e Endpoint) Monitor(params types.EndpointParams, ch chan string) {
	for {
		data, err := consulclient.GetServiceEndpoint(params.ServiceKey)
		if err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
		}
		url := fmt.Sprintf("http://%s:%v%s", data.Address, data.Port, params.Path)
		ch <- url
		time.Sleep(time.Millisecond * time.Duration(params.Interval))
	}
}
