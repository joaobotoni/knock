create schema documento;
create schema pessoas;
create schema usuarios;
create schema vinculos;

create function public.tg_set_updated_at()
    returns trigger
    language plpgsql
as
$$
begin
    new.updated_at := now();
    return new;
end;
$$;

create function documento.somente_digitos(p_texto text)
    returns text
    language sql
    immutable strict parallel safe
as
$$
select regexp_replace(p_texto, '[^0-9]', '', 'g');
$$;

create function documento.somente_alfanumericos(p_texto text)
    returns text
    language sql
    immutable strict parallel safe
as
$$
select regexp_replace(upper(p_texto), '[^0-9A-Z]', '', 'g');
$$;

create function documento.dv_modulo11(p_base text, p_peso_maximo int)
    returns int
    language plpgsql
    immutable strict parallel safe
as
$$
declare
    v_tam   int := length(p_base);
    v_ciclo int := p_peso_maximo - 1;
    v_soma  int := 0;
    v_peso  int;
    v_valor int;
    v_resto int;
    i       int;
begin
    for i in 1..v_tam
        loop
            v_peso  := ((v_tam - i) % v_ciclo) + 2;
            v_valor := ascii(substr(p_base, i, 1)) - 48;
            v_soma  := v_soma + v_valor * v_peso;
        end loop;

    v_resto := v_soma % 11;

    if v_resto < 2 then
        return 0;
    else
        return 11 - v_resto;
    end if;
end;
$$;

create function documento.cpf_valido(p_cpf text)
    returns boolean
    language plpgsql
    immutable parallel safe
as
$$
begin
    if p_cpf is null or p_cpf !~ '^[0-9]{11}$' then
        return false;
    end if;

    if p_cpf = repeat(left(p_cpf, 1), 11) then
        return false;
    end if;

    if documento.dv_modulo11(substr(p_cpf, 1, 9), 11) <> ascii(substr(p_cpf, 10, 1)) - 48 then
        return false;
    end if;

    if documento.dv_modulo11(substr(p_cpf, 1, 10), 11) <> ascii(substr(p_cpf, 11, 1)) - 48 then
        return false;
    end if;

    return true;
end;
$$;

create function documento.cnpj_valido(p_cnpj text)
    returns boolean
    language plpgsql
    immutable parallel safe
as
$$
begin
    if p_cnpj is null or p_cnpj !~ '^[0-9A-Z]{12}[0-9]{2}$' then
        return false;
    end if;

    if p_cnpj = repeat(left(p_cnpj, 1), 14) then
        return false;
    end if;

    if documento.dv_modulo11(substr(p_cnpj, 1, 12), 9) <> ascii(substr(p_cnpj, 13, 1)) - 48 then
        return false;
    end if;

    if documento.dv_modulo11(substr(p_cnpj, 1, 13), 9) <> ascii(substr(p_cnpj, 14, 1)) - 48 then
        return false;
    end if;

    return true;
end;
$$;

create type pessoas.tipo_pessoa as enum ('pf', 'pj');

create table pessoas.pessoa
(
    id_pessoa  int generated always as identity,
    tipo       pessoas.tipo_pessoa not null,
    nome       text                not null,
    created_at timestamptz         not null default now(),
    updated_at timestamptz         not null default now(),
    constraint pk_pessoa primary key (id_pessoa),
    constraint uq_pessoa_id_tipo unique (id_pessoa, tipo),
    constraint ck_pessoa_nome check (nome = btrim(nome) and length(nome) between 1 and 200)
);

create table pessoas.pessoa_fisica
(
    id_pessoa  int primary key,
    tipo       pessoas.tipo_pessoa not null default 'pf',
    cpf        text                not null,
    created_at timestamptz         not null default now(),
    updated_at timestamptz         not null default now(),
    constraint uq_pessoa_fisica_cpf unique (cpf),
    constraint ck_pessoa_fisica_tipo check (tipo = 'pf'),
    constraint ck_pessoa_fisica_cpf check (documento.cpf_valido(cpf)),
    constraint fk_pessoa_fisica_pessoa foreign key (id_pessoa, tipo) references pessoas.pessoa (id_pessoa, tipo) on delete cascade
);

