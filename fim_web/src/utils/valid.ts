interface Payload {
    PayLoad: any
    exp: number
    iss: string
}

export namespace Valid {
    class jwt {
        Parse(token: string): Payload {
            let payload = token.split(".")[1];
            let atob = window.atob(payload.replace(/-/g, "+").replace(/_/g, "/"));
            return JSON.parse(decodeURIComponent(escape(atob)));
        }
    }

    export const Jwt = new jwt()
}

