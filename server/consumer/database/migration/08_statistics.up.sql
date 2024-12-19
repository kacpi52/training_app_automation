CREATE TABLE IF NOT EXISTS public.statistics(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "projectId" UUID,
    day VARCHAR COLLATE pg_catalog."default",
    "endWeight" double precision,
    "downWeight" double precision,
    "sumKg" double precision,
    "avgKg" double precision,
    "sumKcal" integer,
    "typeTraining" VARCHAR COLLATE pg_catalog."default",
    "sumTime" TIME,
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT statistics_pkey PRIMARY KEY (id)
);