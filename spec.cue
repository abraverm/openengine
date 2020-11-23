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
#Property: string | bool | number | [...#Property]
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
  properties: #Properties
  response?: #Properties
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
  script: string | *null   // The script to call
  args: #Properties // Arguments for the script, they could be explicit or linked to implicit properties
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
  response: #Property | *null
}

// Provisioner
#Provisioner: {
  type: #Type
  action: #Action
  system: #System
  properties: #Properties
  provisioner: string
  response: #Property | *null
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
    type: resource.type & #provisioner.type
    properties: resource.properties & #provisioner.properties
    response: #provisioner.response & (resource.response | { ... })
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


// ------------------------------------------------------------------------
// Tests
// ------------------------------------------------------------------------
// The following are tests that the engine will use to validate user definitions.

#Test: {
  #testProvider: {...} | *{
    action: "create"
    type: "test"
    system:type: "test"
  }
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
Example: {
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
    response: []
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
    response: []

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
