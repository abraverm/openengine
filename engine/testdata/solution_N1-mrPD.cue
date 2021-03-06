[[{
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
    name: "R(Minimal)S(Minimal)PD(Implicit2)PR(Minimal)"
    properties: {
        name: [{
            resolved: true
            action:   "create"
            resource: {
                type: "Server"
                name: "Implicit"
            }
            solutions: [{
                name:     "R(Implicit)S(Minimal)PD(Minimal)PR(Minimal)"
                resolved: true
                resource: {
                    type: "Server"
                    name: "Implicit"
                }
                system: {
                    type: "Openstack"
                    name: "Minimal"
                }
                provisioner: "example.sh"
            }, {
                name:     "R(Implicit)S(Minimal)PD(Implicit2)PR(Minimal)"
                resolved: true
                resource: {
                    type: "Server"
                    name: "Implicit"
                }
                system: {
                    type: "Openstack"
                    name: "Minimal"
                }
                provisioner: "example.sh"
                properties: {
                    name: [{
                        resolved: true
                        action:   "create"
                        resource: {
                            type: "Server"
                            name: "Implicit"
                        }
                        solutions: [{
                            name:     "R(Implicit)S(Minimal)PD(Minimal)PR(Minimal)"
                            resolved: true
                            resource: {
                                type: "Server"
                                name: "Implicit"
                            }
                            system: {
                                type: "Openstack"
                                name: "Minimal"
                            }
                            provisioner: "example.sh"
                        }]
                    }]
                }
            }]
        }]
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]
