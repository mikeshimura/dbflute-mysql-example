@echo off

set ANT_OPTS=-Xmx512m

set DBFLUTE_HOME=..\mydbflute\dbflute-go-0.0.1

set MY_PROJECT_NAME=mysql

set MY_PROPERTIES_PATH=build.properties
set JAVA_HOME=d:\pg\Java\jdk1.8.0_20

if "%pause_at_end%"=="" set pause_at_end=y
