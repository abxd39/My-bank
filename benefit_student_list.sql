/*
 Navicat Premium Data Transfer

 Source Server         : oracle
 Source Server Type    : Oracle
 Source Server Version : 122010
 Source Host           : 10.100.100.42
 Source Schema         : C##SUBSIDIES

 Target Server Type    : Oracle
 Target Server Version : 122010
 File Encoding         : utf-8

 Date: 12/12/2019 11:29:01 AM
*/

-- ----------------------------
--  Table structure for benefit_student_list
-- ----------------------------
DROP TABLE "C##SUBSIDIES"."benefit_student_list";
CREATE TABLE "benefit_student_list" (   "id" NUMBER(11,0) NOT NULL, "benefit_project_id" NUMBER(11,0) NOT NULL, "school_name" VARCHAR2(50BYTE) NOT NULL, "school_no" NUMBER(11,0) NOT NULL, "enter_year" NUMBER(11,0) DEFAULT 0  NOT NULL, "grade_no" NUMBER(11,0) NOT NULL, "grade_name" VARCHAR2(50BYTE) DEFAULT ''  NOT NULL, "sex" NUMBER(4,0) DEFAULT 0  NOT NULL, "class_no" NUMBER(11,0) NOT NULL, "class_name" VARCHAR2(50BYTE) NOT NULL, "student_name" VARCHAR2(50BYTE) NOT NULL, "student_no" NUMBER(11,0) NOT NULL, "id_type" NUMBER(4,0) DEFAULT 1  NOT NULL, "card_no" VARCHAR2(50BYTE) NOT NULL, "amount" NUMBER(11,0) NOT NULL, "deleted" NUMBER(4,0) DEFAULT 0  NOT NULL, "graduated" NUMBER(4,0) DEFAULT 0  NOT NULL, "guardian1_name" VARCHAR2(255BYTE) DEFAULT '' , "guardian1_card_no" VARCHAR2(255BYTE) DEFAULT '' , "guardian1_card_type" NUMBER(4,0) DEFAULT 0 , "guardian2_name" VARCHAR2(50BYTE) DEFAULT '' , "guardian2_card_no" VARCHAR2(50BYTE) DEFAULT '' , "guardian2_card_type" NUMBER(4,0) DEFAULT 1 , "created" TIMESTAMP(6) NULL, "updated" TIMESTAMP(6) NULL);

-- ----------------------------
--  Primary key structure for table benefit_student_list
-- ----------------------------
ALTER TABLE "C##SUBSIDIES"."benefit_student_list" ADD CONSTRAINT "SYS_C0060759" PRIMARY KEY("id");

-- ----------------------------
--  Checks structure for table benefit_student_list
-- ----------------------------
ALTER TABLE "C##SUBSIDIES"."benefit_student_list" ADD CONSTRAINT "SYS_C0060742" CHECK ("id" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060743" CHECK ("benefit_project_id" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060744" CHECK ("school_name" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060745" CHECK ("school_no" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060746" CHECK ("enter_year" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060747" CHECK ("grade_no" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060748" CHECK ("grade_name" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060749" CHECK ("sex" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060750" CHECK ("class_no" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060751" CHECK ("class_name" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060752" CHECK ("student_name" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060753" CHECK ("student_no" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060754" CHECK ("id_type" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060755" CHECK ("card_no" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060756" CHECK ("amount" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060757" CHECK ("deleted" IS NOT NULL) ENABLE ADD CONSTRAINT "SYS_C0060758" CHECK ("graduated" IS NOT NULL) ENABLE;

