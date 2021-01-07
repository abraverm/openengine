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
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ID-Minimal"
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
        }]
        dependencies: ["ID-Network"]
        interfacesDependencies: ["ID-Network"]
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
}], [{
    name: "R(ID-Minimal)S(Two)PD(ID-minimal)PR(ID-Minimal)"
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
                name: "R(ID-Network)S(Two)PD(ID-Network)PR(ID-Network)"
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
                    name: "Two"
                    properties: {}
                }
                match: {
                    action: "create"
                    type:   "Network"
                    name:   "ID-Network"
                    system: {
                        type: "Openstack"
                        name: "Two"
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
        name: "ID-Minimal"
        properties: {}
        solutions: [{
            name: "R(ID-Network)S(Two)PD(ID-Network)PR(ID-Network)"
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
                name: "Two"
                properties: {}
            }
            match: {
                action: "create"
                type:   "Network"
                name:   "ID-Network"
                system: {
                    type: "Openstack"
                    name: "Two"
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
        dependencies: ["ID-Network"]
        interfacesDependencies: ["ID-Network"]
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
    name: "R(ID-Network)S(Two)PD(ID-Network)PR(ID-Network)"
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
        name: "Two"
        properties: {}
    }
}]]