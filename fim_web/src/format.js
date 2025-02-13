import fs from 'fs'
import path from 'path'

let state = true;

const MAX_SIZE = 300;
const Tab = {
    form: "el-form-item",
}

function Trim(tag, data) {
    let source = data;
    let form = data.match(new RegExp(`<${tag}([\\s\\S]*?)<\/${tag}>`, 'g'));
    if (form === null) return null;
    form.forEach(match => {
        let content = match.replace(/>\s+</g, "><");
        if (content.length < MAX_SIZE) data = data.replace(match, content);
    });
    if (source === data) return null;
    return data;
}

function vueConfig(filePath) {
    fs.readFile(filePath, 'utf8', (err, data) => {
        let v = data;
        if (v !== null) v = Trim(Tab.form, v);
        
        if (v !== null) {
            console.warn(filePath);
            fs.writeFile(filePath, v, err => {if (err) throw err;});
        }
    });
}

function Traverse(folder) {
    fs.readdir(folder, (err, files) => {

        // 遍历文件夹中的每个项
        files.forEach(item => {
            if (!state) return;

            const itemPath = path.join(folder, item);
            fs.stat(itemPath, (err, stats) => {
                if (stats.isDirectory()) {
                    Traverse(itemPath); // 递归遍历文件夹
                } else if (stats.isFile()) {
                    if (path.extname(item) === '.vue') vueConfig(itemPath);
                }
            });
        });
    });
}

Traverse("views");