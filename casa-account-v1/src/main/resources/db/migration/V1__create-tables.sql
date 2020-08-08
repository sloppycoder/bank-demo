create table casa_account
(
    account_id varchar(255) not null,
    nick_name  varchar(255) not null,
    prod_code  varchar(255) not null,
    prod_name  varchar(255) not null,
    currency   varchar(3)   not null,
    status     int          not null,
    balance    float        not null,
    constraint account_id   unique (account_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;
