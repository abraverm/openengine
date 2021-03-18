import "list"
import "math"
import "crypto/sha256"
import "encoding/json"
import "strings"

// ============================================================================
// OpenEngine Specifications - version 0.1
// ============================================================================
// The following are OpenEngine specifications written in CUE language [1].
// The specifications purpose:
// - API for users to develop their definitions -the input for the OpenEngine
// - Testing and validation of user definitions
// - Core components definition
// The specification is a core component of OpenEngine execution, thus its
// a single point of truth.
//
// [1] https://cuelang.org/

// ----------------------------------------------------------------------------
// Specifications
// ----------------------------------------------------------------------------
// The following are specifications for users to follow when writing their
// difinitions of the different components. Some of the specifications
// are common types or values.

// Type is the initial matching property for finding a solution. Type of resource
// narrows the scope of the matching problem segnificanetly and works well for
// implicit and explicit user requests.
// The current implementation is a simple string.
#Type: string

// Property of resources\system is used define what it is or should be, it acts as user
// interface to define implicit or explicit properties and interface for solutions.
#Property: string | bool | number | [...#Property] | {...} | null
#Properties:[string]: #Property | *null

// Resource are the objects that providers manage. It defines what is
// (how its defined in the provider), or what it should be (user request).
// The resource could be explicit and unique, in its type and\or properties, or
// implicit and could be realized in number of ways. The definition of resources
// is for identifying its type to match with other components, and act as template
// for storing input and output of solutions.
// Example:
// myResource: {
//   type: "Server"
//   properties: {
//     memory: "4g"
//   }
// }
#Resource: {
  type: #Type
  name?: string
  system?: #Properties
  properties:[string]: #Property
  response?: #Properties
  solutions: [...]
  dependencies: [for p, d in dependedProperties { d.resourceName }, for i in interfacesDependencies { i } , ...string] // Explicit dependency and implicit dependency using interfaces
  interfacesDependencies: [...string]
  enabledInterfaces: [...string]
  disabledInterfaces: [...string]
  dependedProperties:[string]: { // Implicit dependency with direcrt reference
    resourceName: string
    response: {...#Property}
    resolved: len(solutions & [...{ resource:name: resourceName, response: response}]) > 0
  }
}



// System defines the properties of an actual system that the user has access to.
// Type of a system is relevant for cases where there might be number of instances
// of similar provider, i.e Openstack, and it will help to narrow the matching process,
// just like with resources.
// mySystem: {
//   type: "Openstack"
//   properties: {
//     nova:version: 1.4
//   }
// }
#System: {
  type: #Type
  name?: string
  properties: #Properties
}

// Only actions available via REST API and by all providers, a.k.a CRUD
// https://en.wikipedia.org/wiki/Create,_read,_update_and_delete
#Action: "create" | "delete" | "update" | "read"

// Implicit to explicit process described in a series of tasks. Each task is
// a wrapper to the actual script that will be executed and being active interface
// to OpenEngine.
#ImplicitTask: #ScriptTask | #ResourceTask

#ResourceTask: {
  action: #Action
  resource: #Resource
}

#ScriptTask: {
  resolved: len([ for a in args if a == null {null}]) == 0
  script: string | *null   // The script to call
  args: #Properties // Arguments for the script, they could be explicit or linked to implicit properties
}

