-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.1.2
-- PostgreSQL version: 16.0
-- Project Site: pgmodeler.io
-- Model Author: Cristian David Espitia Vanegas

-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.

-- object: agenda_personal | type: SCHEMA --
-- DROP SCHEMA IF EXISTS agenda_personal CASCADE;
CREATE SCHEMA agenda_personal;
-- ddl-end --
ALTER SCHEMA agenda_personal OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,agenda_personal;
-- ddl-end --

-- object: agenda_personal.agenda_personal | type: TABLE --
-- DROP TABLE IF EXISTS agenda_personal.agenda_personal CASCADE;
CREATE TABLE agenda_personal.agenda (
	id serial NOT NULL,
	nombre character varying(50),
	descripcion character varying(250),
	numero_documento character varying(20),
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	CONSTRAINT pk_agenda PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda_personal.agenda OWNER TO postgres;
-- ddl-end --

-- object: agenda_personal.contacto | type: TABLE --
-- DROP TABLE IF EXISTS agenda_personal.contacto CASCADE;
CREATE TABLE agenda_personal.contacto (
	id serial NOT NULL,
	nombre character varying(50),
	numero_documento character varying(20),
	direccion character varying(100),
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	agenda_id integer,
	CONSTRAINT pk_contacto PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda_personal.contacto OWNER TO postgres;
-- ddl-end --

-- object: agenda_personal.numero_telefono | type: TABLE --
-- DROP TABLE IF EXISTS agenda_personal.numero_telefono CASCADE;
CREATE TABLE agenda_personal.numero_telefono (
	id serial NOT NULL,
	telefono character varying(20),
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	contacto_id integer,
	CONSTRAINT pk_numero_telefono PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda_personal.numero_telefono OWNER TO postgres;
-- ddl-end --

-- object: agenda_personal.correo_electronico | type: TABLE --
-- DROP TABLE IF EXISTS agenda_personal.correo_electronico CASCADE;
CREATE TABLE agenda_personal.correo_electronico (
	id serial NOT NULL,
	correo character varying(50),
	activo boolean DEFAULT true,
	fecha_creacion timestamp,
	fecha_modificacion timestamp,
	contacto_id integer,
	CONSTRAINT pk_correo_electronico PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda_personal.correo_electronico OWNER TO postgres;
-- ddl-end --

-- object: agenda_personal_fk | type: CONSTRAINT --
-- ALTER TABLE agenda_personal.contacto DROP CONSTRAINT IF EXISTS agenda_personal_fk CASCADE;
ALTER TABLE agenda_personal.contacto ADD CONSTRAINT fk_contacto_agenda FOREIGN KEY (agenda_id)
REFERENCES agenda_personal.agenda (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE CASCADE;
-- ddl-end --

-- object: contacto_fk | type: CONSTRAINT --
-- ALTER TABLE agenda_personal.numero_telefono DROP CONSTRAINT IF EXISTS contacto_fk CASCADE;
ALTER TABLE agenda_personal.numero_telefono ADD CONSTRAINT fk_numero_telefono_contacto FOREIGN KEY (contacto_id)
REFERENCES agenda_personal.contacto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE CASCADE;
-- ddl-end --

-- object: contacto_fk | type: CONSTRAINT --
-- ALTER TABLE agenda_personal.correo_electronico DROP CONSTRAINT IF EXISTS contacto_fk CASCADE;
ALTER TABLE agenda_personal.correo_electronico ADD CONSTRAINT fk_correo_electronico_contacto FOREIGN KEY (contacto_id)
REFERENCES agenda_personal.contacto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE CASCADE;
-- ddl-end --