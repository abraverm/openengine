[[{
    name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
    constrains: [],
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
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}], [{
    name: "R(Minimal)S(Minimal)PD(Implicit4)PR(Minimal)"
    constrains: [],
    properties: {
        name: [{
            resolved: true
            script:   "something"
            args: {
                a: "always"
            }
        }]
    }
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
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]