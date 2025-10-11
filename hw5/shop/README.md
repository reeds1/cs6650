# My Shop API

This project is a Go-based REST API for managing products. 
It is deployed on AWS ECS Fargate using Docker containers and Terraform for infrastructure as code.

## Deployment

### Prerequisites
- Docker
- Terraform
- AWS CLI
- AWS account with permissions for ECS, ECR, VPC, etc.

### Steps
1. Configure AWS credentials
```bash
aws configure

2. clone repo
git clone <your_repo_url>
cd <repo_folder>

3. Initialize Terraform
cd terraform
terraform init
terraform plan
terraform apply

4. Build and push Docker image
cd ../src
docker buildx build --platform linux/amd64 -t <account_id>.dkr.ecr.<region>.amazonaws.com/<repo>:latest --push .
(create new task)
aws ecs update-service \
  --cluster CS6650L2-cluster \
  --service CS6650L2 \
  --force-new-deployment \
  --region us-west-2

5. Test API
get products
curl http://<public_ip>:8080/products

post product
curl -X POST http://54.191.29.157:8080/products \                                                                     
  -H "Content-Type: application/json" \
  -d '{
        "name": "Tablet",
        "price": 499.99,
        "stock": 7
      }'
