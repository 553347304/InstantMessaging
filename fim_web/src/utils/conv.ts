const Method = new class {
    DeleteDesc = (m_content: string): string => {
        return m_content.replace(/\/\/.*$/gm, '');
    }
}
const Json = new class {
    TsInterface = (m_name: string, m_content: string): string => {
        function child(data: any): string {
            let result = "";
            for (const key in data) {
                let value = data[key];
                let type: string = typeof value;

                if (Array.isArray(value)) type = "[]"
                if (!Array.isArray(value) && typeof value === "object") type = child(value);
                result += `\n${key}: ${type}`
            }

            return `{ ${result} \n}`;
        }
        const desc = Method.DeleteDesc(m_content)
        console.log(desc)
        const data = JSON.parse(desc);
        return `export interface ${m_name} ${child(data)}`;
    }
    TsValue = (m_name: string, m_content: string): string => {
        function child(data: any): string {
            let result = "";
            for (const key in data) {
                let value = data[key];
                if (typeof value === "number") value = 0;
                if (typeof value === "string") value = '""';
                if (typeof value === "object" && !Array.isArray(value)) value = child(value);
                if (Array.isArray(value)) value = "[]";
                result += `\n${key}: ${value}, `;
            }
            return `{ ${result} \n}`;
        }

        const desc = Method.DeleteDesc(m_content)
        const data = JSON.parse(desc);
        return `export const ${m_name} = (): ${m_name} => (${child(data)})`;
    }
}

const name = "Response";
const content = `

`;

console.log(Json.TsInterface(name, content));
console.log(Json.TsValue(name, content));