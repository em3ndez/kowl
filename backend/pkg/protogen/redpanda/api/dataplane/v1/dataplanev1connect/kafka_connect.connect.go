// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: redpanda/api/dataplane/v1/kafka_connect.proto

package dataplanev1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/dataplane/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// KafkaConnectServiceName is the fully-qualified name of the KafkaConnectService service.
	KafkaConnectServiceName = "redpanda.api.dataplane.v1.KafkaConnectService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// KafkaConnectServiceListConnectClustersProcedure is the fully-qualified name of the
	// KafkaConnectService's ListConnectClusters RPC.
	KafkaConnectServiceListConnectClustersProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/ListConnectClusters"
	// KafkaConnectServiceGetConnectClusterProcedure is the fully-qualified name of the
	// KafkaConnectService's GetConnectCluster RPC.
	KafkaConnectServiceGetConnectClusterProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/GetConnectCluster"
	// KafkaConnectServiceListConnectorsProcedure is the fully-qualified name of the
	// KafkaConnectService's ListConnectors RPC.
	KafkaConnectServiceListConnectorsProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/ListConnectors"
	// KafkaConnectServiceCreateConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's CreateConnector RPC.
	KafkaConnectServiceCreateConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/CreateConnector"
	// KafkaConnectServiceRestartConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's RestartConnector RPC.
	KafkaConnectServiceRestartConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/RestartConnector"
	// KafkaConnectServiceGetConnectorProcedure is the fully-qualified name of the KafkaConnectService's
	// GetConnector RPC.
	KafkaConnectServiceGetConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/GetConnector"
	// KafkaConnectServiceGetConnectorStatusProcedure is the fully-qualified name of the
	// KafkaConnectService's GetConnectorStatus RPC.
	KafkaConnectServiceGetConnectorStatusProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/GetConnectorStatus"
	// KafkaConnectServicePauseConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's PauseConnector RPC.
	KafkaConnectServicePauseConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/PauseConnector"
	// KafkaConnectServiceResumeConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's ResumeConnector RPC.
	KafkaConnectServiceResumeConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/ResumeConnector"
	// KafkaConnectServiceStopConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's StopConnector RPC.
	KafkaConnectServiceStopConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/StopConnector"
	// KafkaConnectServiceDeleteConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's DeleteConnector RPC.
	KafkaConnectServiceDeleteConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/DeleteConnector"
	// KafkaConnectServiceUpsertConnectorProcedure is the fully-qualified name of the
	// KafkaConnectService's UpsertConnector RPC.
	KafkaConnectServiceUpsertConnectorProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/UpsertConnector"
	// KafkaConnectServiceGetConnectorConfigProcedure is the fully-qualified name of the
	// KafkaConnectService's GetConnectorConfig RPC.
	KafkaConnectServiceGetConnectorConfigProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/GetConnectorConfig"
	// KafkaConnectServiceListConnectorTopicsProcedure is the fully-qualified name of the
	// KafkaConnectService's ListConnectorTopics RPC.
	KafkaConnectServiceListConnectorTopicsProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/ListConnectorTopics"
	// KafkaConnectServiceResetConnectorTopicsProcedure is the fully-qualified name of the
	// KafkaConnectService's ResetConnectorTopics RPC.
	KafkaConnectServiceResetConnectorTopicsProcedure = "/redpanda.api.dataplane.v1.KafkaConnectService/ResetConnectorTopics"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	kafkaConnectServiceServiceDescriptor                    = v1.File_redpanda_api_dataplane_v1_kafka_connect_proto.Services().ByName("KafkaConnectService")
	kafkaConnectServiceListConnectClustersMethodDescriptor  = kafkaConnectServiceServiceDescriptor.Methods().ByName("ListConnectClusters")
	kafkaConnectServiceGetConnectClusterMethodDescriptor    = kafkaConnectServiceServiceDescriptor.Methods().ByName("GetConnectCluster")
	kafkaConnectServiceListConnectorsMethodDescriptor       = kafkaConnectServiceServiceDescriptor.Methods().ByName("ListConnectors")
	kafkaConnectServiceCreateConnectorMethodDescriptor      = kafkaConnectServiceServiceDescriptor.Methods().ByName("CreateConnector")
	kafkaConnectServiceRestartConnectorMethodDescriptor     = kafkaConnectServiceServiceDescriptor.Methods().ByName("RestartConnector")
	kafkaConnectServiceGetConnectorMethodDescriptor         = kafkaConnectServiceServiceDescriptor.Methods().ByName("GetConnector")
	kafkaConnectServiceGetConnectorStatusMethodDescriptor   = kafkaConnectServiceServiceDescriptor.Methods().ByName("GetConnectorStatus")
	kafkaConnectServicePauseConnectorMethodDescriptor       = kafkaConnectServiceServiceDescriptor.Methods().ByName("PauseConnector")
	kafkaConnectServiceResumeConnectorMethodDescriptor      = kafkaConnectServiceServiceDescriptor.Methods().ByName("ResumeConnector")
	kafkaConnectServiceStopConnectorMethodDescriptor        = kafkaConnectServiceServiceDescriptor.Methods().ByName("StopConnector")
	kafkaConnectServiceDeleteConnectorMethodDescriptor      = kafkaConnectServiceServiceDescriptor.Methods().ByName("DeleteConnector")
	kafkaConnectServiceUpsertConnectorMethodDescriptor      = kafkaConnectServiceServiceDescriptor.Methods().ByName("UpsertConnector")
	kafkaConnectServiceGetConnectorConfigMethodDescriptor   = kafkaConnectServiceServiceDescriptor.Methods().ByName("GetConnectorConfig")
	kafkaConnectServiceListConnectorTopicsMethodDescriptor  = kafkaConnectServiceServiceDescriptor.Methods().ByName("ListConnectorTopics")
	kafkaConnectServiceResetConnectorTopicsMethodDescriptor = kafkaConnectServiceServiceDescriptor.Methods().ByName("ResetConnectorTopics")
)

