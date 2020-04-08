///
//  Generated code. Do not modify.
//  source: img.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'img.pbenum.dart';

export 'img.pbenum.dart';

class FetchPicturesRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FetchPicturesRequest', createEmptyInstance: create)
    ..aOS(1, 'UUID', protoName: 'UUID')
    ..a<$core.int>(2, 'Pos', $pb.PbFieldType.O3, protoName: 'Pos')
    ..e<PictureType>(3, 'Ptype', $pb.PbFieldType.OE, protoName: 'Ptype', defaultOrMaker: PictureType.ANIMAL, valueOf: PictureType.valueOf, enumValues: PictureType.values)
    ..hasRequiredFields = false
  ;

  FetchPicturesRequest._() : super();
  factory FetchPicturesRequest() => create();
  factory FetchPicturesRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FetchPicturesRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FetchPicturesRequest clone() => FetchPicturesRequest()..mergeFromMessage(this);
  FetchPicturesRequest copyWith(void Function(FetchPicturesRequest) updates) => super.copyWith((message) => updates(message as FetchPicturesRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FetchPicturesRequest create() => FetchPicturesRequest._();
  FetchPicturesRequest createEmptyInstance() => create();
  static $pb.PbList<FetchPicturesRequest> createRepeated() => $pb.PbList<FetchPicturesRequest>();
  @$core.pragma('dart2js:noInline')
  static FetchPicturesRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FetchPicturesRequest>(create);
  static FetchPicturesRequest _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get uUID => $_getSZ(0);
  @$pb.TagNumber(1)
  set uUID($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUUID() => $_has(0);
  @$pb.TagNumber(1)
  void clearUUID() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get pos => $_getIZ(1);
  @$pb.TagNumber(2)
  set pos($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasPos() => $_has(1);
  @$pb.TagNumber(2)
  void clearPos() => clearField(2);

  @$pb.TagNumber(3)
  PictureType get ptype => $_getN(2);
  @$pb.TagNumber(3)
  set ptype(PictureType v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasPtype() => $_has(2);
  @$pb.TagNumber(3)
  void clearPtype() => clearField(3);
}

class FetchPicturesResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FetchPicturesResponse', createEmptyInstance: create)
    ..e<StatusCode>(1, 'status', $pb.PbFieldType.OE, defaultOrMaker: StatusCode.SUCC, valueOf: StatusCode.valueOf, enumValues: StatusCode.values)
    ..aOS(2, 'msg')
    ..pc<PictureInfo>(3, 'info', $pb.PbFieldType.PM, subBuilder: PictureInfo.create)
    ..hasRequiredFields = false
  ;

  FetchPicturesResponse._() : super();
  factory FetchPicturesResponse() => create();
  factory FetchPicturesResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FetchPicturesResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FetchPicturesResponse clone() => FetchPicturesResponse()..mergeFromMessage(this);
  FetchPicturesResponse copyWith(void Function(FetchPicturesResponse) updates) => super.copyWith((message) => updates(message as FetchPicturesResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FetchPicturesResponse create() => FetchPicturesResponse._();
  FetchPicturesResponse createEmptyInstance() => create();
  static $pb.PbList<FetchPicturesResponse> createRepeated() => $pb.PbList<FetchPicturesResponse>();
  @$core.pragma('dart2js:noInline')
  static FetchPicturesResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FetchPicturesResponse>(create);
  static FetchPicturesResponse _defaultInstance;

  @$pb.TagNumber(1)
  StatusCode get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(StatusCode v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get msg => $_getSZ(1);
  @$pb.TagNumber(2)
  set msg($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMsg() => $_has(1);
  @$pb.TagNumber(2)
  void clearMsg() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<PictureInfo> get info => $_getList(2);
}

class PictureInfo extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('PictureInfo', createEmptyInstance: create)
    ..aOS(1, 'url')
    ..a<$core.int>(2, 'axisX', $pb.PbFieldType.O3, protoName: 'axis_X')
    ..a<$core.int>(3, 'axisY', $pb.PbFieldType.O3, protoName: 'axis_Y')
    ..hasRequiredFields = false
  ;

  PictureInfo._() : super();
  factory PictureInfo() => create();
  factory PictureInfo.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory PictureInfo.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  PictureInfo clone() => PictureInfo()..mergeFromMessage(this);
  PictureInfo copyWith(void Function(PictureInfo) updates) => super.copyWith((message) => updates(message as PictureInfo));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static PictureInfo create() => PictureInfo._();
  PictureInfo createEmptyInstance() => create();
  static $pb.PbList<PictureInfo> createRepeated() => $pb.PbList<PictureInfo>();
  @$core.pragma('dart2js:noInline')
  static PictureInfo getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<PictureInfo>(create);
  static PictureInfo _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get url => $_getSZ(0);
  @$pb.TagNumber(1)
  set url($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasUrl() => $_has(0);
  @$pb.TagNumber(1)
  void clearUrl() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get axisX => $_getIZ(1);
  @$pb.TagNumber(2)
  set axisX($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasAxisX() => $_has(1);
  @$pb.TagNumber(2)
  void clearAxisX() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get axisY => $_getIZ(2);
  @$pb.TagNumber(3)
  set axisY($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasAxisY() => $_has(2);
  @$pb.TagNumber(3)
  void clearAxisY() => clearField(3);
}

