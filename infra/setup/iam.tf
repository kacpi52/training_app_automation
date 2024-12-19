# create IAM user and policies for the cicd 
resource "aws_iam_user" "cd" {
  name = "training-app-cd"
}

resource "aws_iam_access_key" "cd" {
  user = aws_iam_user.cd.name
}

# create permissions for user and policy for s3 and dynamodb access
# permission to list specified bucket with tf state and get data to create policies
data "aws_iam_policy_document" "tf_backend" {
  statement {
    effect    = "Allow"
    actions   = ["s3:ListBucket"]
    resources = ["arn:aws:s3:::${var.tf_state_bucket}"]
  }

  statement {
    effect  = "Allow"
    actions = ["s3:GetObject", "s3:PutObject", "s3:DeleteObject"]
    resources = [
      "arn:aws:s3:::${var.tf_state_bucket}/tf-state-deploy/*",
      "arn:aws:s3:::${var.tf_state_bucket}/tf-state-deploy-env/*",
    ]
  }

  statement {
    effect = "Allow"
    actions = [
      "dynamodb:DescribeTable",
      "dynamodb:GetItem",
      "dynamodb:PutItem",
      "dynamodb:DeleteItem"
    ]
    resources = ["arn:aws:dynamodb:*:*:table/${var.tf_state_lock_table}"]
  }
}
# create policy
resource "aws_iam_policy" "tf_backend" {
  name        = "${aws_iam_user.cd.name}-tf-s3-dynamodb"
  description = "allow user to use s3 and dynamodb for tf backend resources"
  policy      = data.aws_iam_policy_document.tf_backend.json
}

# attach user to the policy
resource "aws_iam_user_policy_attachment" "tf_backend" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.tf_backend.arn
}


# ec2 policy 

