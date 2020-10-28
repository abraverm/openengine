package engine

#action: *"create" | "delete" | "read" | "update"

_action: #action @tag(action,short=create|delete|read|update)

_resources: [
	{
		name:     "my_server"
		imageRef: "Fedora-31"
		key_name: "ccit"
		networks: "ccit-net"
	},
]

_systems: [
	{
		type: "openstack"
		nova: 1.3
	},
]

#Server: _openstack_server

_openstack_server: {
	action: "create"
	system: {
		type: "openstack"
		nova: >1.0
	}
	resource: {
		name:       string
		imageRef:   string
		flavorRef?: string
		key_name:   string
		networks:   string
	}
	...
}

#provisioner: #Linchpin
#provider:    #Server

_solutions: [{
	for _resource in _resources {
		for _system in _systems {
			solution: *( #provisioner & #provider & _solution ) | null
			_solution: {
				action:   _action
				resource: _resource
				system:   _system
			}
		}
	}
}]

solutions: [ for v in _solutions if v.solution != null { v.solution } ]