create table pessoas.pessoa_juridica
(
    id_pessoa     int primary key,
    tipo          pessoas.tipo_pessoa not null default 'pj',
    cnpj          text                not null,
    nome_fantasia text                not null,
    created_at    timestamptz         not null default now(),
    updated_at    timestamptz         not null default now(),
    constraint uq_pessoa_juridica_cnpj unique (cnpj),
    constraint ck_pessoa_juridica_tipo check (tipo = 'pj'),
    constraint ck_pessoa_juridica_cnpj check (documento.cnpj_valido(cnpj)),
    constraint ck_pessoa_juridica_fantasia check (nome_fantasia = btrim(nome_fantasia) and length(nome_fantasia) between 1 and 200),
    constraint fk_pessoa_juridica_pessoa foreign key (id_pessoa, tipo) references pessoas.pessoa (id_pessoa, tipo) on delete cascade
);

create function usuarios.email_valido(p_email text)
    returns boolean
    language sql
    immutable parallel safe
as
$$
select p_email is not null
   and p_email = lower(btrim(p_email))
   and length(p_email) between 6 and 254
   and p_email ~ '^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$';
$$;

create table usuarios.usuario
(
    id_usuario          int generated always as identity primary key,
    id_pessoa           int         not null unique,
    email               text        not null,
    email_verificado_em timestamptz,
    created_at          timestamptz not null default now(),
    updated_at          timestamptz not null default now(),
    constraint uq_usuario_email unique (email),
    constraint ck_usuario_email check (usuarios.email_valido(email)),
    constraint fk_usuario_pessoa foreign key (id_pessoa) references pessoas.pessoa_fisica (id_pessoa) on delete cascade
);

create type usuarios.provedor_social as enum ('google');

create table usuarios.credencial_local
(
    id_usuario          int primary key,
    senha_hash          text        not null,
    senha_atualizada_em timestamptz not null default now(),
    created_at          timestamptz not null default now(),
    updated_at          timestamptz not null default now(),
    constraint ck_credencial_local_hash check (senha_hash ~ '^\$(2[aby]|argon2(i|d|id))\$'),
    constraint fk_credencial_local_usuario foreign key (id_usuario) references usuarios.usuario (id_usuario) on delete cascade
);

create table usuarios.credencial_social
(
    id_usuario      int                      not null,
    provedor        usuarios.provedor_social not null,
    subject_externo text                     not null,
    created_at      timestamptz              not null default now(),
    updated_at      timestamptz              not null default now(),
    constraint pk_credencial_social primary key (id_usuario, provedor),
    constraint uq_credencial_social_identidade unique (provedor, subject_externo),
    constraint fk_credencial_social_usuario foreign key (id_usuario) references usuarios.usuario (id_usuario) on delete cascade
);

create type vinculos.tipo_vinculo as enum ('socio', 'representante');

create table vinculos.vinculo_pessoa_empresa
(
    id_vinculo_pessoa_empresa int generated always as identity primary key,
    id_pessoa_fisica          int                   not null,
    id_pessoa_juridica        int                   not null,
    tipo                      vinculos.tipo_vinculo not null,
    created_at                timestamptz           not null default now(),
    updated_at                timestamptz           not null default now(),
    constraint uq_vinculo unique (id_pessoa_fisica, id_pessoa_juridica, tipo),
    constraint fk_vinculo_pessoa_fisica foreign key (id_pessoa_fisica) references pessoas.pessoa_fisica (id_pessoa) on delete cascade,
    constraint fk_vinculo_pessoa_juridica foreign key (id_pessoa_juridica) references pessoas.pessoa_juridica (id_pessoa) on delete cascade
);


create trigger tg_pessoa_updated_at
    before update
    on pessoas.pessoa
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();

create trigger tg_pessoa_fisica_updated_at
    before update
    on pessoas.pessoa_fisica
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();

create trigger tg_pessoa_juridica_updated_at
    before update
    on pessoas.pessoa_juridica
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();

create trigger tg_usuario_updated_at
    before update
    on usuarios.usuario
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();

create trigger tg_credencial_local_updated_at
    before update
    on usuarios.credencial_local
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();

create trigger tg_credencial_social_updated_at
    before update
    on usuarios.credencial_social
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();

create trigger tg_vinculo_updated_at
    before update
    on vinculos.vinculo_pessoa_empresa
    for each row
    when (old.* is distinct from new.*)
execute function public.tg_set_updated_at();