
CREATE TYPE STATUS AS ENUM ('PENDING', 'CANCELLED', 'DONE');

CREATE TABLE "user"(
   id serial PRIMARY KEY,
   name VARCHAR (150) NOT NULL
);


CREATE TABLE todo(
    id serial PRIMARY KEY,
    description VARCHAR (255) NOT NULL,
    status STATUS NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);


INSERT INTO public.user ("name") VALUES ('tester');

INSERT INTO public.todo ("description", "status", "user_id") VALUES ('First todo', false, 1);

