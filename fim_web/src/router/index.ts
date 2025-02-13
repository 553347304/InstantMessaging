import {createRouter, createWebHistory} from 'vue-router'


const children = {
    contacts: {
        path: "m", name: "contacts", component: () => import("@/views/web/contacts/index.vue"), redirect: "/m/welcome",
        children: [
            {path: "users/:id", name: "user_detail", component: () => import("@/views/web/contacts/user_detail.vue"),},
            {path: "groups/:id", name: "group_detail", component: () => import("@/views/web/contacts/group_detail.vue"),},
            {path: "add_group", name: "add_group", component: () => import("@/views/web/contacts/add_group.vue"),},
            {path: "welcome", name: "welcome", component: () => import("@/views/web/contacts/welcome.vue"),},
            {
                path: "search", name: "search", component: () => import("@/views/web/contacts/search/index.vue"),
                children: [
                    {path: "user", name: "search_user", component: () => import("@/views/web/contacts/search/search_user.vue"),},
                    {path: "group", name: "search_group", component: () => import("@/views/web/contacts/search/search_group.vue"),}
                ]
            }
        ],
    },
    session: {
        path: "", name: "session", component: () => import("@/views/web/session/index.vue"),
        children: [
            {path: "", name: "session_welcome", component: () => import("@/views/web/session/session_welcome.vue"),},
            {
                path: "user/:id", name: "session_user", component: () => import("@/views/web/session/user_chat/index.vue"),
                children: [
                    {path: "", name: "session_user_chat", component: () => import("@/views/web/session/user_chat/chat.vue"),},
                    {path: "info", name: "session_user_info", component: () => import("@/views/web/session/user_chat/info.vue"),},
                ]
            },
            {
                path: "group/:id", name: "session_group", component: () => import("@/views/web/session/group_chat/index.vue"),
                children: [
                    {path: "", name: "session_group_chat", component: () => import("@/views/web/session/group_chat/chat.vue"),},
                    {
                        path: "info", name: "session_group_info", component: () => import("@/views/web/session/group_chat/info/index.vue"),
                        children: [
                            {path: "", name: "group_information", component: () => import("@/views/web/session/group_chat/info/information.vue"),},
                            {path: "member", name: "group_member", component: () => import("@/views/web/session/group_chat/info/member.vue"),},
                            {path: "settings", name: "group_settings", component: () => import("@/views/web/session/group_chat/info/settings.vue"),},
                        ]
                    }
                ]
            }
        ]
    },
    info: {
        path: "info", name: "info", component: () => import("@/views/web/info/index.vue"), redirect: "/info",
        children: [
            {path: "", name: "my_info", component: () => import("@/views/web/info/my_info.vue"),},
            {path: "base", name: "base_info", component: () => import("@/views/web/info/base_info.vue"),},
            {path: "safe", name: "safe_info", component: () => import("@/views/web/info/safe_info.vue"),},
            {path: "role", name: "role_info", component: () => import("@/views/web/info/role_info.vue"),},
            {path: "settings", name: "settings", component: () => import("@/views/web/info/settings.vue"),}
        ]
    },
    notice: {path: "notice", name: "notice", component: () => import("@/views/web/notice/index.vue")},
}


const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {path: "/login", name: "login", component: () => import("@/views/login/index.vue")},
        {
            path: "/", name: "web", component: () => import("@/views/web/index.vue"),
            children: [
                children.contacts,
                children.session,
                children.info,
            ],
            meta: {
                isLogin: true, // 需要登陆验证的
            }
        }
    ],
})

export default router
