resource "aws_lambda_function" "bin_schedule" {
  function_name    = "bin-schedule-api"
  role             = aws_iam_role.bin_schedule_role.arn
  handler          = "main"
  filename         = data.archive_file.zip.output_path
  runtime          = "go1.x"
  source_code_hash = data.archive_file.zip.output_sha
  timeout          = 2
  #  reserved_concurrent_executions = 1
  environment {
    variables = {
      APP_CONTEXT = "lambda"
      GIN_MODE    = "release"
    }
  }
}

resource "aws_lambda_function_url" "bin_schedule_url" {
  authorization_type = "NONE"
  function_name      = aws_lambda_function.bin_schedule.function_name
}

resource "aws_cloudwatch_log_group" "bin_schedule" {
  name              = aws_lambda_function.bin_schedule.function_name
  retention_in_days = 2
}

data "archive_file" "zip" {
  source_file      = "${path.module}/../backend/main"
  output_path      = "${path.module}/bin-schedule.zip"
  type             = "zip"
  output_file_mode = "0666"
}

resource "aws_iam_role" "bin_schedule_role" {
  name                = "bin-schedule-lambda-role"
  assume_role_policy  = data.aws_iam_policy_document.bin_schedule_trust_policy.json
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]
}

data "aws_iam_policy_document" "bin_schedule_trust_policy" {
  statement {
    principals {
      identifiers = ["lambda.amazonaws.com"]
      type        = "Service"
    }
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
  }
}

output "api_url" {
  value = aws_lambda_function_url.bin_schedule_url.function_url
}