# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/admin/node_execution.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2
from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2
from flyteidl.core import catalog_pb2 as flyteidl_dot_core_dot_catalog__pb2
from flyteidl.core import compiler_pb2 as flyteidl_dot_core_dot_compiler__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/admin/node_execution.proto',
  package='flyteidl.admin',
  syntax='proto3',
  serialized_options=_b('Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin'),
  serialized_pb=_b('\n#flyteidl/admin/node_execution.proto\x12\x0e\x66lyteidl.admin\x1a\x1b\x66lyteidl/admin/common.proto\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1b\x66lyteidl/core/catalog.proto\x1a\x1c\x66lyteidl/core/compiler.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"M\n\x17NodeExecutionGetRequest\x12\x32\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifier\"\xd5\x01\n\x18NodeExecutionListRequest\x12I\n\x15workflow_execution_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifier\x12\r\n\x05limit\x18\x02 \x01(\r\x12\r\n\x05token\x18\x03 \x01(\t\x12\x0f\n\x07\x66ilters\x18\x04 \x01(\t\x12%\n\x07sort_by\x18\x05 \x01(\x0b\x32\x14.flyteidl.admin.Sort\x12\x18\n\x10unique_parent_id\x18\x06 \x01(\t\"\xba\x01\n\x1fNodeExecutionForTaskListRequest\x12\x41\n\x11task_execution_id\x18\x01 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifier\x12\r\n\x05limit\x18\x02 \x01(\r\x12\r\n\x05token\x18\x03 \x01(\t\x12\x0f\n\x07\x66ilters\x18\x04 \x01(\t\x12%\n\x07sort_by\x18\x05 \x01(\x0b\x32\x14.flyteidl.admin.Sort\"\xc6\x01\n\rNodeExecution\x12\x32\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifier\x12\x11\n\tinput_uri\x18\x02 \x01(\t\x12\x35\n\x07\x63losure\x18\x03 \x01(\x0b\x32$.flyteidl.admin.NodeExecutionClosure\x12\x37\n\x08metadata\x18\x04 \x01(\x0b\x32%.flyteidl.admin.NodeExecutionMetaData\"Z\n\x15NodeExecutionMetaData\x12\x13\n\x0bretry_group\x18\x01 \x01(\t\x12\x16\n\x0eis_parent_node\x18\x02 \x01(\x08\x12\x14\n\x0cspec_node_id\x18\x03 \x01(\t\"Z\n\x11NodeExecutionList\x12\x36\n\x0fnode_executions\x18\x01 \x03(\x0b\x32\x1d.flyteidl.admin.NodeExecution\x12\r\n\x05token\x18\x02 \x01(\t\"\xf8\x03\n\x14NodeExecutionClosure\x12\x14\n\noutput_uri\x18\x01 \x01(\tH\x00\x12.\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x1d.flyteidl.core.ExecutionErrorH\x00\x12\x31\n\x05phase\x18\x03 \x01(\x0e\x32\".flyteidl.core.NodeExecution.Phase\x12.\n\nstarted_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12+\n\x08\x64uration\x18\x05 \x01(\x0b\x32\x19.google.protobuf.Duration\x12.\n\ncreated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x46\n\x16workflow_node_metadata\x18\x08 \x01(\x0b\x32$.flyteidl.admin.WorkflowNodeMetadataH\x01\x12>\n\x12task_node_metadata\x18\t \x01(\x0b\x32 .flyteidl.admin.TaskNodeMetadataH\x01\x42\x0f\n\routput_resultB\x11\n\x0ftarget_metadata\"W\n\x14WorkflowNodeMetadata\x12?\n\x0b\x65xecutionId\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifier\"\x80\x01\n\x10TaskNodeMetadata\x12\x37\n\x0c\x63\x61\x63he_status\x18\x01 \x01(\x0e\x32!.flyteidl.core.CatalogCacheStatus\x12\x33\n\x0b\x63\x61talog_key\x18\x02 \x01(\x0b\x32\x1e.flyteidl.core.CatalogMetadata\"\x87\x01\n\x1b\x44ynamicWorkflowNodeMetadata\x12%\n\x02id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12\x41\n\x11\x63ompiled_workflow\x18\x02 \x01(\x0b\x32&.flyteidl.core.CompiledWorkflowClosure\"Q\n\x1bNodeExecutionGetDataRequest\x12\x32\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifier\"\x99\x02\n\x1cNodeExecutionGetDataResponse\x12\'\n\x06inputs\x18\x01 \x01(\x0b\x32\x17.flyteidl.admin.UrlBlob\x12(\n\x07outputs\x18\x02 \x01(\x0b\x32\x17.flyteidl.admin.UrlBlob\x12.\n\x0b\x66ull_inputs\x18\x03 \x01(\x0b\x32\x19.flyteidl.core.LiteralMap\x12/\n\x0c\x66ull_outputs\x18\x04 \x01(\x0b\x32\x19.flyteidl.core.LiteralMap\x12\x45\n\x10\x64ynamic_workflow\x18\x10 \x01(\x0b\x32+.flyteidl.admin.DynamicWorkflowNodeMetadata\"Q\n\x1bNodeExecutionRecoverRequest\x12\x32\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifierB7Z5github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/adminb\x06proto3')
  ,
  dependencies=[flyteidl_dot_admin_dot_common__pb2.DESCRIPTOR,flyteidl_dot_core_dot_execution__pb2.DESCRIPTOR,flyteidl_dot_core_dot_catalog__pb2.DESCRIPTOR,flyteidl_dot_core_dot_compiler__pb2.DESCRIPTOR,flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,flyteidl_dot_core_dot_literals__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])




