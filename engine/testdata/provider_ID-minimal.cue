{
    type: "Server"
    action: "create"
    name: "ID-minimal"
    system: {
        type: "Openstack"
    }
    properties: {
        network: string | *null
    }
    interfaces:network:[
        {
            name: "Openstack Server network from Network"
            type: "Network"
            action: "create"
            response: {
                id: string | *null
            }
            field: "id"
        }
    ]
}