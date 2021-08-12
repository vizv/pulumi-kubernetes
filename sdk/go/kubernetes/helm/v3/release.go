// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Release is an instance of a chart running in a Kubernetes cluster.
//
// A Chart is a Helm package. It contains all of the resource definitions necessary to run an application, tool, or service inside of a Kubernetes cluster.
type Release struct {
	pulumi.CustomResourceState

	ReleaseSpec  ReleaseSpecOutput      `pulumi:"releaseSpec"`
	ResourceType pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Status of the deployed release.
	Status ReleaseStatusOutput `pulumi:"status"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReleaseSpec == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseSpec'")
	}
	args.Compat = pulumi.StringPtr("true")
	var resource Release
	err := ctx.RegisterResource("kubernetes:helm.sh/v3:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("kubernetes:helm.sh/v3:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
}

type ReleaseState struct {
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	Compat      *string     `pulumi:"compat"`
	ReleaseSpec ReleaseSpec `pulumi:"releaseSpec"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	Compat      pulumi.StringPtrInput
	ReleaseSpec ReleaseSpecInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((*Release)(nil))
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

func (i *Release) ToReleasePtrOutput() ReleasePtrOutput {
	return i.ToReleasePtrOutputWithContext(context.Background())
}

func (i *Release) ToReleasePtrOutputWithContext(ctx context.Context) ReleasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleasePtrOutput)
}

type ReleasePtrInput interface {
	pulumi.Input

	ToReleasePtrOutput() ReleasePtrOutput
	ToReleasePtrOutputWithContext(ctx context.Context) ReleasePtrOutput
}

type releasePtrType ReleaseArgs

func (*releasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil))
}

func (i *releasePtrType) ToReleasePtrOutput() ReleasePtrOutput {
	return i.ToReleasePtrOutputWithContext(context.Background())
}

func (i *releasePtrType) ToReleasePtrOutputWithContext(ctx context.Context) ReleasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleasePtrOutput)
}

// ReleaseArrayInput is an input type that accepts ReleaseArray and ReleaseArrayOutput values.
// You can construct a concrete instance of `ReleaseArrayInput` via:
//
//          ReleaseArray{ ReleaseArgs{...} }
type ReleaseArrayInput interface {
	pulumi.Input

	ToReleaseArrayOutput() ReleaseArrayOutput
	ToReleaseArrayOutputWithContext(context.Context) ReleaseArrayOutput
}

type ReleaseArray []ReleaseInput

func (ReleaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Release)(nil)).Elem()
}

func (i ReleaseArray) ToReleaseArrayOutput() ReleaseArrayOutput {
	return i.ToReleaseArrayOutputWithContext(context.Background())
}

func (i ReleaseArray) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseArrayOutput)
}

// ReleaseMapInput is an input type that accepts ReleaseMap and ReleaseMapOutput values.
// You can construct a concrete instance of `ReleaseMapInput` via:
//
//          ReleaseMap{ "key": ReleaseArgs{...} }
type ReleaseMapInput interface {
	pulumi.Input

	ToReleaseMapOutput() ReleaseMapOutput
	ToReleaseMapOutputWithContext(context.Context) ReleaseMapOutput
}

type ReleaseMap map[string]ReleaseInput

func (ReleaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Release)(nil)).Elem()
}

func (i ReleaseMap) ToReleaseMapOutput() ReleaseMapOutput {
	return i.ToReleaseMapOutputWithContext(context.Background())
}

func (i ReleaseMap) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseMapOutput)
}

type ReleaseOutput struct {
	*pulumi.OutputState
}

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Release)(nil))
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleasePtrOutput() ReleasePtrOutput {
	return o.ToReleasePtrOutputWithContext(context.Background())
}

func (o ReleaseOutput) ToReleasePtrOutputWithContext(ctx context.Context) ReleasePtrOutput {
	return o.ApplyT(func(v Release) *Release {
		return &v
	}).(ReleasePtrOutput)
}

type ReleasePtrOutput struct {
	*pulumi.OutputState
}

func (ReleasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil))
}

func (o ReleasePtrOutput) ToReleasePtrOutput() ReleasePtrOutput {
	return o
}

func (o ReleasePtrOutput) ToReleasePtrOutputWithContext(ctx context.Context) ReleasePtrOutput {
	return o
}

type ReleaseArrayOutput struct{ *pulumi.OutputState }

func (ReleaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Release)(nil))
}

func (o ReleaseArrayOutput) ToReleaseArrayOutput() ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) ToReleaseArrayOutputWithContext(ctx context.Context) ReleaseArrayOutput {
	return o
}

func (o ReleaseArrayOutput) Index(i pulumi.IntInput) ReleaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Release {
		return vs[0].([]Release)[vs[1].(int)]
	}).(ReleaseOutput)
}

type ReleaseMapOutput struct{ *pulumi.OutputState }

func (ReleaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Release)(nil))
}

func (o ReleaseMapOutput) ToReleaseMapOutput() ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) ToReleaseMapOutputWithContext(ctx context.Context) ReleaseMapOutput {
	return o
}

func (o ReleaseMapOutput) MapIndex(k pulumi.StringInput) ReleaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Release {
		return vs[0].(map[string]Release)[vs[1].(string)]
	}).(ReleaseOutput)
}

func init() {
	pulumi.RegisterOutputType(ReleaseOutput{})
	pulumi.RegisterOutputType(ReleasePtrOutput{})
	pulumi.RegisterOutputType(ReleaseArrayOutput{})
	pulumi.RegisterOutputType(ReleaseMapOutput{})
}
