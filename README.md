# payment-gateway-thuang86714
## Content Table
### [1. How to run your solution?](#How-to-run-your-solution?)
### [2. Any assumption you made](#Any-assumptions-you-made?)
### [3. Areas for improvenment](#Areas-for-improvenment)
### [4. What cloud technologies you'd use and why?](#What-cloud-technologies-you'd-use-and-why?)
### [5. Overall Architecture for future](#Overall-Architecture-for-future)
### [6. Links to other documents](#Links-to-other-documents)
##  1. How to run your solution?
  Whichever approach you choose, clone this repository first.<br>
  a. Docker && Docker-Compose<br>
      This approach assumes you already have docker installed. 
    <ol>
          <li>Use your terminal, go to root level directory(payment-gateway-thuang86714), enter
             ```
             docker-compose up -d
             ```
          <li>There should be 5 services + 1 network up and running</li>
          <li>In the same terminal, enter             
          ```
             docker attach merchant-service
         ```</li>
          <li>Press Enter, and then welcome message and instructions should be printed out</li>
          <li>Follow the instructions. 2 things to note: <br>
                --There are input check for every field you enter. For example, for expiration date, it will take now to now + 5yr,<br> eg. now = Aug/24, it will accepts any date till Aug/29. Any date before or after this time period is invalid.<br>
                --For card number, if you put "1234567812345678", you can test connections between merchant and gateway. You will receive a dummy response, and no fund will be moved in the bank
                </li>
          </ol>
        b. Good old            ```
             go run main.go
         ```<br>
         <ol>
          <li>This approach will require you to spin up 2 postgreSQL instance yourself. Docker is suggested: <br>
          for gateway_db:<br>
         ```
         docker run --name my-post-g -e POSTGRES_USER=Tommy -e POSTGRES_PASSWORD=test123 -e POSTGRES_DB=gateway_dev -p 5432:5432 -d postgres:13
         ```<br>
          for bank_db:<br>
         ```
         docker run --name my-postgre -e POSTGRES_USER=Tommy -e POSTGRES_PASSWORD=test123 -e POSTGRES_DB=gateway_dev -p 5433:5432 -d postgres:13
         ```<br>
          Please refer to bank/.env and gateway/.env for more details about environment variables.<br></li>
          <li>Then open 3 terminal. cd into {service}/cmd(eg. gateway/cmd), then do ```
             go run main.go
         ```<br>
         Turn to merchant's terminal, follow the instructions.<br></li>
        c. How to test? <br>
        At root level, run ```
             go test ./... -v -cover
         ```<br>
         or you could run every CI pipelines if you have act installed<br>
         </ol>
##  2. Any assumptions you made?
  <ol>
        <li>Customers always have sufficient funds for each transaction: transactions won't fail because of lack of fund</li>
        <li>It's always Merchant Initialized Transaction(MIT): gateway doesn't turn to customer for card authorization</li>
        <li>This gateway serves only one merchant: no authentication nor authorization mechnaism</li>
  </ol>
  
##  3. Areas for improvenment
<ol>
  <li>Dependency Injection: current implementation of repository brings several problem:<br>
    <ol>
      <li>Testing can be more challenging because you can't easily inject a mock database for unit tests.
      <li>It creates a tight coupling between your repository and the global DB instance.
      <li>It's less flexible if you ever need to use different DB connections for different operations.
    </ol>
  </li>
  <li>Authentication and Authorization for both merchant and customers</li>
  <li>Platform Implementation: there are two functions in gateway/service decidePSP() and decideServiceFeePercentage() which simulates what ProcessOut is doing. </li>
  <li>How Bank process transaction: In bank/service, there's a functions validateTransaction() simulates the time it takes to validate a transaction</li>
  <li>Deployment: this repos doesn't include any deployment onto cloud</li>
  
</ol>

4. What cloud technologies you'd use and why?
<ol>
  <li>AWS EKS for merchant/gateway/bank:<br> Amazon Elastic Kubernetes Service (EKS) would be used to deploy and manage the merchant, gateway, and bank services. EKS provides a managed Kubernetes platform, allowing for easy orchestration and scaling of containerized applications. It offers high availability, security, and seamless integration with other AWS services. EKS is ideal for these microservices as it enables efficient resource utilization, easy updates, and robust monitoring capabilities.</li>
  <li>AWS RDS with read replica:<br> Amazon Relational Database Service (RDS) with PostgreSQL engine would be employed for the primary database, with a read replica for improved read performance and high availability. RDS offers automated backups, patching, and scaling, reducing administrative overhead. The read replica helps distribute read traffic, improving overall application performance and providing a failover option in case of primary database issues.</li>
  <li>Nginx:<br> Nginx would be deployed as an Ingress Controller within the EKS cluster, handling incoming traffic, load balancing, and SSL termination for the services. It provides efficient request routing and can be easily configured for various traffic management scenarios. </li>
</ol>

##  5. Overall Architecture for future
![Architecture](https://github.com/user-attachments/assets/ac6560de-a41a-4f43-b41b-c7b43b456c76)

##  6. Links to other documents
####  [Shared]()
####  [merchant]()
####  [gateway]()
####  [bank]()
