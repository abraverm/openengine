{
    type: "Server"
    action: "create"
    name: "nrms"
    system: {
        type: "Openstack"
    }
    properties: {
        name: string | *null
    }
    implicit: {
        name: *[
            {script: "something", args: { a: "always" }},
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