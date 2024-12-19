CREATE TABLE IF NOT EXISTS public.project(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT project_pkey PRIMARY KEY (id)
);


INSERT INTO public.project (id, "userId", "createdUp", "updateUp")
VALUES ('b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', '2024-05-24 12:00:00', '2024-05-24 12:00:00');

