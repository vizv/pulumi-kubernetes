// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
type Event struct {
	pulumi.CustomResourceState

	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Action pulumi.StringPtrOutput `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount pulumi.IntPtrOutput `pulumi:"deprecatedCount"`
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp pulumi.StringPtrOutput `pulumi:"deprecatedFirstTimestamp"`
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp pulumi.StringPtrOutput `pulumi:"deprecatedLastTimestamp"`
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource corev1.EventSourcePtrOutput `pulumi:"deprecatedSource"`
	// eventTime is the time when this Event was first observed. It is required.
	EventTime pulumi.StringOutput `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrOutput `pulumi:"note"`
	// reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Reason pulumi.StringPtrOutput `pulumi:"reason"`
	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePtrOutput `pulumi:"regarding"`
	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePtrOutput `pulumi:"related"`
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController pulumi.StringPtrOutput `pulumi:"reportingController"`
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance pulumi.StringPtrOutput `pulumi:"reportingInstance"`
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPtrOutput `pulumi:"series"`
	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewEvent registers a new resource with the given unique name, arguments, and options.
func NewEvent(ctx *pulumi.Context,
	name string, args *EventArgs, opts ...pulumi.ResourceOption) (*Event, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventTime == nil {
		return nil, errors.New("invalid value for required argument 'EventTime'")
	}
	args.ApiVersion = pulumi.StringPtr("events.k8s.io/v1")
	args.Kind = pulumi.StringPtr("Event")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:core/v1:Event"),
		},
		{
			Type: pulumi.String("kubernetes:events.k8s.io/v1beta1:Event"),
		},
	})
	opts = append(opts, aliases)
	var resource Event
	err := ctx.RegisterResource("kubernetes:events.k8s.io/v1:Event", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEvent gets an existing Event resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEvent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventState, opts ...pulumi.ResourceOption) (*Event, error) {
	var resource Event
	err := ctx.ReadResource("kubernetes:events.k8s.io/v1:Event", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Event resources.
type eventState struct {
}

type EventState struct {
}

func (EventState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventState)(nil)).Elem()
}

type eventArgs struct {
	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Action *string `pulumi:"action"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount *int `pulumi:"deprecatedCount"`
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp *string `pulumi:"deprecatedFirstTimestamp"`
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp *string `pulumi:"deprecatedLastTimestamp"`
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource *corev1.EventSource `pulumi:"deprecatedSource"`
	// eventTime is the time when this Event was first observed. It is required.
	EventTime string `pulumi:"eventTime"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note *string `pulumi:"note"`
	// reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Reason *string `pulumi:"reason"`
	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *corev1.ObjectReference `pulumi:"regarding"`
	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *corev1.ObjectReference `pulumi:"related"`
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController *string `pulumi:"reportingController"`
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance *string `pulumi:"reportingInstance"`
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `pulumi:"series"`
	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Event resource.
type EventArgs struct {
	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Action pulumi.StringPtrInput
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount pulumi.IntPtrInput
	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp pulumi.StringPtrInput
	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp pulumi.StringPtrInput
	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource corev1.EventSourcePtrInput
	// eventTime is the time when this Event was first observed. It is required.
	EventTime pulumi.StringInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note pulumi.StringPtrInput
	// reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
	Reason pulumi.StringPtrInput
	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding corev1.ObjectReferencePtrInput
	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related corev1.ObjectReferencePtrInput
	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController pulumi.StringPtrInput
	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance pulumi.StringPtrInput
	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series EventSeriesPtrInput
	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
	Type pulumi.StringPtrInput
}

func (EventArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventArgs)(nil)).Elem()
}

type EventInput interface {
	pulumi.Input

	ToEventOutput() EventOutput
	ToEventOutputWithContext(ctx context.Context) EventOutput
}

func (*Event) ElementType() reflect.Type {
	return reflect.TypeOf((*Event)(nil))
}

func (i *Event) ToEventOutput() EventOutput {
	return i.ToEventOutputWithContext(context.Background())
}

func (i *Event) ToEventOutputWithContext(ctx context.Context) EventOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventOutput)
}

func (i *Event) ToEventPtrOutput() EventPtrOutput {
	return i.ToEventPtrOutputWithContext(context.Background())
}

