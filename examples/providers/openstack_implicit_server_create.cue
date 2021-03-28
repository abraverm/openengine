{
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
        properties: {
            nova: number | *2.19
            nova: >=2.19
        }
    }
    properties: {
      name: string | *null
      flavorRef: string | *null
      imageRef: string | *null
       _image: string | *null
       _memory: string | *null
    }
    implicit: {
        _image: properties._image
        _memory: properties._memory
        name: [{script: "random"}]
        flavorRef: *[
            { script: "convert_to_mb", args: {value: _memory }},
            {
                resource: {
                    type: "Flavor"
                    properties: {
                        minRam: "{{ result }}"
                    }
                }
                action: "read"
            },
            { script: "first_flavor" }
        ] | null
        imageRef: *[
            {
                resource: {
                    type: "Image"
                    properties: {
                        name: _image
                    }
                }
                action: "read"
            },
            { script: "first_image" }
        ] | null
    }
}