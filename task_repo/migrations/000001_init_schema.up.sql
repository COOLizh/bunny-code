CREATE TABLE "users"(
                        "id" SERIAL PRIMARY KEY,
                        "login" VARCHAR(128) UNIQUE NOT NULL,
                        "password" VARCHAR(64) NOT NULL
);


CREATE TABLE "languages"(
                            "id" SERIAL PRIMARY KEY,
                            "name" VARCHAR UNIQUE NOT NULL
);

CREATE TABLE "tasks"(
                        "id" SERIAL PRIMARY KEY,
                        "name" VARCHAR NOT NULL,
                        "description" VARCHAR NOT NULL,
                        "time_limit" INTEGER NOT NULL,
                        "memory" INTEGER NOT NULL
);

CREATE TABLE "test_cases"(
                             "id" SERIAL PRIMARY KEY,
                             "task_id" SERIAL,
                             "test_data" TEXT NOT NULL,
                             "answer" TEXT NOT NULL,
                             FOREIGN KEY ("task_id") REFERENCES "tasks"("id") ON DELETE CASCADE
);

CREATE TABLE "history_user"(
                               "uuid" UUID PRIMARY KEY,
                               "user_id" INTEGER,
                               "task_id" INTEGER,
                               "language_id" INTEGER,
                               "solution" TEXT NOT NULL,
                               "result" JSONB,
                               "timestamp" TIMESTAMP ,
                               FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
                               FOREIGN KEY ("task_id") REFERENCES "tasks"("id") ON DELETE CASCADE,
                               FOREIGN KEY ("language_id") REFERENCES "languages"("id") ON DELETE CASCADE
);
