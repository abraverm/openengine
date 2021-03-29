{
    type: "Server"
    action: "create"
    system: {
        type: "AWS"
        properties: {
            api_versions: {
                ec2: int | *20210301
                ec2: >20150301
                ec2: <=20210301
             }
        }
    }
    properties: {
      ImageId?: string | *null
      KeyName?: string | *null
      InstanceType?: string | *null
      MaxCount: int | *1
      MinCount: int | *1
      _image: string | *null
      _memory: string | *null
    }
    implicit: {
        _image: properties._image
        ImageId: *[{
            resource: {
                type: "DescribeImages"
                properties: {
                    name: _image
                }
            }
            action: "read"
        },
        { script: "first_image" }
        ] | null
        InstanceType: *[{ script: "aws_flavor", args: {flavor: properties._memory }}] | null
    }
}