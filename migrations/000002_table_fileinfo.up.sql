CREATE TABLE fileinfo(
    id int NOT NULL AUTO_INCREMENT,
    filename varchar(255),
    mime varchar(255),
    path varchar(255),
    PRIMARY KEY (id),
    UNIQUE (path)
);
