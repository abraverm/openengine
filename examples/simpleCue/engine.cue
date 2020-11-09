package engine

#action: *"create" | "delete" | "read" | "update"

_action: #action @tag(action,short=create|delete|read|update)

_resources: [
	{
		name:     "my_server"
        type:     "Server"
		imageRef: "Fedora-31"
		key_name: "ccit"
		networks: "ccit-net"
        flavorRef: "Sfd"
        _memory: "4g"
        _disk: "10g"
	},
]

_systems: [
	{
		type: "openstack"
		nova: 1.3
	},
]

#Server: #_openstack_server

#Flavor: [#_openstack_flavor_get]

#implicit_task: {
  script: string | *""
  args: {...}
  resolved: bool | *false
  ...
}


#implicit_tasks: [...#implicit_task]

#_openstack_flavor_get: {
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
}

#_openstack_server: {
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
		imageRef:   string
        flavorRef: string | *(*_flavorRef | null)
        _flavorRef: ([
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
              flavors: [1]
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
          ] & [...{resolved?: true, ...}])
        key_name:   string
		networks:   string
	}
    provisioner: string
}


#provisioner: #Linchpin
#provider: #Server

#solution: {
  #do: #action
  #resource: {...}
  #system: {...}
  out: [{
    for _provider in #provider {
      solution: *(#provisioner & _provider & { action: #do, resource: #resource, system: #system }) | null
      resolved: len([for key, value in solution.resource if value == null { null }]) == 0 | false
    }
  }]
}

_solutions: [
	for _resource in _resources {
      for _system in _systems {
          solution: #provider & #provisioner & {action: _action, resource: _resource ,  system: _system}
          resolved: len([for key, value in solution.resource if value == null { null }]) == 0
      }
    }
  ]

solutions: [ for s in _solutions if s.resolved == true { s.solution } ]
