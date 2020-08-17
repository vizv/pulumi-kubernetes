# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from ... import core as _core
from ... import meta as _meta

__all__ = [
    'EventArgs',
    'EventSeriesArgs',
]

@pulumi.input_type
class EventArgs:
    def __init__(__self__, *,
                 event_time: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 deprecated_count: Optional[pulumi.Input[float]] = None,
                 deprecated_first_timestamp: Optional[pulumi.Input[str]] = None,
                 deprecated_last_timestamp: Optional[pulumi.Input[str]] = None,
                 deprecated_source: Optional[pulumi.Input['_core.v1.EventSourceArgs']] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']] = None,
                 note: Optional[pulumi.Input[str]] = None,
                 reason: Optional[pulumi.Input[str]] = None,
                 regarding: Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']] = None,
                 related: Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']] = None,
                 reporting_controller: Optional[pulumi.Input[str]] = None,
                 reporting_instance: Optional[pulumi.Input[str]] = None,
                 series: Optional[pulumi.Input['EventSeriesArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
        :param pulumi.Input[str] event_time: Required. Time when this Event was first observed.
        :param pulumi.Input[str] action: What action was taken/failed regarding to the regarding object.
        :param pulumi.Input[str] api_version: APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        :param pulumi.Input[float] deprecated_count: Deprecated field assuring backward compatibility with core.v1 Event type
        :param pulumi.Input[str] deprecated_first_timestamp: Deprecated field assuring backward compatibility with core.v1 Event type
        :param pulumi.Input[str] deprecated_last_timestamp: Deprecated field assuring backward compatibility with core.v1 Event type
        :param pulumi.Input['_core.v1.EventSourceArgs'] deprecated_source: Deprecated field assuring backward compatibility with core.v1 Event type
        :param pulumi.Input[str] kind: Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        :param pulumi.Input[str] note: Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
        :param pulumi.Input[str] reason: Why the action was taken.
        :param pulumi.Input['_core.v1.ObjectReferenceArgs'] regarding: The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
        :param pulumi.Input['_core.v1.ObjectReferenceArgs'] related: Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
        :param pulumi.Input[str] reporting_controller: Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        :param pulumi.Input[str] reporting_instance: ID of the controller instance, e.g. `kubelet-xyzf`.
        :param pulumi.Input['EventSeriesArgs'] series: Data about the Event series this event represents or nil if it's a singleton Event.
        :param pulumi.Input[str] type: Type of this event (Normal, Warning), new types could be added in the future.
        """
        pulumi.set(__self__, "event_time", event_time)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if api_version is not None:
            pulumi.set(__self__, "api_version", 'events.k8s.io/v1beta1')
        if deprecated_count is not None:
            pulumi.set(__self__, "deprecated_count", deprecated_count)
        if deprecated_first_timestamp is not None:
            pulumi.set(__self__, "deprecated_first_timestamp", deprecated_first_timestamp)
        if deprecated_last_timestamp is not None:
            pulumi.set(__self__, "deprecated_last_timestamp", deprecated_last_timestamp)
        if deprecated_source is not None:
            pulumi.set(__self__, "deprecated_source", deprecated_source)
        if kind is not None:
            pulumi.set(__self__, "kind", 'Event')
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if note is not None:
            pulumi.set(__self__, "note", note)
        if reason is not None:
            pulumi.set(__self__, "reason", reason)
        if regarding is not None:
            pulumi.set(__self__, "regarding", regarding)
        if related is not None:
            pulumi.set(__self__, "related", related)
        if reporting_controller is not None:
            pulumi.set(__self__, "reporting_controller", reporting_controller)
        if reporting_instance is not None:
            pulumi.set(__self__, "reporting_instance", reporting_instance)
        if series is not None:
            pulumi.set(__self__, "series", series)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="eventTime")
    def event_time(self) -> pulumi.Input[str]:
        """
        Required. Time when this Event was first observed.
        """
        ...

    @event_time.setter
    def event_time(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        What action was taken/failed regarding to the regarding object.
        """
        ...

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        """
        ...

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="deprecatedCount")
    def deprecated_count(self) -> Optional[pulumi.Input[float]]:
        """
        Deprecated field assuring backward compatibility with core.v1 Event type
        """
        ...

    @deprecated_count.setter
    def deprecated_count(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="deprecatedFirstTimestamp")
    def deprecated_first_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated field assuring backward compatibility with core.v1 Event type
        """
        ...

    @deprecated_first_timestamp.setter
    def deprecated_first_timestamp(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="deprecatedLastTimestamp")
    def deprecated_last_timestamp(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated field assuring backward compatibility with core.v1 Event type
        """
        ...

    @deprecated_last_timestamp.setter
    def deprecated_last_timestamp(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="deprecatedSource")
    def deprecated_source(self) -> Optional[pulumi.Input['_core.v1.EventSourceArgs']]:
        """
        Deprecated field assuring backward compatibility with core.v1 Event type
        """
        ...

    @deprecated_source.setter
    def deprecated_source(self, value: Optional[pulumi.Input['_core.v1.EventSourceArgs']]):
        ...

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input[str]]:
        """
        Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        """
        ...

    @kind.setter
    def kind(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]:
        ...

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input['_meta.v1.ObjectMetaArgs']]):
        ...

    @property
    @pulumi.getter
    def note(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
        """
        ...

    @note.setter
    def note(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def reason(self) -> Optional[pulumi.Input[str]]:
        """
        Why the action was taken.
        """
        ...

    @reason.setter
    def reason(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def regarding(self) -> Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']]:
        """
        The object this Event is about. In most cases it's an Object reporting controller implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
        """
        ...

    @regarding.setter
    def regarding(self, value: Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']]):
        ...

    @property
    @pulumi.getter
    def related(self) -> Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']]:
        """
        Optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
        """
        ...

    @related.setter
    def related(self, value: Optional[pulumi.Input['_core.v1.ObjectReferenceArgs']]):
        ...

    @property
    @pulumi.getter(name="reportingController")
    def reporting_controller(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        """
        ...

    @reporting_controller.setter
    def reporting_controller(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="reportingInstance")
    def reporting_instance(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the controller instance, e.g. `kubelet-xyzf`.
        """
        ...

    @reporting_instance.setter
    def reporting_instance(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def series(self) -> Optional[pulumi.Input['EventSeriesArgs']]:
        """
        Data about the Event series this event represents or nil if it's a singleton Event.
        """
        ...

    @series.setter
    def series(self, value: Optional[pulumi.Input['EventSeriesArgs']]):
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of this event (Normal, Warning), new types could be added in the future.
        """
        ...

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class EventSeriesArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float],
                 last_observed_time: pulumi.Input[str],
                 state: pulumi.Input[str]):
        """
        EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
        :param pulumi.Input[float] count: Number of occurrences in this series up to the last heartbeat time
        :param pulumi.Input[str] last_observed_time: Time when last Event from the series was seen before last heartbeat.
        :param pulumi.Input[str] state: Information whether this series is ongoing or finished. Deprecated. Planned removal for 1.18
        """
        pulumi.set(__self__, "count", count)
        pulumi.set(__self__, "last_observed_time", last_observed_time)
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        Number of occurrences in this series up to the last heartbeat time
        """
        ...

    @count.setter
    def count(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="lastObservedTime")
    def last_observed_time(self) -> pulumi.Input[str]:
        """
        Time when last Event from the series was seen before last heartbeat.
        """
        ...

    @last_observed_time.setter
    def last_observed_time(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter
    def state(self) -> pulumi.Input[str]:
        """
        Information whether this series is ongoing or finished. Deprecated. Planned removal for 1.18
        """
        ...

    @state.setter
    def state(self, value: pulumi.Input[str]):
        ...


