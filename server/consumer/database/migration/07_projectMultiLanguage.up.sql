CREATE TABLE IF NOT EXISTS public.project_multi_language(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "idProject" UUID,
    "idLanguage" UUID,
    title VARCHAR COLLATE pg_catalog."default",
    description VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT project_multi_language_pkey PRIMARY KEY (id)
);


INSERT INTO public.project_multi_language (id, "idProject", "idLanguage", title, description)
VALUES ('a1b2c3d4-e5f6-7890-abcd-ef1234567890', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 'e09bd685-aaf8-4d65-bcdd-aadca85670bc', 'Sample Project Title', 'This is a sample description for the project.');
