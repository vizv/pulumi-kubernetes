// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list of actions may be incomplete depending on the server's authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization decisions to the API server.
type SelfSubjectRulesReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates the set of actions a user can perform.
	Status SubjectRulesReviewStatusPtrOutput `pulumi:"status"`
}

// NewSelfSubjectRulesReview registers a new resource with the given unique name, arguments, and options.
func NewSelfSubjectRulesReview(ctx *pulumi.Context,
	name string, args *SelfSubjectRulesReviewArgs, opts ...pulumi.ResourceOption) (*SelfSubjectRulesReview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("SelfSubjectRulesReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectRulesReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SelfSubjectRulesReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectRulesReview gets an existing SelfSubjectRulesReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSubjectRulesReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSubjectRulesReviewState, opts ...pulumi.ResourceOption) (*SelfSubjectRulesReview, error) {
	var resource SelfSubjectRulesReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1:SelfSubjectRulesReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSubjectRulesReview resources.
type selfSubjectRulesReviewState struct {
}

type SelfSubjectRulesReviewState struct {
}

func (SelfSubjectRulesReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectRulesReviewState)(nil)).Elem()
}

type selfSubjectRulesReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a SelfSubjectRulesReview resource.
type SelfSubjectRulesReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.
	Spec SelfSubjectRulesReviewSpecInput
}

func (SelfSubjectRulesReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectRulesReviewArgs)(nil)).Elem()
}

type SelfSubjectRulesReviewInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewOutput() SelfSubjectRulesReviewOutput
	ToSelfSubjectRulesReviewOutputWithContext(ctx context.Context) SelfSubjectRulesReviewOutput
}

func (*SelfSubjectRulesReview) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectRulesReview)(nil))
}

func (i *SelfSubjectRulesReview) ToSelfSubjectRulesReviewOutput() SelfSubjectRulesReviewOutput {
	return i.ToSelfSubjectRulesReviewOutputWithContext(context.Background())
}

func (i *SelfSubjectRulesReview) ToSelfSubjectRulesReviewOutputWithContext(ctx context.Context) SelfSubjectRulesReviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewOutput)
}

func (i *SelfSubjectRulesReview) ToSelfSubjectRulesReviewPtrOutput() SelfSubjectRulesReviewPtrOutput {
	return i.ToSelfSubjectRulesReviewPtrOutputWithContext(context.Background())
}

func (i *SelfSubjectRulesReview) ToSelfSubjectRulesReviewPtrOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewPtrOutput)
}

type SelfSubjectRulesReviewPtrInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewPtrOutput() SelfSubjectRulesReviewPtrOutput
	ToSelfSubjectRulesReviewPtrOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPtrOutput
}

type selfSubjectRulesReviewPtrType SelfSubjectRulesReviewArgs

func (*selfSubjectRulesReviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectRulesReview)(nil))
}

func (i *selfSubjectRulesReviewPtrType) ToSelfSubjectRulesReviewPtrOutput() SelfSubjectRulesReviewPtrOutput {
	return i.ToSelfSubjectRulesReviewPtrOutputWithContext(context.Background())
}

func (i *selfSubjectRulesReviewPtrType) ToSelfSubjectRulesReviewPtrOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewPtrOutput)
}

// SelfSubjectRulesReviewArrayInput is an input type that accepts SelfSubjectRulesReviewArray and SelfSubjectRulesReviewArrayOutput values.
// You can construct a concrete instance of `SelfSubjectRulesReviewArrayInput` via:
//
//          SelfSubjectRulesReviewArray{ SelfSubjectRulesReviewArgs{...} }
type SelfSubjectRulesReviewArrayInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewArrayOutput() SelfSubjectRulesReviewArrayOutput
	ToSelfSubjectRulesReviewArrayOutputWithContext(context.Context) SelfSubjectRulesReviewArrayOutput
}

type SelfSubjectRulesReviewArray []SelfSubjectRulesReviewInput

func (SelfSubjectRulesReviewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectRulesReview)(nil)).Elem()
}

func (i SelfSubjectRulesReviewArray) ToSelfSubjectRulesReviewArrayOutput() SelfSubjectRulesReviewArrayOutput {
	return i.ToSelfSubjectRulesReviewArrayOutputWithContext(context.Background())
}

func (i SelfSubjectRulesReviewArray) ToSelfSubjectRulesReviewArrayOutputWithContext(ctx context.Context) SelfSubjectRulesReviewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewArrayOutput)
}

