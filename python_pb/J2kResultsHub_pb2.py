# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: J2kResultsHub.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13J2kResultsHub.proto\x12\rJ2KResultsHub\"M\n\nVarResults\x12\x12\n\ncellNumber\x18\x01 \x01(\r\x12+\n\tvarResuls\x18\x02 \x03(\x0b\x32\x18.J2KResultsHub.VarResult\"v\n\tVarResult\x12\x0f\n\x07varName\x18\x01 \x01(\t\x12\x0f\n\x07varType\x18\x02 \x01(\t\x12\x10\n\x08varBytes\x18\x03 \x01(\x0c\x12\x11\n\tavailable\x18\x04 \x01(\x08\x12\x0e\n\x06isJson\x18\x05 \x01(\x08\x12\x12\n\njsonString\x18\x06 \x01(\t\"A\n\x15\x46\x65tchVarResultRequest\x12\x0f\n\x07varName\x18\x01 \x01(\t\x12\x17\n\x0fvarAncestorCell\x18\x02 \x01(\r\"\"\n\x0fWaitCellRequest\x12\x0f\n\x07waitFor\x18\x01 \x01(\r\"\x07\n\x05\x45mpty\"1\n\x0cHelloRequest\x12\x10\n\x08senderId\x18\x01 \x01(\t\x12\x0f\n\x07message\x18\x02 \x01(\t\"\x1d\n\nHelloReply\x12\x0f\n\x07message\x18\x01 \x01(\t2\xb5\x02\n\nResultsHub\x12\x46\n\x11\x43laimCellFinished\x12\x19.J2KResultsHub.VarResults\x1a\x14.J2KResultsHub.Empty\"\x00\x12R\n\x0e\x46\x65tchVarResult\x12$.J2KResultsHub.FetchVarResultRequest\x1a\x18.J2KResultsHub.VarResult\"\x00\x12\x45\n\x0bWaitForCell\x12\x1e.J2KResultsHub.WaitCellRequest\x1a\x14.J2KResultsHub.Empty\"\x00\x12\x44\n\x08SayHello\x12\x1b.J2KResultsHub.HelloRequest\x1a\x19.J2KResultsHub.HelloReply\"\x00\x42\x04Z\x02./b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'J2kResultsHub_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\002./'
  _globals['_VARRESULTS']._serialized_start=38
  _globals['_VARRESULTS']._serialized_end=115
  _globals['_VARRESULT']._serialized_start=117
  _globals['_VARRESULT']._serialized_end=235
  _globals['_FETCHVARRESULTREQUEST']._serialized_start=237
  _globals['_FETCHVARRESULTREQUEST']._serialized_end=302
  _globals['_WAITCELLREQUEST']._serialized_start=304
  _globals['_WAITCELLREQUEST']._serialized_end=338
  _globals['_EMPTY']._serialized_start=340
  _globals['_EMPTY']._serialized_end=347
  _globals['_HELLOREQUEST']._serialized_start=349
  _globals['_HELLOREQUEST']._serialized_end=398
  _globals['_HELLOREPLY']._serialized_start=400
  _globals['_HELLOREPLY']._serialized_end=429
  _globals['_RESULTSHUB']._serialized_start=432
  _globals['_RESULTSHUB']._serialized_end=741
# @@protoc_insertion_point(module_scope)
