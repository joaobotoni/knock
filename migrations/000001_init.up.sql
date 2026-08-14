create function public.updated_at()
    returns trigger
    language plpgsql
as
$$
begin
    new.updated_at := now();
    return new;
end;
$$;

create schema documento;

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

create function documento.cpf_normalizado(p_cpf text)
    returns text
    language sql
    immutable strict parallel safe
as
$$
select case
           when valor ~ '^([0-9])\1{10}$' then null
           when valor ~ '^[0-9]{11}$' then valor
           end
from (select documento.somente_digitos(p_cpf) as valor) t;
$$;

create function documento.cnpj_normalizado(p_cnpj text)
    returns text
    language sql
    immutable strict parallel safe
as
$$
select case
           when valor ~ '^(.)\1{13}$' then null
           when valor ~ '^[0-9A-Z]{12}[0-9]{2}$' then valor
           end
from (select documento.somente_alfanumericos(p_cnpj) as valor) t;
$$;

create function documento.dv_modulo11(p_base text, p_pesos integer[])
    returns integer
    language plpgsql
    immutable strict parallel safe
as
$$
declare
    v_tam   int := array_length(p_pesos, 1);
    v_sum   int := 0;
    v_resto int;
    i       int;
begin

    if length(p_base) < v_tam then
        return null;
    end if;

    for i in 1..v_tam
        loop
            v_sum := v_sum + ((ascii(substr(p_base, i, 1)) - 48) * p_pesos[i]);
        end loop;

    v_resto := v_sum % 11;

    if v_resto < 2 then
        return 0;
    else
        return 11 - v_resto;
    end if;
end;
$$;

create function documento.cpf_valido(p_cpf text)
    returns boolean
    language sql
    immutable parallel safe
as
$$
select coalesce(
               documento.dv_modulo11(cpf, '{10,9,8,7,6,5,4,3,2}') = substr(cpf, 10, 1)::integer
                   and documento.dv_modulo11(cpf, '{11,10,9,8,7,6,5,4,3,2}') = substr(cpf, 11, 1)::integer,
               false
       )
from (select documento.cpf_normalizado(p_cpf) as cpf) t;
$$;

create function documento.cnpj_valido(p_cnpj text)
    returns boolean
    language sql
    immutable parallel safe
as
$$
select coalesce(
               documento.dv_modulo11(cnpj, '{5,4,3,2,9,8,7,6,5,4,3,2}') = substr(cnpj, 13, 1)::integer
                   and documento.dv_modulo11(cnpj, '{6,5,4,3,2,9,8,7,6,5,4,3,2}') = substr(cnpj, 14, 1)::integer,
               false
       )
from (select documento.cnpj_normalizado(p_cnpj) as cnpj) t;
$$;

create schema pessoas;

create type pessoas.tipo_pessoa as enum ('pf', 'pj');

create table pessoas.pessoa
(
    id_pessoa  int generated always as identity,
    tipo       pessoas.tipo_pessoa not null,
    nome       text                not null,
    created_at timestamptz         not null default now(),
    updated_at timestamptz         not null default now(),
    constraint pk_pessoa primary key (id_pessoa, tipo),
    constraint ck_pessoa_nome check (nome = btrim(nome) and length(nome) > 0)
);

create table pessoas.pessoa_fisica
(
    id_pessoa  int primary key,
    tipo       pessoas.tipo_pessoa not null generated always as ('pf') stored,
    cpf        text                not null,
    created_at timestamptz         not null default now(),
    updated_at timestamptz         not null default now(),
    constraint uq_pessoa_fisica_cpf unique (cpf),
    constraint ck_pessoa_fisica_cpf check (cpf ~ '^[0-9]{11}$' AND documento.cpf_valido(cpf)),
    constraint fk_pessoa_fisica_pessoa foreign key (id_pessoa, tipo) references pessoas.pessoa (id_pessoa, tipo) on delete cascade
);

create table pessoas.pessoa_juridica
(
    id_pessoa     int primary key,
    tipo          pessoas.tipo_pessoa not null generated always as ('pj') stored,
    cnpj          text                not null,
    nome_fantasia text,
    created_at    timestamptz         not null default now(),
    updated_at    timestamptz         not null default now(),
    constraint uq_pessoa_juridica_cnpj unique (cnpj),
    constraint ck_pessoa_juridica_cnpj check (cnpj ~ '^[0-9]{14}$' AND documento.cnpj_valido(cnpj)),
    constraint fk_pessoa_juridica_pessoa foreign key (id_pessoa, tipo) references pessoas.pessoa (id_pessoa, tipo) on delete cascade
);


create schema usuarios;

create function usuarios.email_normalizado(p_email text)
    returns text
    language sql
    immutable parallel safe
as
$$
select case
           when valor ~ '^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$' then valor
           end
from (select lower(btrim(p_email)) as valor) t;
$$;

create type usuarios.provedor as enum ('local', 'google');

create table usuarios.usuario
(
    id_usuario          int generated always as identity primary key,
    id_pessoa           int         not null unique,
    email               text        not null,
    email_verificado_em timestamptz,
    created_at          timestamptz not null default now(),
    updated_at          timestamptz not null default now(),
    constraint uq_usuario_email unique (email),
    constraint ck_usuario_email check (email = usuarios.email_normalizado(email)),
    constraint fk_usuario_pessoa foreign key (id_pessoa) references pessoas.pessoa_fisica (id_pessoa) on delete cascade
);

create table usuarios.credencial
(
    id_credencial int generated always as identity primary key,
    id_usuario    int               not null,
    provedor      usuarios.provedor not null default 'local',
    segredo       text              not null,
    created_at    timestamptz       not null default now(),
    updated_at    timestamptz       not null default now(),
    constraint uq_credencial_usuario_provedor unique (id_usuario, provedor),
    constraint fk_credencial_usuario foreign key (id_usuario) references usuarios.usuario (id_usuario) on delete cascade
);

create schema vinculos;

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
execute function public.updated_at();

create trigger tg_pessoa_fisica_updated_at
    before update
    on pessoas.pessoa_fisica
    for each row
execute function public.updated_at();

create trigger tg_pessoa_juridica_updated_at
    before update
    on pessoas.pessoa_juridica
    for each row
execute function public.updated_at();

create trigger tg_usuario_updated_at
    before update
    on usuarios.usuario
    for each row
execute function public.updated_at();

create trigger tg_credencial_updated_at
    before update
    on usuarios.credencial
    for each row
execute function public.updated_at();

create trigger tg_vinculo_updated_at
    before update
    on vinculos.vinculo_pessoa_empresa
    for each row
execute function public.updated_at();