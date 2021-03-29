[[{
    name: "R(ID-Minimal)S(Minimal)PD(ID-minimal)PR(ID-Minimal)"
    properties: {
        network: {
            response: {
                id: ""
            }
            action: "create"
            name:   "Openstack Server network from Network"
            type:   "Network"
            field:  "id"
            solution: {
                name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
                resource: {
                    type: "Network"
                    name: "ID-Network"
                }
                system: {
                    type: "Openstack"
                    name: "Minimal"
                }
                provisioner: "example.sh"
                resolved: true
            }
        }
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ID-Minimal"
        solutions: [{
            name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
            resource: {
                type: "Network"
                name: "ID-Network"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["ID-Network"]
        interfacesDependencies: ["ID-Network"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
    name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
    provisioner: "example.sh"
    resource: {
        type: "Network"
        name: "ID-Network"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}], [{
    name: "R(ID-Minimal)S(Minimal)PD(ID-minimal)PR(ID-Minimal2)"
    properties: {
        network: {
            response: {
                id: ""
            }
            action: "create"
            name:   "Openstack Server network from Network"
            type:   "Network"
            field:  "id"
            solution: {
                name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
                resource: {
                    type: "Network"
                    name: "ID-Network"
                }
                system: {
                    type: "Openstack"
                    name: "Minimal"
                }
                provisioner: "example.sh"
                resolved: true
            }
        }
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ID-Minimal"
        solutions: [{
            name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
            resource: {
                type: "Network"
                name: "ID-Network"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["ID-Network"]
        interfacesDependencies: ["ID-Network"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
    name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
    provisioner: "example.sh"
    resource: {
        type: "Network"
        name: "ID-Network"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]