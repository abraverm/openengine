import "list"
import "math"

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
#Property: string | bool | number | [...#Property] | {...}
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
  system?: #Properties
  properties:[string]: #Property
  response?: #Properties
  dependencies?:[string]: {
    solutions?: [...]
    resource: #Resource
    response: string
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
  properties: #Properties
}

// Only actions available via REST API and by all providers, a.k.a CRUD
// https://en.wikipedia.org/wiki/Create,_read,_update_and_delete
#Action: "create" | "delete" | "update" | "read"

// Implict to explicit process described in a series of tasks. Each task is
// a wrapper to the actual script that will be executed and being active interface
// to OpenEngine.
#ImplicitTask: {
  resolved: bool
  script?: string | *null   // The script to call
  args?: #Properties // Arguments for the script, they could be explicit or linked to implicit properties
  solutions?: [...#Solution]
}


// Provider defines interface and implicit process for resource action.
// Null has a special meaning of unmatched property and the engine drops
// such unresolved solutions.
#Provider: {
  type: #Type
  action: #Action
  system: #System
  properties: #Properties
  implicit: [string]: [#ImplicitTask, ...#ImplicitTask]
  response: #Properties
}

// Provisioner
#Provisioner: {
  type: #Type
  action: #Action
  system: #System
  properties: #Properties
  provisioner: string
  response: #Properties
}
// ----------------------------------------------------------------------------
// OpenEngine mechanics
// ----------------------------------------------------------------------------
// TBD - is this part required for testinng the user definitions?

#Solution: {
  #provider: #Provider
  #provisioner: #Provisioner
  resource: #Resource
  System: #System
  #provisioner: properties: { ... }
  resource: properties: { ... }
  match: #provider & {
    system: #provisioner.system & System & (resource.system | { ... })
    type: (resource.type & #provisioner.type) | null
    properties: resource.properties & #provisioner.properties
    response: #provisioner.response & (*resource.response | { ... })
  }
  _joined: {
    for x, y in match.properties { "\(x)":explicit: y }
    for x, y in match.implicit if y != null { "\(x)":implicit: *(y & [...{resolved: true, ...}]) | null }
  }
  properties: {
    for x, y in _joined {
      if (*y.implicit | null) == null { "\(x)": y.explicit }
      if y.explicit == null { "\(x)": y.implicit | null }
      if (*y.implicit | null) == null && y.explicit == null { "\(x)":  null }
      if (*y.implicit | null) != null && y.explicit != null { "\(x)":  y.explicit }
    }
  }
  resolved: len([for p in properties if p == null {null}]) == 0 && len(match) > 0
}

