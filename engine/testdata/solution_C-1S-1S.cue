[[{
    name: "R(Minimal)S(Minimal)PD(C-1S-1S)PR(Minimal)"
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
        name: "1S-1S"
        pre: [{
            resolved: true
            script:   "pre.sh"
            args: {}
        }]
        post: [{
            resolved: true
            script:   "pre.sh"
            args: {}
        }]
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
        properties: {}
    }
}]]