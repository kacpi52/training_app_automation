CREATE TABLE IF NOT EXISTS public.type_training (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    name VARCHAR COLLATE pg_catalog."default",
   "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT type_training_pkey PRIMARY KEY (id)
);