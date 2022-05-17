// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/project-radius/radius/pkg/corerp/servicecontext"
	"github.com/project-radius/radius/pkg/radrp/armerrors"
	"github.com/project-radius/radius/pkg/radrp/rest"
)

// ARMRequestCtx is the middleware to inject ARMRequestContext to the http request.
func ARMRequestCtx(pathBase string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			rpcContext, err := servicecontext.FromARMRequest(r, pathBase)
			if err != nil {
				resp := rest.NewBadRequestARMResponse(armerrors.ErrorResponse{
					Error: armerrors.ErrorDetails{
						Code:    strconv.Itoa(http.StatusBadRequest),
						Message: fmt.Sprintf("unexpected error: %v", err),
					},
				})

				_ = resp.Apply(r.Context(), w, r)
				return
			}

			r = r.WithContext(servicecontext.WithARMRequestContext(r.Context(), rpcContext))
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
