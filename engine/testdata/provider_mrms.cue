{
    type: "Server"
    action: "create"
    name: "mrms"
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
                    type: "Server"
                    name: "Implicit"
                }
                action: "create"
            }
        ] | null
    }
}