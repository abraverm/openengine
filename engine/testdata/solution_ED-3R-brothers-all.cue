[[{
    name: "R(ED-Minimal2)S(Minimal)PD(Minimal)PR(Minimal)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal2"
        properties: {}
        solutions: [{
            name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
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
            System: {
                type: "Openstack"
                name: "Minimal"
                properties: {}
            }
            match: {
                action: "create"
                type:   "Server"
                name:   "Minimal"
                system: {
                    type: "Openstack"
                    name: "Minimal"
                    properties: {}
                }
                properties: {}
                implicit: {}
                interfaces: {}
                response: {}
                constrains: []
            }
            implicit: {}
            joined: {}
            provisioner: "example.sh"
            properties: {}
            resolved: true
        }, {
            name: "R(Two)S(Minimal)PD(Minimal)PR(Minimal)"
            resource: {
                type: "Server"
                name: "Two"
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
                type:   "Server"
                name:   "Minimal"
                system: {
                    type: "Openstack"
                    name: "Minimal"
                    properties: {}
                }
                properties: {}
                implicit: {}
                interfaces: {}
                response: {}
                constrains: []
            }
            implicit: {}
            joined: {}
            provisioner: "example.sh"
            properties: {}
            resolved: true
        }]
        dependencies: ["Minimal", "Two"]
        interfacesDependencies: []
        enabledInterfaces: []
        disabledInterfaces: []
        dependedProperties: {}
    }
    constrains: []
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}, {
    name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
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
    constrains: []
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}, {
    name: "R(Two)S(Minimal)PD(Minimal)PR(Minimal)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Two"
        properties: {}
        solutions: []
        dependencies: []
        interfacesDependencies: []
        enabledInterfaces: []
        disabledInterfaces: []
        dependedProperties: {}
    }
    constrains: []
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]