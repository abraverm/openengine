[[{
    name: "R(Minimal)S(Minimal)PD(C-1R-1R)PR(Minimal)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
        properties: {}
        solutions: []
        dependencies: []
        interfacesDependencies: []
        enabledInterfaces: []
        disabledInterfaces: []
        dependedProperties: {}
    }
    constrains: [{
        name: "1R-1R"
        pre: [{
            resolved: true
            action:   "create"
            resource: {
                type: "Storage"
                properties: {}
                solutions: []
                dependencies: []
                interfacesDependencies: []
                enabledInterfaces: []
                disabledInterfaces: []
                dependedProperties: {}
            }
            solutions: [{
                name:     "S(Minimal)PD(ID-Storage)PR(ID-Storage)"
                resolved: true
                resource: {
                    type: "Storage"
                    properties: {}
                    solutions: []
                    dependencies: []
                    interfacesDependencies: []
                    enabledInterfaces: []
                    disabledInterfaces: []
                    dependedProperties: {}
                }
                System: {
                    type: "Openstack"
                    name: "Minimal"
                    properties: {}
                }
                match: {
                    action: "create"
                    type:   "Storage"
                    name:   "ID-Storage"
                    system: {
                        type: "Openstack"
                        name: "Minimal"
                        properties: {}
                    }
                    properties: {}
                    implicit: {}
                    interfaces: {}
                    response: {
                        id: ""
                    }
                    constrains: []
                }
                interfaces: {}
                constrains: []
                implicit: {}
                joined: {}
                provisioner: "example.sh"
                properties: {}
            }]
        }]
        post: [{
            resolved: true
            action:   "create"
            resource: {
                type: "Storage"
                properties: {}
                solutions: []
                dependencies: []
                interfacesDependencies: []
                enabledInterfaces: []
                disabledInterfaces: []
                dependedProperties: {}
            }
            solutions: [{
                name:     "S(Minimal)PD(ID-Storage)PR(ID-Storage)"
                resolved: true
                resource: {
                    type: "Storage"
                    properties: {}
                    solutions: []
                    dependencies: []
                    interfacesDependencies: []
                    enabledInterfaces: []
                    disabledInterfaces: []
                    dependedProperties: {}
                }
                System: {
                    type: "Openstack"
                    name: "Minimal"
                    properties: {}
                }
                match: {
                    action: "create"
                    type:   "Storage"
                    name:   "ID-Storage"
                    system: {
                        type: "Openstack"
                        name: "Minimal"
                        properties: {}
                    }
                    properties: {}
                    implicit: {}
                    interfaces: {}
                    response: {
                        id: ""
                    }
                    constrains: []
                }
                interfaces: {}
                constrains: []
                implicit: {}
                joined: {}
                provisioner: "example.sh"
                properties: {}
            }]
        }]
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]