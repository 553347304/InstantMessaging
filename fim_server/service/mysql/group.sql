select group_id, user_id, role, created_at, member_name,
       (select group_message_models.created_at from group_message_models
                 where group_member_models.group_id = 1
                 and group_message_models.send_user_id = user_id limit 1) as new_message_date
from group_member_models
where group_id = 1;


select group_id,max(created_at),
       (select message_preview
        from group_message_models as g
        where g.group_id = g.group_id
        order by g.created_at desc limit 1)
       as new_message_date
from group_message_models
where group_id in (1)
group by group_id;