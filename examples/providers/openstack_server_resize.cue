{
    type: "Server"
    action: "update"
    system: {
         type: "Openstack"
    }
    properties: {
        name: string | *null
        flavorRef: string | *null
    }
    _properties: {
        name: properties.name
    }
    constrains: [
        {
            name: "Resize constrain"
            pre: [
                {
                    action: "update"
                    resource: {
                        type: "Server"
                        properties: {
                            name: _properties.name
                            stop: true
                        }
                    }
                }
            ]
            post: [
                {
                    action: "update"
                    resource: {
                        type: "Server"
                        properties: {
                            name: _properties.name
                            start: true
                        }
                    }
                },
            ]
            properties: {
                name: string | *null
                flavorRef: string | *null
            }
        }
    ]
}