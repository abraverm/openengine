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
                no: string | *null
            }
            field: "no"
        }
    ]
}