_NODEEXECUTIONGETREQUEST = _descriptor.Descriptor(
  name='NodeExecutionGetRequest',
  full_name='flyteidl.admin.NodeExecutionGetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.NodeExecutionGetRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=301,
  serialized_end=378,
)


_NODEEXECUTIONLISTREQUEST = _descriptor.Descriptor(
  name='NodeExecutionListRequest',
  full_name='flyteidl.admin.NodeExecutionListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workflow_execution_id', full_name='flyteidl.admin.NodeExecutionListRequest.workflow_execution_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='flyteidl.admin.NodeExecutionListRequest.limit', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='flyteidl.admin.NodeExecutionListRequest.token', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='filters', full_name='flyteidl.admin.NodeExecutionListRequest.filters', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sort_by', full_name='flyteidl.admin.NodeExecutionListRequest.sort_by', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='unique_parent_id', full_name='flyteidl.admin.NodeExecutionListRequest.unique_parent_id', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=381,
  serialized_end=594,
)


_NODEEXECUTIONFORTASKLISTREQUEST = _descriptor.Descriptor(
  name='NodeExecutionForTaskListRequest',
  full_name='flyteidl.admin.NodeExecutionForTaskListRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='task_execution_id', full_name='flyteidl.admin.NodeExecutionForTaskListRequest.task_execution_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='flyteidl.admin.NodeExecutionForTaskListRequest.limit', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='flyteidl.admin.NodeExecutionForTaskListRequest.token', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='filters', full_name='flyteidl.admin.NodeExecutionForTaskListRequest.filters', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sort_by', full_name='flyteidl.admin.NodeExecutionForTaskListRequest.sort_by', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=597,
  serialized_end=783,
)


_NODEEXECUTION = _descriptor.Descriptor(
  name='NodeExecution',
  full_name='flyteidl.admin.NodeExecution',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.NodeExecution.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='input_uri', full_name='flyteidl.admin.NodeExecution.input_uri', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='closure', full_name='flyteidl.admin.NodeExecution.closure', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='flyteidl.admin.NodeExecution.metadata', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=786,
  serialized_end=984,
)


_NODEEXECUTIONMETADATA = _descriptor.Descriptor(
  name='NodeExecutionMetaData',
  full_name='flyteidl.admin.NodeExecutionMetaData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='retry_group', full_name='flyteidl.admin.NodeExecutionMetaData.retry_group', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_parent_node', full_name='flyteidl.admin.NodeExecutionMetaData.is_parent_node', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='spec_node_id', full_name='flyteidl.admin.NodeExecutionMetaData.spec_node_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=986,
  serialized_end=1076,
)


_NODEEXECUTIONLIST = _descriptor.Descriptor(
  name='NodeExecutionList',
  full_name='flyteidl.admin.NodeExecutionList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='node_executions', full_name='flyteidl.admin.NodeExecutionList.node_executions', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='token', full_name='flyteidl.admin.NodeExecutionList.token', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=1078,
  serialized_end=1168,
)


