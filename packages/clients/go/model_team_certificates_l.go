/*
 * Rainbond Open API
 *
 * Rainbond open api
 *
 * API version: v1
 * Contact: barnett@goodrain.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package openapi

type TeamCertificatesL struct {
	List []CertificatesR `json:"list"`
	Page int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Total int32 `json:"total"`
}
