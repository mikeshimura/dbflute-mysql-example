-- #df:entity#
-- !df:pmb!
-- !!df.NullString name!!
SELECT 
  member.member_id, 
  member.member_name, 
  member.member_account, 
  member.birthdate, 
  member.formalized_datetime, 
  member.member_status_code, 
  member_status.member_status_name, 
  member_status.description
FROM 
  member, 
  member_status
WHERE 
/*IF pmb.name != null*/
  member.member_name like /*pmb.name*/'S%'
-- ELSE member.member_name like 'M%'
/*END*/
and member.member_status_code = member_status.member_status_code