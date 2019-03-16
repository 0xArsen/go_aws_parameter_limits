# go_aws_parameter_limits

# Results of limit testing yielded

AWS SSM Parameter Store limits are approximately in the 40 req/sec range.
Testing was done by executing parameter store GetParameter calls from multiple concurrent AWS Lambdas


AWS Secrets Manager Limits are much higher. Testing was done as follows:

Exp #1: 2000 Invocations within 95 seconds, every invocation makes 3 Secrets manager requests(GetSecretValue)
6,000 req / 95 sec = 63 req/sec

Exp #2: 2000 Invocations within 98 seconds, every invocation makes 10 Secrets Manager Requests
20,000 req / 98 sec = 204 req/sec

Exp #3: 2000 Invocations within 101 seconds, every invocation makes 60 Secrets Manager Requests
120,000 req / 101 sec = 1188 req/sec
Throttling Rate was exceeded for this experiment.

Exp #4: 2000 Invocations within 99 seconds, every invocation makes 50 Secrets Manager Requests
100,000 req / 99 sec = 1010 req/sec
Throttling Rate was NOT exceeded for this experiment.
There are deeper timing issues that may skew the results.
