# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: watchdog.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='watchdog.proto',
  package='watchdog',
  syntax='proto3',
  serialized_options=_b('Z\032github.com/jchorl/watchdog'),
  serialized_pb=_b('\n\x0ewatchdog.proto\x12\x08watchdog\"y\n\x05Watch\x12\x0c\n\x04name\x18\x01 \x01(\t\x12,\n\tfrequency\x18\x02 \x01(\x0e\x32\x19.watchdog.Watch.Frequency\x12\x10\n\x08LastSeen\x18\x03 \x01(\x03\"\"\n\tFrequency\x12\t\n\x05\x44\x41ILY\x10\x00\x12\n\n\x06WEEKLY\x10\x01\x42\x1cZ\x1agithub.com/jchorl/watchdogb\x06proto3')
)



_WATCH_FREQUENCY = _descriptor.EnumDescriptor(
  name='Frequency',
  full_name='watchdog.Watch.Frequency',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DAILY', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WEEKLY', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=115,
  serialized_end=149,
)
_sym_db.RegisterEnumDescriptor(_WATCH_FREQUENCY)


_WATCH = _descriptor.Descriptor(
  name='Watch',
  full_name='watchdog.Watch',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='watchdog.Watch.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='frequency', full_name='watchdog.Watch.frequency', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='LastSeen', full_name='watchdog.Watch.LastSeen', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _WATCH_FREQUENCY,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=149,
)

_WATCH.fields_by_name['frequency'].enum_type = _WATCH_FREQUENCY
_WATCH_FREQUENCY.containing_type = _WATCH
DESCRIPTOR.message_types_by_name['Watch'] = _WATCH
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Watch = _reflection.GeneratedProtocolMessageType('Watch', (_message.Message,), dict(
  DESCRIPTOR = _WATCH,
  __module__ = 'watchdog_pb2'
  # @@protoc_insertion_point(class_scope:watchdog.Watch)
  ))
_sym_db.RegisterMessage(Watch)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)