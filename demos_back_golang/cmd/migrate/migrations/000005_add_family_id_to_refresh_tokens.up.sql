ALTER TABLE refresh_tokens ADD COLUMN family_id UUID DEFAULT gen_random_uuid();

ALTER TABLE refresh_tokens ADD COLUMN used BOOLEAN DEFAULT FALSE;

CREATE INDEX idx_refresh_tokens_family_id ON refresh_tokens(family_id);

UPDATE refresh_tokens SET family_id = gen_random_uuid() WHERE family_id IS NULL;