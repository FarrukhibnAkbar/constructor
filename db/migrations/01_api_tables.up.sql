create table api(
    id uuid primary key not null,
    state numeric(2) not null default 1,
    name varchar(100) not null,
    method varchar(10) not null,
    description varchar(1000),
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now(),
    unique(name, method)
);

create table api_detail(
    id uuid primary key not null,
    api_id uuid not null references api(id) on delete cascade on update cascade,
    state numeric(2) not null default 1,
    name varchar(100) not null,
    query_type varchar(64) not null,
    query text not null,
    description varchar(1000),
    created_at timestamp with time zone not null default now(),
    updated_at timestamp with time zone not null default now(),
    unique(api_id, name)
);