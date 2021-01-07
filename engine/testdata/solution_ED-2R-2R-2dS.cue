[[{
    name: "R(ED-Other2)S(Other)PD(Other)PR(Other)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Other"
        name: "ED-Other2"
        properties: {}
        solutions: [{
            name: "R(Other)S(Other)PD(Other)PR(Other)"
            resource: {
                type: "Other"
                name: "Other"
                properties: {}
                solutions: []
                dependencies: []
                interfacesDependencies: []
                enabledInterfaces: []
                disabledInterfaces: []
                dependedProperties: {}
            }
            System: {
                type: "Other"
                name: "Other"
                properties: {}
            }
            match: {
                action: "create"
                type:   "Other"
                name:   "Other"
                system: {
                    type: "Other"
                    name: "Other"
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
        dependencies: ["Other"]
        interfacesDependencies: []
        enabledInterfaces: []
        disabledInterfaces: []
        dependedProperties: {}
    }
    constrains: []
    system: {
        type: "Other"
        name: "Other"
        properties: {}
    }
}, {
    name: "R(Other)S(Other)PD(Other)PR(Other)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Other"
        name: "Other"
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
        type: "Other"
        name: "Other"
        properties: {}
    }
}], [{
    name: "R(ED-Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
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
        }]
        dependencies: ["Minimal"]
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
}]]