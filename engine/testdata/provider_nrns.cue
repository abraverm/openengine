{
    type: "Server"
    action: "create"
    name: "nrns"
    system: {
        type: "Openstack"
    }
    properties: {
        name: string | *null
        _random: bool | *null
    }
    implicit: {
        name: *[
            {script: "something", args: { a: properties._random }},
            {
                resource: {
                    type: "Other"
                    name: "Implicit"
                }
                action: "create"
            }
        ] | null
    }
}