_NODEEXECUTIONCLOSURE = _descriptor.Descriptor(
  name='NodeExecutionClosure',
  full_name='flyteidl.admin.NodeExecutionClosure',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='output_uri', full_name='flyteidl.admin.NodeExecutionClosure.output_uri', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='flyteidl.admin.NodeExecutionClosure.error', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase', full_name='flyteidl.admin.NodeExecutionClosure.phase', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='started_at', full_name='flyteidl.admin.NodeExecutionClosure.started_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='duration', full_name='flyteidl.admin.NodeExecutionClosure.duration', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='flyteidl.admin.NodeExecutionClosure.created_at', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='flyteidl.admin.NodeExecutionClosure.updated_at', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='workflow_node_metadata', full_name='flyteidl.admin.NodeExecutionClosure.workflow_node_metadata', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='task_node_metadata', full_name='flyteidl.admin.NodeExecutionClosure.task_node_metadata', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
    _descriptor.OneofDescriptor(
      name='output_result', full_name='flyteidl.admin.NodeExecutionClosure.output_result',
      index=0, containing_type=None, fields=[]),
    _descriptor.OneofDescriptor(
      name='target_metadata', full_name='flyteidl.admin.NodeExecutionClosure.target_metadata',
      index=1, containing_type=None, fields=[]),
  ],
  serialized_start=1171,
  serialized_end=1675,
)


_WORKFLOWNODEMETADATA = _descriptor.Descriptor(
  name='WorkflowNodeMetadata',
  full_name='flyteidl.admin.WorkflowNodeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='executionId', full_name='flyteidl.admin.WorkflowNodeMetadata.executionId', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1677,
  serialized_end=1764,
)


_TASKNODEMETADATA = _descriptor.Descriptor(
  name='TaskNodeMetadata',
  full_name='flyteidl.admin.TaskNodeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cache_status', full_name='flyteidl.admin.TaskNodeMetadata.cache_status', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='catalog_key', full_name='flyteidl.admin.TaskNodeMetadata.catalog_key', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1767,
  serialized_end=1895,
)


_DYNAMICWORKFLOWNODEMETADATA = _descriptor.Descriptor(
  name='DynamicWorkflowNodeMetadata',
  full_name='flyteidl.admin.DynamicWorkflowNodeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.DynamicWorkflowNodeMetadata.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='compiled_workflow', full_name='flyteidl.admin.DynamicWorkflowNodeMetadata.compiled_workflow', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=1898,
  serialized_end=2033,
)


_NODEEXECUTIONGETDATAREQUEST = _descriptor.Descriptor(
  name='NodeExecutionGetDataRequest',
  full_name='flyteidl.admin.NodeExecutionGetDataRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.NodeExecutionGetDataRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=2035,
  serialized_end=2116,
)


_NODEEXECUTIONGETDATARESPONSE = _descriptor.Descriptor(
  name='NodeExecutionGetDataResponse',
  full_name='flyteidl.admin.NodeExecutionGetDataResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='inputs', full_name='flyteidl.admin.NodeExecutionGetDataResponse.inputs', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='flyteidl.admin.NodeExecutionGetDataResponse.outputs', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='full_inputs', full_name='flyteidl.admin.NodeExecutionGetDataResponse.full_inputs', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='full_outputs', full_name='flyteidl.admin.NodeExecutionGetDataResponse.full_outputs', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dynamic_workflow', full_name='flyteidl.admin.NodeExecutionGetDataResponse.dynamic_workflow', index=4,
      number=16, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=2119,
  serialized_end=2400,
)


_NODEEXECUTIONRECOVERREQUEST = _descriptor.Descriptor(
  name='NodeExecutionRecoverRequest',
  full_name='flyteidl.admin.NodeExecutionRecoverRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.admin.NodeExecutionRecoverRequest.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=2402,
  serialized_end=2483,
)

