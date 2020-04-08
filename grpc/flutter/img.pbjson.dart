///
//  Generated code. Do not modify.
//  source: img.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const StatusCode$json = const {
  '1': 'StatusCode',
  '2': const [
    const {'1': 'SUCC', '2': 0},
    const {'1': 'FAILE', '2': 1},
  ],
};

const PictureType$json = const {
  '1': 'PictureType',
  '2': const [
    const {'1': 'ANIMAL', '2': 0},
  ],
};

const FetchPicturesRequest$json = const {
  '1': 'FetchPicturesRequest',
  '2': const [
    const {'1': 'UUID', '3': 1, '4': 1, '5': 9, '10': 'UUID'},
    const {'1': 'Pos', '3': 2, '4': 1, '5': 5, '10': 'Pos'},
    const {'1': 'Ptype', '3': 3, '4': 1, '5': 14, '6': '.PictureType', '10': 'Ptype'},
  ],
};

const FetchPicturesResponse$json = const {
  '1': 'FetchPicturesResponse',
  '2': const [
    const {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.StatusCode', '10': 'status'},
    const {'1': 'msg', '3': 2, '4': 1, '5': 9, '10': 'msg'},
    const {'1': 'info', '3': 3, '4': 3, '5': 11, '6': '.PictureInfo', '10': 'info'},
  ],
};

const PictureInfo$json = const {
  '1': 'PictureInfo',
  '2': const [
    const {'1': 'url', '3': 1, '4': 1, '5': 9, '10': 'url'},
    const {'1': 'axis_X', '3': 2, '4': 1, '5': 5, '10': 'axisX'},
    const {'1': 'axis_Y', '3': 3, '4': 1, '5': 5, '10': 'axisY'},
  ],
};

