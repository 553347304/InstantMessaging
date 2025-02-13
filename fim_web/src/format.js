import fs from 'fs'
import path from 'path'


function vueConfig(filePath) {
    if (filePath !== "views\\login\\index.vue") return;

    const tag = {
        form: "el-form-item",
    }

    fs.readFile(filePath, 'utf8', (err, data) => {
        let newData = data;
        let form = newData.match(new RegExp(`<${tag.form}>([\\s\\S]*?)<\/${tag.form}>`, 'g'));
        form.forEach(match => {
            let replace = match.replace(new RegExp(`<${tag.form}>|<\/${tag.form}>`, 'g'), '').trim();
            let content = `<${tag.form}>${replace}</${tag.form}>`;
            newData = newData.replace(match, content);
        });

        fs.writeFile(filePath, newData, err => {if (err) throw err;});
    });
}


let state = true;

function Traverse(folder) {
    fs.readdir(folder, (err, files) => {
        
        // 遍历文件夹中的每个项
        files.forEach(item => {
            if (!state) return;

            const itemPath = path.join(folder, item);
            console.log(itemPath)
            
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