// KafkaConnectServiceClient is a client for the redpanda.api.dataplane.v1.KafkaConnectService
// service.
type KafkaConnectServiceClient interface {
	// ListConnectClusters implements the list clusters method, list connect
	// clusters available in the console configuration
	ListConnectClusters(context.Context, *connect.Request[v1.ListConnectClustersRequest]) (*connect.Response[v1.ListConnectClustersResponse], error)
	// GetConnectCluster implements the get cluster info method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnectCluster(context.Context, *connect.Request[v1.GetConnectClusterRequest]) (*connect.Response[v1.GetConnectClusterResponse], error)
	// ListConnectors implements the list connectors method, exposes a Kafka
	// Connect equivalent REST endpoint
	ListConnectors(context.Context, *connect.Request[v1.ListConnectorsRequest]) (*connect.Response[v1.ListConnectorsResponse], error)
	// CreateConnector implements the create connector method, and exposes an
	// equivalent REST endpoint as the Kafka connect API endpoint
	CreateConnector(context.Context, *connect.Request[v1.CreateConnectorRequest]) (*connect.Response[v1.CreateConnectorResponse], error)
	// RestartConnector implements the restart connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	RestartConnector(context.Context, *connect.Request[v1.RestartConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// GetConnector implements the get connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnector(context.Context, *connect.Request[v1.GetConnectorRequest]) (*connect.Response[v1.GetConnectorResponse], error)
	// GetConnectorStatus implement the get status method, Gets the current status of the connector, including:
	// Whether it is running or restarting, or if it has failed or paused
	// Which worker it is assigned to
	// Error information if it has failed
	// The state of all its tasks
	GetConnectorStatus(context.Context, *connect.Request[v1.GetConnectorStatusRequest]) (*connect.Response[v1.GetConnectorStatusResponse], error)
	// PauseConnector implements the pause connector method, exposes a Kafka
	// connect equivalent REST endpoint
	PauseConnector(context.Context, *connect.Request[v1.PauseConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// ResumeConnector implements the resume connector method, exposes a Kafka
	// connect equivalent REST endpoint
	ResumeConnector(context.Context, *connect.Request[v1.ResumeConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// StopConnector implements the stop connector method, exposes a Kafka
	// connect equivalent REST endpoint it stops the connector but does not
	// delete the connector. All tasks for the connector are shut down completely
	StopConnector(context.Context, *connect.Request[v1.StopConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// DeleteConnector implements the delete connector method, exposes a Kafka
	// connect equivalent REST endpoint
	DeleteConnector(context.Context, *connect.Request[v1.DeleteConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// UpsertConector implements the update or create connector method, it
	// exposes a kafka connect equivalent REST endpoint
	UpsertConnector(context.Context, *connect.Request[v1.UpsertConnectorRequest]) (*connect.Response[v1.UpsertConnectorResponse], error)
	// GetConnectorConfig implements the get connector configuration method, expose a kafka connect equivalent REST endpoint
	GetConnectorConfig(context.Context, *connect.Request[v1.GetConnectorConfigRequest]) (*connect.Response[v1.GetConnectorConfigResponse], error)
	// ListConnectorTopics implements the list connector topics method, expose a kafka connect equivalent REST endpoint
	ListConnectorTopics(context.Context, *connect.Request[v1.ListConnectorTopicsRequest]) (*connect.Response[v1.ListConnectorTopicsResponse], error)
	// ResetConnectorTopics implements the reset connector topics method, expose a kafka connect equivalent REST endpoint
	// the request body is empty.
	ResetConnectorTopics(context.Context, *connect.Request[v1.ResetConnectorTopicsRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewKafkaConnectServiceClient constructs a client for the
// redpanda.api.dataplane.v1.KafkaConnectService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewKafkaConnectServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) KafkaConnectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &kafkaConnectServiceClient{
		listConnectClusters: connect.NewClient[v1.ListConnectClustersRequest, v1.ListConnectClustersResponse](
			httpClient,
			baseURL+KafkaConnectServiceListConnectClustersProcedure,
			connect.WithSchema(kafkaConnectServiceListConnectClustersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getConnectCluster: connect.NewClient[v1.GetConnectClusterRequest, v1.GetConnectClusterResponse](
			httpClient,
			baseURL+KafkaConnectServiceGetConnectClusterProcedure,
			connect.WithSchema(kafkaConnectServiceGetConnectClusterMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listConnectors: connect.NewClient[v1.ListConnectorsRequest, v1.ListConnectorsResponse](
			httpClient,
			baseURL+KafkaConnectServiceListConnectorsProcedure,
			connect.WithSchema(kafkaConnectServiceListConnectorsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createConnector: connect.NewClient[v1.CreateConnectorRequest, v1.CreateConnectorResponse](
			httpClient,
			baseURL+KafkaConnectServiceCreateConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceCreateConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		restartConnector: connect.NewClient[v1.RestartConnectorRequest, emptypb.Empty](
			httpClient,
			baseURL+KafkaConnectServiceRestartConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceRestartConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getConnector: connect.NewClient[v1.GetConnectorRequest, v1.GetConnectorResponse](
			httpClient,
			baseURL+KafkaConnectServiceGetConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceGetConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getConnectorStatus: connect.NewClient[v1.GetConnectorStatusRequest, v1.GetConnectorStatusResponse](
			httpClient,
			baseURL+KafkaConnectServiceGetConnectorStatusProcedure,
			connect.WithSchema(kafkaConnectServiceGetConnectorStatusMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		pauseConnector: connect.NewClient[v1.PauseConnectorRequest, emptypb.Empty](
			httpClient,
			baseURL+KafkaConnectServicePauseConnectorProcedure,
			connect.WithSchema(kafkaConnectServicePauseConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resumeConnector: connect.NewClient[v1.ResumeConnectorRequest, emptypb.Empty](
			httpClient,
			baseURL+KafkaConnectServiceResumeConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceResumeConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		stopConnector: connect.NewClient[v1.StopConnectorRequest, emptypb.Empty](
			httpClient,
			baseURL+KafkaConnectServiceStopConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceStopConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteConnector: connect.NewClient[v1.DeleteConnectorRequest, emptypb.Empty](
			httpClient,
			baseURL+KafkaConnectServiceDeleteConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceDeleteConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		upsertConnector: connect.NewClient[v1.UpsertConnectorRequest, v1.UpsertConnectorResponse](
			httpClient,
			baseURL+KafkaConnectServiceUpsertConnectorProcedure,
			connect.WithSchema(kafkaConnectServiceUpsertConnectorMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getConnectorConfig: connect.NewClient[v1.GetConnectorConfigRequest, v1.GetConnectorConfigResponse](
			httpClient,
			baseURL+KafkaConnectServiceGetConnectorConfigProcedure,
			connect.WithSchema(kafkaConnectServiceGetConnectorConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listConnectorTopics: connect.NewClient[v1.ListConnectorTopicsRequest, v1.ListConnectorTopicsResponse](
			httpClient,
			baseURL+KafkaConnectServiceListConnectorTopicsProcedure,
			connect.WithSchema(kafkaConnectServiceListConnectorTopicsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		resetConnectorTopics: connect.NewClient[v1.ResetConnectorTopicsRequest, emptypb.Empty](
			httpClient,
			baseURL+KafkaConnectServiceResetConnectorTopicsProcedure,
			connect.WithSchema(kafkaConnectServiceResetConnectorTopicsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// kafkaConnectServiceClient implements KafkaConnectServiceClient.
type kafkaConnectServiceClient struct {
	listConnectClusters  *connect.Client[v1.ListConnectClustersRequest, v1.ListConnectClustersResponse]
	getConnectCluster    *connect.Client[v1.GetConnectClusterRequest, v1.GetConnectClusterResponse]
	listConnectors       *connect.Client[v1.ListConnectorsRequest, v1.ListConnectorsResponse]
	createConnector      *connect.Client[v1.CreateConnectorRequest, v1.CreateConnectorResponse]
	restartConnector     *connect.Client[v1.RestartConnectorRequest, emptypb.Empty]
	getConnector         *connect.Client[v1.GetConnectorRequest, v1.GetConnectorResponse]
	getConnectorStatus   *connect.Client[v1.GetConnectorStatusRequest, v1.GetConnectorStatusResponse]
	pauseConnector       *connect.Client[v1.PauseConnectorRequest, emptypb.Empty]
	resumeConnector      *connect.Client[v1.ResumeConnectorRequest, emptypb.Empty]
	stopConnector        *connect.Client[v1.StopConnectorRequest, emptypb.Empty]
	deleteConnector      *connect.Client[v1.DeleteConnectorRequest, emptypb.Empty]
	upsertConnector      *connect.Client[v1.UpsertConnectorRequest, v1.UpsertConnectorResponse]
	getConnectorConfig   *connect.Client[v1.GetConnectorConfigRequest, v1.GetConnectorConfigResponse]
	listConnectorTopics  *connect.Client[v1.ListConnectorTopicsRequest, v1.ListConnectorTopicsResponse]
	resetConnectorTopics *connect.Client[v1.ResetConnectorTopicsRequest, emptypb.Empty]
}

// ListConnectClusters calls redpanda.api.dataplane.v1.KafkaConnectService.ListConnectClusters.
func (c *kafkaConnectServiceClient) ListConnectClusters(ctx context.Context, req *connect.Request[v1.ListConnectClustersRequest]) (*connect.Response[v1.ListConnectClustersResponse], error) {
	return c.listConnectClusters.CallUnary(ctx, req)
}

// GetConnectCluster calls redpanda.api.dataplane.v1.KafkaConnectService.GetConnectCluster.
func (c *kafkaConnectServiceClient) GetConnectCluster(ctx context.Context, req *connect.Request[v1.GetConnectClusterRequest]) (*connect.Response[v1.GetConnectClusterResponse], error) {
	return c.getConnectCluster.CallUnary(ctx, req)
}

// ListConnectors calls redpanda.api.dataplane.v1.KafkaConnectService.ListConnectors.
func (c *kafkaConnectServiceClient) ListConnectors(ctx context.Context, req *connect.Request[v1.ListConnectorsRequest]) (*connect.Response[v1.ListConnectorsResponse], error) {
	return c.listConnectors.CallUnary(ctx, req)
}

// CreateConnector calls redpanda.api.dataplane.v1.KafkaConnectService.CreateConnector.
func (c *kafkaConnectServiceClient) CreateConnector(ctx context.Context, req *connect.Request[v1.CreateConnectorRequest]) (*connect.Response[v1.CreateConnectorResponse], error) {
	return c.createConnector.CallUnary(ctx, req)
}

// RestartConnector calls redpanda.api.dataplane.v1.KafkaConnectService.RestartConnector.
func (c *kafkaConnectServiceClient) RestartConnector(ctx context.Context, req *connect.Request[v1.RestartConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.restartConnector.CallUnary(ctx, req)
}

// GetConnector calls redpanda.api.dataplane.v1.KafkaConnectService.GetConnector.
func (c *kafkaConnectServiceClient) GetConnector(ctx context.Context, req *connect.Request[v1.GetConnectorRequest]) (*connect.Response[v1.GetConnectorResponse], error) {
	return c.getConnector.CallUnary(ctx, req)
}

// GetConnectorStatus calls redpanda.api.dataplane.v1.KafkaConnectService.GetConnectorStatus.
func (c *kafkaConnectServiceClient) GetConnectorStatus(ctx context.Context, req *connect.Request[v1.GetConnectorStatusRequest]) (*connect.Response[v1.GetConnectorStatusResponse], error) {
	return c.getConnectorStatus.CallUnary(ctx, req)
}

// PauseConnector calls redpanda.api.dataplane.v1.KafkaConnectService.PauseConnector.
func (c *kafkaConnectServiceClient) PauseConnector(ctx context.Context, req *connect.Request[v1.PauseConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.pauseConnector.CallUnary(ctx, req)
}

// ResumeConnector calls redpanda.api.dataplane.v1.KafkaConnectService.ResumeConnector.
func (c *kafkaConnectServiceClient) ResumeConnector(ctx context.Context, req *connect.Request[v1.ResumeConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.resumeConnector.CallUnary(ctx, req)
}

// StopConnector calls redpanda.api.dataplane.v1.KafkaConnectService.StopConnector.
func (c *kafkaConnectServiceClient) StopConnector(ctx context.Context, req *connect.Request[v1.StopConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.stopConnector.CallUnary(ctx, req)
}

// DeleteConnector calls redpanda.api.dataplane.v1.KafkaConnectService.DeleteConnector.
func (c *kafkaConnectServiceClient) DeleteConnector(ctx context.Context, req *connect.Request[v1.DeleteConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteConnector.CallUnary(ctx, req)
}

// UpsertConnector calls redpanda.api.dataplane.v1.KafkaConnectService.UpsertConnector.
func (c *kafkaConnectServiceClient) UpsertConnector(ctx context.Context, req *connect.Request[v1.UpsertConnectorRequest]) (*connect.Response[v1.UpsertConnectorResponse], error) {
	return c.upsertConnector.CallUnary(ctx, req)
}

// GetConnectorConfig calls redpanda.api.dataplane.v1.KafkaConnectService.GetConnectorConfig.
func (c *kafkaConnectServiceClient) GetConnectorConfig(ctx context.Context, req *connect.Request[v1.GetConnectorConfigRequest]) (*connect.Response[v1.GetConnectorConfigResponse], error) {
	return c.getConnectorConfig.CallUnary(ctx, req)
}

// ListConnectorTopics calls redpanda.api.dataplane.v1.KafkaConnectService.ListConnectorTopics.
func (c *kafkaConnectServiceClient) ListConnectorTopics(ctx context.Context, req *connect.Request[v1.ListConnectorTopicsRequest]) (*connect.Response[v1.ListConnectorTopicsResponse], error) {
	return c.listConnectorTopics.CallUnary(ctx, req)
}

// ResetConnectorTopics calls redpanda.api.dataplane.v1.KafkaConnectService.ResetConnectorTopics.
func (c *kafkaConnectServiceClient) ResetConnectorTopics(ctx context.Context, req *connect.Request[v1.ResetConnectorTopicsRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.resetConnectorTopics.CallUnary(ctx, req)
}

// KafkaConnectServiceHandler is an implementation of the
// redpanda.api.dataplane.v1.KafkaConnectService service.
type KafkaConnectServiceHandler interface {
	// ListConnectClusters implements the list clusters method, list connect
	// clusters available in the console configuration
	ListConnectClusters(context.Context, *connect.Request[v1.ListConnectClustersRequest]) (*connect.Response[v1.ListConnectClustersResponse], error)
	// GetConnectCluster implements the get cluster info method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnectCluster(context.Context, *connect.Request[v1.GetConnectClusterRequest]) (*connect.Response[v1.GetConnectClusterResponse], error)
	// ListConnectors implements the list connectors method, exposes a Kafka
	// Connect equivalent REST endpoint
	ListConnectors(context.Context, *connect.Request[v1.ListConnectorsRequest]) (*connect.Response[v1.ListConnectorsResponse], error)
	// CreateConnector implements the create connector method, and exposes an
	// equivalent REST endpoint as the Kafka connect API endpoint
	CreateConnector(context.Context, *connect.Request[v1.CreateConnectorRequest]) (*connect.Response[v1.CreateConnectorResponse], error)
	// RestartConnector implements the restart connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	RestartConnector(context.Context, *connect.Request[v1.RestartConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// GetConnector implements the get connector method, exposes a Kafka
	// Connect equivalent REST endpoint
	GetConnector(context.Context, *connect.Request[v1.GetConnectorRequest]) (*connect.Response[v1.GetConnectorResponse], error)
	// GetConnectorStatus implement the get status method, Gets the current status of the connector, including:
	// Whether it is running or restarting, or if it has failed or paused
	// Which worker it is assigned to
	// Error information if it has failed
	// The state of all its tasks
	GetConnectorStatus(context.Context, *connect.Request[v1.GetConnectorStatusRequest]) (*connect.Response[v1.GetConnectorStatusResponse], error)
	// PauseConnector implements the pause connector method, exposes a Kafka
	// connect equivalent REST endpoint
	PauseConnector(context.Context, *connect.Request[v1.PauseConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// ResumeConnector implements the resume connector method, exposes a Kafka
	// connect equivalent REST endpoint
	ResumeConnector(context.Context, *connect.Request[v1.ResumeConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// StopConnector implements the stop connector method, exposes a Kafka
	// connect equivalent REST endpoint it stops the connector but does not
	// delete the connector. All tasks for the connector are shut down completely
	StopConnector(context.Context, *connect.Request[v1.StopConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// DeleteConnector implements the delete connector method, exposes a Kafka
	// connect equivalent REST endpoint
	DeleteConnector(context.Context, *connect.Request[v1.DeleteConnectorRequest]) (*connect.Response[emptypb.Empty], error)
	// UpsertConector implements the update or create connector method, it
	// exposes a kafka connect equivalent REST endpoint
	UpsertConnector(context.Context, *connect.Request[v1.UpsertConnectorRequest]) (*connect.Response[v1.UpsertConnectorResponse], error)
	// GetConnectorConfig implements the get connector configuration method, expose a kafka connect equivalent REST endpoint
	GetConnectorConfig(context.Context, *connect.Request[v1.GetConnectorConfigRequest]) (*connect.Response[v1.GetConnectorConfigResponse], error)
	// ListConnectorTopics implements the list connector topics method, expose a kafka connect equivalent REST endpoint
	ListConnectorTopics(context.Context, *connect.Request[v1.ListConnectorTopicsRequest]) (*connect.Response[v1.ListConnectorTopicsResponse], error)
	// ResetConnectorTopics implements the reset connector topics method, expose a kafka connect equivalent REST endpoint
	// the request body is empty.
	ResetConnectorTopics(context.Context, *connect.Request[v1.ResetConnectorTopicsRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewKafkaConnectServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewKafkaConnectServiceHandler(svc KafkaConnectServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	kafkaConnectServiceListConnectClustersHandler := connect.NewUnaryHandler(
		KafkaConnectServiceListConnectClustersProcedure,
		svc.ListConnectClusters,
		connect.WithSchema(kafkaConnectServiceListConnectClustersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceGetConnectClusterHandler := connect.NewUnaryHandler(
		KafkaConnectServiceGetConnectClusterProcedure,
		svc.GetConnectCluster,
		connect.WithSchema(kafkaConnectServiceGetConnectClusterMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceListConnectorsHandler := connect.NewUnaryHandler(
		KafkaConnectServiceListConnectorsProcedure,
		svc.ListConnectors,
		connect.WithSchema(kafkaConnectServiceListConnectorsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceCreateConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceCreateConnectorProcedure,
		svc.CreateConnector,
		connect.WithSchema(kafkaConnectServiceCreateConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceRestartConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceRestartConnectorProcedure,
		svc.RestartConnector,
		connect.WithSchema(kafkaConnectServiceRestartConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceGetConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceGetConnectorProcedure,
		svc.GetConnector,
		connect.WithSchema(kafkaConnectServiceGetConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceGetConnectorStatusHandler := connect.NewUnaryHandler(
		KafkaConnectServiceGetConnectorStatusProcedure,
		svc.GetConnectorStatus,
		connect.WithSchema(kafkaConnectServiceGetConnectorStatusMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServicePauseConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServicePauseConnectorProcedure,
		svc.PauseConnector,
		connect.WithSchema(kafkaConnectServicePauseConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceResumeConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceResumeConnectorProcedure,
		svc.ResumeConnector,
		connect.WithSchema(kafkaConnectServiceResumeConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceStopConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceStopConnectorProcedure,
		svc.StopConnector,
		connect.WithSchema(kafkaConnectServiceStopConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceDeleteConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceDeleteConnectorProcedure,
		svc.DeleteConnector,
		connect.WithSchema(kafkaConnectServiceDeleteConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceUpsertConnectorHandler := connect.NewUnaryHandler(
		KafkaConnectServiceUpsertConnectorProcedure,
		svc.UpsertConnector,
		connect.WithSchema(kafkaConnectServiceUpsertConnectorMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceGetConnectorConfigHandler := connect.NewUnaryHandler(
		KafkaConnectServiceGetConnectorConfigProcedure,
		svc.GetConnectorConfig,
		connect.WithSchema(kafkaConnectServiceGetConnectorConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceListConnectorTopicsHandler := connect.NewUnaryHandler(
		KafkaConnectServiceListConnectorTopicsProcedure,
		svc.ListConnectorTopics,
		connect.WithSchema(kafkaConnectServiceListConnectorTopicsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	kafkaConnectServiceResetConnectorTopicsHandler := connect.NewUnaryHandler(
		KafkaConnectServiceResetConnectorTopicsProcedure,
		svc.ResetConnectorTopics,
		connect.WithSchema(kafkaConnectServiceResetConnectorTopicsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/redpanda.api.dataplane.v1.KafkaConnectService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case KafkaConnectServiceListConnectClustersProcedure:
			kafkaConnectServiceListConnectClustersHandler.ServeHTTP(w, r)
		case KafkaConnectServiceGetConnectClusterProcedure:
			kafkaConnectServiceGetConnectClusterHandler.ServeHTTP(w, r)
		case KafkaConnectServiceListConnectorsProcedure:
			kafkaConnectServiceListConnectorsHandler.ServeHTTP(w, r)
		case KafkaConnectServiceCreateConnectorProcedure:
			kafkaConnectServiceCreateConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceRestartConnectorProcedure:
			kafkaConnectServiceRestartConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceGetConnectorProcedure:
			kafkaConnectServiceGetConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceGetConnectorStatusProcedure:
			kafkaConnectServiceGetConnectorStatusHandler.ServeHTTP(w, r)
		case KafkaConnectServicePauseConnectorProcedure:
			kafkaConnectServicePauseConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceResumeConnectorProcedure:
			kafkaConnectServiceResumeConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceStopConnectorProcedure:
			kafkaConnectServiceStopConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceDeleteConnectorProcedure:
			kafkaConnectServiceDeleteConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceUpsertConnectorProcedure:
			kafkaConnectServiceUpsertConnectorHandler.ServeHTTP(w, r)
		case KafkaConnectServiceGetConnectorConfigProcedure:
			kafkaConnectServiceGetConnectorConfigHandler.ServeHTTP(w, r)
		case KafkaConnectServiceListConnectorTopicsProcedure:
			kafkaConnectServiceListConnectorTopicsHandler.ServeHTTP(w, r)
		case KafkaConnectServiceResetConnectorTopicsProcedure:
			kafkaConnectServiceResetConnectorTopicsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedKafkaConnectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedKafkaConnectServiceHandler struct{}

func (UnimplementedKafkaConnectServiceHandler) ListConnectClusters(context.Context, *connect.Request[v1.ListConnectClustersRequest]) (*connect.Response[v1.ListConnectClustersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.ListConnectClusters is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) GetConnectCluster(context.Context, *connect.Request[v1.GetConnectClusterRequest]) (*connect.Response[v1.GetConnectClusterResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.GetConnectCluster is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) ListConnectors(context.Context, *connect.Request[v1.ListConnectorsRequest]) (*connect.Response[v1.ListConnectorsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.ListConnectors is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) CreateConnector(context.Context, *connect.Request[v1.CreateConnectorRequest]) (*connect.Response[v1.CreateConnectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.CreateConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) RestartConnector(context.Context, *connect.Request[v1.RestartConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.RestartConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) GetConnector(context.Context, *connect.Request[v1.GetConnectorRequest]) (*connect.Response[v1.GetConnectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.GetConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) GetConnectorStatus(context.Context, *connect.Request[v1.GetConnectorStatusRequest]) (*connect.Response[v1.GetConnectorStatusResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.GetConnectorStatus is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) PauseConnector(context.Context, *connect.Request[v1.PauseConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.PauseConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) ResumeConnector(context.Context, *connect.Request[v1.ResumeConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.ResumeConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) StopConnector(context.Context, *connect.Request[v1.StopConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.StopConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) DeleteConnector(context.Context, *connect.Request[v1.DeleteConnectorRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.DeleteConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) UpsertConnector(context.Context, *connect.Request[v1.UpsertConnectorRequest]) (*connect.Response[v1.UpsertConnectorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.UpsertConnector is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) GetConnectorConfig(context.Context, *connect.Request[v1.GetConnectorConfigRequest]) (*connect.Response[v1.GetConnectorConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.GetConnectorConfig is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) ListConnectorTopics(context.Context, *connect.Request[v1.ListConnectorTopicsRequest]) (*connect.Response[v1.ListConnectorTopicsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.ListConnectorTopics is not implemented"))
}

func (UnimplementedKafkaConnectServiceHandler) ResetConnectorTopics(context.Context, *connect.Request[v1.ResetConnectorTopicsRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("redpanda.api.dataplane.v1.KafkaConnectService.ResetConnectorTopics is not implemented"))
}
