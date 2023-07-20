package types

import "net/http"

type RequestModifier func(r *http.Request)