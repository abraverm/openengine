[[{
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
    name: "R(Minimal)S(Minimal)PD(Implicit2)PR(Minimal)"
    properties: {
        name: [{
            resolved: true
            action:   "create"
            resource: {
                type: "Server"
                name: "Implicit"
                properties: {}
                solutions: []
                dependencies: []
                interfacesDependencies: []
                enabledInterfaces: []
                disabledInterfaces: []
                dependedProperties: {}
            }
            solutions: [{
                name:     "R(Implicit)S(Minimal)PD(Minimal)PR(Minimal)"
                resolved: true
                resource: {
                    type: "Server"
                    name: "Implicit"
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
                interfaces: {}
                constrains: []
                implicit: {}
                joined: {}
                provisioner: "example.sh"
                properties: {}
            }, {
                name:     "R(Implicit)S(Minimal)PD(Implicit2)PR(Minimal)"
                resolved: true
                resource: {
                    type: "Server"
                    name: "Implicit"
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
                    name:   "Implicit2"
                    system: {
                        type: "Openstack"
                        name: "Minimal"
                        properties: {}
                    }
                    properties: {
                        name: null
                    }
                    implicit: {
                        name: [{
                            resource: {
                                type: "Server"
                                name: "Implicit"
                                properties: {}
                                solutions: []
                                dependencies: []
                                interfacesDependencies: []
                                enabledInterfaces: []
                                disabledInterfaces: []
                                dependedProperties: {}
                            }
                            action: "create"
                        }]
                    }
                    interfaces: {}
                    response: {}
                    constrains: []
                }
                interfaces: {}
                constrains: []
                implicit: {
                    name: [{
                        resolved: true
                        action:   "create"
                        resource: {
                            type: "Server"
                            name: "Implicit"
                            properties: {}
                            solutions: []
                            dependencies: []
                            interfacesDependencies: []
                            enabledInterfaces: []
                            disabledInterfaces: []
                            dependedProperties: {}
                        }
                        solutions: [{
                            name:     "R(Implicit)S(Minimal)PD(Minimal)PR(Minimal)"
                            resolved: true
                            resource: {
                                type: "Server"
                                name: "Implicit"
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
                            interfaces: {}
                            constrains: []
                            implicit: {}
                            joined: {}
                            provisioner: "example.sh"
                            properties: {}
                        }]
                    }]
                }
                joined: {
                    name: {
                        explicit: null
                        implicit: [{
                            resolved: true
                            action:   "create"
                            resource: {
                                type: "Server"
                                name: "Implicit"
                                properties: {}
                                solutions: []
                                dependencies: []
                                interfacesDependencies: []
                                enabledInterfaces: []
                                disabledInterfaces: []
                                dependedProperties: {}
                            }
                            solutions: [{
                                name:     "R(Implicit)S(Minimal)PD(Minimal)PR(Minimal)"
                                resolved: true
                                resource: {
                                    type: "Server"
                                    name: "Implicit"
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
                                interfaces: {}
                                constrains: []
                                implicit: {}
                                joined: {}
                                provisioner: "example.sh"
                                properties: {}
                            }]
                        }]
                    }
                }
                provisioner: "example.sh"
                properties: {
                    name: [{
                        resolved: true
                        action:   "create"
                        resource: {
                            type: "Server"
                            name: "Implicit"
                            properties: {}
                            solutions: []
                            dependencies: []
                            interfacesDependencies: []
                            enabledInterfaces: []
                            disabledInterfaces: []
                            dependedProperties: {}
                        }
                        solutions: [{
                            name:     "R(Implicit)S(Minimal)PD(Minimal)PR(Minimal)"
                            resolved: true
                            resource: {
                                type: "Server"
                                name: "Implicit"
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
                            interfaces: {}
                            constrains: []
                            implicit: {}
                            joined: {}
                            provisioner: "example.sh"
                            properties: {}
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