# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cheese.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='cheese.proto',
  package='cheesefarm',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0c\x63heese.proto\x12\ncheesefarm\";\n\x06\x43heese\x12\x0b\n\x03\x61ge\x18\x01 \x01(\x05\x12$\n\x04type\x18\x02 \x01(\x0e\x32\x16.cheesefarm.CheeseType\"5\n\rCheeseRequest\x12$\n\x04type\x18\x01 \x01(\x0e\x32\x16.cheesefarm.CheeseType*O\n\nCheeseType\x12\x0c\n\x08\x45MMENTAL\x10\x00\x12\x08\n\x04\x42RIE\x10\x01\x12\x0c\n\x08PECORINO\x10\x02\x12\r\n\tROQUEFORT\x10\x03\x12\x0c\n\x08\x43\x41NASTRA\x10\x04\x32G\n\rCheeseService\x12\x36\n\x05Order\x12\x19.cheesefarm.CheeseRequest\x1a\x12.cheesefarm.Cheeseb\x06proto3')
)

_CHEESETYPE = _descriptor.EnumDescriptor(
  name='CheeseType',
  full_name='cheesefarm.CheeseType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EMMENTAL', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BRIE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PECORINO', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ROQUEFORT', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CANASTRA', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=144,
  serialized_end=223,
)
_sym_db.RegisterEnumDescriptor(_CHEESETYPE)

CheeseType = enum_type_wrapper.EnumTypeWrapper(_CHEESETYPE)
EMMENTAL = 0
BRIE = 1
PECORINO = 2
ROQUEFORT = 3
CANASTRA = 4



_CHEESE = _descriptor.Descriptor(
  name='Cheese',
  full_name='cheesefarm.Cheese',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='age', full_name='cheesefarm.Cheese.age', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='cheesefarm.Cheese.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=87,
)


_CHEESEREQUEST = _descriptor.Descriptor(
  name='CheeseRequest',
  full_name='cheesefarm.CheeseRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='cheesefarm.CheeseRequest.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=89,
  serialized_end=142,
)

_CHEESE.fields_by_name['type'].enum_type = _CHEESETYPE
_CHEESEREQUEST.fields_by_name['type'].enum_type = _CHEESETYPE
DESCRIPTOR.message_types_by_name['Cheese'] = _CHEESE
DESCRIPTOR.message_types_by_name['CheeseRequest'] = _CHEESEREQUEST
DESCRIPTOR.enum_types_by_name['CheeseType'] = _CHEESETYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Cheese = _reflection.GeneratedProtocolMessageType('Cheese', (_message.Message,), {
  'DESCRIPTOR' : _CHEESE,
  '__module__' : 'cheese_pb2'
  # @@protoc_insertion_point(class_scope:cheesefarm.Cheese)
  })
_sym_db.RegisterMessage(Cheese)

CheeseRequest = _reflection.GeneratedProtocolMessageType('CheeseRequest', (_message.Message,), {
  'DESCRIPTOR' : _CHEESEREQUEST,
  '__module__' : 'cheese_pb2'
  # @@protoc_insertion_point(class_scope:cheesefarm.CheeseRequest)
  })
_sym_db.RegisterMessage(CheeseRequest)



_CHEESESERVICE = _descriptor.ServiceDescriptor(
  name='CheeseService',
  full_name='cheesefarm.CheeseService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=225,
  serialized_end=296,
  methods=[
  _descriptor.MethodDescriptor(
    name='Order',
    full_name='cheesefarm.CheeseService.Order',
    index=0,
    containing_service=None,
    input_type=_CHEESEREQUEST,
    output_type=_CHEESE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CHEESESERVICE)

DESCRIPTOR.services_by_name['CheeseService'] = _CHEESESERVICE

# @@protoc_insertion_point(module_scope)
