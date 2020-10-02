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

from flyteadmin.models.admin_cron_schedule_with_offset import AdminCronScheduleWithOffset  # noqa: F401,E501
from flyteadmin.models.admin_fixed_rate import AdminFixedRate  # noqa: F401,E501


class AdminSchedule(object):
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
        'cron_expression': 'str',
        'rate': 'AdminFixedRate',
        'cron_schedule_with_offset': 'AdminCronScheduleWithOffset',
        'kickoff_time_input_arg': 'str'
    }

    attribute_map = {
        'cron_expression': 'cron_expression',
        'rate': 'rate',
        'cron_schedule_with_offset': 'cron_schedule_with_offset',
        'kickoff_time_input_arg': 'kickoff_time_input_arg'
    }

    def __init__(self, cron_expression=None, rate=None, cron_schedule_with_offset=None, kickoff_time_input_arg=None):  # noqa: E501
        """AdminSchedule - a model defined in Swagger"""  # noqa: E501

        self._cron_expression = None
        self._rate = None
        self._cron_schedule_with_offset = None
        self._kickoff_time_input_arg = None
        self.discriminator = None

        if cron_expression is not None:
            self.cron_expression = cron_expression
        if rate is not None:
            self.rate = rate
        if cron_schedule_with_offset is not None:
            self.cron_schedule_with_offset = cron_schedule_with_offset
        if kickoff_time_input_arg is not None:
            self.kickoff_time_input_arg = kickoff_time_input_arg

    @property
    def cron_expression(self):
        """Gets the cron_expression of this AdminSchedule.  # noqa: E501


        :return: The cron_expression of this AdminSchedule.  # noqa: E501
        :rtype: str
        """
        return self._cron_expression

    @cron_expression.setter
    def cron_expression(self, cron_expression):
        """Sets the cron_expression of this AdminSchedule.


        :param cron_expression: The cron_expression of this AdminSchedule.  # noqa: E501
        :type: str
        """

        self._cron_expression = cron_expression

    @property
    def rate(self):
        """Gets the rate of this AdminSchedule.  # noqa: E501


        :return: The rate of this AdminSchedule.  # noqa: E501
        :rtype: AdminFixedRate
        """
        return self._rate

    @rate.setter
    def rate(self, rate):
        """Sets the rate of this AdminSchedule.


        :param rate: The rate of this AdminSchedule.  # noqa: E501
        :type: AdminFixedRate
        """

        self._rate = rate

    @property
    def cron_schedule_with_offset(self):
        """Gets the cron_schedule_with_offset of this AdminSchedule.  # noqa: E501


        :return: The cron_schedule_with_offset of this AdminSchedule.  # noqa: E501
        :rtype: AdminCronScheduleWithOffset
        """
        return self._cron_schedule_with_offset

    @cron_schedule_with_offset.setter
    def cron_schedule_with_offset(self, cron_schedule_with_offset):
        """Sets the cron_schedule_with_offset of this AdminSchedule.


        :param cron_schedule_with_offset: The cron_schedule_with_offset of this AdminSchedule.  # noqa: E501
        :type: AdminCronScheduleWithOffset
        """

        self._cron_schedule_with_offset = cron_schedule_with_offset

    @property
    def kickoff_time_input_arg(self):
        """Gets the kickoff_time_input_arg of this AdminSchedule.  # noqa: E501

        Name of the input variable that the kickoff time will be supplied to when the workflow is kicked off.  # noqa: E501

        :return: The kickoff_time_input_arg of this AdminSchedule.  # noqa: E501
        :rtype: str
        """
        return self._kickoff_time_input_arg

    @kickoff_time_input_arg.setter
    def kickoff_time_input_arg(self, kickoff_time_input_arg):
        """Sets the kickoff_time_input_arg of this AdminSchedule.

        Name of the input variable that the kickoff time will be supplied to when the workflow is kicked off.  # noqa: E501

        :param kickoff_time_input_arg: The kickoff_time_input_arg of this AdminSchedule.  # noqa: E501
        :type: str
        """

        self._kickoff_time_input_arg = kickoff_time_input_arg

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
        if issubclass(AdminSchedule, dict):
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
        if not isinstance(other, AdminSchedule):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
