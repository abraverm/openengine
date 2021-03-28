# OpenEngine Examples

The following examples will help you to get started with OpenEngine and learn how it handles different
cases. The examples are practical and can be experimented with [OpenEngine CLI](../cli/oe/README.md):

```bash
oe -n <action> <example combination file>
```

    Note: '-n' option is short for 'noop', no operation, only showing found solutions

Before customizing any example, it is highly recommended reading about [combination file](../cli/oe/README.md#combination-file).

Open Engine is provider/provisioner/system agnostic, `generic` examples will emphasise it with non-existing types of
providers/provisioners/systems/resources types. The generic examples should help you understand how the example
case is being solved by the OpenEngine. Some cases might have multiple examples to show how it could look on different
providers.


## Use Cases
### Getting started

The simplest type of cases are when there is only one to one matches - single solution found for the requested resource
with only one provider and system are suitable and there is a provisioner capable to fulfil the request.

Examples:

| Use Case                                    | Command                                       | Results                                      | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| ------------------------------------------- | --------------------------------------------- | -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| [Generic](getting_started_generic.yaml)     | `oe -n create getting_started_generic.yaml`   | [log](results/getting_started_generic.log)   | The generic examples should help you understand how the example case is being solved by the OpenEngine                                                                                                                                                                                                                                                                                                                                                 |
| [AWS](getting_started_aws.yaml)             | `oe -n create getting_started_aws.yaml`       | [log](results/getting_started_aws.log)       | Notice how Cue syntax plays a role in the matching process: <br> - [bounds](cue_bounds) of AWS api_version for ec2  <br> - [optional fields](cue_optional) in the provider Vs. provisioner required fields  <br> - [default value](cue_default) for number of instances <br><br> Provider definition doesn't have all [RunInstance parameters](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html#API_RunInstances_Examples) |
| [Openstack](getting_started_openstack.yaml) | `oe -n create getting_started_openstack.yaml` | [log](results/getting_started_openstack.log) | [Openstack Server create API](https://docs.openstack.org/api-ref/compute/?expanded=create-server-detail,get-availability-zone-information-detail,list-flavors-with-details-detail,show-flavor-details-detail,list-keypairs-detail,create-flavor-detail,list-flavors-detail,show-keypair-details-detail#create-server)                                                                                                                                  |
| [Beaker](getting_started_beaker.yaml)       | `oe -n create getting_started_beaker.yaml`    | [log](results/getting_started_beaker.log)    |                                                                                                                                                                                                                                                                                                                                                                                                                                                        |

### Implicit and back again

One of the hardest problems in provisioning domain is abstraction of providers and creating a single interface for
resources. The problem has three parts, standardizing the parameters names, values and supporting explicit values.
Standardizing values is the hardest part as it requires logic and sometimes additional interaction with the provider to
find the right value. For example some providers don't allow explicit server memory size and have a concept of "flavors",
predefined templates with fix memory sizes while other providers that allow custom memory sizes, and not to mention
different memory scales.

OpenEngine handles this case with two concepts: implicit parameters and resolution process. The implicit parameters are
additional parameters that a provider would have defined. The implicit parameters are declared locally in each provider,
and there is no "global" definition of their meaning. Provider definition has required parameters, if the user didn't
provide an explicit value, then the resolution process will kick in. **The resolution process might use implicit parameters
to resolve an explicit value**. In some cases it could be simple as generating random string, in other cases it could
have steps that require fetching information from a provider.

| Use case                                                                | Command                                 | Results                                | Notes                                                                                                                                                                                                                                                    |
| ----------------------------------------------------------------------- | --------------------------------------- | -------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Generic](implicit_generic.yaml)                                        | `oe -n create implicit_generic.yaml`    | [log](results/implicit_generic.log)    | Basic implicit parameters and resolution process                                                                                                                                                                                                         |
| [Server provided by AWS, Openstack or Beaker](implicit_aws_os_bkr.yaml) | `oe -n create implicit_aws_os_bkr.yaml` | [log](results/implicit_aws_os_bkr.log) | Basic concepts are demonstrated in a more realistic example requesting a general resource such as server using different providers (AWS, Openstack and Beaker)</br>**There is a bug that causes a memory leak and long execution time of this example.** |

### Pre/Post actions

One common post deployment scenario is doing something with resource(s) that are related or depended somehow on other
resources that cause constrains on what can be done or how. For example, to resize a disk on Openstack, one must first
shutdown the instance that uses it. Such constrain would require one to find the instance that uses the disk and shut it
down (pre action), and turn it back on after the resize (post action). With OpenEngine, providers of a resource might
have one or more constrains, that will be activated under some conditions. Active constrains must be resolved too for
that provider to be a valid option.

| Use case                                           | Command                                  | Results                                 | Notes |
| -------------------------------------------------- | ---------------------------------------- | --------------------------------------- | ----- |
| [Generic](constrains_generic.yaml)                 | `oe -n update constrains_generic.yaml`   | [log](results/constrains_generic.log)   |       |
| [Openstack resize disk](constrains_os_resize.yaml) | `oe -n update constrains_os_resize.yaml` | [log](results/constrains_os_resize.log) |       |

### Dependencies

Previous examples had a dependencies in the solutions OpenEngine found. However, a simpler case is when user defines
a dependency in their requested resources. The dependency could be explicit result of one resource to be used in another
as a parameter, or could implicit without knowing what exact information is needed and how to use it.

For example, let say creating a custom disk that would be attached to a new server - the important part about the disk
customization might be something like its type or size, and the new server depends on the disk to be created first.
Knowing which provider to use or maybe implicit parameters exist than the user could use the id of the created disk as a
value for a specific parameter. Another option is if the provider has interface that support such relationship, that is
the server definition knows how to use dependencies without an explicit parameter or response specification.

| Use case                                                         | Command                                           | Results                                          | Notes                                                                                                            |
| ---------------------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------- |
| [Generic explicit](dependencies_generic_explicit.yaml)           | `oe -n update dependencies_generic_explicit.yaml` | [log](results/dependencies_generic_explicit.log) | when only the dependency structure is important, everything else is explicitly defined                           |
| [Generic explicit](dependencies_generic_implicit.yaml)           | `oe -n update dependencies_generic_implicit.yaml` | [log](results/dependencies_generic_implicit.log) | When dependency provides required information but what and how it be used depends on the provider and dependency |
| [Openstack custom disk](dependencies_os_disk.yaml)               | `oe -n update dependencies_os_disk.yaml`          | [log](results/dependencies_os_disk.log)          |                                                                                                                  |

[cue_bounds]: https://cuelang.org/docs/references/spec/#bounds
[cue_optional]: https://cuelang.org/docs/references/spec/#structs
[cue_default]: https://cuelang.org/docs/references/spec/#default-values