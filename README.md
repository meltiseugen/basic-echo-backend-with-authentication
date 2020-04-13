# Basic-Echo-backend-with-authentication

### 1. Description

This is a simple GoLang server implemented using Echo meant to be a starting point when
implementing a backend.

### 2. Functionality

It offers the following functionality:
* JWT authentication that can be configured inside the code.
Also you can define which routes to be excluded from the check by changing `auth.Excluded` 
```go
e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &auth.JWTClaims{},
		SigningKey:    []byte("<CHANGE ME>"),
		SigningMethod: "HS512",
		Skipper: func(c echo.Context) bool {
			return utils.Contains(auth.Excluded, c.Path())
		},
	}))
```

* CORS implementation. Consider changing `AllowOrigins` in production. 
```go
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
    Skipper: func(c echo.Context) bool {
        return true
    },
}))
```

* CSRF implementation to prevent cross site request forgery.
```go
e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
    TokenLength:  32,
    TokenLookup:  "header:" + echo.HeaderXCSRFToken,
    ContextKey:   "csrf",
    CookieName:   "_csrf",
    CookieMaxAge: 86400,
    Skipper: func(c echo.Context) bool {
        return true
    },
}))
```

* Logging of information.
```go
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: "method=${method}, uri=${uri}, status=${status}\n",
}))
```

### 3. Basic Routes

Below you can find the basic routes of the server:
```go
INDEX         = "/"
METRICS       = INDEX + "metrics"
ROOT          = INDEX + "v1/"

MESSAGE = ROOT + "message"
```

### 4. Starting the service

In order ot start the service run the file `cmd/backend/main.go`. Please note
that the server runs on port `:1323`
