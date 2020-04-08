///
//  Generated code. Do not modify.
//  source: img.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'img.pb.dart' as $0;
export 'img.pb.dart';

class PictureServiceClient extends $grpc.Client {
  static final _$fetchPicture =
      $grpc.ClientMethod<$0.FetchPicturesRequest, $0.FetchPicturesResponse>(
          '/PictureService/FetchPicture',
          ($0.FetchPicturesRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.FetchPicturesResponse.fromBuffer(value));

  PictureServiceClient($grpc.ClientChannel channel, {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$0.FetchPicturesResponse> fetchPicture(
      $0.FetchPicturesRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$fetchPicture, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class PictureServiceBase extends $grpc.Service {
  $core.String get $name => 'PictureService';

  PictureServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$0.FetchPicturesRequest, $0.FetchPicturesResponse>(
            'FetchPicture',
            fetchPicture_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $0.FetchPicturesRequest.fromBuffer(value),
            ($0.FetchPicturesResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.FetchPicturesResponse> fetchPicture_Pre(
      $grpc.ServiceCall call,
      $async.Future<$0.FetchPicturesRequest> request) async {
    return fetchPicture(call, await request);
  }

  $async.Future<$0.FetchPicturesResponse> fetchPicture(
      $grpc.ServiceCall call, $0.FetchPicturesRequest request);
}
