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
                    type:   "Network"
                    name:   "ID-Network"
                    system: {
                        type: "Openstack"
                        name: "Minimal"
                        properties: {}
                    }
                    properties: {}
                    implicit: {}
                    interfaces: {}
                    response: {
                        id: ""
                    }
                    constrains: []
                }
                implicit: {}
                joined: {}
                provisioner: "example.sh"
                properties: {}
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
                    type:   "Storage"
                    name:   "ID-Storage"
                    system: {
                        type: "Openstack"
                        name: "Minimal"
                        properties: {}
                    }
                    properties: {}
                    implicit: {}
                    interfaces: {}
                    response: {
                        id: ""
                    }
                    constrains: []
                }
                implicit: {}
                joined: {}
                provisioner: "example.sh"
                properties: {}
                resolved: true
            }
        }
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ID-Minimal3"
        properties: {}
        solutions: [{
            name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
            resource: {
                type: "Network"
                name: "ID-Network"
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
                type:   "Network"
                name:   "ID-Network"
                system: {
                    type: "Openstack"
                    name: "Minimal"
                    properties: {}
                }
                properties: {}
                implicit: {}
                interfaces: {}
                response: {
                    id: ""
                }
                constrains: []
            }
            implicit: {}
            joined: {}
            provisioner: "example.sh"
            properties: {}
            resolved: true
        }, {
            name: "R(ID-Storage)S(Minimal)PD(ID-Storage)PR(ID-Storage)"
            resource: {
                type: "Storage"
                name: "ID-Storage"
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
                type:   "Storage"
                name:   "ID-Storage"
                system: {
                    type: "Openstack"
                    name: "Minimal"
                    properties: {}
                }
                properties: {}
                implicit: {}
                interfaces: {}
                response: {
                    id: ""
                }
                constrains: []
            }
            implicit: {}
            joined: {}
            provisioner: "example.sh"
            properties: {}
            resolved: true
        }]
        dependencies: ["ID-Network", "ID-Storage"]
        interfacesDependencies: ["ID-Network", "ID-Storage"]
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
    name: "R(ID-Network)S(Minimal)PD(ID-Network)PR(ID-Network)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Network"
        name: "ID-Network"
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
}, {
    name: "R(ID-Storage)S(Minimal)PD(ID-Storage)PR(ID-Storage)"
    properties: {}
    provisioner: "example.sh"
    resource: {
        type: "Storage"
        name: "ID-Storage"
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