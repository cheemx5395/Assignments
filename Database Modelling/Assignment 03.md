Why is normalization important? What problems occur if a database is not normalized?

Basically when we write data on the database we're using up space and when similar data is stored multiple times which causes redundancy which can lead to slow retrieval of data from database, that's why its necessary to keep our database as fastly accessible as possible, which can be achieved by Normalization.

Problems that can occur if a database is not normalized are as in slow data retrieval, more space used up by similar data and redundant data causing developer to use more storage leading to more costs of database servers.

1NF for the problem table:

| order_id | customer_name | product |
| -------: | ------------- | ------- |
|      101 | Rahul         | Laptop  |
|      101 | Rahul         | Mouse   |
