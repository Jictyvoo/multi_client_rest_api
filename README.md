# Multi-Client REST

This is a test server that works to enable the access to multiple clients using a single endpoint. So this can work as
an API gateway.

## Flow

- The API will receive a JSON via POST containing the name and cell phone;
- The customer must be authenticated to enter the contact in the base
- The contact must be entered in the customer's database following the rules of each customer

## REST-Api specifications

- Authentication will be through a JWT token in the Authorization Header
- Each customer has a unique key

## Client specifications

Here are described the clients rules and databases

1. **Client ABZ_1**
    - Mysql database
    - Name format is uppercase only
    - Phone format follows standard `+55 (41) 93030-6905`
2. **Client XYC_2**
    - Postgresql database
    - Name format is free
    - Phone format follows standard `554130306905`
