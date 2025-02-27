// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// RuntimeClassList is a list of RuntimeClass objects.
type RuntimeClassList struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items RuntimeClassTypeArrayOutput `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrOutput `pulumi:"metadata"`
}

// NewRuntimeClassList registers a new resource with the given unique name, arguments, and options.
func NewRuntimeClassList(ctx *pulumi.Context,
	name string, args *RuntimeClassListArgs, opts ...pulumi.ResourceOption) (*RuntimeClassList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Items == nil {
		return nil, errors.New("invalid value for required argument 'Items'")
	}
	args.ApiVersion = pulumi.StringPtr("node.k8s.io/v1beta1")
	args.Kind = pulumi.StringPtr("RuntimeClassList")
	var resource RuntimeClassList
	err := ctx.RegisterResource("kubernetes:node.k8s.io/v1beta1:RuntimeClassList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeClassList gets an existing RuntimeClassList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeClassList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeClassListState, opts ...pulumi.ResourceOption) (*RuntimeClassList, error) {
	var resource RuntimeClassList
	err := ctx.ReadResource("kubernetes:node.k8s.io/v1beta1:RuntimeClassList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeClassList resources.
type runtimeClassListState struct {
}

type RuntimeClassListState struct {
}

func (RuntimeClassListState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassListState)(nil)).Elem()
}

type runtimeClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Items is a list of schema objects.
	Items []RuntimeClassType `pulumi:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ListMeta `pulumi:"metadata"`
}

// The set of arguments for constructing a RuntimeClassList resource.
type RuntimeClassListArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Items is a list of schema objects.
	Items RuntimeClassTypeArrayInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ListMetaPtrInput
}

func (RuntimeClassListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeClassListArgs)(nil)).Elem()
}

type RuntimeClassListInput interface {
	pulumi.Input

	ToRuntimeClassListOutput() RuntimeClassListOutput
	ToRuntimeClassListOutputWithContext(ctx context.Context) RuntimeClassListOutput
}

func (*RuntimeClassList) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassList)(nil))
}

func (i *RuntimeClassList) ToRuntimeClassListOutput() RuntimeClassListOutput {
	return i.ToRuntimeClassListOutputWithContext(context.Background())
}

func (i *RuntimeClassList) ToRuntimeClassListOutputWithContext(ctx context.Context) RuntimeClassListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListOutput)
}

func (i *RuntimeClassList) ToRuntimeClassListPtrOutput() RuntimeClassListPtrOutput {
	return i.ToRuntimeClassListPtrOutputWithContext(context.Background())
}

func (i *RuntimeClassList) ToRuntimeClassListPtrOutputWithContext(ctx context.Context) RuntimeClassListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListPtrOutput)
}

type RuntimeClassListPtrInput interface {
	pulumi.Input

	ToRuntimeClassListPtrOutput() RuntimeClassListPtrOutput
	ToRuntimeClassListPtrOutputWithContext(ctx context.Context) RuntimeClassListPtrOutput
}

type runtimeClassListPtrType RuntimeClassListArgs

func (*runtimeClassListPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClassList)(nil))
}

func (i *runtimeClassListPtrType) ToRuntimeClassListPtrOutput() RuntimeClassListPtrOutput {
	return i.ToRuntimeClassListPtrOutputWithContext(context.Background())
}

func (i *runtimeClassListPtrType) ToRuntimeClassListPtrOutputWithContext(ctx context.Context) RuntimeClassListPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListPtrOutput)
}

// RuntimeClassListArrayInput is an input type that accepts RuntimeClassListArray and RuntimeClassListArrayOutput values.
// You can construct a concrete instance of `RuntimeClassListArrayInput` via:
//
//          RuntimeClassListArray{ RuntimeClassListArgs{...} }
type RuntimeClassListArrayInput interface {
	pulumi.Input

	ToRuntimeClassListArrayOutput() RuntimeClassListArrayOutput
	ToRuntimeClassListArrayOutputWithContext(context.Context) RuntimeClassListArrayOutput
}

type RuntimeClassListArray []RuntimeClassListInput

func (RuntimeClassListArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuntimeClassList)(nil)).Elem()
}

func (i RuntimeClassListArray) ToRuntimeClassListArrayOutput() RuntimeClassListArrayOutput {
	return i.ToRuntimeClassListArrayOutputWithContext(context.Background())
}