_NODEEXECUTIONGETREQUEST.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._NODEEXECUTIONIDENTIFIER
_NODEEXECUTIONLISTREQUEST.fields_by_name['workflow_execution_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._WORKFLOWEXECUTIONIDENTIFIER
_NODEEXECUTIONLISTREQUEST.fields_by_name['sort_by'].message_type = flyteidl_dot_admin_dot_common__pb2._SORT
_NODEEXECUTIONFORTASKLISTREQUEST.fields_by_name['task_execution_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._TASKEXECUTIONIDENTIFIER
_NODEEXECUTIONFORTASKLISTREQUEST.fields_by_name['sort_by'].message_type = flyteidl_dot_admin_dot_common__pb2._SORT
_NODEEXECUTION.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._NODEEXECUTIONIDENTIFIER
_NODEEXECUTION.fields_by_name['closure'].message_type = _NODEEXECUTIONCLOSURE
_NODEEXECUTION.fields_by_name['metadata'].message_type = _NODEEXECUTIONMETADATA
_NODEEXECUTIONLIST.fields_by_name['node_executions'].message_type = _NODEEXECUTION
_NODEEXECUTIONCLOSURE.fields_by_name['error'].message_type = flyteidl_dot_core_dot_execution__pb2._EXECUTIONERROR
_NODEEXECUTIONCLOSURE.fields_by_name['phase'].enum_type = flyteidl_dot_core_dot_execution__pb2._NODEEXECUTION_PHASE
_NODEEXECUTIONCLOSURE.fields_by_name['started_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NODEEXECUTIONCLOSURE.fields_by_name['duration'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_NODEEXECUTIONCLOSURE.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NODEEXECUTIONCLOSURE.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NODEEXECUTIONCLOSURE.fields_by_name['workflow_node_metadata'].message_type = _WORKFLOWNODEMETADATA
_NODEEXECUTIONCLOSURE.fields_by_name['task_node_metadata'].message_type = _TASKNODEMETADATA
_NODEEXECUTIONCLOSURE.oneofs_by_name['output_result'].fields.append(
  _NODEEXECUTIONCLOSURE.fields_by_name['output_uri'])
_NODEEXECUTIONCLOSURE.fields_by_name['output_uri'].containing_oneof = _NODEEXECUTIONCLOSURE.oneofs_by_name['output_result']
_NODEEXECUTIONCLOSURE.oneofs_by_name['output_result'].fields.append(
  _NODEEXECUTIONCLOSURE.fields_by_name['error'])
_NODEEXECUTIONCLOSURE.fields_by_name['error'].containing_oneof = _NODEEXECUTIONCLOSURE.oneofs_by_name['output_result']
_NODEEXECUTIONCLOSURE.oneofs_by_name['target_metadata'].fields.append(
  _NODEEXECUTIONCLOSURE.fields_by_name['workflow_node_metadata'])
_NODEEXECUTIONCLOSURE.fields_by_name['workflow_node_metadata'].containing_oneof = _NODEEXECUTIONCLOSURE.oneofs_by_name['target_metadata']
_NODEEXECUTIONCLOSURE.oneofs_by_name['target_metadata'].fields.append(
  _NODEEXECUTIONCLOSURE.fields_by_name['task_node_metadata'])
_NODEEXECUTIONCLOSURE.fields_by_name['task_node_metadata'].containing_oneof = _NODEEXECUTIONCLOSURE.oneofs_by_name['target_metadata']
_WORKFLOWNODEMETADATA.fields_by_name['executionId'].message_type = flyteidl_dot_core_dot_identifier__pb2._WORKFLOWEXECUTIONIDENTIFIER
_TASKNODEMETADATA.fields_by_name['cache_status'].enum_type = flyteidl_dot_core_dot_catalog__pb2._CATALOGCACHESTATUS
_TASKNODEMETADATA.fields_by_name['catalog_key'].message_type = flyteidl_dot_core_dot_catalog__pb2._CATALOGMETADATA
_DYNAMICWORKFLOWNODEMETADATA.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_DYNAMICWORKFLOWNODEMETADATA.fields_by_name['compiled_workflow'].message_type = flyteidl_dot_core_dot_compiler__pb2._COMPILEDWORKFLOWCLOSURE
_NODEEXECUTIONGETDATAREQUEST.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._NODEEXECUTIONIDENTIFIER
_NODEEXECUTIONGETDATARESPONSE.fields_by_name['inputs'].message_type = flyteidl_dot_admin_dot_common__pb2._URLBLOB
_NODEEXECUTIONGETDATARESPONSE.fields_by_name['outputs'].message_type = flyteidl_dot_admin_dot_common__pb2._URLBLOB
_NODEEXECUTIONGETDATARESPONSE.fields_by_name['full_inputs'].message_type = flyteidl_dot_core_dot_literals__pb2._LITERALMAP
_NODEEXECUTIONGETDATARESPONSE.fields_by_name['full_outputs'].message_type = flyteidl_dot_core_dot_literals__pb2._LITERALMAP
_NODEEXECUTIONGETDATARESPONSE.fields_by_name['dynamic_workflow'].message_type = _DYNAMICWORKFLOWNODEMETADATA
_NODEEXECUTIONRECOVERREQUEST.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._NODEEXECUTIONIDENTIFIER
DESCRIPTOR.message_types_by_name['NodeExecutionGetRequest'] = _NODEEXECUTIONGETREQUEST
DESCRIPTOR.message_types_by_name['NodeExecutionListRequest'] = _NODEEXECUTIONLISTREQUEST
DESCRIPTOR.message_types_by_name['NodeExecutionForTaskListRequest'] = _NODEEXECUTIONFORTASKLISTREQUEST
DESCRIPTOR.message_types_by_name['NodeExecution'] = _NODEEXECUTION
DESCRIPTOR.message_types_by_name['NodeExecutionMetaData'] = _NODEEXECUTIONMETADATA
DESCRIPTOR.message_types_by_name['NodeExecutionList'] = _NODEEXECUTIONLIST
DESCRIPTOR.message_types_by_name['NodeExecutionClosure'] = _NODEEXECUTIONCLOSURE
DESCRIPTOR.message_types_by_name['WorkflowNodeMetadata'] = _WORKFLOWNODEMETADATA
DESCRIPTOR.message_types_by_name['TaskNodeMetadata'] = _TASKNODEMETADATA
DESCRIPTOR.message_types_by_name['DynamicWorkflowNodeMetadata'] = _DYNAMICWORKFLOWNODEMETADATA
DESCRIPTOR.message_types_by_name['NodeExecutionGetDataRequest'] = _NODEEXECUTIONGETDATAREQUEST
DESCRIPTOR.message_types_by_name['NodeExecutionGetDataResponse'] = _NODEEXECUTIONGETDATARESPONSE
DESCRIPTOR.message_types_by_name['NodeExecutionRecoverRequest'] = _NODEEXECUTIONRECOVERREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

NodeExecutionGetRequest = _reflection.GeneratedProtocolMessageType('NodeExecutionGetRequest', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONGETREQUEST,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionGetRequest)
  ))
_sym_db.RegisterMessage(NodeExecutionGetRequest)

NodeExecutionListRequest = _reflection.GeneratedProtocolMessageType('NodeExecutionListRequest', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONLISTREQUEST,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionListRequest)
  ))
