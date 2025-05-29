CREATE TABLE frogs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    species VARCHAR(255) NOT NULL,
    habitat VARCHAR(255),
    age INT CHECK (age >= 0)
);

CREATE OR REPLACE PROCEDURE delete_frog_by_id(frog_id INT)
LANGUAGE plpgsql
AS $$
BEGIN
  DELETE FROM frogs WHERE id = frog_id;

  IF NOT FOUND THEN
    RAISE NOTICE 'Жаба с id % не найдена.', frog_id;
  ELSE
    RAISE NOTICE 'Жаба с id % удалена.', frog_id;
  END IF;
END;
$$;