// Provider defines interface and implicit process for resource action.
// Null has a special meaning of unmatched property and the engine drops
// such unresolved solutions.
#Provider: {
  type: #Type
  name?: string
  action: #Action
  system: #System
  properties: #Properties
  implicit:[string]: [#ImplicitTask, ...#ImplicitTask] | null
  interfaces:[string]: [...#Interface]
  response: #Properties
  constrains: [...#Constrain]
}

// Interface is similar to depended property, it provides explicit value from
// other solution response. However, interface is defined by the provider
// and matches *any* dependency. Interface solves the problem of having
// two resources depnded on each other but doesn't explain how exactly it help.
// The ambiguity resolved by the provider expert. Each property might have more
// than one interface. The provider expert will have to decide which resources
// to bridge with interfaces and the interface order. Interfaces by default
// replace implicit workflows (unless disabled), even if no actual solution was
// matched - for example A depends on B and there is an interface that matches
// B type but not the response of the solution, that would mean that property x
// in A that has that interface would resolve to Null and not to implicit
// workflow.
//
// Example 1: two resources - a Server that depends on a Network. Two possible
// systems to deploy on: AWS and Openstack. The systems have different providers
// and provisioners, but in both cases, the Network will be created first and using
// the interface it will pick the the field from the provisioner response and
// use it in the right property when the Server is created.
//
// Interfaces allows sub-domains in provisionining - for example interfaces
// (and their order) that are more suitable for deployment of a complex system
// such as Openshift - in such cases the dependencies between resources would
// make more sense to the user than a different set of interfaces.
//
// User has some control over interfaces by defining which dependencies can
// be matched, and picking interfaces (disable\enable). That said, interfaces
// could cause ambiguity, that should be solved with a different tool
// such as a graphic interface to plan effective dependencies.
#Interface: {
  name: string
  type: #Type
  action: #Action
  response: #Properties
  field: string
}

// Constrain is a set of additional pre and post steps that are required to
// execute an action upon a resource. For example to add a disk to existing
// Server (Update action), it might require to shutdown it first (in some
// systems). Constrain must match specific conditions to activate. There
// could be multiple active constrains, this also might mean conflicts (out of
// scope for the OpenEngine). Steps could be action on the resource or other
// resources and\or script execution - just like implicit workflow.
#Constrain: {
  name: string
  pre: [...#ImplicitTask]
  post: [...#ImplicitTask]
  properties: #Properties
  response: #Properties
}

// Provisioner
#Provisioner: {
  type: #Type
  name?: string
  action: #Action
  system: #System
  properties: #Properties
  provisioner: string
  response: #Properties
}

Resources: [...#Resource]

// ----------------------------------------------------------------------------
// OpenEngine mechanics
// ----------------------------------------------------------------------------
#structToHash: {
  #in: {...}
  out: strings.Replace(json.Marshal(sha256.Sum256(json.Marshal(#in))[0:6]), "\"", "", -1)
}

#resourceToHash: {
  #resource: #Resource
  new: {
    type: *#resource.type | ""
    name: *#resource.name | ""
    system: *#resource.system | {}
    properties: *#resource.properties | {}
    dependencies: *#resource.dependencies | []
    dependedProperties: *#resource.dependedProperties | []
    response: *#resource.response | {}
    interfacesDependencies: *#resource.interfacesDependencies | []
    enabledInterfaces: *#resource.enabledInterfaces | []
    disabledInterfaces: *#resource.disabledInterfaces | []
  }
  out: (#structToHash & { #in: new }).out
}

#solutionId: {
    #provider: {...}
    #provisioner: {...}
    #system: {...}
    #resource: {...}
    provider_id: *(#structToHash & {#in: #provider}).out | ""
    provisioner_id: *(#structToHash & {#in: #provisioner}).out | ""
    system_id: *(#structToHash & {#in: #system}).out | ""
    resource_id: *(#resourceToHash & { #resource: #resource }).out | ""
    out: provider_id + provisioner_id + system_id + resource_id
}
#Solution: {
  name: (*("R("+resource.name + ")")| "" ) + (*("S("+System.name+")")| "" ) + (*("PD("+#provider.name+")")| "" ) + (*("PR("+#provisioner.name+")")| "" )
  _id: (#solutionId & {#provider: #provider, #provisioner: #provisioner, #system: System, #resource: resource}).out
  #provider: #Provider
  #provisioner: #Provisioner
  resource: #Resource
  System: #System
  #xids: [...]
  #xsolutions: [...]
  #provisioner: properties: { ... }
  resource: properties: { ... }
  match: *(#provider & {
    system: #provisioner.system & System & (resource.system | { ... })
    type: (resource.type & #provisioner.type) | null
    properties: resource.properties & #provisioner.properties
    response: #provisioner.response & (*resource.response | { ... })
  }) | { action: string | *"" }
  interfaces: {
    for param, options in (*match.interfaces | [] ){
      "\(param)": [ for i in [
        for o in options
        for s in resource.solutions
        {
          *{
            response: { for p, v in o.response { "\(p)": *(v & s.match.response[p]) | null } }
            action: s.match.action & o.action
			name: o.name
			type: s.match.type & o.type
            field: o.field
            solution: s
          } | null
        }
      ] if i != null && len(*[ for r in i.response if r == null { r } ] | []) == 0  { i } ]
    }
  }
  constrains: [ for x in [
    for c in (*match.constrains | [] ) if ( *{
       conditions: close(match.properties) & close({ for p, v in c.properties  {"\(p)": v & match.properties[p]}})
       response: close(match.response) & close({ for p, v in c.response  {"\(p)": v & match.response[p]}})
     } | null) != null {
        name: c.name
        pre: ([
          for task in c.pre {
             if (*task.action | "") != "" {
                resolved: true
                action: task.action
                resource: task.resource
				solutions: (#Solutions & {
					action: action
					xresources: [resource]
					systems: [System]
					#providers: Providers
					#provisioners: Provisioners
					xids: #xids + [_id]
					xsolutions: #xsolutions
					...
					}).results & [...{resolved: true}]
             }
            if (*task.script | "") != "" { task }
          }
        ] & [...{resolved: true}])
        post:  ([
          for task in c.post {
             if (*task.action | "") != "" {
                resolved: true
                action: task.action
                resource: task.resource
				solutions: (#Solutions & {
					action: action
					xresources: [resource]
					systems: [System]
					#providers: Providers
					#provisioners: Provisioners
					xids: #xids + [_id]
					xsolutions: #xsolutions
					...
					}).results & [...{resolved: true}]
             }
            if (*task.script | "") != "" { task }
          }
        ] & [...{resolved: true}])
       }
    ] if (x.pre != null && x.post != null) || true { x }
  ]
  implicit: {
     for param, workflow in (*match.implicit| []) {
		"\(param)": *([
          for task in workflow {
             if (*task.action | "") != "" {
                resolved: true
                action: task.action
                resource: task.resource
				solutions: (#Solutions & {
					action: action
					xresources: [resource]
					systems: [System]
					#providers: Providers
					#provisioners: Provisioners
					xids: #xids + [_id]
					xsolutions: #xsolutions
					...
					}).out & [...{resolved: true}]
             }
            if (*task.script | "") != "" { task }
          }
        ] & [...{resolved: true}] ) | null
     }
  }
  joined: {
    for x, y in (*match.properties | [] ) { "\(x)":explicit: y }
    for x, y in implicit if (*resource.dependedProperties[x]| null) == null && len(interfaces[x] | []) == 0 {"\(x)":implicit: y }
    for x, y in interfaces if len(y) > 0 {"\(x)":interface: *y[0] | null }
  }
  provisioner: #provisioner.provisioner
  properties: {
    for x, y in joined {
      if y.explicit != null { "\(x)": y.explicit }
      if y.explicit == null {
        if (*y.interface | null) != null { "\(x)": y.interface }
        if (*y.interface | null) == null {
          if (*y.implicit | null) == null { "\(x)": null }
          if (*y.implicit | null) != null { "\(x)": y.implicit }
        }
      }
    }
  }
  resolved: len([for p in properties if p == null {null}]) == 0 && len(match) > 1
}

#dependlessResources: {
  #resources: [...#Resource]
  out: [ for r in #resources if len(r.dependencies) == 0 { r } ]
}

#dependedResources: {
  #resources: [...#Resource]
  out: [ for r in #resources if len(r.dependencies) > 0 { r } ]
}

#filterSolutionsByDependencies: {
  s: [...]
  d: [...]
  tmp: { for x in d { "\(x)": s & [...{resource:name: x}] } }
  out: [ for x, y in tmp {y} ]
}

#allDependencies: {
  solutions: [...]
  dependencies: [...]
  tmp: { for d in dependencies { "\(d)": [for s in solutions if s.resource.name == d { s} ] } }
  out: len(dependencies) == len([ for d, s in tmp if len(s) > 0 {d} ])
}

#dependedUnresolved: {
  #resources: [...#Resource]
  #solutions: [...]
  out: [ for r in #resources if ! (#allDependencies & {solutions: #solutions, dependencies: r.dependencies}).out { r } ]
}

#dependedResolved: {
  #resources: [...#Resource]
  #solutions: [...]
  out: [ for r in #resources if (#allDependencies & {solutions: #solutions, dependencies: r.dependencies}).out { r } ]
}
// recursion...
// need to group solutions by system
// solutions with dependencies must be in the same system
// group of solutions must be complete - all resources resolved
#Solutions: {
  action: #Action
  xresources: [...#Resource]
  systems: [...#System]
  #providers: [...#Provider]
  #provisioners: [...#Provisioner]
  xsolutions: [...]
  xids: [...]
  XIDS: [ for x in xsolutions { x._id }, for x in xids {x}]
  dependless_solutions:[ for s in [
    for xresource in (#dependlessResources & { #resources: xresources }).out
    for system in systems
    for provider in #providers
    for provisioner in #provisioners
    if ! list.Contains( XIDS , (#solutionId & {#provider: provider, #provisioner: provisioner, #system: system, #resource: xresource}).out )
        {*(#Solution & { #xsolutions: xsolutions, #xids: XIDS + [(#solutionId & {#provider: provider, #provisioner: provisioner, #system: system, #resource: xresource}).out], match:action: action, resource: close(xresource), System: close(system), #provider: close(provider), #provisioner: close(provisioner) }) | {resolved: false}}
  ] if s.resolved { s } ]

  // Resources that can't be satisfied with current known solutions
  depended_unresolved: (#dependedUnresolved & {
    #resources: (#dependedResources & { #resources: xresources }).out
    #solutions: xsolutions + dependless_solutions
  }).out
  // Resources that can be satisfied with current known solutions
  depended_resolved: (#dependedResolved & {
    #resources: (#dependedResources & { #resources: xresources }).out
    #solutions: xsolutions + dependless_solutions
  }).out
  // resource dependencies might have more than one solution
  // Explode will create combinations of dependencies and different solutions
  // For each resource will be a set of solutions to use
  #explode: {
    #solutions: [...]
    #dependencies: [...string]
    input: { for d in #dependencies { "\(d)": [ for s in #solutions if s.resource.name == d {
      name: s.name
      _id: s._id
      resource: s.resource
      System: s.System
      match: s.match
      implicit: s.implicit
      joined: s.joined
      provisioner: s.provisioner
      properties: s.properties
      resolved: s.resolved
    }]} }
    combos: (#combo & {#input: input}).out
    out: [ for c in combos {[for d, s in c { s }]}]
  }
  depended_resolved_decoupled: [
    for resource in depended_resolved
    for set in (#explode & {#solutions: dependless_solutions + xsolutions, #dependencies: resource.dependencies}).out
    { resource & { solutions: set } }
  ]
  // Filter the solutions that don't satisfy dependency properties response requirement
  depended_resolved_decoupled_properties: [
    for resource in depended_resolved_decoupled
    if len([for d in resource.dependedProperties if ! d.resolved { false }]) == 0
    { resource }
  ]

  // Find solutions for resources who have resolved dependencies
  depended_resolved_solutions: [ for s in [
    for xresource in depended_resolved_decoupled_properties
    for system in systems
    for provider in #providers
    for provisioner in #provisioners
    if ! list.Contains( XIDS , (#solutionId & {#provider: provider, #provisioner: provisioner, #system: system, #resource: xresource}).out )
      {*(#Solution & { #xsolutions: xsolutions, #xids: XIDS, match:action: action, resource: xresource, System: system, #provider: provider, #provisioner: provisioner }) | {resolved: false}}
  ] if s.resolved { s }]

  // Recursion:
  Saction: action
  Ssystems: systems
  Sproviders: #providers
  Sprovisioners: #provisioners
  unresolved_solutions: list.FlattenN([ if len(depended_unresolved) > 0 {[ for s in (*(#Solutions & {
    action: Saction
    xresources: depended_unresolved
    systems: Ssystems
    #providers: Sproviders
    #provisioners: Sprovisioners
    xsolutions: xsolutions + dependless_solutions + depended_resolved_solutions
    ...
  }).out | []) if s.resolved { s } ]} ], 2 )
  _ids: [for s in (xsolutions + dependless_solutions + depended_resolved_solutions) { s._id } ]
  unresolved_solutions_filtered: [ for s in unresolved_solutions if ! list.Contains(_ids, s._id) { s } ]
  out: xsolutions + dependless_solutions + depended_resolved_solutions + unresolved_solutions_filtered
  implicit_workaround: [ for i in out {*( {
    name: i.name
    _id: i._id
    properties: i.properties
	provisioner: i.provisioner
    resolved: i.resolved
    constrains: i.constrains
	resource: i.resource
	system: i.System
    joined: i.joined
  }& { joined:[string]:implicit: [...{ solutions?: [{...}, ...], ...}], ...}) | null}]
  results: [ for o in implicit_workaround if o != null {
    name: o.name
    _id: o._id
    properties: o.properties
	provisioner: o.provisioner
    constrains: o.constrains
    resolved: o.resolved
	resource: o.resource
	system: o.system
  }]
}


#combo: { // generates combinations of all possible solutions
  #input:[string]: [...]
  _size: math.Round(list.Product([ for key, set in #input { len(set) } ]))
  _c: {
    for key, set in #input { "\(key)": *(set * (_size div len(set))) | 0 }
  }
  out: [
    for index in list.Range(0, _size, 1) {
      for key, set in #input { "\(key)": _c["\(key)"][math.Round(index)] }
    }
  ]
}

resources:[string]:#Resource
Resources: [ for _, r in resources { r } ]
systems:[string]:#System
Systems: [ for _, s in systems { s } ]
providers:[string]:#Provider
Providers: [ for _, p in providers { p } ]
provisioners:[string]:#Provisioner
Provisioners: [ for _, p in provisioners { p } ]

#FlattenUniqueGroup: {
  input: [...[...string]]
  flatten: list.FlattenN(input, 2)
  unique: { ... }
  { for i in flatten { unique:"\(i)":1 } }
  out: list.SortStrings([ for v, _ in unique { v } ])
}

#CreateGroups: {
  input: [...[...string]]
  start: list.SortStrings(*input[0] | [])
  others: *input[1:] | []
  start_matching: [ for other in others if ! list.UniqueItems(list.SortStrings(other + start)) { other } ] + [start]
  matched: *(#FlattenUniqueGroup & { input: [for s in start_matching { [for x in s  {x} ] } ]}).out | []
  not_matched: [ for other in others if list.UniqueItems(list.SortStrings(other + start)) { other } ]
  next: [matched] + not_matched
  complete_start: list.FlattenN([ for c in [
    if len(next) == len(input) || len(next) == 1 { matched },
    if len(next) != len(input) && len(next) > 1 {(*(#CreateGroups & { input: next}).matched | []) },
  ] if len(c) > 0 { c } ], 2)
  ungrouped: [ for other in others if list.UniqueItems(list.SortStrings(other + (*complete_start | []))) { other }]
  other_groups: *(#CreateGroups & { input: ungrouped }).out | []
  out: [for o in other_groups if len(o) > 0 {o} ] + [complete_start]
}


#globalNames: list.SortStrings([ for name, resource in resources {  strings.TrimSpace(*resource.name | name) } ])
globalUniqueNames: list.UniqueItems(#globalNames)
#GroupsByNames: (#CreateGroups & {
    input: [ for name, resource in resources {
         [strings.TrimSpace(*resource.name | name)] + resource.dependencies
      }
    ]
  }).out


#DependencyGroups: [
  for group in #GroupsByNames {[
    for gn in group {
       for name, r in resources if list.Contains([r.name, name], gn) { r }
    }
  ]}
]

ACTION: #Action

// Each group will have a set of solutions per system, each set is complete:
// User dependencies, constrains and implicit tasks are all resolved
// Note: set might contain multiple solutions for one resource
#DependencyGroupsSolutions: [
	for group in #DependencyGroups {[
		for System in Systems {
           (#Solutions & {
				action: ACTION
				xresources: group
				systems: [System]
				#providers: Providers
				#provisioners: Provisioners
				xsolutions: []
				...
			}).results }
	]}
]

// The set of solutions (for group and system) is decoupled to sub-sets that is complete but each resource has only
// one solution. Dependency solution tree live within the resources, so the sub-set must also match the dependency solution
// tree within the resource. This is the final result of OpenEngine work.

#getDependencySolutionsIds: {
  #solution: {...}
  ids: [
    [ #solution._id ],
    [for s in #solution.resource.solutions { s._id }],
    for s in #solution.resource.solutions { *(#getDependencySolutionsIds & { #solution: s }).out | []}
  ]
  out: (#FlattenUniqueGroup & { input: ids }).out
}

#minimalSet: {
  #set: [...]
  base: [ for s in #set { s._id } ]
  full: (#FlattenUniqueGroup & { input: [ for s in #set { (#getDependencySolutionsIds & { #solution: s } ).out} ]}).out
  out: len(base) == len(full)
}

#DecoupleGroup: {
  #set: [...]
  #resources: [...#Resource]
  rids: [ for r in #resources { (#resourceToHash & { #resource: r } ).out } ]
  solutions: { for r in rids { "\(r)": [ for s in #set if (#resourceToHash & { #resource: s.resource } ).out == r { s } ] } }
  combos: [ for combo in (#combo & { #input: solutions } ).out {[for n, s in combo { s } ]} ]
  out: [ for combo in combos if (#minimalSet & {#set: combo }).out {combo}]
}


DependencyGroupsSolutionsDecoupled: list.FlattenN([
	for gid, group in #DependencyGroupsSolutions {[
		for set in group if len(set) > 0 {[
 		  for c in (#DecoupleGroup & { #set: set, #resources: #DependencyGroups[gid], ... }).out if len(c) > 0 {[
			for s in c{
              name: s.name
              properties: s.properties & {...}
              provisioner: s.provisioner
              resource: s.resource
              constrains: s.constrains
              system: s.system
            }
          ]}
		]}
	]}
], 2)
