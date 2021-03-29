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
    }
}