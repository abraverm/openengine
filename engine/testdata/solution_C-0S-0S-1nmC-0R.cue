[[{
    name: "S(Minimal)PD(C-empty)PR(Minimal)"
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
    constrains: []
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]