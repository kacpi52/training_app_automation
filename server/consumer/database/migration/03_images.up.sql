CREATE TABLE IF NOT EXISTS public.images (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "projectId" UUID,
    name VARCHAR COLLATE pg_catalog."default",
    folder VARCHAR COLLATE pg_catalog."default",
    "folderPath" VARCHAR COLLATE pg_catalog."default",
    path VARCHAR COLLATE pg_catalog."default",
    url VARCHAR COLLATE pg_catalog."default",
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT images_pkey PRIMARY KEY (id)
);