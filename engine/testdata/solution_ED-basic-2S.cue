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
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
        solutions: [{
            name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
            resource: {
                type: "Server"
                name: "Minimal"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["Minimal"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
    name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}], [{
    name: "R(ED-Minimal)S(Two)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
        solutions: [{
            name: "R(Minimal)S(Two)PD(Minimal)PR(Minimal)"
            resource: {
                type: "Server"
                name: "Minimal"
            }
            system: {
                type: "Openstack"
                name: "Two"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["Minimal"]
    }
    system: {
        type: "Openstack"
        name: "Two"
    }
}, {
    name: "R(Minimal)S(Two)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Two"
    }
}]]