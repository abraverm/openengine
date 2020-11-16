package example

import "crypto/sha256"
import "encoding/json"
import "strings"

#action: *"create" | "delete" | "read" | "update"

_action: #action @tag(action,short=create|delete|read|update)

#implicit_task: {
  script: string | *""
  args: {...}
  resolved: bool | *false
  ...
}

#solution: {
  action: #action
  resource: {
    type: string
    ...
  }
  system: {
    type: string
    ...
  }
  response?: {...}
  provisioner: string
}


#implicit_tasks: [...#implicit_task]

_solution: {
  xaction: string
  xresource: {...}
  xsystem: {...}
  _out: [
    for i, _provider in #providers
    for j, _provisioner in #provisioners
    {
      _name: *(strings.Replace(json.Marshal(sha256.Sum256(json.Marshal({action: xaction, resource: xresource, system: xsystem, provider: i, provisioner: j}))), "\"", "", -1)) | ""
       "\(_name)": *(_provider & _provisioner & {action: xaction, resource: xresource,  system: xsystem}) | null
    }
  ]
  out: [
    for batch in _out
    for s in batch
    if s != null { s }
  ]
  //out: [ for s in _tmp if len([for key, value in s.resource if value == null { null }]) == 0 { s } ] | []
}

//_solutions: [
//  for _resource in _resources
//  for _system in _systems
//  {
//    _name: *(strings.Replace(json.Marshal(sha256.Sum256(json.Marshal({action: _action, resource: _resource, system: _system}))), "\"", "", -1)) | ""
//    "\(_name)": (_solution & {xaction: _action, xresource: _resource ,  xsystem: _system}).out
//  }
//]

_dependent_solutions: [
  for _resource in _resources_dependent
  for _system in _systems
  {
    _name: *(strings.Replace(json.Marshal(sha256.Sum256(json.Marshal({action: _action, resource: _resource, system: _system}))), "\"", "", -1)) | ""
    "\(_name)": (_solution & {xaction: _action, xresource: _resource ,  xsystem: _system}).out
  }
]

_dependentless_solutions: [
  for _resource in _resources_dependentless
  for _system in _systems
  {
    _name: *(strings.Replace(json.Marshal(sha256.Sum256(json.Marshal({action: _action, resource: _resource, system: _system}))), "\"", "", -1)) | ""
    "\(_name)": (_solution & {xaction: _action, xresource: _resource ,  xsystem: _system}).out
  }
]

// solutions: [
//  for batch in _solutions
//  for s in batch
//  for r in s
//  if len([for key, value in r.resource if value == null { null }]) == 0 { r }
//]

solutions: [
  for batch in _dependent_solutions
  for s in batch
  for r in s
  if len([for key, value in r.resource if value == null { null }]) == 0 { r }
]

dependentless_solutions: [
  for batch in _dependentless_solutions
  for s in batch
  for r in s
  if len([for key, value in r.resource if value == null { null }]) == 0 { r }
]

_resources_dependentless: [
  for resource in _resources
  if len([for param in resource if param == null { true }]) == 0 { resource }
]

_resources_dependent: [
  for resource in _resources
  if len([for param in resource if param == null { true }]) > 0 { resource }
]

#dependency: {
  #resource: {...}
  #solutions: [...]
  #response: {...}
  out: [
    for s in #solutions
    if (*(s.resource & #resource) | null ) != null && (*(s.response & #response) | null ) != null { s }
  ]
}
