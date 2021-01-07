{
    type: "Server"
    action: "create"
    name: "ID-cycle"
    system: {
        type: "Openstack"
    }
    properties: {
        chicken: string | *null
    }
    interfaces:chicken:[
        {
            name: "Openstack Server chicken from Server"
            type: "Server"
            action: "create"
            response: {
                id: string | *null
            }
            field: "id"
        }
    ]
}