// SelfSubjectRulesReviewMapInput is an input type that accepts SelfSubjectRulesReviewMap and SelfSubjectRulesReviewMapOutput values.
// You can construct a concrete instance of `SelfSubjectRulesReviewMapInput` via:
//
//          SelfSubjectRulesReviewMap{ "key": SelfSubjectRulesReviewArgs{...} }
type SelfSubjectRulesReviewMapInput interface {
	pulumi.Input

	ToSelfSubjectRulesReviewMapOutput() SelfSubjectRulesReviewMapOutput
	ToSelfSubjectRulesReviewMapOutputWithContext(context.Context) SelfSubjectRulesReviewMapOutput
}

type SelfSubjectRulesReviewMap map[string]SelfSubjectRulesReviewInput

func (SelfSubjectRulesReviewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectRulesReview)(nil)).Elem()
}

func (i SelfSubjectRulesReviewMap) ToSelfSubjectRulesReviewMapOutput() SelfSubjectRulesReviewMapOutput {
	return i.ToSelfSubjectRulesReviewMapOutputWithContext(context.Background())
}

func (i SelfSubjectRulesReviewMap) ToSelfSubjectRulesReviewMapOutputWithContext(ctx context.Context) SelfSubjectRulesReviewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectRulesReviewMapOutput)
}

type SelfSubjectRulesReviewOutput struct {
	*pulumi.OutputState
}

func (SelfSubjectRulesReviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectRulesReview)(nil))
}

func (o SelfSubjectRulesReviewOutput) ToSelfSubjectRulesReviewOutput() SelfSubjectRulesReviewOutput {
	return o
}

func (o SelfSubjectRulesReviewOutput) ToSelfSubjectRulesReviewOutputWithContext(ctx context.Context) SelfSubjectRulesReviewOutput {
	return o
}

func (o SelfSubjectRulesReviewOutput) ToSelfSubjectRulesReviewPtrOutput() SelfSubjectRulesReviewPtrOutput {
	return o.ToSelfSubjectRulesReviewPtrOutputWithContext(context.Background())
}

func (o SelfSubjectRulesReviewOutput) ToSelfSubjectRulesReviewPtrOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPtrOutput {
	return o.ApplyT(func(v SelfSubjectRulesReview) *SelfSubjectRulesReview {
		return &v
	}).(SelfSubjectRulesReviewPtrOutput)
}

type SelfSubjectRulesReviewPtrOutput struct {
	*pulumi.OutputState
}

func (SelfSubjectRulesReviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectRulesReview)(nil))
}

func (o SelfSubjectRulesReviewPtrOutput) ToSelfSubjectRulesReviewPtrOutput() SelfSubjectRulesReviewPtrOutput {
	return o
}

func (o SelfSubjectRulesReviewPtrOutput) ToSelfSubjectRulesReviewPtrOutputWithContext(ctx context.Context) SelfSubjectRulesReviewPtrOutput {
	return o
}

type SelfSubjectRulesReviewArrayOutput struct{ *pulumi.OutputState }

func (SelfSubjectRulesReviewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelfSubjectRulesReview)(nil))
}

func (o SelfSubjectRulesReviewArrayOutput) ToSelfSubjectRulesReviewArrayOutput() SelfSubjectRulesReviewArrayOutput {
	return o
}

func (o SelfSubjectRulesReviewArrayOutput) ToSelfSubjectRulesReviewArrayOutputWithContext(ctx context.Context) SelfSubjectRulesReviewArrayOutput {
	return o
}

func (o SelfSubjectRulesReviewArrayOutput) Index(i pulumi.IntInput) SelfSubjectRulesReviewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelfSubjectRulesReview {
		return vs[0].([]SelfSubjectRulesReview)[vs[1].(int)]
	}).(SelfSubjectRulesReviewOutput)
}

type SelfSubjectRulesReviewMapOutput struct{ *pulumi.OutputState }

func (SelfSubjectRulesReviewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SelfSubjectRulesReview)(nil))
}

func (o SelfSubjectRulesReviewMapOutput) ToSelfSubjectRulesReviewMapOutput() SelfSubjectRulesReviewMapOutput {
	return o
}

func (o SelfSubjectRulesReviewMapOutput) ToSelfSubjectRulesReviewMapOutputWithContext(ctx context.Context) SelfSubjectRulesReviewMapOutput {
	return o
}

func (o SelfSubjectRulesReviewMapOutput) MapIndex(k pulumi.StringInput) SelfSubjectRulesReviewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SelfSubjectRulesReview {
		return vs[0].(map[string]SelfSubjectRulesReview)[vs[1].(string)]
	}).(SelfSubjectRulesReviewOutput)
}

func init() {
	pulumi.RegisterOutputType(SelfSubjectRulesReviewOutput{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewPtrOutput{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewArrayOutput{})
	pulumi.RegisterOutputType(SelfSubjectRulesReviewMapOutput{})
}
