{
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
    }
    properties: {
        volume: string | *null
    }
    interfaces:volume:[
        {
            name: "Resolved from depended volume"
            type: "Volume"
            action: "create"
            response: {
                id: string | *null
            }
            field: "id"
        }
    ]
}