// recursion...
// need to group solutions by system
// solutions with dependencies must be in the same system
// group of solutions must be complete - all resources resolved
#Solutions: {
  action: #Action
  xresources: [...#Resource]
  systems: [...#System]
  providers: [...#Provider]
  provisioners: [...#Provisioner]
  solutions?: [...]
  _dependless_resources: [ for r in xresources if len(r.dependencies) == 0 { r } ]
  _depended_resources: [ for r in xresources if len(r.dependencies) > 0 { r } ]
  dependless_tmp: [
    for xresource in _dependless_resources
    for system in systems
    for provider in providers
    for provisioner in provisioners
    {*(#Solution & { match:action: action, resource: xresource, System: system, #provider: provider, #provisioner: provisioner }) | {resolved: false}}
  ]
  dependless_solutions: [for s in dependless_tmp if s.resolved == true { s }]
  _depended_resolved: [
    for r in _depended_resources
    for d, v in r.dependencies
    { *(r & { dependencies: { "\(d)":solutions: [for x in [for s in dependless_solutions + solutions { *(s & { resource: v.resource }) | {} } ] if len(x) > 0 {x}] } }) | {type: ""}}
  ]
  _resolved_decouple: [ for r in _depended_resolved if r.type != "" { #Explode & { resource: r } }]
  _resolved_decouple_tmp: [
    for xresource in _resolved_decouple
    for system in systems
    for provider in providers
    for provisioner in provisioners
    {*(#Solution & { match:action: action, resource: xresource, System: system, #provider: provider, #provisioner: provisioner }) | {resolved: false}}
  ]
  resolved_decoupled_solutions: [for s in _resolved_decouple_tmp if s.resolved == true { s }]
  _unresolved: [ for r in _depended_resolved if r.type == "" { r } ]
  unresolved_tmp: (#Solutions & {action: action, xresources: _unresolved, systems: systems, provisioners: provisioners, solutions: solutions + dependless_solutions + resolved_decoupled_solutions}).out | [] // empty list required to avoid structual cycle error
  unresolved_solutions: [for s in unresolved_tmp if s.resolved == true { s }]
  out: solutions + dependless_solution + unresolved_solutions + resolved_decoupled_solutions 
}

#combo: { // generates combinations of all possible solutions
  #input:[string]: [...]
  _size: math.Round(list.Product([ for key, set in #input { len(set) } ]))
  _c: {
    for key, set in #input { "\(key)": set * (_size div len(set)) }
  }
  out: [
    for index in list.Range(0, _size, 1) {
      for key, set in #input { "\(key)": _c["\(key)"][math.Round(index)] }
    }
  ]
}

#Explode: {
  resource: #Resource
  _solutions: { for x, y in resource.dependencies { "\(x)": y.solutions } }
  _response: { for x, y in resource.dependencies { "\(x)": y.response } }
  _combos: (#combo & {#input: _solutions}).out
  _dependless: {
    type: resource.type
    properties: resource.properties
    system?: resource.system | *{...}
    response?: resource.response | *{...}
  }
  resolved: [
    for c in _combos { 
      solutions: _dependless & { properties: c } 
      response: _dependless & { properties: { for key, solution in c { "\(key)": *solution.response[_response[key]] | null } } }
    }
  ]
}


// ------------------------------------------------------------------------
// Tests
// ------------------------------------------------------------------------
// The following are tests that the engine will use to validate user definitions.

#Test: {
  #matchSystem: {...}
  #testProvider: {system: #matchSystem, response: *{[string]: 1 | "string" | true } | [], ...}
  test:[
    // provider must match Provider definition and can't ahve additional fields.
    #testProvider & #Provider,

    // provider explicit properties must have 'null' as default option, example:
    // myProvider: properties: memory: string | *null
    #testProvider & {properties:[string]:null, ...},

    // provider explicit properties should accept Property type values only.
    // with current definition of Property, this test prevents a property of null type.
    #testProvider & {properties:[string]:1 | "string" | true, ...}
  ]
}

// ------------------------------------------------------------------------
// Examples
// ------------------------------------------------------------------------
_Example: {
  #ExampleProvider: {
    action: "create"
    type: "example"
    system: {
      type: "example"
      properties: {...}
    }
    properties:{
      i: string | *null
      j: string | *null
      _k: string | *null
    }
    implicit: {
      j: [{script: "something", args: { a: properties._k }, resolved: len([ for a in args if a == null {null}]) == 0}]
    }
    response: {
      name: string
    }
  }
  ExampleResource: {
    type: "example"
    properties: {
      i: "i"
      _k: "j"
    }
  }
  ExampleSystem: {
    type: "example"
    properties: {...}
  }
  #ExampleProvisioner: {
    type: "example"
    action: "create"
    system: {
      type: "example"
      properties: {...}
    }
    properties: {
      i: string | *null
      j: string | *null
    }
    provisioner: string
    response: {
      name: string
    }

  }
  Solution: #Solution & {
    #provider: #ExampleProvider
    #provisioner: #ExampleProvisioner
    resource: ExampleResource
    System: ExampleSystem
    ...
  }
  // How to test your component:
  test: (#Test & { #testProvider: #ExampleProvider }).test
}

Example2: {
    R2: {
      type: "Keypair"
      dependencies:
        name: {
          resource: R1
          response: "name"
        }
    }
    R1: {
      type: "Keypair"
      properties: {
        name: "ccit"
      }
    }

  all_resources: [R1, R2]
  Resources: all_resources
  Systems: [
	{
		type: "openstack"
		properties:nova: 1.3
	}
  ]
  #Provisioners: [
    {
      action: "create"
      type: "Keypair"
      system: {
        type: "openstack"
        properties:nova: >1.0
      }
      properties: {
        name: string | *null
      }
      response: {
        name: ""
      }
      provisioner: "linchpin_keypair.tmpl"
    },
    {
      action: "create"
      type: "Server"
      system: {
        type: "openstack"
        properties:nova: >1.0
      }
      properties: {
        name:      string | *null
        imageRef:  string | *null
        flavorRef: string | *null
        key_name:  string | *null
        networks:  string | *null
      }
      provisioner: "linchpin_server.tmpl"
      response: {
        name: ""
      }
    },
    {
      action: "get"
      type: "Flavor"
      system: {
        type: "openstack"
        properties:nova: >1.0
      }
      properties: {
        minDisk:   string | *null
        minMemory: string | *null
      }
      provisioner: "linchpin_flavor.tmpl"
      response: {
        name: ""
      }
    }
  ]
  #Providers: [
    {
      action: "create"
      type: "Keypair"
      system: {
        type: "openstack"
        properties:nova: >1.0
      }
      properties:name: string | *null
      implicit: {...}
      response: {
        name: string
      }
    },
    {
      action: "read"
      type: "Flavor"
      system: {
        type: "openstack"
        properties:nova: >1.0
      }
      properties: {
        minDisk:   string | *null
        minMemory: string | *null
      }
      implicit: {...}
      response: {
        name: string
      }
    },
    {
        action: "create"
        type: "Server"
        system: {
            type: "openstack"
            properties:nova: >1.0
        }
        properties: {
          _memory:   string | *null
          _disk:     string | *null
          name:      string | *null
          imageRef:  string | *null
          flavorRef: string | *null
          key_name:  string | *null
          networks:  string | *null
        }
        implicit: {
          flavorRef: [
            {
              script: "scripts/size_converter.sh"
              args: {
                from: properties._memory
                to: "MiB"
              }
              resolved: len([ for a in args if a == null {null}]) == 0
            },
            {
              script: "scripts/size_converter.sh"
              args: {
                from: properties._disk
                to: "GiB"
              }
              resolved: len([ for a in args if a == null {null}]) == 0
            },
            {
              solutions: (#Solutions & {
                xresources: [{
                  type: "Flavor"
                  properties: {
                    minDisk:  flavorRef[1]
                    minMemory: flavorRef[0]
                  }
                }]
                action: "get"
                systems: [system]
              }).out
              resolved: len(solutions) > 0
            },
            {
              script: "scripts/json_first.sh"
              args: {
                data: flavorRef[2]
              }
              resolved: true
            },
            {
              script: "scripts/json_path.sh"
              args: {
                path: ".name"
                data: flavorRef[3]
              }
              resolved: true
            }
          ]
        }
        response: {
          name: string
        }
    }
  ]
  _test: [
    for provider in #Providers {(#Test & {
      #testProvider: provider
      #matchSystem: { type: "openstack", properties:nova:1.3 }
    }).test}
  ]
  solutions: (#Solutions & {
    action: "create"
    xresources: Resources
    systems: Systems
    providers: #Providers
    provisioners: #Provisioners
  }).out
}
