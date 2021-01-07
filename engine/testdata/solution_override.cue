[[{
    name: "S(Minimal)"
    constrains: [],
    properties: {
        name: "explicit"
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        properties: {
            name: "explicit"
        }
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