///
//  Generated code. Do not modify.
//  source: img.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class StatusCode extends $pb.ProtobufEnum {
  static const StatusCode SUCC = StatusCode._(0, 'SUCC');
  static const StatusCode FAILE = StatusCode._(1, 'FAILE');

  static const $core.List<StatusCode> values = <StatusCode> [
    SUCC,
    FAILE,
  ];

  static final $core.Map<$core.int, StatusCode> _byValue = $pb.ProtobufEnum.initByValue(values);
  static StatusCode valueOf($core.int value) => _byValue[value];

  const StatusCode._($core.int v, $core.String n) : super(v, n);
}

class PictureType extends $pb.ProtobufEnum {
  static const PictureType ANIMAL = PictureType._(0, 'ANIMAL');

  static const $core.List<PictureType> values = <PictureType> [
    ANIMAL,
  ];

  static final $core.Map<$core.int, PictureType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static PictureType valueOf($core.int value) => _byValue[value];

  const PictureType._($core.int v, $core.String n) : super(v, n);
}

