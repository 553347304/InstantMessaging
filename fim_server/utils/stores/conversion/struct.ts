class Convert {
    private _max_length: number = 0;

    private Repeat(char: string, length: number): string {
        let result = '';
        for (let i = 0; i < length; i++) {
            result += char;
        }
        return result;
    }

    private Align(content: string): string {
        const list = content.split("\n")
        let result = "";
        for (const s of list) {
            const text = `    ${s}${this.Repeat(" ", parseInt(this._max_length - s.length))}    ,`;
            result += text + "\n";
        }
        return result
    }

    GoStructToFieldType(goStruct: string): string {
        const json: string = /\{(?:(?!\{).)*}/s.exec(goStruct).toString();
        const name: string = /type (\w+) struct/.exec(goStruct.replace(json, ""))[1];
        const all = /(\w+)\s+(\w+)(?:\s+`([^`]+)`)?/g;
        let result = "";
        let match: any;
        while ((match = all.exec(json)) !== null) {
            const text: string = `${match[1]}:\n`;
            result += text;       // All Name Type Tag
            if (text.length > this._max_length) this._max_length = text.length;
        }
        result = result.slice(0, -1);   // 删除尾部最后一个 \n
        result = `${name} {\n${this.Align(result)}}`;
        return result;
    }
}

const c = new Convert();

console.log(c.GoStructToFieldType(`
type GroupMeInfo struct {
	Id          string \`json:"id"\`
	Name        string \`json:"name"\`
	Avatar      string \`json:"avatar"\`
	MemberTatal int64  \`json:"member_tatal"\`
	Role        int8   \`json:"role"\`
	Mode        int8   \`json:"mode"\`
}
`))