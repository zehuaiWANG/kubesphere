/*
Copyright 2019 The KubeSphere Authors.

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

package operations

import (
	k8serr "k8s.io/apimachinery/pkg/api/errors"
	"kubesphere.io/kubesphere/pkg/models/workloads"
	"kubesphere.io/kubesphere/pkg/server/errors"
	"net/http"

	"github.com/emicklei/go-restful"

	"fmt"
)

func RerunJob(req *restful.Request, resp *restful.Response) {
	var err error

	job := req.PathParameter("job")
	namespace := req.PathParameter("namespace")
	action := req.QueryParameter("action")
	resourceVersion := req.QueryParameter("resourceVersion")

	switch action {
	case "rerun":
		err = workloads.JobReRun(namespace, job, resourceVersion)
	default:
		resp.WriteHeaderAndEntity(http.StatusBadRequest, errors.Wrap(fmt.Errorf("invalid operation %s", action)))
		return
	}
	if err != nil {
		if k8serr.IsConflict(err) {
			resp.WriteHeaderAndEntity(http.StatusConflict, errors.Wrap(err))
			return
		}
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, errors.Wrap(err))
		return
	}

	resp.WriteAsJson(errors.None)
}
