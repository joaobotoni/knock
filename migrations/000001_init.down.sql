drop trigger if exists tg_credencial_social_updated_at on usuarios.credencial_social;
drop trigger if exists tg_credencial_local_updated_at on usuarios.credencial_local;
drop trigger if exists tg_usuario_updated_at on usuarios.usuario;
drop trigger if exists tg_pessoa_juridica_updated_at on pessoas.pessoa_juridica;
drop trigger if exists tg_pessoa_fisica_updated_at on pessoas.pessoa_fisica;
drop trigger if exists tg_pessoa_updated_at on pessoas.pessoa;

drop function if exists public.tg_set_updated_at();

drop table if exists vinculos.vinculo_pessoa_empresa;
drop type if exists vinculos.tipo_vinculo;
drop schema if exists vinculos;

drop table if exists usuarios.credencial_social;
drop table if exists usuarios.credencial_local;
drop type if exists usuarios.provedor_social;
drop table if exists usuarios.usuario;
drop function if exists usuarios.email_valido(text);
drop schema if exists usuarios;

drop table if exists pessoas.pessoa_juridica;
drop table if exists pessoas.pessoa_fisica;
drop table if exists pessoas.pessoa;
drop type if exists pessoas.tipo_pessoa;
drop schema if exists pessoas;

drop function if exists documento.cnpj_valido(text);
drop function if exists documento.cpf_valido(text);
drop function if exists documento.dv_modulo11(text, int);
drop function if exists documento.somente_alfanumericos(text);
drop function if exists documento.somente_digitos(text);
drop schema if exists documento;
