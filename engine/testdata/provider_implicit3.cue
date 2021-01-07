{
    type: "Server"
    name: "Implicit2"
    action: "create"
    system: {
        type: "Openstack"
    }
    properties: {
        name: string | *null
        _random: bool | *null
    }
    implicit: {
        name: *[{
            resource: {
                type:"Other"
                name: "Implicit"
            }
            action: "create"
        }] | null
    }
}