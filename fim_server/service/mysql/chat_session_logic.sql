select *
from (select least(send_user_id, receive_user_id) as s_u,
             greatest(send_user_id, receive_user_id) as r_u,
             max(created_at) as max_date,
             (select preview
              from chat_models
              where ((send_user_id = s_u and receive_user_id = r_u)
                  or (send_user_id = r_u and receive_user_id = s_u))
                  and id not in (select chat_id from user_chat_delete_models where user_id = 1)
              order by created_at desc limit 1) as max_preview,
             if((select 1 from top_user_models where user_id = 1
                 and (top_user_id = s_u or top_user_id = r_u) limit 1), 1, 0) as is_top
      from chat_models
      where (send_user_id = 1 or receive_user_id = 1)
        and id not in (select chat_id from user_chat_delete_models where user_id = 1)
      group by least(send_user_id, receive_user_id),
               greatest(send_user_id, receive_user_id)
     ) as u
order by is_top desc, max_date desc;



select *
from chat_models
where send_user_id = 1 and receive_user_id = 2
   and id not in (select chat_id from user_chat_delete_models where user_id = 1);


# 用户发送消息总数
select send_user_id,count(id) from chat_models group by send_user_id;
select receive_user_id,count(id) from chat_models group by receive_user_id;