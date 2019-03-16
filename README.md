# go_aws_parameter_limits

# Results of limit testing yielded

AWS SSM Parameter Store limits are approximately in the 40 req/sec range.
Testing was done by executing parameter store GetParameter calls from multiple concurrent AWS Lambdas
