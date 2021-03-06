package batchaiapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/batchai/mgmt/2018-05-01/batchai"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result batchai.OperationListResultPage, err error)
}

var _ OperationsClientAPI = (*batchai.OperationsClient)(nil)

// UsagesClientAPI contains the set of methods on the UsagesClient type.
type UsagesClientAPI interface {
	List(ctx context.Context, location string) (result batchai.ListUsagesResultPage, err error)
}

var _ UsagesClientAPI = (*batchai.UsagesClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, parameters batchai.WorkspaceCreateParameters) (result batchai.WorkspacesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result batchai.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result batchai.Workspace, err error)
	List(ctx context.Context, maxResults *int32) (result batchai.WorkspaceListResultPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, maxResults *int32) (result batchai.WorkspaceListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters batchai.WorkspaceUpdateParameters) (result batchai.Workspace, err error)
}

var _ WorkspacesClientAPI = (*batchai.WorkspacesClient)(nil)

// ExperimentsClientAPI contains the set of methods on the ExperimentsClient type.
type ExperimentsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (result batchai.ExperimentsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (result batchai.ExperimentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string) (result batchai.Experiment, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result batchai.ExperimentListResultPage, err error)
}

var _ ExperimentsClientAPI = (*batchai.ExperimentsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, jobName string, parameters batchai.JobCreateParameters) (result batchai.JobsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, jobName string) (result batchai.JobsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, jobName string) (result batchai.Job, err error)
	ListByExperiment(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, maxResults *int32) (result batchai.JobListResultPage, err error)
	ListOutputFiles(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, jobName string, outputdirectoryid string, directory string, linkexpiryinminutes *int32, maxResults *int32) (result batchai.FileListResultPage, err error)
	ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, jobName string) (result batchai.RemoteLoginInformationListResultPage, err error)
	Terminate(ctx context.Context, resourceGroupName string, workspaceName string, experimentName string, jobName string) (result batchai.JobsTerminateFuture, err error)
}

var _ JobsClientAPI = (*batchai.JobsClient)(nil)

// FileServersClientAPI contains the set of methods on the FileServersClient type.
type FileServersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, fileServerName string, parameters batchai.FileServerCreateParameters) (result batchai.FileServersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, fileServerName string) (result batchai.FileServersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, fileServerName string) (result batchai.FileServer, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result batchai.FileServerListResultPage, err error)
}

var _ FileServersClientAPI = (*batchai.FileServersClient)(nil)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string, parameters batchai.ClusterCreateParameters) (result batchai.ClustersCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result batchai.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result batchai.Cluster, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, maxResults *int32) (result batchai.ClusterListResultPage, err error)
	ListRemoteLoginInformation(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string) (result batchai.RemoteLoginInformationListResultPage, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, clusterName string, parameters batchai.ClusterUpdateParameters) (result batchai.Cluster, err error)
}

var _ ClustersClientAPI = (*batchai.ClustersClient)(nil)
