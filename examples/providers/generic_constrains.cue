{
    name: "Generic Explicit"
    type: "Example Resource Type"
    action: "create"
    system: {
        type: "Example Provider"
    }
    properties: {
        name: string | *null
    }
    constrains: [
        {
            name: "Example of active constrain"
            pre: [{ script: "pre.sh" }]
            post: [{ script: "post.sh" }]
            properties: {
                name: =~ "Example" | *null
            }
        },
        {
            name: "Example of disabled constrain" // this won't show in the solutions
            pre: [{ script: "pre.sh" }]
            post: [{ script: "post.sh" }]
            properties: {
                name: =~ "Other" | *null
            }
        }
    ]
}