data "aws_iam_policy_document" "ec2" {
  statement {
    effect = "Allow"
    actions = [
      "ec2:DescribeVpcs",
      "ec2:CreateTags",
      "ec2:CreateVpc",
      "ec2:DeleteVpc",
      "ec2:DescribeSecurityGroups",
      "ec2:DeleteSubnet",
      "ec2:DeleteSecurityGroup",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DetachInternetGateway",
      "ec2:DescribeInternetGateways",
      "ec2:DeleteInternetGateway",
      "ec2:DetachNetworkInterface",
      "ec2:DescribeVpcEndpoints",
      "ec2:DescribeRouteTables",
      "ec2:DeleteRouteTable",
      "ec2:DeleteVpcEndpoints",
      "ec2:DisassociateRouteTable",
      "ec2:DeleteRoute",
      "ec2:DescribePrefixLists",
      "ec2:DescribeSubnets",
      "ec2:DescribeVpcAttribute",
      "ec2:DescribeNetworkAcls",
      "ec2:AssociateRouteTable",
      "ec2:AuthorizeSecurityGroupIngress",
      "ec2:RevokeSecurityGroupEgress",
      "ec2:CreateSecurityGroup",
      "ec2:AuthorizeSecurityGroupEgress",
      "ec2:CreateVpcEndpoint",
      "ec2:ModifySubnetAttribute",
      "ec2:CreateSubnet",
      "ec2:CreateRoute",
      "ec2:CreateRouteTable",
      "ec2:CreateInternetGateway",
      "ec2:AttachInternetGateway",
      "ec2:ModifyVpcAttribute",
      "ec2:RevokeSecurityGroupIngress",
      "ec2:DescribeAvailabilityZones",
      "ec2:DescribeAccountAttributes",
      "ec2:AllocateAddress",
      "ec2:ReleaseAddress",
      "ec2:DescribeAddresses",
      "ec2:CreateNatGateway",
      "ec2:DeleteNatGateway",
      "ec2:DescribeNatGateways"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "ec2" {
  name        = "${aws_iam_user.cd.name}-ec2"
  description = "Allow user to manage EC2 resources."
  policy      = data.aws_iam_policy_document.ec2.json
}

resource "aws_iam_user_policy_attachment" "ec2" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.ec2.arn
}

# rds access policy for user 

data "aws_iam_policy_document" "rds" {
  statement {
    effect = "Allow"
    actions = [
      "rds:DescribeDBSubnetGroups",
      "rds:DescribeDBInstances",
      "rds:CreateDBSubnetGroup",
      "rds:DeleteDBSubnetGroup",
      "rds:CreateDBInstance",
      "rds:DeleteDBInstance",
      "rds:ListTagsForResource",
      "rds:ModifyDBInstance",
      "rds:AddTagsToResource"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "rds" {
  name        = "${aws_iam_user.cd.name}-rds"
  description = "Allow user to manage RDS resources."
  policy      = data.aws_iam_policy_document.rds.json
}

resource "aws_iam_user_policy_attachment" "rds" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.rds.arn
}

data "aws_iam_policy_document" "service_linked_rds" {
  statement {
    effect    = "Allow"
    actions   = ["iam:CreateServiceLinkedRole"]
    resources = ["arn:aws:iam::*:role/aws-service-role/rds.amazonaws.com/AWSServiceRoleForRDS"]

    condition {
      test     = "StringLike"
      variable = "iam:AWSServiceName"
      values   = ["rds.amazonaws.com"]
    }
  }
}

resource "aws_iam_policy" "service_linked_rds" {
  name        = "${aws_iam_user.cd.name}-service_linked_rds"
  description = "Allow Amazon RDS to call AWS services on behalf of your DB instances."
  policy      = data.aws_iam_policy_document.service_linked_rds.json
}

resource "aws_iam_user_policy_attachment" "service_linked_rds" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.service_linked_rds.arn
}


# policy for iam access

data "aws_iam_policy_document" "iam" {
  statement {
    effect = "Allow"
    actions = [
      "iam:ListInstanceProfilesForRole",
      "iam:ListAttachedRolePolicies",
      "iam:DeleteRole",
      "iam:ListPolicyVersions",
      "iam:DeletePolicy",
      "iam:DetachRolePolicy",
      "iam:ListRolePolicies",
      "iam:GetRole",
      "iam:GetPolicyVersion",
      "iam:GetPolicy",
      "iam:CreateRole",
      "iam:CreatePolicy",
      "iam:AttachRolePolicy",
      "iam:TagRole",
      "iam:TagPolicy",
      "iam:PassRole", "iam:CreateServiceLinkedRole",
      "iam:DeleteServiceLinkedRole",
      "iam:GetServiceLinkedRoleDeletionStatus"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "iam" {
  name        = "${aws_iam_user.cd.name}-iam"
  description = "Allow user to manage IAM resources."
  policy      = data.aws_iam_policy_document.iam.json
}

resource "aws_iam_user_policy_attachment" "iam" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.iam.arn
}

# policy for cloudwatch

data "aws_iam_policy_document" "logs" {
  statement {
    effect = "Allow"
    actions = [
      "logs:DeleteLogGroup",
      "logs:DescribeLogGroups",
      "logs:CreateLogGroup",
      "logs:TagResource",
      "logs:ListTagsLogGroup"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "logs" {
  name        = "${aws_iam_user.cd.name}-logs"
  description = "Allow user to manage CloudWatch resources."
  policy      = data.aws_iam_policy_document.logs.json
}

resource "aws_iam_user_policy_attachment" "logs" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.logs.arn
}

# policy for elastic load balancer

data "aws_iam_policy_document" "elb" {
  statement {
    effect = "Allow"
    actions = [
      "elasticloadbalancing:DeleteLoadBalancer",
      "elasticloadbalancing:DeleteTargetGroup",
      "elasticloadbalancing:DeleteListener",
      "elasticloadbalancing:DescribeListeners",
      "elasticloadbalancing:DescribeLoadBalancerAttributes",
      "elasticloadbalancing:DescribeTargetGroups",
      "elasticloadbalancing:DescribeTargetGroupAttributes",
      "elasticloadbalancing:DescribeLoadBalancers",
      "elasticloadbalancing:CreateListener",
      "elasticloadbalancing:SetSecurityGroups",
      "elasticloadbalancing:ModifyLoadBalancerAttributes",
      "elasticloadbalancing:CreateLoadBalancer",
      "elasticloadbalancing:ModifyTargetGroupAttributes",
      "elasticloadbalancing:CreateTargetGroup",
      "elasticloadbalancing:AddTags",
      "elasticloadbalancing:DescribeTags",
      "elasticloadbalancing:ModifyListener"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "elb" {
  name        = "${aws_iam_user.cd.name}-elb"
  description = "Allow user to manage ELB resources."
  policy      = data.aws_iam_policy_document.elb.json
}

resource "aws_iam_user_policy_attachment" "elb" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.elb.arn
}

# Policy for Route53 access 

data "aws_iam_policy_document" "route53" {
  statement {
    effect = "Allow"
    actions = [
      "route53:ListHostedZones",
      "route53:ListHostedZones",
      "route53:ChangeTagsForResource",
      "route53:GetHostedZone",
      "route53:ListTagsForResource",
      "route53:ChangeResourceRecordSets",
      "route53:GetChange",
      "route53:ListResourceRecordSets",
      "acm:RequestCertificate",
      "acm:AddTagsToCertificate",
      "acm:DescribeCertificate",
      "acm:ListTagsForCertificate",
      "acm:DeleteCertificate",
      "acm:CreateCertificate"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_policy" "route53" {
  name        = "${aws_iam_user.cd.name}-route53"
  description = "Allow user to manage Route53 resources."
  policy      = data.aws_iam_policy_document.route53.json
}

resource "aws_iam_user_policy_attachment" "route53" {
  user       = aws_iam_user.cd.name
  policy_arn = aws_iam_policy.route53.arn
}

