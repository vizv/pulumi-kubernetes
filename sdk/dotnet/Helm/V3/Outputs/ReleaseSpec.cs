// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Helm.V3
{

    /// <summary>
    /// Specification defining the Helm Release to install.
    /// </summary>
    [OutputType]
    public sealed class ReleaseSpec
    {
        /// <summary>
        /// If set, installation process purges chart on fail. The wait flag will be set automatically if atomic is used
        /// </summary>
        public readonly bool Atomic;
        /// <summary>
        /// Chart name to be installed. A path may be used.
        /// </summary>
        public readonly string Chart;
        /// <summary>
        /// Allow deletion of new resources created in this upgrade when upgrade fails
        /// </summary>
        public readonly bool CleanupOnFail;
        /// <summary>
        /// Create the namespace if it does not exist
        /// </summary>
        public readonly bool CreateNamespace;
        /// <summary>
        /// Run helm dependency update before installing the chart
        /// </summary>
        public readonly bool DependencyUpdate;
        /// <summary>
        /// Add a custom description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Use chart development versions, too. Equivalent to version '&gt;0.0.0-0'. If `version` is set, this is ignored
        /// </summary>
        public readonly bool Devel;
        /// <summary>
        /// Prevent CRD hooks from, running, but run other hooks.  See helm install --no-crd-hook
        /// </summary>
        public readonly bool DisableCRDHooks;
        /// <summary>
        /// If set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
        /// </summary>
        public readonly bool DisableOpenapiValidation;
        /// <summary>
        /// Prevent hooks from running.
        /// </summary>
        public readonly bool DisableWebhooks;
        /// <summary>
        /// Force resource update through delete/recreate if needed.
        /// </summary>
        public readonly bool ForceUpdate;
        /// <summary>
        /// Location of public keys used for verification. Used only if `verify` is true
        /// </summary>
        public readonly string Keyring;
        /// <summary>
        /// Run helm lint when planning
        /// </summary>
        public readonly bool Lint;
        /// <summary>
        /// The rendered manifest as JSON.
        /// </summary>
        public readonly string Manifest;
        /// <summary>
        /// Limit the maximum number of revisions saved per release. Use 0 for no limit
        /// </summary>
        public readonly int MaxHistory;
        /// <summary>
        /// Release name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace to install the release into.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// Postrender command to run.
        /// </summary>
        public readonly string Postrender;
        /// <summary>
        /// Perform pods restart during upgrade/rollback
        /// </summary>
        public readonly bool RecreatePods;
        /// <summary>
        /// If set, render subchart notes along with the parent
        /// </summary>
        public readonly bool RenderSubchartNotes;
        /// <summary>
        /// Re-use the given name, even if that name is already used. This is unsafe in production
        /// </summary>
        public readonly bool Replace;
        /// <summary>
        /// Specification defining the Helm chart repository to use.
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Helm.V3.RepositorySpec RepositorySpec;
        /// <summary>
        /// When upgrading, reset the values to the ones built into the chart
        /// </summary>
        public readonly bool ResetValues;
        /// <summary>
        /// When upgrading, reuse the last release's values and merge in any overrides. If 'resetValues' is specified, this is ignored
        /// </summary>
        public readonly bool ReuseValues;
        /// <summary>
        /// Custom values to be merged with items loaded from values.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Set;
        /// <summary>
        /// By default, the provider waits until all resources are in a ready state before marking the release as successful. Setting this to true will skip such await logic.
        /// </summary>
        public readonly bool SkipAwait;
        /// <summary>
        /// If set, no CRDs will be installed. By default, CRDs are installed if not already present
        /// </summary>
        public readonly bool SkipCrds;
        /// <summary>
        /// Time in seconds to wait for any individual kubernetes operation.
        /// </summary>
        public readonly int Timeout;
        /// <summary>
        /// List of assets (raw yaml files) to pass to helm.
        /// </summary>
        public readonly ImmutableArray<AssetOrArchive> Values;
        /// <summary>
        /// Verify the package before installing it.
        /// </summary>
        public readonly bool Verify;
        /// <summary>
        /// Specify the exact chart version to install. If this is not specified, the latest version is installed.
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// Will wait until all Jobs have been completed before marking the release as successful. This is ignored if `skipWait` is enabled.
        /// </summary>
        public readonly bool WaitForJobs;

        [OutputConstructor]
        private ReleaseSpec(
            bool atomic,

            string chart,

            bool cleanupOnFail,

            bool createNamespace,

            bool dependencyUpdate,

            string description,

            bool devel,

            bool disableCRDHooks,

            bool disableOpenapiValidation,

            bool disableWebhooks,

            bool forceUpdate,

            string keyring,

            bool lint,

            string manifest,

            int maxHistory,

            string name,

            string @namespace,

            string postrender,

            bool recreatePods,

            bool renderSubchartNotes,

            bool replace,

            Pulumi.Kubernetes.Types.Outputs.Helm.V3.RepositorySpec repositorySpec,

            bool resetValues,

            bool reuseValues,

            ImmutableDictionary<string, object> set,

            bool skipAwait,

            bool skipCrds,

            int timeout,

            ImmutableArray<AssetOrArchive> values,

            bool verify,

            string version,

            bool waitForJobs)
        {
            Atomic = atomic;
            Chart = chart;
            CleanupOnFail = cleanupOnFail;
            CreateNamespace = createNamespace;
            DependencyUpdate = dependencyUpdate;
            Description = description;
            Devel = devel;
            DisableCRDHooks = disableCRDHooks;
            DisableOpenapiValidation = disableOpenapiValidation;
            DisableWebhooks = disableWebhooks;
            ForceUpdate = forceUpdate;
            Keyring = keyring;
            Lint = lint;
            Manifest = manifest;
            MaxHistory = maxHistory;
            Name = name;
            Namespace = @namespace;
            Postrender = postrender;
            RecreatePods = recreatePods;
            RenderSubchartNotes = renderSubchartNotes;
            Replace = replace;
            RepositorySpec = repositorySpec;
            ResetValues = resetValues;
            ReuseValues = reuseValues;
            Set = set;
            SkipAwait = skipAwait;
            SkipCrds = skipCrds;
            Timeout = timeout;
            Values = values;
            Verify = verify;
            Version = version;
            WaitForJobs = waitForJobs;
        }
    }
}
