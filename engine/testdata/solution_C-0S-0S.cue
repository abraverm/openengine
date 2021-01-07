[[{
    name: "R(Minimal)S(Minimal)PD(C-0S-0S)PR(Minimal)"
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
        name: "0S-0S"
        pre: []
        post: []
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]