_sym_db.RegisterMessage(NodeExecutionListRequest)

NodeExecutionForTaskListRequest = _reflection.GeneratedProtocolMessageType('NodeExecutionForTaskListRequest', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONFORTASKLISTREQUEST,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionForTaskListRequest)
  ))
_sym_db.RegisterMessage(NodeExecutionForTaskListRequest)

NodeExecution = _reflection.GeneratedProtocolMessageType('NodeExecution', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTION,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecution)
  ))
_sym_db.RegisterMessage(NodeExecution)

NodeExecutionMetaData = _reflection.GeneratedProtocolMessageType('NodeExecutionMetaData', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONMETADATA,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionMetaData)
  ))
_sym_db.RegisterMessage(NodeExecutionMetaData)

NodeExecutionList = _reflection.GeneratedProtocolMessageType('NodeExecutionList', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONLIST,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionList)
  ))
_sym_db.RegisterMessage(NodeExecutionList)

NodeExecutionClosure = _reflection.GeneratedProtocolMessageType('NodeExecutionClosure', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONCLOSURE,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionClosure)
  ))
_sym_db.RegisterMessage(NodeExecutionClosure)

WorkflowNodeMetadata = _reflection.GeneratedProtocolMessageType('WorkflowNodeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWNODEMETADATA,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.WorkflowNodeMetadata)
  ))
_sym_db.RegisterMessage(WorkflowNodeMetadata)

TaskNodeMetadata = _reflection.GeneratedProtocolMessageType('TaskNodeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _TASKNODEMETADATA,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.TaskNodeMetadata)
  ))
_sym_db.RegisterMessage(TaskNodeMetadata)

DynamicWorkflowNodeMetadata = _reflection.GeneratedProtocolMessageType('DynamicWorkflowNodeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _DYNAMICWORKFLOWNODEMETADATA,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.DynamicWorkflowNodeMetadata)
  ))
_sym_db.RegisterMessage(DynamicWorkflowNodeMetadata)

NodeExecutionGetDataRequest = _reflection.GeneratedProtocolMessageType('NodeExecutionGetDataRequest', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONGETDATAREQUEST,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionGetDataRequest)
  ))
_sym_db.RegisterMessage(NodeExecutionGetDataRequest)

NodeExecutionGetDataResponse = _reflection.GeneratedProtocolMessageType('NodeExecutionGetDataResponse', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONGETDATARESPONSE,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionGetDataResponse)
  ))
_sym_db.RegisterMessage(NodeExecutionGetDataResponse)

NodeExecutionRecoverRequest = _reflection.GeneratedProtocolMessageType('NodeExecutionRecoverRequest', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONRECOVERREQUEST,
  __module__ = 'flyteidl.admin.node_execution_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.admin.NodeExecutionRecoverRequest)
  ))
_sym_db.RegisterMessage(NodeExecutionRecoverRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
