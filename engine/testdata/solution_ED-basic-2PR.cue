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
    name: "R(ED-Minimal)S(Minimal)PD(Minimal)PR(Two)"
    provisioner: "example2.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
        solutions: [{
            name: "R(Minimal)S(Minimal)PD(Minimal)PR(Two)"
            resource: {
                type: "Server"
                name: "Minimal"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example2.sh"
            resolved: true
        }]
        dependencies: ["Minimal"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
    name: "R(Minimal)S(Minimal)PD(Minimal)PR(Two)"
    provisioner: "example2.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}], [{
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
    name: "R(ED-Minimal)S(Minimal)PD(Minimal)PR(Two)"
    provisioner: "example2.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
        solutions: [{
            name: "R(Minimal)S(Minimal)PD(Minimal)PR(Two)"
            resource: {
                type: "Server"
                name: "Minimal"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example2.sh"
            resolved: true
        }]
        dependencies: ["Minimal"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
    name: "R(Minimal)S(Minimal)PD(Minimal)PR(Two)"
    provisioner: "example2.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]