// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// User defines model for user.
type User struct {
	// CreatedAt CreatedAt
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Id User ID
	Id   *openapi_types.UUID `json:"id,omitempty"`
	Info *UserInfo           `json:"info,omitempty"`

	// Name User name
	Name *string `json:"name,omitempty"`

	// UpdatedAt UpdatedAt
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// UserInfo defines model for userInfo.
type UserInfo struct {
	// Tel Telephone number
	Tel *string `json:"tel,omitempty"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody struct {
	// CreatedAt CreatedAt
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Id User ID
	Id   *openapi_types.UUID `json:"id,omitempty"`
	Info interface{}         `json:"info"`

	// Name User name
	Name string `json:"name"`

	// UpdatedAt UpdatedAt
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RUzW7bPBB8FWO/7yhZlCw7Dm/9AYq0h1ySQxsYBS2uYwYSyfDHaGDw3QtSUWLVMdwe",
	"ivpgWzvkzu7MrvbQqE4ridJZoHuwzRY7lv56iyb+aqM0GicwRRuDzCH/zlx84mgbI7QTSgKFDz32zkEG",
	"+IN1ukWgUJFqnpMyJ9UNmVFSUzKfLi6Wl6SsZvU3yGCjTBfTAWcOcyc6hAzck46XrTNC3kPIQPBjwluL",
	"ZnL1cURXVjOs54uLHJeX67ys+Cxn9XyR19ViUdblRU0IOST1XvA3+eRGRcb/DW6Awn/Fq1DFs0pFlOgq",
	"ngsZSNbhiQoTdFjjZ9Y8epx89W926jU/KfFtj/0NicNLRK0fsHGplKHBozlw2B5Xd4Mt6q2SOJG+W6MZ",
	"FUnmJCeE9F+/Qx8OXBjz4E7I+63yNrUhXMq/UWrNIuUOje3PlVMSu1AaJdMCKMymZDqDDDRz29RGwbQo",
	"dmVyMgW0sicne5JWIuUzLCJX/AW77SGDjx6te6/4U9oWJR3KlJBp3YomXSsebMw6bFtC2/Z6A/Tu/LxB",
	"yPaJRhjkQO/gebqSUqsIjo0aFDy8Es1bhRBWSeNXxBmPKWC1kra/XxHyR52crz9yjuW9/hKjIRvbUezT",
	"+PEQ097jG7aw4XNkyid0z45oZliHLtl7d/oFIuJjHAsYVrkffg6/CpQdNHvmLRJW/0rMDKzvOmaeei0m",
	"r4ctmt0ghjctUNg6p2lRtKph7VZZR5dkeQlhFX4GAAD//8fG/IMdBgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
