[[{
    name: "S(Minimal)"
    constrains: [],
    properties: {
        name: [{
            resolved: true
            script:   "something"
            args: {
                a: true
            }
        }]
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
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