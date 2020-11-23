#ImplicitTask: {
  script: string | *null
  args: {...}
  resolved: bool
  ...
}

#solution: {
  provider: #provider
  resource: properties: { ... }
  _match: provider & {
    properties: resource.properties
  }
  _joined: {
    for x, y in _match.properties { "\(x)":explicit: y }
    for x, y in _match.implicit if y != null { "\(x)":implicit: *(y & [...{resolved: true, ...}]) | null }
  }
  properties: {
    for x, y in _joined {
      if (*y.implicit | null) == null { "\(x)": y.explicit }
      if y.explicit == null { "\(x)": y.implicit | null }
      if (*y.implicit | null) == null && y.explicit == null { "\(x)":  null }
      if (*y.implicit | null) != null && y.explicit != null { "\(x)":  y.explicit }
    }
  }
  resolved: len([for z in properties if z == null {null}]) == 0
}

#provider : {
  properties: [string]: string | number | bool | *null
  implicit: [string]: [#ImplicitTask, ...#ImplicitTask]
}

// OE validiation:
#null: {
  properties:[string]: null
  ...
}

#standardProperty: 1 | "string" | true


tests:[
  _subject & #null,     // explicit properties must have `null` as default value
  _subject & #provider, // valid provider fields
  _subject & { properties:[string]:#standardProperty } // valid properties type
]

_subject: #openstack

// Example provider definition

#openstack: {
  properties:i: string | *null 
  properties:_k: string | *null
  properties:_g: string | *null
  properties:j: string | *null
  implicit:j: [{script: "something", args: { a: properties._k }, resolved: len([ for a in args if a == null {null}]) == 0}]
//  implicit:i: [{script: "something", args: { a: properties._g }, resolved: len([ for a in args if a == null {null}]) == 0}]
}

// Example user requested resource matcehd with the provider

_resource: properties:i: "one"
//_resource: properties:_g: "one"
//_resource: properties:j: "one"
_resource: properties:_k: "one"
//_resource: {}
solution: #solution & { provider: #openstack, resource: _resource, ... }
