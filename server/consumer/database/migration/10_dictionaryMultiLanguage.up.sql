CREATE TABLE IF NOT EXISTS public.dictionary_multi_language(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "dictionaryId" UUID,
    key VARCHAR COLLATE pg_catalog."default",
    translation VARCHAR COLLATE pg_catalog."default",
    CONSTRAINT dictionary_multi_language_pkey PRIMARY KEY (id)
);


INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('d2aab915-7a45-4925-8be4-74d3c6a02c69', 'pl', 'e09bd685-aaf8-4d65-bcdd-aadca85670bc', 'Polski');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('eca11514-9fd9-4ac4-9ef9-4762d4e5140b', 'en', 'e09bd685-aaf8-4d65-bcdd-aadca85670bc', 'Angielski');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('ff230db4-91ee-43c8-b31b-337ffccb4983', 'ger', 'e09bd685-aaf8-4d65-bcdd-aadca85670bc', 'Niemiecki');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('9508f0d1-151f-4f1b-a94b-265b32bba1fc', 'pl', 'df88c32f-f71d-41bc-84ac-7cc36e37305f', 'Polish');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('5b172965-8093-4800-a992-c461f0b81d1e', 'en', 'df88c32f-f71d-41bc-84ac-7cc36e37305f', 'English');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('900ac75b-209a-43f4-befa-7686e54e9c4a', 'ger', 'df88c32f-f71d-41bc-84ac-7cc36e37305f', 'German');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('2f6ed612-d2d1-44a3-88a6-922d85037f0e', 'pl', 'f63ce14b-af7a-4fc9-abc5-c68bc3254e2c', 'Polieren');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('9a59825a-897a-4b1a-83b6-d8726c031a65', 'en', 'f63ce14b-af7a-4fc9-abc5-c68bc3254e2c', 'Englisch');

INSERT INTO public.dictionary_multi_language (id, key, "dictionaryId", translation) VALUES
('ead37394-eae2-40de-b8a7-9afa429e7b90', 'ger', 'f63ce14b-af7a-4fc9-abc5-c68bc3254e2c', 'Deutsch');