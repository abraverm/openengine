{
    type: "Server"
    action: "create"
    system: {
        type: "AWS"
        properties: {
            api_versions: {
                ec2: int | *20160102
                ec2: >20160101
            }
        }
    }
    provisioner: "aws ec2 run-instances --image-id {{ ImageID }} --count {{ MaxCount }} --instance-type {{ InstanceType }} --key-name  {{ KeyName }}"
    properties: {
      ImageId: string | *null
      KeyName: string | *null
      InstanceType: string | *null
      MaxCount: int | *1
      MinCount: int | *1
    }
}