// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Events.V1Beta1
{
    /// <summary>
    /// Event is a report of an event somewhere in the cluster. It generally denotes some state
    /// change in the system.
    /// </summary>
    public partial class Event : KubernetesResource
    {
        /// <summary>
        /// What action was taken/failed regarding to the regarding object.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Deprecated field assuring backward compatibility with core.v1 Event type
        /// </summary>
        [Output("deprecatedCount")]
        public Output<int> DeprecatedCount { get; private set; } = null!;

        /// <summary>
        /// Deprecated field assuring backward compatibility with core.v1 Event type
        /// </summary>
        [Output("deprecatedFirstTimestamp")]
        public Output<string> DeprecatedFirstTimestamp { get; private set; } = null!;

        /// <summary>
        /// Deprecated field assuring backward compatibility with core.v1 Event type
        /// </summary>
        [Output("deprecatedLastTimestamp")]
        public Output<string> DeprecatedLastTimestamp { get; private set; } = null!;

        /// <summary>
        /// Deprecated field assuring backward compatibility with core.v1 Event type
        /// </summary>
        [Output("deprecatedSource")]
        public Output<Types.Outputs.Core.V1.EventSource> DeprecatedSource { get; private set; } = null!;

        /// <summary>
        /// Required. Time when this Event was first observed.
        /// </summary>
        [Output("eventTime")]
        public Output<string> EventTime { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Optional. A human-readable description of the status of this operation. Maximal length
        /// of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
        /// </summary>
        [Output("note")]
        public Output<string> Note { get; private set; } = null!;

        /// <summary>
        /// Why the action was taken.
        /// </summary>
        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        /// <summary>
        /// The object this Event is about. In most cases it's an Object reporting controller
        /// implements. E.g. ReplicaSetController implements ReplicaSets and this event is emitted
        /// because it acts on some changes in a ReplicaSet object.
        /// </summary>
        [Output("regarding")]
        public Output<Types.Outputs.Core.V1.ObjectReference> Regarding { get; private set; } = null!;

        /// <summary>
        /// Optional secondary object for more complex actions. E.g. when regarding object triggers
        /// a creation or deletion of related object.
        /// </summary>
        [Output("related")]
        public Output<Types.Outputs.Core.V1.ObjectReference> Related { get; private set; } = null!;

        /// <summary>
        /// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
        /// </summary>
        [Output("reportingController")]
        public Output<string> ReportingController { get; private set; } = null!;

        /// <summary>
        /// ID of the controller instance, e.g. `kubelet-xyzf`.
        /// </summary>
        [Output("reportingInstance")]
        public Output<string> ReportingInstance { get; private set; } = null!;

        /// <summary>
        /// Data about the Event series this event represents or nil if it's a singleton Event.
        /// </summary>
        [Output("series")]
        public Output<Types.Outputs.Events.V1Beta1.EventSeries> Series { get; private set; } = null!;

        /// <summary>
        /// Type of this event (Normal, Warning), new types could be added in the future.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Event resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Event(string name, Types.Inputs.Events.V1Beta1.EventArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:events.k8s.io/v1beta1:Event", name, SetAPIKindAndVersion(args), options)
        {
        }

        internal Event(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:events.k8s.io/v1beta1:Event", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Events.V1Beta1.EventArgs? args)
        {
            args ??= new Types.Inputs.Events.V1Beta1.EventArgs();
            args.ApiVersion = "events.k8s.io/v1beta1";
            args.Kind = "Event";
            return args;
        }

        /// <summary>
        /// Get an existing Event resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Event Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Event(name, default(Types.Inputs.Events.V1Beta1.EventArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }

    }
}
