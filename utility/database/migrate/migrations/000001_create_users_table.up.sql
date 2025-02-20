CREATE TABLE users
(
    user_id bigserial NOT NULL,
    first_name character varying NOT NULL,
    last_name character varying NOT NULL,
    email_id character varying(255) NOT NULL,
    password character varying(200) NOT NULL,
    created_at timestamp without time zone DEFAULT NOW(),
    updated_at timestamp without time zone DEFAULT NOW(),
    PRIMARY KEY (user_id)
);
