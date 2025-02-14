const Json = new class {
    private child(data: any) {
        let result = "";
        for (const key in data) {
            let value = data[key];
            let type: string = typeof value;

            if (Array.isArray(value)) {
                type = "[]";
            } else if (typeof value === "object") {
                type = `{\n${this.child(value)}}`;
            }
            result += `${key}: ${type}\n`
        }

        return result;
    }

    TsInterface(json: string): string {
        const data = JSON.parse(json);
        return `interface Response {\n${this.child(data)}}`;
    }
}

console.log(Json.TsInterface(`
{
    "user_id": 1,
    "name": "白音",
    "sign": "签名",
    "avatar": "/",
    "recall_message": "消息",
    "friend_online": true,
    "sound": false,
    "secure_link": false,
    "save_password": false,
    "search_user": 2,
    "valid": 2,
    "valid_info": {
        "issue": [
            "问1",
            "问2",
            "问3"
        ],
        "answer": [
            "答1",
            "答2",
            "答3"
        ]
    }
}
`))