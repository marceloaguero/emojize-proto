/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./emoji_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.EmojiServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.EmojiServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.EmojiRequest,
 *   !proto.proto.EmojiResponse>}
 */
const methodDescriptor_EmojiService_InsertEmojis = new grpc.web.MethodDescriptor(
  '/proto.EmojiService/InsertEmojis',
  grpc.web.MethodType.UNARY,
  proto.proto.EmojiRequest,
  proto.proto.EmojiResponse,
  /**
   * @param {!proto.proto.EmojiRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.EmojiResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.proto.EmojiRequest,
 *   !proto.proto.EmojiResponse>}
 */
const methodInfo_EmojiService_InsertEmojis = new grpc.web.AbstractClientBase.MethodInfo(
  proto.proto.EmojiResponse,
  /**
   * @param {!proto.proto.EmojiRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.EmojiResponse.deserializeBinary
);


/**
 * @param {!proto.proto.EmojiRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.proto.EmojiResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.EmojiResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.EmojiServiceClient.prototype.insertEmojis =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.EmojiService/InsertEmojis',
      request,
      metadata || {},
      methodDescriptor_EmojiService_InsertEmojis,
      callback);
};


/**
 * @param {!proto.proto.EmojiRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.EmojiResponse>}
 *     A native promise that resolves to the response
 */
proto.proto.EmojiServicePromiseClient.prototype.insertEmojis =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.EmojiService/InsertEmojis',
      request,
      metadata || {},
      methodDescriptor_EmojiService_InsertEmojis);
};


module.exports = proto.proto;

