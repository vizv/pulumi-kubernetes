// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Core.V1
{

    /// <summary>
    /// Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
    /// </summary>
    [OutputType]
    public sealed class FCVolumeSource
    {
        /// <summary>
        /// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
        /// </summary>
        public readonly string FsType;
        /// <summary>
        /// Optional: FC target lun number
        /// </summary>
        public readonly int Lun;
        /// <summary>
        /// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
        /// </summary>
        public readonly bool ReadOnly;
        /// <summary>
        /// Optional: FC target worldwide names (WWNs)
        /// </summary>
        public readonly ImmutableArray<string> TargetWWNs;
        /// <summary>
        /// Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
        /// </summary>
        public readonly ImmutableArray<string> Wwids;

        [OutputConstructor]
        private FCVolumeSource(
            string fsType,

            int lun,

            bool readOnly,

            ImmutableArray<string> targetWWNs,

            ImmutableArray<string> wwids)
        {
            FsType = fsType;
            Lun = lun;
            ReadOnly = readOnly;
            TargetWWNs = targetWWNs;
            Wwids = wwids;
        }
    }
}