func (i RuntimeClassListArray) ToRuntimeClassListArrayOutputWithContext(ctx context.Context) RuntimeClassListArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListArrayOutput)
}

// RuntimeClassListMapInput is an input type that accepts RuntimeClassListMap and RuntimeClassListMapOutput values.
// You can construct a concrete instance of `RuntimeClassListMapInput` via:
//
//          RuntimeClassListMap{ "key": RuntimeClassListArgs{...} }
type RuntimeClassListMapInput interface {
	pulumi.Input

	ToRuntimeClassListMapOutput() RuntimeClassListMapOutput
	ToRuntimeClassListMapOutputWithContext(context.Context) RuntimeClassListMapOutput
}

type RuntimeClassListMap map[string]RuntimeClassListInput

func (RuntimeClassListMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuntimeClassList)(nil)).Elem()
}

func (i RuntimeClassListMap) ToRuntimeClassListMapOutput() RuntimeClassListMapOutput {
	return i.ToRuntimeClassListMapOutputWithContext(context.Background())
}

func (i RuntimeClassListMap) ToRuntimeClassListMapOutputWithContext(ctx context.Context) RuntimeClassListMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeClassListMapOutput)
}

type RuntimeClassListOutput struct {
	*pulumi.OutputState
}

func (RuntimeClassListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuntimeClassList)(nil))
}

func (o RuntimeClassListOutput) ToRuntimeClassListOutput() RuntimeClassListOutput {
	return o
}

func (o RuntimeClassListOutput) ToRuntimeClassListOutputWithContext(ctx context.Context) RuntimeClassListOutput {
	return o
}

func (o RuntimeClassListOutput) ToRuntimeClassListPtrOutput() RuntimeClassListPtrOutput {
	return o.ToRuntimeClassListPtrOutputWithContext(context.Background())
}

func (o RuntimeClassListOutput) ToRuntimeClassListPtrOutputWithContext(ctx context.Context) RuntimeClassListPtrOutput {
	return o.ApplyT(func(v RuntimeClassList) *RuntimeClassList {
		return &v
	}).(RuntimeClassListPtrOutput)
}

type RuntimeClassListPtrOutput struct {
	*pulumi.OutputState
}

func (RuntimeClassListPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeClassList)(nil))
}

func (o RuntimeClassListPtrOutput) ToRuntimeClassListPtrOutput() RuntimeClassListPtrOutput {
	return o
}

func (o RuntimeClassListPtrOutput) ToRuntimeClassListPtrOutputWithContext(ctx context.Context) RuntimeClassListPtrOutput {
	return o
}

type RuntimeClassListArrayOutput struct{ *pulumi.OutputState }

func (RuntimeClassListArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuntimeClassList)(nil))
}

func (o RuntimeClassListArrayOutput) ToRuntimeClassListArrayOutput() RuntimeClassListArrayOutput {
	return o
}

func (o RuntimeClassListArrayOutput) ToRuntimeClassListArrayOutputWithContext(ctx context.Context) RuntimeClassListArrayOutput {
	return o
}

func (o RuntimeClassListArrayOutput) Index(i pulumi.IntInput) RuntimeClassListOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuntimeClassList {
		return vs[0].([]RuntimeClassList)[vs[1].(int)]
	}).(RuntimeClassListOutput)
}

type RuntimeClassListMapOutput struct{ *pulumi.OutputState }

func (RuntimeClassListMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RuntimeClassList)(nil))
}

func (o RuntimeClassListMapOutput) ToRuntimeClassListMapOutput() RuntimeClassListMapOutput {
	return o
}

func (o RuntimeClassListMapOutput) ToRuntimeClassListMapOutputWithContext(ctx context.Context) RuntimeClassListMapOutput {
	return o
}

func (o RuntimeClassListMapOutput) MapIndex(k pulumi.StringInput) RuntimeClassListOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RuntimeClassList {
		return vs[0].(map[string]RuntimeClassList)[vs[1].(string)]
	}).(RuntimeClassListOutput)
}

func init() {
	pulumi.RegisterOutputType(RuntimeClassListOutput{})
	pulumi.RegisterOutputType(RuntimeClassListPtrOutput{})
	pulumi.RegisterOutputType(RuntimeClassListArrayOutput{})
	pulumi.RegisterOutputType(RuntimeClassListMapOutput{})
}