func (i *Event) ToEventPtrOutputWithContext(ctx context.Context) EventPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventPtrOutput)
}

type EventPtrInput interface {
	pulumi.Input

	ToEventPtrOutput() EventPtrOutput
	ToEventPtrOutputWithContext(ctx context.Context) EventPtrOutput
}

type eventPtrType EventArgs

func (*eventPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Event)(nil))
}

func (i *eventPtrType) ToEventPtrOutput() EventPtrOutput {
	return i.ToEventPtrOutputWithContext(context.Background())
}

func (i *eventPtrType) ToEventPtrOutputWithContext(ctx context.Context) EventPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventPtrOutput)
}

// EventArrayInput is an input type that accepts EventArray and EventArrayOutput values.
// You can construct a concrete instance of `EventArrayInput` via:
//
//          EventArray{ EventArgs{...} }
type EventArrayInput interface {
	pulumi.Input

	ToEventArrayOutput() EventArrayOutput
	ToEventArrayOutputWithContext(context.Context) EventArrayOutput
}

type EventArray []EventInput

func (EventArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Event)(nil)).Elem()
}

func (i EventArray) ToEventArrayOutput() EventArrayOutput {
	return i.ToEventArrayOutputWithContext(context.Background())
}

func (i EventArray) ToEventArrayOutputWithContext(ctx context.Context) EventArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventArrayOutput)
}

// EventMapInput is an input type that accepts EventMap and EventMapOutput values.
// You can construct a concrete instance of `EventMapInput` via:
//
//          EventMap{ "key": EventArgs{...} }
type EventMapInput interface {
	pulumi.Input

	ToEventMapOutput() EventMapOutput
	ToEventMapOutputWithContext(context.Context) EventMapOutput
}

type EventMap map[string]EventInput

func (EventMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Event)(nil)).Elem()
}

func (i EventMap) ToEventMapOutput() EventMapOutput {
	return i.ToEventMapOutputWithContext(context.Background())
}

func (i EventMap) ToEventMapOutputWithContext(ctx context.Context) EventMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventMapOutput)
}

type EventOutput struct {
	*pulumi.OutputState
}

func (EventOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Event)(nil))
}

func (o EventOutput) ToEventOutput() EventOutput {
	return o
}

func (o EventOutput) ToEventOutputWithContext(ctx context.Context) EventOutput {
	return o
}

func (o EventOutput) ToEventPtrOutput() EventPtrOutput {
	return o.ToEventPtrOutputWithContext(context.Background())
}

func (o EventOutput) ToEventPtrOutputWithContext(ctx context.Context) EventPtrOutput {
	return o.ApplyT(func(v Event) *Event {
		return &v
	}).(EventPtrOutput)
}

type EventPtrOutput struct {
	*pulumi.OutputState
}

func (EventPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Event)(nil))
}

func (o EventPtrOutput) ToEventPtrOutput() EventPtrOutput {
	return o
}

func (o EventPtrOutput) ToEventPtrOutputWithContext(ctx context.Context) EventPtrOutput {
	return o
}

type EventArrayOutput struct{ *pulumi.OutputState }

func (EventArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Event)(nil))
}

func (o EventArrayOutput) ToEventArrayOutput() EventArrayOutput {
	return o
}

func (o EventArrayOutput) ToEventArrayOutputWithContext(ctx context.Context) EventArrayOutput {
	return o
}

func (o EventArrayOutput) Index(i pulumi.IntInput) EventOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Event {
		return vs[0].([]Event)[vs[1].(int)]
	}).(EventOutput)
}

type EventMapOutput struct{ *pulumi.OutputState }

func (EventMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Event)(nil))
}

func (o EventMapOutput) ToEventMapOutput() EventMapOutput {
	return o
}

func (o EventMapOutput) ToEventMapOutputWithContext(ctx context.Context) EventMapOutput {
	return o
}

func (o EventMapOutput) MapIndex(k pulumi.StringInput) EventOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Event {
		return vs[0].(map[string]Event)[vs[1].(string)]
	}).(EventOutput)
}

func init() {
	pulumi.RegisterOutputType(EventOutput{})
	pulumi.RegisterOutputType(EventPtrOutput{})
	pulumi.RegisterOutputType(EventArrayOutput{})
	pulumi.RegisterOutputType(EventMapOutput{})
}
