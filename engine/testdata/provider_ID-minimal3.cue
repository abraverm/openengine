{
    type: "Server"
    action: "create"
    name: "ID-minimal3"
    system: {
        type: "Openstack"
    }
    properties: {
        network: string | *null
        storage: string | *null
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
   interfaces:storage:[
        {
            name: "Openstack Server storage from Storage"
            type: "Storage"
            action: "create"
            response: {
                id: string | *null
            }
            field: "id"
        }
    ]
}