# Products API
Come with pagination and simple CRUD Products made with Golang and Postgres


## Script For Database 


```sql
CREATE TABLE IF NOT EXISTS public.products
(
    id integer NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    product_code character varying(20) COLLATE pg_catalog."default",
    name character varying(255) COLLATE pg_catalog."default",
    subcategory character varying(25) COLLATE pg_catalog."default",
    brand character varying(25) COLLATE pg_catalog."default",
    price numeric(12,2) DEFAULT 0.00,
    status boolean DEFAULT true,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    sub_category text COLLATE pg_catalog."default",
    CONSTRAINT products_pkey PRIMARY KEY (id)
)
```


## Script For Populate Data

```sql
INSERT INTO products (product_code, name, subcategory, brand, price)
SELECT substr(md5(random()::text), 1, 10) AS product_code, md5(random()::text) AS name, substr(md5(random()::text), 1, 10) AS subcategory, substr(md5(random()::text), 1, 10) AS brand, FLOOR(RANDOM()*(1000000)::float) AS price FROM generate_series(1,200000)
```


## API Contract

#### /product
* `GET` : Get all products with default limit is 8.

#### /product?limit=?&page=?
* `GET` : Get all products with pagination and limit and page from client request. Replace (?) with the desired number 

#### /product
* `POST` : Create new product.

#### /project/{id}
* `GET` : Get a detail product
* `POST` : Update a product
* `DELETE` : Delete a product