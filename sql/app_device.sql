DROP TABLE if exists device_type cascade;
CREATE TABLE device_type
(
  device_type_id serial NOT NULL,
  created_date_time timestamp without time zone,
  description character varying(255),
  is_active boolean,
  name character varying(255) NOT NULL,
  updated_date_time timestamp without time zone,
  CONSTRAINT device_type_pkey PRIMARY KEY (device_type_id)
);

DROP TABLE if exists app_dist_channel cascade;
CREATE TABLE app_dist_channel
(
  dist_channel_id serial NOT NULL,
  created_date_time timestamp without time zone,
  description character varying(255),
  is_active boolean,
  name character varying(255) NOT NULL,
  updated_date_time timestamp without time zone,
  CONSTRAINT app_dist_channel_pkey PRIMARY KEY (dist_channel_id)
);

DROP TABLE if exists device_status cascade;
CREATE TABLE device_status
(
  device_status_id serial NOT NULL,
  created_date_time timestamp without time zone,
  description character varying(255),
  is_active boolean,
  name character varying(255) NOT NULL,
  updated_date_time timestamp without time zone,
  CONSTRAINT device_status_pkey PRIMARY KEY (device_status_id)
);

DROP TABLE if exists device_info cascade;
CREATE TABLE device_info
(
  device_info_id serial NOT NULL,
  android_device_id character varying(255),
  android_version character varying(255),
  api_level character varying(255),
  brand character varying(255),
  build_number character varying(255),
  cpu_hardware character varying(255),
  created_date_time timestamp without time zone,
  display_density character varying(255),
  display_physical_size character varying(255),
  display_resolution character varying(255),
  hardware_serial_no character varying(255),
  imei character varying(255) NOT NULL,
  instruction_sets character varying(255),
  is_active boolean,
  manufacturer character varying(255),
  model character varying(255),
  updated_date_time timestamp without time zone,
  device_type_id integer,
  CONSTRAINT device_info_pkey PRIMARY KEY (device_info_id),
  CONSTRAINT fk_device_info_device_type FOREIGN KEY (device_type_id)
      REFERENCES device_type (device_type_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

DROP TABLE if exists device_request cascade;
CREATE TABLE device_request
(
  device_request_id serial NOT NULL,
  created_date_time timestamp without time zone,
  device_status_id integer,
  is_active boolean,
  request_by character varying(255),
  updated_date_time timestamp without time zone,
  device_info_id integer,
  CONSTRAINT device_request_pkey PRIMARY KEY (device_request_id),
  CONSTRAINT fk_device_request_device_type FOREIGN KEY (device_info_id)
      REFERENCES device_info (device_info_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

DROP TABLE if exists app cascade;
CREATE TABLE app
(
  app_id serial NOT NULL,
  android_package character varying(255),
  android_version character varying(255),
  apk_data bytea NOT NULL,
  apk_name character varying(255) NOT NULL,
  created_date_time timestamp without time zone,
  description character varying(255),
  is_active boolean,
  name character varying(255) NOT NULL,
  updated_date_time timestamp without time zone,
  dist_channel_id integer,
  device_type_id integer,
  CONSTRAINT app_pkey PRIMARY KEY (app_id),
  CONSTRAINT fk_app_device_type_id FOREIGN KEY (device_type_id)
      REFERENCES device_type (device_type_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_app_dist_channel_id FOREIGN KEY (dist_channel_id)
      REFERENCES app_dist_channel (dist_channel_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

DROP TABLE if exists app_dependency cascade;
CREATE TABLE app_dependency
(
  app_main_id integer NOT NULL,
  app_dependency_id integer NOT NULL,
  CONSTRAINT fk_app_dependency_app_dependency_id FOREIGN KEY (app_dependency_id)
      REFERENCES app (app_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_app_dependency_app_main_id FOREIGN KEY (app_main_id)
      REFERENCES app (app_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

	  
DROP TABLE if exists device_app cascade;
CREATE TABLE device_app
(
  device_app_id serial NOT NULL,
  app_id integer,
  device_info_id integer,
  device_status_id integer,
  CONSTRAINT device_app_pkey PRIMARY KEY (device_app_id),
  CONSTRAINT fk55ap8hrgsrbt4941oqc4fnw4i FOREIGN KEY (app_id)
      REFERENCES app (app_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_device_app_device_status FOREIGN KEY (device_status_id)
      REFERENCES device_status (device_status_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT fk_device_app_device_info FOREIGN KEY (device_info_id)
      REFERENCES device_info (device_info_id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE NO ACTION
);

drop table if exists app_user cascade;
create table app_user
(
  id serial not null,
  created_date_time timestamp without time zone,
  first_name character varying(255),
  is_active boolean,
  last_name character varying(255),
  password character varying(255),
  updated_date_time timestamp without time zone,
  username character varying(255),
  constraint user_pkey primary key (id)
);
create unique index username_idx on app_user using btree(username);

drop table if exists role cascade;
create table role
(
  id serial not null,
  description character varying(255),
  role_name character varying(255),
  constraint role_pkey primary key (id)
);
create unique index role_idx on role using btree(role_name);

drop table if exists user_role cascade;
create table user_role
(
  user_id bigint not null,
  role_id integer not null,
  constraint fk_role_id foreign key (role_id)
      references role (id) match simple
      on update no action on delete no action,
  constraint fk_user_id foreign key (user_id)
      references app_user (id) match simple
      on update no action on delete no action
);

insert into app_dist_channel(dist_channel_id, name, description)
values (1, 'DEMO', 'Demo channel');
insert into app_dist_channel(dist_channel_id, name, description)
values (2, 'PRODUCT', 'Product channel');
insert into app_dist_channel(dist_channel_id, name, description)
values (3, 'TEST', 'Test channel');
insert into app_dist_channel(dist_channel_id, name, description)
values (4, 'SUPPORT', 'Support channel');

insert into device_type(device_type_id, name, description, created_date_time, is_active)
values(1, 'Phone', 'Phone device type', CURRENT_TIMESTAMP, true);
insert into device_type(device_type_id, name, description, created_date_time, is_active)
values(2, 'Telpo', 'Telpo device type', CURRENT_TIMESTAMP, true);
insert into device_type(device_type_id, name, description, created_date_time, is_active)
values(3, 'AsperisX7', 'AsperisX7 device type', CURRENT_TIMESTAMP, true);

insert into app(dist_channel_id, device_type_id, name, apk_name, apk_data, created_date_time, is_active)
values(1, 1, 'test-Demo-Jan18a', 'test.apk', decode('DEADBEEF', 'hex'), CURRENT_TIMESTAMP, true);
insert into app(dist_channel_id, device_type_id, name, apk_name, apk_data, created_date_time, is_active)
values(2, 2, 'test-Demo-Jan18a', 'test.apk', decode('DEADBEEF', 'hex'), CURRENT_TIMESTAMP, true);
insert into app(dist_channel_id, device_type_id, name, apk_name, apk_data, created_date_time, is_active)
values(3, 3, 'test-Demo-Jan18a', 'test.apk', decode('DEADBEEF', 'hex'), CURRENT_TIMESTAMP, true);

-- Device status
insert into device_status (name, description, created_date_time, updated_date_time, is_active) values ('APPROVED', 'The request has been approved', current_timestamp, null, true);
insert into device_status (name, description, created_date_time, updated_date_time, is_active) values ('UNAPPROVED', 'The request has been disapproved.', current_timestamp, null, true);
insert into device_status (name, description, created_date_time, updated_date_time, is_active) values ('APPROVING', 'The request is in the process of approving.', current_timestamp, null, true);
insert into device_status (name, description, created_date_time, updated_date_time, is_active) values ('REQUEST_APPROVAL', 'The request has been sent.', current_timestamp, null, true);

-- Device info
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J7 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY101', '10840923FDSLK180392840928495', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J8 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY102', '10840923FDSLK180392840928496', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J9 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY103', '10840923FDSLK180392840928497', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J10 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY104', '10840923FDSLK180392840928498', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J11 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY105', '10840923FDSLK180392840928499', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J12 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY106', '10840923FDSLK180392840928100', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);
insert into device_info (
  device_type_id,
  model,
  manufacturer,
  brand,
  android_version,
  api_level,
  build_number,
  android_device_id,
  imei,
  hardware_serial_no,
  instruction_sets,
  cpu_hardware,
  display_resolution,
  display_density,
  display_physical_size,
  created_date_time,
  updated_date_time,
  is_active
) values (1, 'Samsung Galaxy J13 Pro', 'Samsung', 'Samsung', 'Android 7.0 (Nougat)', '2.0', '1.0.1.01', '50434927897AE23TY107', '10840923FDSLK180392840928101', '899001002395004889301', 'ARMv8-A', 'Octa-core 1.6 GHz Cortex-A53', '1080 x 1920 pixels', '~401 ppi', '5.5 inches 83.4 cm2 (~73.1% screen-to-body ratio)', current_timestamp, null, true);

-- Device request
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928495'), 'Mark', (select device_status_id from device_status where name = 'APPROVING'), current_timestamp, null, true);
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928496'), 'John', (select device_status_id from device_status where name = 'UNAPPROVED'), current_timestamp, null, true);
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928497'), 'Mark', (select device_status_id from device_status where name = 'REQUEST_APPROVAL'), current_timestamp, null, true);
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928498'), 'Luis', (select device_status_id from device_status where name = 'APPROVING'), current_timestamp, null, true);
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928499'), 'Aragorn',(select device_status_id from device_status where name = 'UNAPPROVED'), current_timestamp, null, true);
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928100'), 'Jason', (select device_status_id from device_status where name = 'REQUEST_APPROVAL'), current_timestamp, null, true);
insert into device_request (device_info_id, request_by, device_status_id, created_date_time, updated_date_time, is_active) values ((select device_info_id from device_info where imei = '10840923FDSLK180392840928101'), 'Slavi', (select device_status_id from device_status where name = 'APPROVING'), current_timestamp, null, true);

insert into role (id, role_name, description) values (1, 'STANDARD_USER', 'Standard User - Has no admin rights');
insert into role (id, role_name, description) values (2, 'ADMIN_USER', 'Admin User - Has permission to perform admin tasks');

-- users: admin (ADMIN_USER), test(STANDARD_USER)
-- non-encrypted password: thepass argon2$4$32768$4$32$1okUqMPgN4a0lzMmxyl+lg==$3Tk3zyvbKQuikHkBWq0dr9TsX5cv6EaRCyUiZphshFA=
insert into app_user (first_name, last_name, password, username, is_active, created_date_time, updated_date_time) values ('Test', 'Test', 'argon2$4$32768$4$32$1okUqMPgN4a0lzMmxyl+lg==$3Tk3zyvbKQuikHkBWq0dr9TsX5cv6EaRCyUiZphshFA=', 'test', true, current_timestamp, null);
-- non-encrypted password: adminpass argon2$4$32768$4$32$/HD7Ip8vh4GBhdpQOkYG9A==$AiN7mMFCf+DWsOub0pyTUPtwxwpoZM14EIpgI4IY4PI=
insert into app_user (first_name, last_name, password, username, is_active, created_date_time, updated_date_time) values ('Admin', 'Admin', 'argon2$4$32768$4$32$/HD7Ip8vh4GBhdpQOkYG9A==$AiN7mMFCf+DWsOub0pyTUPtwxwpoZM14EIpgI4IY4PI=', 'admin', true, current_timestamp, null);

insert into user_role (user_id, role_id) values((select id from app_user where username = 'test'),1);
insert into user_role (user_id, role_id) values((select id from app_user where username = 'admin'),1);
insert into user_role (user_id, role_id) values((select id from app_user where username = 'admin'),2);
