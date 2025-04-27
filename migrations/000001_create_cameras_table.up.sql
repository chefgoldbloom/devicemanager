create table if not exists cameras (
    id bigserial primary key,
    created_at timestamp(0) with time zone not null default now(),
    ip_address inet not null,
    mac_address macaddr not null,
    model text not null,
    firmware text not null,
    site text not null,
    name text not null,
    version integer not null default 1
);
