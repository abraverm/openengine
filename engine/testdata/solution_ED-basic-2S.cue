// Resource "ED-Minimal" has explicit dependency for "Minimal":
//   - The resource solutions are grouped into one set
//   - The set is complete
//      * Has solutions for all resources
//      * Each resource has a full solution with all the additional dependencies: constrains, implicit properties, explicit dependencies, etc
//   - The set is minimal
//      * All solutions and nested are in the same system
//      * Dependency solutions (explicit) are matching to root level solutions
//      * One solution for each declared resource
// Two system are matching, thus two sets of solutions - one per each system
[[{
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
}], [{
    name: "R(ED-Minimal)S(Two)PD(Minimal)PR(Minimal)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
        properties: {}
        solutions: [{
            name: "R(Minimal)S(Two)PD(Minimal)PR(Minimal)"
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
                name: "Two"
                properties: {}
            }
            match: {
                action: "create"
                type:   "Server"
                name:   "Minimal"
                system: {
                    type: "Openstack"
                    name: "Two"
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
        name: "Two"
        properties: {}
    }
}, {
    name: "R(Minimal)S(Two)PD(Minimal)PR(Minimal)"
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
        name: "Two"
        properties: {}
    }
}]]