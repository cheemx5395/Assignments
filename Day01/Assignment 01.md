Real-Time Applications and their Databases and why they use it!

1. Whatsapp:

- SQL Databases:
Postgres: For user profiles, settings and contact information
Why ?: cause the data is strucutured and retrieval becomes easier. 

- NoSQL Databases:
Cassandra: For storing media and backups and large datasets
Why ?: easy to scale horizontally and fault tolerance is great as server crashings happen it stores replicas to avoid unavailability of data.

RocksDB: storing message state management
Why ?: offers horizontal scalability mostly.

2. Uber:

- SQL Databases:
MySQL: For user profiles and ride information mostly
Why ?: to have the possibility of using joins since we're storing relational data

- NoSQL Databases:
Apache Cassandra: storing real-time events and trip information
Why ?: high-availability, low-latency needs

Redis: Essential for caching(popular routes and low-latency needs)
Why ?: Since low-latency is expected Redis being key-value based database hence faster access is available.

3. Tinder: 

- SQL Databases:
Postgres: Potential Source-of-truth database
Why ?: Because its SQL database it follows ACID properties so it can be considered as source of truth since it will have to follow Consistency rule which will oblige it to be actually trustable

- NoSQL Databases:
MongoDB: Core data storage
Why ?: Easier to scale, can deploy cluster in different region

Redis: For caching frequently accessed data, improving speed and reducing load on primary databases.
Why ?:  As stated above.

ElasticSearch: Utilized for caching frequently accessed data
Why ?: helps with recommendation features and matching users based on their likeness.