# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.admin_execution_metadata import AdminExecutionMetadata  # noqa: F401,E501
from flyteadmin.models.core_workflow_execution_identifier import CoreWorkflowExecutionIdentifier  # noqa: F401,E501


class AdminExecutionRecoverRequest(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'id': 'CoreWorkflowExecutionIdentifier',
        'name': 'str',
        'metadata': 'AdminExecutionMetadata'
    }

    attribute_map = {
        'id': 'id',
        'name': 'name',
        'metadata': 'metadata'
    }

    def __init__(self, id=None, name=None, metadata=None):  # noqa: E501
        """AdminExecutionRecoverRequest - a model defined in Swagger"""  # noqa: E501

        self._id = None
        self._name = None
        self._metadata = None
        self.discriminator = None

        if id is not None:
            self.id = id
        if name is not None:
            self.name = name
        if metadata is not None:
            self.metadata = metadata

    @property
    def id(self):
        """Gets the id of this AdminExecutionRecoverRequest.  # noqa: E501

        Identifier of the workflow execution to recover.  # noqa: E501

        :return: The id of this AdminExecutionRecoverRequest.  # noqa: E501
        :rtype: CoreWorkflowExecutionIdentifier
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this AdminExecutionRecoverRequest.

        Identifier of the workflow execution to recover.  # noqa: E501

        :param id: The id of this AdminExecutionRecoverRequest.  # noqa: E501
        :type: CoreWorkflowExecutionIdentifier
        """

        self._id = id

    @property
    def name(self):
        """Gets the name of this AdminExecutionRecoverRequest.  # noqa: E501


        :return: The name of this AdminExecutionRecoverRequest.  # noqa: E501
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """Sets the name of this AdminExecutionRecoverRequest.


        :param name: The name of this AdminExecutionRecoverRequest.  # noqa: E501
        :type: str
        """

        self._name = name

    @property
    def metadata(self):
        """Gets the metadata of this AdminExecutionRecoverRequest.  # noqa: E501

        Additional metadata which will be used to overwrite any metadata in the reference execution when triggering a recovery execution.  # noqa: E501

        :return: The metadata of this AdminExecutionRecoverRequest.  # noqa: E501
        :rtype: AdminExecutionMetadata
        """
        return self._metadata

    @metadata.setter
    def metadata(self, metadata):
        """Sets the metadata of this AdminExecutionRecoverRequest.

        Additional metadata which will be used to overwrite any metadata in the reference execution when triggering a recovery execution.  # noqa: E501

        :param metadata: The metadata of this AdminExecutionRecoverRequest.  # noqa: E501
        :type: AdminExecutionMetadata
        """

        self._metadata = metadata

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminExecutionRecoverRequest, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminExecutionRecoverRequest):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
