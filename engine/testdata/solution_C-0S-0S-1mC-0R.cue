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
    constrains: [{
        name: "0S-0S-1nmC-0R"
        pre: []
        post: []
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]