{
    type: "Server"
    action: "create"
    name: "C-0S-0S"
    system: {
        type: "Openstack"
    }
    constrains: [
        {
            name: "0S-0S"
            pre: []
            post: []
            response: {
                id: string | *null
            }
        }
    ]
    response: {
        id: string | *null
    }
}