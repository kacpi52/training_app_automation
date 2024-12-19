CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS public.post(
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    "userId" UUID,
    "projectId" UUID,
    day integer,
    weight double precision,
    kcal integer,
    "createdUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updateUp" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT post_pkey PRIMARY KEY (id)
);


INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000001', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 1, 90.0, 2800, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000002', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 2, 89.4, 2750, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000003', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 3, 88.8, 2900, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000004', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 4, 88.2, 2650, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000005', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 5, 87.6, 2850, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000006', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 6, 87.0, 2600, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000007', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 7, 86.4, 2750, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000008', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 8, 85.8, 2900, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000009', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 9, 85.2, 2500, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-00000000000a', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 10, 84.6, 2650, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-00000000000b', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 11, 84.0, 2800, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-00000000000c', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 12, 83.4, 2900, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-00000000000d', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 13, 82.8, 2650, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-00000000000e', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 14, 82.2, 2750, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-00000000000f', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 15, 81.6, 2850, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000010', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 16, 81.0, 2600, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000011', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 17, 80.4, 2750, '2024-05-24 12:00:00', '2024-05-24 12:00:00');

INSERT INTO public.post (id, "userId", "projectId", day, weight, kcal, "createdUp", "updateUp")
VALUES ('c0a80100-0000-0000-0000-000000000012', '1738a3f7-5ab0-47fe-a0ea-18df6d282641', 'b3f8c1d1-4c8b-4e45-9b7e-f8c3b912d3e1', 18, 80.0, 2900, '2024-05-24 12:00:00', '2024-05-24 12:00:00');
