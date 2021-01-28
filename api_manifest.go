/*
 * Quay Frontend
 *
 * This API allows you to perform many of the operations required to work with Quay repositories, users, and organizations. You can find out more at <a href=\"https://quay.io\">Quay</a>.
 *
 * API version: v1
 * Contact: support@quay.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ManifestApiService ManifestApi service
type ManifestApiService service

type ApiAddManifestLabelRequest struct {
	ctx         _context.Context
	ApiService  *ManifestApiService
	manifestref string
	repository  string
	body        *AddLabel
}

func (r ApiAddManifestLabelRequest) Body(body AddLabel) ApiAddManifestLabelRequest {
	r.body = &body
	return r
}

func (r ApiAddManifestLabelRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AddManifestLabelExecute(r)
}

/*
 * AddManifestLabel Method for AddManifestLabel
 * Adds a new label into the tag manifest.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param manifestref The digest of the manifest
 * @param repository The full path of the repository. e.g. namespace/name
 * @return ApiAddManifestLabelRequest
 */
func (a *ManifestApiService) AddManifestLabel(ctx _context.Context, manifestref string, repository string) ApiAddManifestLabelRequest {
	return ApiAddManifestLabelRequest{
		ApiService:  a,
		ctx:         ctx,
		manifestref: manifestref,
		repository:  repository,
	}
}

/*
 * Execute executes the request
 */
func (a *ManifestApiService) AddManifestLabelExecute(r ApiAddManifestLabelRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManifestApiService.AddManifestLabel")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repository/{repository}/manifest/{manifestref}/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"manifestref"+"}", _neturl.PathEscape(parameterToString(r.manifestref, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository"+"}", _neturl.PathEscape(parameterToString(r.repository, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteManifestLabelRequest struct {
	ctx         _context.Context
	ApiService  *ManifestApiService
	labelid     string
	manifestref string
	repository  string
}

func (r ApiDeleteManifestLabelRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteManifestLabelExecute(r)
}

/*
 * DeleteManifestLabel Method for DeleteManifestLabel
 * Deletes an existing label from a manifest.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param labelid The ID of the label
 * @param manifestref The digest of the manifest
 * @param repository The full path of the repository. e.g. namespace/name
 * @return ApiDeleteManifestLabelRequest
 */
func (a *ManifestApiService) DeleteManifestLabel(ctx _context.Context, labelid string, manifestref string, repository string) ApiDeleteManifestLabelRequest {
	return ApiDeleteManifestLabelRequest{
		ApiService:  a,
		ctx:         ctx,
		labelid:     labelid,
		manifestref: manifestref,
		repository:  repository,
	}
}

/*
 * Execute executes the request
 */
func (a *ManifestApiService) DeleteManifestLabelExecute(r ApiDeleteManifestLabelRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManifestApiService.DeleteManifestLabel")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}"
	localVarPath = strings.Replace(localVarPath, "{"+"labelid"+"}", _neturl.PathEscape(parameterToString(r.labelid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"manifestref"+"}", _neturl.PathEscape(parameterToString(r.manifestref, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository"+"}", _neturl.PathEscape(parameterToString(r.repository, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetManifestLabelRequest struct {
	ctx         _context.Context
	ApiService  *ManifestApiService
	labelid     string
	manifestref string
	repository  string
}

func (r ApiGetManifestLabelRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetManifestLabelExecute(r)
}

/*
 * GetManifestLabel Method for GetManifestLabel
 * Retrieves the label with the specific ID under the manifest.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param labelid The ID of the label
 * @param manifestref The digest of the manifest
 * @param repository The full path of the repository. e.g. namespace/name
 * @return ApiGetManifestLabelRequest
 */
func (a *ManifestApiService) GetManifestLabel(ctx _context.Context, labelid string, manifestref string, repository string) ApiGetManifestLabelRequest {
	return ApiGetManifestLabelRequest{
		ApiService:  a,
		ctx:         ctx,
		labelid:     labelid,
		manifestref: manifestref,
		repository:  repository,
	}
}

/*
 * Execute executes the request
 */
func (a *ManifestApiService) GetManifestLabelExecute(r ApiGetManifestLabelRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManifestApiService.GetManifestLabel")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid}"
	localVarPath = strings.Replace(localVarPath, "{"+"labelid"+"}", _neturl.PathEscape(parameterToString(r.labelid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"manifestref"+"}", _neturl.PathEscape(parameterToString(r.manifestref, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository"+"}", _neturl.PathEscape(parameterToString(r.repository, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetRepoManifestRequest struct {
	ctx         _context.Context
	ApiService  *ManifestApiService
	manifestref string
	repository  string
}

func (r ApiGetRepoManifestRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetRepoManifestExecute(r)
}

/*
 * GetRepoManifest Method for GetRepoManifest
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param manifestref The digest of the manifest
 * @param repository The full path of the repository. e.g. namespace/name
 * @return ApiGetRepoManifestRequest
 */
func (a *ManifestApiService) GetRepoManifest(ctx _context.Context, manifestref string, repository string) ApiGetRepoManifestRequest {
	return ApiGetRepoManifestRequest{
		ApiService:  a,
		ctx:         ctx,
		manifestref: manifestref,
		repository:  repository,
	}
}

/*
 * Execute executes the request
 */
func (a *ManifestApiService) GetRepoManifestExecute(r ApiGetRepoManifestRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManifestApiService.GetRepoManifest")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repository/{repository}/manifest/{manifestref}"
	localVarPath = strings.Replace(localVarPath, "{"+"manifestref"+"}", _neturl.PathEscape(parameterToString(r.manifestref, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository"+"}", _neturl.PathEscape(parameterToString(r.repository, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListManifestLabelsRequest struct {
	ctx         _context.Context
	ApiService  *ManifestApiService
	manifestref string
	repository  string
	filter      *string
}

func (r ApiListManifestLabelsRequest) Filter(filter string) ApiListManifestLabelsRequest {
	r.filter = &filter
	return r
}

func (r ApiListManifestLabelsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ListManifestLabelsExecute(r)
}

/*
 * ListManifestLabels Method for ListManifestLabels
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param manifestref The digest of the manifest
 * @param repository The full path of the repository. e.g. namespace/name
 * @return ApiListManifestLabelsRequest
 */
func (a *ManifestApiService) ListManifestLabels(ctx _context.Context, manifestref string, repository string) ApiListManifestLabelsRequest {
	return ApiListManifestLabelsRequest{
		ApiService:  a,
		ctx:         ctx,
		manifestref: manifestref,
		repository:  repository,
	}
}

/*
 * Execute executes the request
 */
func (a *ManifestApiService) ListManifestLabelsExecute(r ApiListManifestLabelsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ManifestApiService.ListManifestLabels")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/repository/{repository}/manifest/{manifestref}/labels"
	localVarPath = strings.Replace(localVarPath, "{"+"manifestref"+"}", _neturl.PathEscape(parameterToString(r.manifestref, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repository"+"}", _neturl.PathEscape(parameterToString(r.repository, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
