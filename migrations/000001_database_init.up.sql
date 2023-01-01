CREATE TABLE templates(
    id int NOT NULL AUTO_INCREMENT,
    active tinyint,
    code varchar(255),
    name varchar(255),
    body longtext,
    title text,
    PRIMARY KEY (id),
    UNIQUE (code)
);
