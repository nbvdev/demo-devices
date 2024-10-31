
CREATE TABLE device (
    id int not null auto_increment PRIMARY KEY,
    name varchar(255) not null,
    brand varchar(255) not null,
    created datetime not null default current_timestamp,
    CONSTRAINT ct_name UNIQUE INDEX u_name (name),
    INDEX i_brand (brand)
);