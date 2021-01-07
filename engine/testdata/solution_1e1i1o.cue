[[{
    name: "R(1e1i1o)S(Minimal)PD(Mixed)PR(Mixed)"
    constrains: [],
    properties: {
        name: [{
            script: "something"
            args: {
                a: true
            }
            resolved: true
        }]
        disk:   "10g"
        memory: "10g"
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "1e1i1o"
        properties: {
            disk:   "10g"
            memory: "10g"
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