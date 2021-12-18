/*
 * Task Execution Service
 *
 * ## Executive Summary The Task Execution Service (TES) API is a standardized schema and API for describing and executing batch execution tasks. A task defines a set of input files, a set of containers and commands to run, a set of output files and some other logging and metadata.  TES servers accept task documents and execute them asynchronously on available compute resources. A TES server could be built on top of a traditional HPC queuing system, such as Grid Engine, Slurm or cloud style compute systems such as AWS Batch or Kubernetes. ## Introduction This document describes the TES API and provides details on the specific endpoints, request formats, and responses. It is intended to provide key information for developers of TES-compatible services as well as clients that will call these TES services. Use cases include:    - Deploying existing workflow engines on new infrastructure. Workflow engines   such as CWL-Tes and Cromwell have extentions for using TES. This will allow   a system engineer to deploy them onto a new infrastructure using a job scheduling   system not previously supported by the engine.    - Developing a custom workflow management system. This API provides a common   interface to asynchronous batch processing capabilities. A developer can write   new tools against this interface and expect them to work using a variety of   backend solutions that all support the same specification.   ## Standards The TES API specification is written in OpenAPI and embodies a RESTful service philosophy. It uses JSON in requests and responses and standard HTTP/HTTPS for information transport. HTTPS should be used rather than plain HTTP except for testing or internal-only purposes. ### Authentication and Authorization Is is envisaged that most TES API instances will require users to authenticate to use the endpoints. However, the decision if authentication is required should be taken by TES API implementers.  If authentication is required, we recommend that TES implementations use an OAuth2  bearer token, although they can choose other mechanisms if appropriate.  Checking that a user is authorized to submit TES requests is a responsibility of TES implementations. ### CORS If TES API implementation is to be used by another website or domain it must implement Cross Origin Resource Sharing (CORS). Please refer to https://w3id.org/ga4gh/product-approval-support/cors for more information about GA4GH’s recommendations and how to implement CORS. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// TaskServiceApiService is a service that implents the logic for the TaskServiceApiServicer
// This service should implement the business logic for every endpoint for the TaskServiceApi API. 
// Include any external packages or services that will be required by this service.
type TaskServiceApiService struct {
}

// NewTaskServiceApiService creates a default api service
func NewTaskServiceApiService() TaskServiceApiServicer {
	return &TaskServiceApiService{}
}

// CancelTask - CancelTask
func (s *TaskServiceApiService) CancelTask(ctx context.Context, id string) (ImplResponse, error) {
	// TODO - update CancelTask with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, map[string]interface{}{}) or use other options such as http.Ok ...
	//return Response(200, map[string]interface{}{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CancelTask method not implemented")
}

// CreateTask - CreateTask
func (s *TaskServiceApiService) CreateTask(ctx context.Context, body TesTask) (ImplResponse, error) {
	// TODO - update CreateTask with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesCreateTaskResponse{}) or use other options such as http.Ok ...
	//return Response(200, TesCreateTaskResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateTask method not implemented")
}

// GetServiceInfo - GetServiceInfo
func (s *TaskServiceApiService) GetServiceInfo(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetServiceInfo with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesServiceInfo{}) or use other options such as http.Ok ...
	//return Response(200, TesServiceInfo{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetServiceInfo method not implemented")
}

// GetTask - GetTask
func (s *TaskServiceApiService) GetTask(ctx context.Context, id string, view string) (ImplResponse, error) {
	// TODO - update GetTask with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesTask{}) or use other options such as http.Ok ...
	//return Response(200, TesTask{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetTask method not implemented")
}

// ListTasks - ListTasks
func (s *TaskServiceApiService) ListTasks(ctx context.Context, namePrefix string, state TesState, tags map[string]string, pageSize int64, pageToken string, view string) (ImplResponse, error) {
	// TODO - update ListTasks with the required logic for this service method.
	// Add api_task_service_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, TesListTasksResponse{}) or use other options such as http.Ok ...
	//return Response(200, TesListTasksResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ListTasks method not implemented")
}

