/*
Copyright 2018 The Kubernetes Authors.

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

package filters

import (
	"fmt"
	"net/http"

	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/apiserver/pkg/endpoints/metrics"
	apirequest "k8s.io/apiserver/pkg/endpoints/request"
)

func WithMaxLongrunningLimit(
	handler http.Handler,
	longRunningLimit int,
	longRunningRequestCheck apirequest.LongRunningRequestCheck,
) http.Handler {
	if longRunningLimit == 0 {
		return handler
	}
	var longRunningChan chan bool
	longRunningChan = make(chan bool, longRunningLimit)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestInfo, ok := apirequest.RequestInfoFrom(ctx)
		if !ok {
			handleError(w, r, fmt.Errorf("no RequestInfo found in context, handler chain must be wrong"))
			return
		}
		// Skip tracking non-long-running events.
		if longRunningRequestCheck != nil && !longRunningRequestCheck(r, requestInfo) {
			handler.ServeHTTP(w, r)
			return
		}

		if longRunningChan == nil {
			handler.ServeHTTP(w, r)
		} else {

			select {
			case longRunningChan <- true:
				defer func() {
					<-longRunningChan
				}()
				handler.ServeHTTP(w, r)

			default:
				// We need to split this data between buckets used for throttling.
				metrics.DroppedRequests.WithLabelValues(metrics.LongRunningKind).Inc()

				// at this point we're about to return a 429, BUT not all actors should be rate limited.  A system:master is so powerful
				// that they should always get an answer.  It's a super-admin or a loopback connection.
				if currUser, ok := apirequest.UserFrom(ctx); ok {
					for _, group := range currUser.GetGroups() {
						if group == user.SystemPrivilegedGroup {
							handler.ServeHTTP(w, r)
							return
						}
					}
				}
				metrics.Record(r, requestInfo, "", http.StatusTooManyRequests, 0, 0)
				tooManyRequests(r, w)
			}
		}
	})
}
