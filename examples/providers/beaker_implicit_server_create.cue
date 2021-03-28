{
    type: "Server"
    action: "create"
    system: {
        type: "Beaker"
        properties: {
            version: number | *24
            version: >=24
        }
    }
    properties: {
      name: string | *null
      family: string | *null
      method: *"nfs" | string
      arch: *"x86_64" | string
      memory?: string | *null
      memory_op: *">=" | string
      _image: string | *null
      _memory: string | *null
    }
    implicit: {
        _image: properties._image
        _memory: properties._memory
        name: [{script: "random"}]
        family: *[
            {
                script: "create_string"
                args: {
                    value: _image
                    template: ".*{{ value }}.*"
                }
            },
            {
                resource: {
                    type: "Distro"
                    properties:name: "{{ result }}"
                }
                action: "read"
            },
            { script: "bkr_list_distro_to_family.sh" }
        ] | null
        memory: *[{script: "convert_to_mb", args: {value: _memory }}] | null
    }
}