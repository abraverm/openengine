[[{
    name: "R(ID-Minimal3)S(Minimal)PD(ID-minimal3)PR(ID-Minimal2)"
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
        storage: {
            response: {
                id: ""
            }
            action: "create"
            name:   "Openstack Server storage from Storage"
            type:   "Storage"
            field:  "id"
            solution: {
                name: "R(ID-Storage)S(Minimal)PD(ID-Storage)PR(ID-Storage)"
                resource: {
                    type: "Storage"
                    name: "ID-Storage"
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
        name: "ID-Minimal3"
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
        }, {
            name: "R(ID-Storage)S(Minimal)PD(ID-Storage)PR(ID-Storage)"
            resource: {
                type: "Storage"
                name: "ID-Storage"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["ID-Network", "ID-Storage"]
        interfacesDependencies: ["ID-Network", "ID-Storage"]
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
}, {
    name: "R(ID-Storage)S(Minimal)PD(ID-Storage)PR(ID-Storage)"
    provisioner: "example.sh"
    resource: {
        type: "Storage"
        name: "ID-Storage"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]