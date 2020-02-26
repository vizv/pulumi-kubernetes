// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.FlowControl.V1Alpha1
{
    /// <summary>
    /// FlowSchema defines the schema of a group of flows. Note that a flow is made up of a set of
    /// inbound API requests with similar attributes and is identified by a pair of strings: the
    /// name of the FlowSchema and a "flow distinguisher".
    /// </summary>
    public partial class FlowSchema : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// `metadata` is the standard object's metadata. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// `spec` is the specification of the desired behavior of a FlowSchema. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("spec")]
        public Output<Types.Outputs.FlowControl.V1Alpha1.FlowSchemaSpec> Spec { get; private set; } = null!;

        /// <summary>
        /// `status` is the current status of a FlowSchema. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
        /// </summary>
        [Output("status")]
        public Output<Types.Outputs.FlowControl.V1Alpha1.FlowSchemaStatus> Status { get; private set; } = null!;


        /// <summary>
        /// Create a FlowSchema resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlowSchema(string name, Types.Inputs.FlowControl.V1Alpha1.FlowSchemaArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema", name, SetAPIKindAndVersion(args), options)
        {
        }

        internal FlowSchema(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:flowcontrol.apiserver.k8s.io/v1alpha1:FlowSchema", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.FlowControl.V1Alpha1.FlowSchemaArgs? args)
        {
            args ??= new Types.Inputs.FlowControl.V1Alpha1.FlowSchemaArgs();
            args.ApiVersion = "flowcontrol.apiserver.k8s.io/v1alpha1";
            args.Kind = "FlowSchema";
            return args;
        }

        /// <summary>
        /// Get an existing FlowSchema resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlowSchema Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FlowSchema(name, default(Types.Inputs.FlowControl.V1Alpha1.FlowSchemaArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }

    }
}
