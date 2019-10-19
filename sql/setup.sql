DROP TABLE IF EXISTS views;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS donations;
DROP TABLE IF EXISTS verification;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS charity_posts;
DROP TABLE IF EXISTS charity_users;
DROP TABLE IF EXISTS charities;

CREATE TABLE users
(
    id            bigserial primary key,
    email         varchar(128) unique   not null,
    password_hash varchar(64)           not null, -- hash size is fixed,
    name          varchar(48)           not null,
    phone_number  varchar(13) unique    not null,
    verified      boolean default false not null
);

CREATE TABLE verification ( -- verification table for email
    user_id bigserial references users(id),
    email_verification_code varchar(12),
    phone_verification_code varchar(8)
);

CREATE TABLE charities (
    id bigserial unique not null,
    short_name varchar(24) not null, -- duplicates are ok here
    long_name varchar(92) unique not null, -- not here
    description varchar(180), -- make not null?
    total_donated bigint default 0,
    queued_donations bigint default 0,
    ceo varchar(36),
    profile_url varchar(128),
    password_hash varchar(64) not null -- hash size is fixed, password to admin charity screen
);

CREATE TABLE charity_users (
    id bigserial unique not null,
    charity_id bigserial unique not null references charities(id),
    display_name varchar(64) not null,
    bio varchar(240),
    password_hash varchar(64) not null -- hash size is fixed,
);

CREATE TABLE charity_posts (
    id bigserial unique not null,
    charity_id bigserial not null references charities(id),
    author_id bigserial not null references charity_users(id),
    title varchar(96) not null,
    content text not null,
    thumbnail varchar(464),
    post_time timestamp default current_timestamp,
    last_edit timestamp default current_timestamp
);

CREATE TABLE donations (
    donation_id bigserial,
    user_id bigserial not null references users(id),
    charity_id bigserial not null references charities(id),
    post_id bigserial references charity_posts(id),
    amount money,
    donation_created timestamp,
    donation_executed timestamp, -- if before now then it will be executed
    flushed boolean
);

CREATE TABLE views (
    user_id bigserial references users(id),
    charity_post_id bigserial references charity_posts(id),
    viewed_at timestamp default current_timestamp
);

CREATE TABLE followers (
    user_id bigserial references users(id),
    charity_id bigserial references charities(id)
);

SELECT * FROM users;