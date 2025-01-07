/**
 * @fileoverview gRPC-Web generated client stub for viam.service.discovery.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: service/discovery/v1/discovery.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var app_v1_robot_pb = require('../../../app/v1/robot_pb.js')

var common_v1_common_pb = require('../../../common/v1/common_pb.js')

var google_api_annotations_pb = require('../../../google/api/annotations_pb.js')

var google_protobuf_struct_pb = require('google-protobuf/google/protobuf/struct_pb.js')
const proto = {};
proto.viam = {};
proto.viam.service = {};
proto.viam.service.discovery = {};
proto.viam.service.discovery.v1 = require('./discovery_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.viam.service.discovery.v1.DiscoveryServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.viam.service.discovery.v1.DiscoveryServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.service.discovery.v1.DiscoverResourcesRequest,
 *   !proto.viam.service.discovery.v1.DiscoverResourcesResponse>}
 */
const methodDescriptor_DiscoveryService_DiscoverResources = new grpc.web.MethodDescriptor(
  '/viam.service.discovery.v1.DiscoveryService/DiscoverResources',
  grpc.web.MethodType.UNARY,
  proto.viam.service.discovery.v1.DiscoverResourcesRequest,
  proto.viam.service.discovery.v1.DiscoverResourcesResponse,
  /**
   * @param {!proto.viam.service.discovery.v1.DiscoverResourcesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.viam.service.discovery.v1.DiscoverResourcesResponse.deserializeBinary
);


/**
 * @param {!proto.viam.service.discovery.v1.DiscoverResourcesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.service.discovery.v1.DiscoverResourcesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.service.discovery.v1.DiscoverResourcesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.service.discovery.v1.DiscoveryServiceClient.prototype.discoverResources =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.service.discovery.v1.DiscoveryService/DiscoverResources',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_DiscoverResources,
      callback);
};


/**
 * @param {!proto.viam.service.discovery.v1.DiscoverResourcesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.service.discovery.v1.DiscoverResourcesResponse>}
 *     Promise that resolves to the response
 */
proto.viam.service.discovery.v1.DiscoveryServicePromiseClient.prototype.discoverResources =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.service.discovery.v1.DiscoveryService/DiscoverResources',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_DiscoverResources);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.viam.common.v1.DoCommandRequest,
 *   !proto.viam.common.v1.DoCommandResponse>}
 */
const methodDescriptor_DiscoveryService_DoCommand = new grpc.web.MethodDescriptor(
  '/viam.service.discovery.v1.DiscoveryService/DoCommand',
  grpc.web.MethodType.UNARY,
  common_v1_common_pb.DoCommandRequest,
  common_v1_common_pb.DoCommandResponse,
  /**
   * @param {!proto.viam.common.v1.DoCommandRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  common_v1_common_pb.DoCommandResponse.deserializeBinary
);


/**
 * @param {!proto.viam.common.v1.DoCommandRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.viam.common.v1.DoCommandResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.viam.common.v1.DoCommandResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.viam.service.discovery.v1.DiscoveryServiceClient.prototype.doCommand =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/viam.service.discovery.v1.DiscoveryService/DoCommand',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_DoCommand,
      callback);
};


/**
 * @param {!proto.viam.common.v1.DoCommandRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.viam.common.v1.DoCommandResponse>}
 *     Promise that resolves to the response
 */
proto.viam.service.discovery.v1.DiscoveryServicePromiseClient.prototype.doCommand =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/viam.service.discovery.v1.DiscoveryService/DoCommand',
      request,
      metadata || {},
      methodDescriptor_DiscoveryService_DoCommand);
};


module.exports = proto.viam.service.discovery.v1;

