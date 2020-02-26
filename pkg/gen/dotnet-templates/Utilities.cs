// Copyright 2016-2020, Pulumi Corporation
// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;

namespace Pulumi.Kubernetes
{
    internal static class Utilities
    {
        public static string Version { get; }

        static Utilities()
        {
            var assembly = typeof(Utilities).GetTypeInfo().Assembly;
            using var stream = assembly.GetManifestResourceStream("Pulumi.Kubernetes.version.txt");
            if (stream == null)
                throw new Exception("Manifest file 'version.txt' not found");
            using var reader = new StreamReader(stream);
            Version = reader.ReadToEnd().Trim();
        }

        public static InvokeOptions WithVersion(this InvokeOptions? options)
        {
            if (options?.Version != null)
            {
                return options;
            }
            return new InvokeOptions
            {
                Parent = options?.Parent,
                Provider = options?.Provider,
                Version = Version,
            };
        }

        public static string ExecuteCommand(string command, string[] flags)
        {
            using var process = new Process
            {
                StartInfo =
                {
                    FileName = command,
                    Arguments = EscapeArguments(flags),
                    RedirectStandardOutput = true,
                    RedirectStandardError = true
                }
            };
            process.Start();
            string output = process.StandardOutput.ReadToEnd();
            process.WaitForExit();
            if (process.ExitCode > 0)
            {
                string error = process.StandardError.ReadToEnd();
                throw new Exception(error);
            }
            return output;
        }

        /// <summary>
        /// Convert an argument array to an argument string for using with Process.StartInfo.Arguments.
        /// </summary>
        private static string EscapeArguments(params string[] args)
            => string.Join(" ", args.Select(EscapeArguments));

        /// <summary>
        /// Convert an argument array to an argument string for using with Process.StartInfo.Arguments.
        /// </summary>
        private static string EscapeArguments(string argument)
        {
            var escapedArgument = new StringBuilder();
            var backslashCount = 0;
            var needsQuotes = false;

            foreach (var character in argument)
            {
                switch (character)
                {
                    case '\\':
                        // Backslashes are simply passed through, except when they need
                        // to be escaped when followed by a \", e.g. the argument string
                        // \", which would be encoded to \\\"
                        backslashCount++;
                        escapedArgument.Append('\\');
                        break;

                    case '\"':
                        // Escape any preceding backslashes
                        escapedArgument.Append(new string('\\', backslashCount));

                        // Append an escaped double quote.
                        escapedArgument.Append("\\\"");

                        // Reset the backslash counter.
                        backslashCount = 0;
                        break;

                    case ' ':
                    case '\t':
                        // White spaces are escaped by surrounding the entire string with
                        // double quotes, which should be done at the end to prevent
                        // multiple wrappings.
                        needsQuotes = true;

                        // Append the whitespace
                        escapedArgument.Append(character);

                        // Reset the backslash counter.
                        backslashCount = 0;
                        break;

                    default:
                        // Reset the backslash counter.
                        backslashCount = 0;

                        // Append the current character
                        escapedArgument.Append(character);
                        break;
                }
            }

            // No need to wrap in quotes
            if (!needsQuotes)
            {
                return escapedArgument.ToString();
            }

            // Prepend the "
            escapedArgument.Insert(0, '"');

            // Escape any preceding backslashes before appending the "
            escapedArgument.Append(new string('\\', backslashCount));

            // Append the final "
            escapedArgument.Append('\"');

            return escapedArgument.ToString();
        }
    }
}
