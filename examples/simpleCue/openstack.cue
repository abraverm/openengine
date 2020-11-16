package example

#Openstack: [
for resource in _openstack
for action in resource
for system_match in action
{ system_match }
]

_openstack: keypair: create: [{
  action: "create"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    type:       "Keypair"
    name: string | [...#solution]
  }
  response: {
    name: string
  }
  provisioner: string
}]

_openstack: flavor: get: [{
  action: "get"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    type:       "Flavor"
    minDisk: string
    minMemory: string
  }
  provisioner: string
}]

_openstack: server: create: [{
	action: "create"
	system: {
		type: "openstack"
		nova: >1.0
	}
	resource: {
        type:       "Server"
        _memory:   string | *""
        _disk:     string | *""
		name:       string
		imageRef:   string| [...#solution]
        flavorRef: string | *(*_flavorRef | null)
        _flavorRef: *(([
            {
              script: "scripts/size_converter.sh"
              args: {
                from: resource._memory
                to: "MiB"
              }
              result: string | *""
              resolved: len(args.from) > 0
            },
            {
              script: "scripts/size_converter.sh"
              args: {
                from: resource._disk
                to: "GiB"
              }
              resolved: len(args.from) > 0
              result: string | *""
            },
            {
              flavors: (_solution & {
                xaction: "get"
                xresource: {
                  type: "Flavor"
                  minDisk:  _flavorRef[1].result
                  minMemory: _flavorRef[0].result
                }
                xsystem: system
              }).out 
              resolved: len(flavors) > 0
              result: string | *""
            },
            {
              script: "scripts/json_first.sh"
              args: {
                data: _flavorRef[2].result
              }
              resolved: true
              result: string | *""
            },
            {
              script: "scripts/json_path.sh"
              args: {
                path: ".name"
                data: _flavorRef[3].result
              }
              resolved: true
              result: string | *""
            }
          ] & [...{resolved?: true, ...}])) | null
        key_name:   string | [...#solution]
		networks:   string
	}
    provisioner: string
    response: {
      name: string
    }
}]
