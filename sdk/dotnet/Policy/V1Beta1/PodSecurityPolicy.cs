// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Policy.V1Beta1
{
    /// <summary>
    /// PodSecurityPolicy governs the ability to make requests that affect the Security Context that
    /// will be applied to a pod and container.
    /// </summary>
    public partial class PodSecurityPolicy : KubernetesResource
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
        /// Standard object's metadata. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// spec defines the policy enforced.
        /// </summary>
        [Output("spec")]
        public Output<Types.Outputs.Policy.V1Beta1.PodSecurityPolicySpec> Spec { get; private set; } = null!;


        /// <summary>
        /// Create a PodSecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PodSecurityPolicy(string name, Types.Inputs.Policy.V1Beta1.PodSecurityPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:policy/v1beta1:PodSecurityPolicy", name, SetAPIKindAndVersion(args), options)
        {
        }

        internal PodSecurityPolicy(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:policy/v1beta1:PodSecurityPolicy", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Policy.V1Beta1.PodSecurityPolicyArgs? args)
        {
            args ??= new Types.Inputs.Policy.V1Beta1.PodSecurityPolicyArgs();
            args.ApiVersion = "policy/v1beta1";
            args.Kind = "PodSecurityPolicy";
            return args;
        }

        /// <summary>
        /// Get an existing PodSecurityPolicy resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PodSecurityPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PodSecurityPolicy(name, default(Types.Inputs.Policy.V1Beta1.PodSecurityPolicyArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }

    }
}
