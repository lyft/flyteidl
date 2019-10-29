# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/service/admin.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from flyteidl.admin import project_pb2 as flyteidl_dot_admin_dot_project__pb2
from flyteidl.admin import task_pb2 as flyteidl_dot_admin_dot_task__pb2
from flyteidl.admin import workflow_pb2 as flyteidl_dot_admin_dot_workflow__pb2
from flyteidl.admin import launch_plan_pb2 as flyteidl_dot_admin_dot_launch__plan__pb2
from flyteidl.admin import event_pb2 as flyteidl_dot_admin_dot_event__pb2
from flyteidl.admin import execution_pb2 as flyteidl_dot_admin_dot_execution__pb2
from flyteidl.admin import node_execution_pb2 as flyteidl_dot_admin_dot_node__execution__pb2
from flyteidl.admin import task_execution_pb2 as flyteidl_dot_admin_dot_task__execution__pb2
from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2
from protoc_gen_swagger.options import annotations_pb2 as protoc__gen__swagger_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/service/admin.proto',
  package='flyteidl.service',
  syntax='proto3',
  serialized_options=_b('Z3github.com/lyft/flyteidl/gen/pb-go/flyteidl/service'),
  serialized_pb=_b('\n\x1c\x66lyteidl/service/admin.proto\x12\x10\x66lyteidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1c\x66lyteidl/admin/project.proto\x1a\x19\x66lyteidl/admin/task.proto\x1a\x1d\x66lyteidl/admin/workflow.proto\x1a flyteidl/admin/launch_plan.proto\x1a\x1a\x66lyteidl/admin/event.proto\x1a\x1e\x66lyteidl/admin/execution.proto\x1a#flyteidl/admin/node_execution.proto\x1a#flyteidl/admin/task_execution.proto\x1a\x1b\x66lyteidl/admin/common.proto\x1a,protoc-gen-swagger/options/annotations.proto2\xbeN\n\x0c\x41\x64minService\x12\xc5\x02\n\nCreateTask\x12!.flyteidl.admin.TaskCreateRequest\x1a\".flyteidl.admin.TaskCreateResponse\"\xef\x01\x82\xd3\xe4\x93\x02\x12\"\r/api/v1/tasks:\x01*\x92\x41\xd3\x01\x1a&Create and register a task definition.JB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\xb2\x01\n\x07GetTask\x12 .flyteidl.admin.ObjectGetRequest\x1a\x14.flyteidl.admin.Task\"o\x82\xd3\xe4\x93\x02?\x12=/api/v1/tasks/{id.project}/{id.domain}/{id.name}/{id.version}\x92\x41\'\x1a%Retrieve an existing task definition.\x12\xde\x01\n\x0bListTaskIds\x12\x30.flyteidl.admin.NamedEntityIdentifierListRequest\x1a).flyteidl.admin.NamedEntityIdentifierList\"r\x82\xd3\xe4\x93\x02%\x12#/api/v1/task_ids/{project}/{domain}\x92\x41\x44\x1a\x42\x46\x65tch existing task definition identifiers matching input filters.\x12\xeb\x01\n\tListTasks\x12#.flyteidl.admin.ResourceListRequest\x1a\x18.flyteidl.admin.TaskList\"\x9e\x01\x82\xd3\xe4\x93\x02\\\x12\x30/api/v1/tasks/{id.project}/{id.domain}/{id.name}Z(\x12&/api/v1/tasks/{id.project}/{id.domain}\x92\x41\x39\x1a\x37\x46\x65tch existing task definitions matching input filters.\x12\xd9\x02\n\x0e\x43reateWorkflow\x12%.flyteidl.admin.WorkflowCreateRequest\x1a&.flyteidl.admin.WorkflowCreateResponse\"\xf7\x01\x82\xd3\xe4\x93\x02\x16\"\x11/api/v1/workflows:\x01*\x92\x41\xd7\x01\x1a*Create and register a workflow definition.JB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\xc2\x01\n\x0bGetWorkflow\x12 .flyteidl.admin.ObjectGetRequest\x1a\x18.flyteidl.admin.Workflow\"w\x82\xd3\xe4\x93\x02\x43\x12\x41/api/v1/workflows/{id.project}/{id.domain}/{id.name}/{id.version}\x92\x41+\x1a)Retrieve an existing workflow definition.\x12\xed\x01\n\x0fListWorkflowIds\x12\x30.flyteidl.admin.NamedEntityIdentifierListRequest\x1a).flyteidl.admin.NamedEntityIdentifierList\"}\x82\xd3\xe4\x93\x02)\x12\'/api/v1/workflow_ids/{project}/{domain}\x92\x41K\x1aIFetch an existing workflow definition identifiers matching input filters.\x12\xff\x01\n\rListWorkflows\x12#.flyteidl.admin.ResourceListRequest\x1a\x1c.flyteidl.admin.WorkflowList\"\xaa\x01\x82\xd3\xe4\x93\x02\x64\x12\x34/api/v1/workflows/{id.project}/{id.domain}/{id.name}Z,\x12*/api/v1/workflows/{id.project}/{id.domain}\x92\x41=\x1a;Fetch existing workflow definitions matching input filters.\x12\xe5\x02\n\x10\x43reateLaunchPlan\x12\'.flyteidl.admin.LaunchPlanCreateRequest\x1a(.flyteidl.admin.LaunchPlanCreateResponse\"\xfd\x01\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/launch_plans:\x01*\x92\x41\xda\x01\x1a-Create and register a launch plan definition.JB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\xcc\x01\n\rGetLaunchPlan\x12 .flyteidl.admin.ObjectGetRequest\x1a\x1a.flyteidl.admin.LaunchPlan\"}\x82\xd3\xe4\x93\x02\x46\x12\x44/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}\x92\x41.\x1a,Retrieve an existing launch plan definition.\x12\xf3\x01\n\x13GetActiveLaunchPlan\x12\'.flyteidl.admin.ActiveLaunchPlanRequest\x1a\x1a.flyteidl.admin.LaunchPlan\"\x96\x01\x82\xd3\xe4\x93\x02@\x12>/api/v1/active_launch_plans/{id.project}/{id.domain}/{id.name}\x92\x41M\x1aKRetrieve the active launch plan version specified by input request filters.\x12\xeb\x01\n\x15ListActiveLaunchPlans\x12+.flyteidl.admin.ActiveLaunchPlanListRequest\x1a\x1e.flyteidl.admin.LaunchPlanList\"\x84\x01\x82\xd3\xe4\x93\x02\x30\x12./api/v1/active_launch_plans/{project}/{domain}\x92\x41K\x1aIFetch the active launch plan versions specified by input request filters.\x12\xf3\x01\n\x11ListLaunchPlanIds\x12\x30.flyteidl.admin.NamedEntityIdentifierListRequest\x1a).flyteidl.admin.NamedEntityIdentifierList\"\x80\x01\x82\xd3\xe4\x93\x02,\x12*/api/v1/launch_plan_ids/{project}/{domain}\x92\x41K\x1aIFetch existing launch plan definition identifiers matching input filters.\x12\x8c\x02\n\x0fListLaunchPlans\x12#.flyteidl.admin.ResourceListRequest\x1a\x1e.flyteidl.admin.LaunchPlanList\"\xb3\x01\x82\xd3\xe4\x93\x02j\x12\x37/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}Z/\x12-/api/v1/launch_plans/{id.project}/{id.domain}\x92\x41@\x1a>Fetch existing launch plan definitions matching input filters.\x12\xc0\x06\n\x10UpdateLaunchPlan\x12\'.flyteidl.admin.LaunchPlanUpdateRequest\x1a(.flyteidl.admin.LaunchPlanUpdateResponse\"\xd8\x05\x82\xd3\xe4\x93\x02I\x1a\x44/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}:\x01*\x92\x41\x85\x05\x1a\x82\x05Update the status of an existing launch plan definition. At most one launch plan version for a given {project, domain, name} can be active at a time. If this call sets a launch plan to active and existing version is already active, the result of this call will be that the formerly active launch plan will be made inactive and specified launch plan in this request will be made active. In the event that the formerly active launch plan had a schedule associated it with it, this schedule will be disabled. If the reference launch plan in this request is being set to active and has a schedule associated with it, the schedule will be enabled.\x12\xa2\x01\n\x0f\x43reateExecution\x12&.flyteidl.admin.ExecutionCreateRequest\x1a\'.flyteidl.admin.ExecutionCreateResponse\">\x82\xd3\xe4\x93\x02\x17\"\x12/api/v1/executions:\x01*\x92\x41\x1e\x1a\x1c\x43reate a workflow execution.\x12\xb1\x01\n\x11RelaunchExecution\x12(.flyteidl.admin.ExecutionRelaunchRequest\x1a\'.flyteidl.admin.ExecutionCreateResponse\"I\x82\xd3\xe4\x93\x02 \"\x1b/api/v1/executions/relaunch:\x01*\x92\x41 \x1a\x1eRelaunch a workflow execution.\x12\xc2\x01\n\x0cGetExecution\x12+.flyteidl.admin.WorkflowExecutionGetRequest\x1a\x19.flyteidl.admin.Execution\"j\x82\xd3\xe4\x93\x02\x37\x12\x35/api/v1/executions/{id.project}/{id.domain}/{id.name}\x92\x41*\x1a(Retrieve an existing workflow execution.\x12\x82\x02\n\x10GetExecutionData\x12/.flyteidl.admin.WorkflowExecutionGetDataRequest\x1a\x30.flyteidl.admin.WorkflowExecutionGetDataResponse\"\x8a\x01\x82\xd3\xe4\x93\x02<\x12:/api/v1/data/executions/{id.project}/{id.domain}/{id.name}\x92\x41\x45\x1a\x43Retrieve input and output data from an existing workflow execution.\x12\xc8\x01\n\x0eListExecutions\x12#.flyteidl.admin.ResourceListRequest\x1a\x1d.flyteidl.admin.ExecutionList\"r\x82\xd3\xe4\x93\x02-\x12+/api/v1/executions/{id.project}/{id.domain}\x92\x41<\x1a:Fetch existing workflow executions matching input filters.\x12\xf4\x01\n\x12TerminateExecution\x12).flyteidl.admin.ExecutionTerminateRequest\x1a*.flyteidl.admin.ExecutionTerminateResponse\"\x86\x01\x82\xd3\xe4\x93\x02:*5/api/v1/executions/{id.project}/{id.domain}/{id.name}:\x01*\x92\x41\x43\x1a\x41Terminate the active workflow execution specified in the request.\x12\xfc\x01\n\x10GetNodeExecution\x12\'.flyteidl.admin.NodeExecutionGetRequest\x1a\x1d.flyteidl.admin.NodeExecution\"\x9f\x01\x82\xd3\xe4\x93\x02p\x12n/api/v1/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}\x92\x41&\x1a$Retrieve an existing node execution.\x12\x9a\x02\n\x12ListNodeExecutions\x12(.flyteidl.admin.NodeExecutionListRequest\x1a!.flyteidl.admin.NodeExecutionList\"\xb6\x01\x82\xd3\xe4\x93\x02u\x12s/api/v1/node_executions/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\x92\x41\x38\x1a\x36\x46\x65tch existing node executions matching input filters.\x12\xef\x04\n\x19ListNodeExecutionsForTask\x12/.flyteidl.admin.NodeExecutionForTaskListRequest\x1a!.flyteidl.admin.NodeExecutionList\"\xfd\x03\x82\xd3\xe4\x93\x02\xac\x03\x12\xa9\x03/api/v1/children/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt}\x92\x41G\x1a\x45\x46\x65tch child node executions launched by the specified task execution.\x12\xb3\x02\n\x14GetNodeExecutionData\x12+.flyteidl.admin.NodeExecutionGetDataRequest\x1a,.flyteidl.admin.NodeExecutionGetDataResponse\"\xbf\x01\x82\xd3\xe4\x93\x02u\x12s/api/v1/data/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}\x92\x41\x41\x1a?Retrieve input and output data from an existing node execution.\x12\x97\x01\n\x0fRegisterProject\x12&.flyteidl.admin.ProjectRegisterRequest\x1a\'.flyteidl.admin.ProjectRegisterResponse\"3\x82\xd3\xe4\x93\x02\x15\"\x10/api/v1/projects:\x01*\x92\x41\x15\x1a\x13Register a project.\x12\x85\x01\n\x0cListProjects\x12\".flyteidl.admin.ProjectListRequest\x1a\x18.flyteidl.admin.Projects\"7\x82\xd3\xe4\x93\x02\x12\x12\x10/api/v1/projects\x92\x41\x1c\x1a\x1a\x46\x65tch registered projects.\x12\xdd\x01\n\x13\x43reateWorkflowEvent\x12-.flyteidl.admin.WorkflowExecutionEventRequest\x1a..flyteidl.admin.WorkflowExecutionEventResponse\"g\x82\xd3\xe4\x93\x02\x1d\"\x18/api/v1/events/workflows:\x01*\x92\x41\x41\x1a?Create a workflow execution event recording a phase transition.\x12\xc9\x01\n\x0f\x43reateNodeEvent\x12).flyteidl.admin.NodeExecutionEventRequest\x1a*.flyteidl.admin.NodeExecutionEventResponse\"_\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/events/nodes:\x01*\x92\x41=\x1a;Create a node execution event recording a phase transition.\x12\xc9\x01\n\x0f\x43reateTaskEvent\x12).flyteidl.admin.TaskExecutionEventRequest\x1a*.flyteidl.admin.TaskExecutionEventResponse\"_\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/events/tasks:\x01*\x92\x41=\x1a;Create a task execution event recording a phase transition.\x12\xa9\x03\n\x10GetTaskExecution\x12\'.flyteidl.admin.TaskExecutionGetRequest\x1a\x1d.flyteidl.admin.TaskExecution\"\xcc\x02\x82\xd3\xe4\x93\x02\x9c\x02\x12\x99\x02/api/v1/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}\x92\x41&\x1a$Retrieve an existing task execution.\x12\xd3\x02\n\x12ListTaskExecutions\x12(.flyteidl.admin.TaskExecutionListRequest\x1a!.flyteidl.admin.TaskExecutionList\"\xef\x01\x82\xd3\xe4\x93\x02\xad\x01\x12\xaa\x01/api/v1/task_executions/{node_execution_id.execution_id.project}/{node_execution_id.execution_id.domain}/{node_execution_id.execution_id.name}/{node_execution_id.node_id}\x92\x41\x38\x1a\x36\x46\x65tch existing task executions matching input filters.\x12\xe0\x03\n\x14GetTaskExecutionData\x12+.flyteidl.admin.TaskExecutionGetDataRequest\x1a,.flyteidl.admin.TaskExecutionGetDataResponse\"\xec\x02\x82\xd3\xe4\x93\x02\xa1\x02\x12\x9e\x02/api/v1/data/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}\x92\x41\x41\x1a?Retrieve input and output data from an existing task execution.\x12\x80\x02\n\x11ListNamedEntities\x12&.flyteidl.admin.NamedEntityListRequest\x1a\x1f.flyteidl.admin.NamedEntityList\"\xa1\x01\x82\xd3\xe4\x93\x02;\x12\x39/api/v1/named_entities/{resource_type}/{project}/{domain}\x92\x41]\x1a[Retrieve a list of NamedEntity objects sharing a common resource type, project, and domain.\x12\xca\x01\n\x0eGetNamedEntity\x12%.flyteidl.admin.NamedEntityGetRequest\x1a\x1b.flyteidl.admin.NamedEntity\"t\x82\xd3\xe4\x93\x02K\x12I/api/v1/named_entities/{resource_type}/{id.project}/{id.domain}/{id.name}\x92\x41 \x1a\x1eRetrieve a NamedEntity object.\x12\xf3\x01\n\x11UpdateNamedEntity\x12(.flyteidl.admin.NamedEntityUpdateRequest\x1a).flyteidl.admin.NamedEntityUpdateResponse\"\x88\x01\x82\xd3\xe4\x93\x02N\x1aI/api/v1/named_entities/{resource_type}/{id.project}/{id.domain}/{id.name}:\x01*\x92\x41\x31\x1a/Update the fields associated with a NamedEntityB5Z3github.com/lyft/flyteidl/gen/pb-go/flyteidl/serviceb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_project__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_task__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_workflow__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_launch__plan__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_event__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_execution__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_node__execution__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_task__execution__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_common__pb2.DESCRIPTOR,protoc__gen__swagger_dot_options_dot_annotations__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None

_ADMINSERVICE = _descriptor.ServiceDescriptor(
  name='AdminService',
  full_name='flyteidl.service.AdminService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=412,
  serialized_end=10458,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateTask',
    full_name='flyteidl.service.AdminService.CreateTask',
    index=0,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__pb2._TASKCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_task__pb2._TASKCREATERESPONSE,
    serialized_options=_b('\202\323\344\223\002\022\"\r/api/v1/tasks:\001*\222A\323\001\032&Create and register a task definition.JB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetTask',
    full_name='flyteidl.service.AdminService.GetTask',
    index=1,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._OBJECTGETREQUEST,
    output_type=flyteidl_dot_admin_dot_task__pb2._TASK,
    serialized_options=_b('\202\323\344\223\002?\022=/api/v1/tasks/{id.project}/{id.domain}/{id.name}/{id.version}\222A\'\032%Retrieve an existing task definition.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListTaskIds',
    full_name='flyteidl.service.AdminService.ListTaskIds',
    index=2,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLIST,
    serialized_options=_b('\202\323\344\223\002%\022#/api/v1/task_ids/{project}/{domain}\222AD\032BFetch existing task definition identifiers matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListTasks',
    full_name='flyteidl.service.AdminService.ListTasks',
    index=3,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_task__pb2._TASKLIST,
    serialized_options=_b('\202\323\344\223\002\\\0220/api/v1/tasks/{id.project}/{id.domain}/{id.name}Z(\022&/api/v1/tasks/{id.project}/{id.domain}\222A9\0327Fetch existing task definitions matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateWorkflow',
    full_name='flyteidl.service.AdminService.CreateWorkflow',
    index=4,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOWCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOWCREATERESPONSE,
    serialized_options=_b('\202\323\344\223\002\026\"\021/api/v1/workflows:\001*\222A\327\001\032*Create and register a workflow definition.JB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetWorkflow',
    full_name='flyteidl.service.AdminService.GetWorkflow',
    index=5,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._OBJECTGETREQUEST,
    output_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOW,
    serialized_options=_b('\202\323\344\223\002C\022A/api/v1/workflows/{id.project}/{id.domain}/{id.name}/{id.version}\222A+\032)Retrieve an existing workflow definition.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListWorkflowIds',
    full_name='flyteidl.service.AdminService.ListWorkflowIds',
    index=6,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLIST,
    serialized_options=_b('\202\323\344\223\002)\022\'/api/v1/workflow_ids/{project}/{domain}\222AK\032IFetch an existing workflow definition identifiers matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListWorkflows',
    full_name='flyteidl.service.AdminService.ListWorkflows',
    index=7,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOWLIST,
    serialized_options=_b('\202\323\344\223\002d\0224/api/v1/workflows/{id.project}/{id.domain}/{id.name}Z,\022*/api/v1/workflows/{id.project}/{id.domain}\222A=\032;Fetch existing workflow definitions matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateLaunchPlan',
    full_name='flyteidl.service.AdminService.CreateLaunchPlan',
    index=8,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANCREATERESPONSE,
    serialized_options=_b('\202\323\344\223\002\031\"\024/api/v1/launch_plans:\001*\222A\332\001\032-Create and register a launch plan definition.JB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetLaunchPlan',
    full_name='flyteidl.service.AdminService.GetLaunchPlan',
    index=9,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._OBJECTGETREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLAN,
    serialized_options=_b('\202\323\344\223\002F\022D/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}\222A.\032,Retrieve an existing launch plan definition.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetActiveLaunchPlan',
    full_name='flyteidl.service.AdminService.GetActiveLaunchPlan',
    index=10,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._ACTIVELAUNCHPLANREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLAN,
    serialized_options=_b('\202\323\344\223\002@\022>/api/v1/active_launch_plans/{id.project}/{id.domain}/{id.name}\222AM\032KRetrieve the active launch plan version specified by input request filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListActiveLaunchPlans',
    full_name='flyteidl.service.AdminService.ListActiveLaunchPlans',
    index=11,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._ACTIVELAUNCHPLANLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANLIST,
    serialized_options=_b('\202\323\344\223\0020\022./api/v1/active_launch_plans/{project}/{domain}\222AK\032IFetch the active launch plan versions specified by input request filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListLaunchPlanIds',
    full_name='flyteidl.service.AdminService.ListLaunchPlanIds',
    index=12,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLIST,
    serialized_options=_b('\202\323\344\223\002,\022*/api/v1/launch_plan_ids/{project}/{domain}\222AK\032IFetch existing launch plan definition identifiers matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListLaunchPlans',
    full_name='flyteidl.service.AdminService.ListLaunchPlans',
    index=13,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANLIST,
    serialized_options=_b('\202\323\344\223\002j\0227/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}Z/\022-/api/v1/launch_plans/{id.project}/{id.domain}\222A@\032>Fetch existing launch plan definitions matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='UpdateLaunchPlan',
    full_name='flyteidl.service.AdminService.UpdateLaunchPlan',
    index=14,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANUPDATEREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANUPDATERESPONSE,
    serialized_options=_b('\202\323\344\223\002I\032D/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}:\001*\222A\205\005\032\202\005Update the status of an existing launch plan definition. At most one launch plan version for a given {project, domain, name} can be active at a time. If this call sets a launch plan to active and existing version is already active, the result of this call will be that the formerly active launch plan will be made inactive and specified launch plan in this request will be made active. In the event that the formerly active launch plan had a schedule associated it with it, this schedule will be disabled. If the reference launch plan in this request is being set to active and has a schedule associated with it, the schedule will be enabled.'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateExecution',
    full_name='flyteidl.service.AdminService.CreateExecution',
    index=15,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONCREATERESPONSE,
    serialized_options=_b('\202\323\344\223\002\027\"\022/api/v1/executions:\001*\222A\036\032\034Create a workflow execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='RelaunchExecution',
    full_name='flyteidl.service.AdminService.RelaunchExecution',
    index=16,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONRELAUNCHREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONCREATERESPONSE,
    serialized_options=_b('\202\323\344\223\002 \"\033/api/v1/executions/relaunch:\001*\222A \032\036Relaunch a workflow execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetExecution',
    full_name='flyteidl.service.AdminService.GetExecution',
    index=17,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._WORKFLOWEXECUTIONGETREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTION,
    serialized_options=_b('\202\323\344\223\0027\0225/api/v1/executions/{id.project}/{id.domain}/{id.name}\222A*\032(Retrieve an existing workflow execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetExecutionData',
    full_name='flyteidl.service.AdminService.GetExecutionData',
    index=18,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._WORKFLOWEXECUTIONGETDATAREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._WORKFLOWEXECUTIONGETDATARESPONSE,
    serialized_options=_b('\202\323\344\223\002<\022:/api/v1/data/executions/{id.project}/{id.domain}/{id.name}\222AE\032CRetrieve input and output data from an existing workflow execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListExecutions',
    full_name='flyteidl.service.AdminService.ListExecutions',
    index=19,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONLIST,
    serialized_options=_b('\202\323\344\223\002-\022+/api/v1/executions/{id.project}/{id.domain}\222A<\032:Fetch existing workflow executions matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='TerminateExecution',
    full_name='flyteidl.service.AdminService.TerminateExecution',
    index=20,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONTERMINATEREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONTERMINATERESPONSE,
    serialized_options=_b('\202\323\344\223\002:*5/api/v1/executions/{id.project}/{id.domain}/{id.name}:\001*\222AC\032ATerminate the active workflow execution specified in the request.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetNodeExecution',
    full_name='flyteidl.service.AdminService.GetNodeExecution',
    index=21,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONGETREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTION,
    serialized_options=_b('\202\323\344\223\002p\022n/api/v1/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}\222A&\032$Retrieve an existing node execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListNodeExecutions',
    full_name='flyteidl.service.AdminService.ListNodeExecutions',
    index=22,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONLIST,
    serialized_options=_b('\202\323\344\223\002u\022s/api/v1/node_executions/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\222A8\0326Fetch existing node executions matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListNodeExecutionsForTask',
    full_name='flyteidl.service.AdminService.ListNodeExecutionsForTask',
    index=23,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONFORTASKLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONLIST,
    serialized_options=_b('\202\323\344\223\002\254\003\022\251\003/api/v1/children/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt}\222AG\032EFetch child node executions launched by the specified task execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetNodeExecutionData',
    full_name='flyteidl.service.AdminService.GetNodeExecutionData',
    index=24,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONGETDATAREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONGETDATARESPONSE,
    serialized_options=_b('\202\323\344\223\002u\022s/api/v1/data/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}\222AA\032?Retrieve input and output data from an existing node execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='RegisterProject',
    full_name='flyteidl.service.AdminService.RegisterProject',
    index=25,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_project__pb2._PROJECTREGISTERREQUEST,
    output_type=flyteidl_dot_admin_dot_project__pb2._PROJECTREGISTERRESPONSE,
    serialized_options=_b('\202\323\344\223\002\025\"\020/api/v1/projects:\001*\222A\025\032\023Register a project.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListProjects',
    full_name='flyteidl.service.AdminService.ListProjects',
    index=26,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_project__pb2._PROJECTLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_project__pb2._PROJECTS,
    serialized_options=_b('\202\323\344\223\002\022\022\020/api/v1/projects\222A\034\032\032Fetch registered projects.'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateWorkflowEvent',
    full_name='flyteidl.service.AdminService.CreateWorkflowEvent',
    index=27,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_event__pb2._WORKFLOWEXECUTIONEVENTREQUEST,
    output_type=flyteidl_dot_admin_dot_event__pb2._WORKFLOWEXECUTIONEVENTRESPONSE,
    serialized_options=_b('\202\323\344\223\002\035\"\030/api/v1/events/workflows:\001*\222AA\032?Create a workflow execution event recording a phase transition.'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateNodeEvent',
    full_name='flyteidl.service.AdminService.CreateNodeEvent',
    index=28,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_event__pb2._NODEEXECUTIONEVENTREQUEST,
    output_type=flyteidl_dot_admin_dot_event__pb2._NODEEXECUTIONEVENTRESPONSE,
    serialized_options=_b('\202\323\344\223\002\031\"\024/api/v1/events/nodes:\001*\222A=\032;Create a node execution event recording a phase transition.'),
  ),
  _descriptor.MethodDescriptor(
    name='CreateTaskEvent',
    full_name='flyteidl.service.AdminService.CreateTaskEvent',
    index=29,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_event__pb2._TASKEXECUTIONEVENTREQUEST,
    output_type=flyteidl_dot_admin_dot_event__pb2._TASKEXECUTIONEVENTRESPONSE,
    serialized_options=_b('\202\323\344\223\002\031\"\024/api/v1/events/tasks:\001*\222A=\032;Create a task execution event recording a phase transition.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetTaskExecution',
    full_name='flyteidl.service.AdminService.GetTaskExecution',
    index=30,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONGETREQUEST,
    output_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTION,
    serialized_options=_b('\202\323\344\223\002\234\002\022\231\002/api/v1/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}\222A&\032$Retrieve an existing task execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListTaskExecutions',
    full_name='flyteidl.service.AdminService.ListTaskExecutions',
    index=31,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONLIST,
    serialized_options=_b('\202\323\344\223\002\255\001\022\252\001/api/v1/task_executions/{node_execution_id.execution_id.project}/{node_execution_id.execution_id.domain}/{node_execution_id.execution_id.name}/{node_execution_id.node_id}\222A8\0326Fetch existing task executions matching input filters.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetTaskExecutionData',
    full_name='flyteidl.service.AdminService.GetTaskExecutionData',
    index=32,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONGETDATAREQUEST,
    output_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONGETDATARESPONSE,
    serialized_options=_b('\202\323\344\223\002\241\002\022\236\002/api/v1/data/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}\222AA\032?Retrieve input and output data from an existing task execution.'),
  ),
  _descriptor.MethodDescriptor(
    name='ListNamedEntities',
    full_name='flyteidl.service.AdminService.ListNamedEntities',
    index=33,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYLIST,
    serialized_options=_b('\202\323\344\223\002;\0229/api/v1/named_entities/{resource_type}/{project}/{domain}\222A]\032[Retrieve a list of NamedEntity objects sharing a common resource type, project, and domain.'),
  ),
  _descriptor.MethodDescriptor(
    name='GetNamedEntity',
    full_name='flyteidl.service.AdminService.GetNamedEntity',
    index=34,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYGETREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITY,
    serialized_options=_b('\202\323\344\223\002K\022I/api/v1/named_entities/{resource_type}/{id.project}/{id.domain}/{id.name}\222A \032\036Retrieve a NamedEntity object.'),
  ),
  _descriptor.MethodDescriptor(
    name='UpdateNamedEntity',
    full_name='flyteidl.service.AdminService.UpdateNamedEntity',
    index=35,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYUPDATEREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYUPDATERESPONSE,
    serialized_options=_b('\202\323\344\223\002N\032I/api/v1/named_entities/{resource_type}/{id.project}/{id.domain}/{id.name}:\001*\222A1\032/Update the fields associated with a NamedEntity'),
  ),
])
_sym_db.RegisterServiceDescriptor(_ADMINSERVICE)

DESCRIPTOR.services_by_name['AdminService'] = _ADMINSERVICE

# @@protoc_insertion_point(module_scope)
