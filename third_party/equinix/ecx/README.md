## Modified Swagger spec for ECX Fabric

* Fixed token_timeout field to allow string (yep, API can answer with a string instead of a int)
* Fixed Metro API response object _GETCommonMetroResp_ to represent the correct API response model 

:warning: if you plan to generate SDK from spec file remember to change _models/o_auth_response.go_ TokenTikmeout definition

````go
        // Add string to definition
        TokenTimeout int64 `json:"token_timeout, string,omitempty"`
````