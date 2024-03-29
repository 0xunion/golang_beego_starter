#!/usr/local/bin/node
import { fileURLToPath } from 'url'
import path, { dirname, join } from 'path'
import fs from 'fs'
const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

import JSONC from 'jsonc-parser';
import minimist from 'minimist';
import { engine } from './engine.mjs';
import Helper from './helper.mjs';
const helper = new Helper();

const project_dir = path.join(__dirname, '..');
const meta_dir = path.join(project_dir, '.metatoy');


// 首先是要分析输入的命令
const args = minimist(process.argv.slice(2));
const { _: command, ...options } = args;
const command_name = command[0];
const command_args = command.slice(1);

if( !command_name )
{
    console.log('usage: exec new <template_dir> --name=<name> --data-file=<data_file>');
    process.exit();
}

// 加载元数据
let db_data = {};
const data_file_name = options['data-file'] || 'metatoy.data.jsonc';
const data_file = path.join(__dirname, data_file_name);
if( fs.existsSync(data_file) )
{
    const parsed = JSONC.parse(fs.readFileSync(data_file, 'utf8'));
    if( parsed ) db_data = parsed;
}

const tpl_path = path.join(meta_dir,'_template', ...command_args);

if (command_name?.toLocaleLowerCase() == 'custom')
{
    // 加载自定义接口配置文件
    let custom_data = {};
    const custom_file_name = options['custom-file'] || 'metatoy.custom.jsonc';
    const custom_file = path.join(__dirname, custom_file_name);
    if( fs.existsSync(custom_file) )
    {
        const parsed =  JSONC.parse(fs.readFileSync(custom_file, 'utf8'));
        if( parsed ) custom_data = parsed;
    }

    // 循环每一个自定义接口配置
    for( let key in custom_data ) {
        const custom = custom_data[key];
        if( custom )
        {
            // check validate has required
            for (let [key, value] of Object.entries(custom.args)) {
                if( !value.validate.required && value.type != 'int' && !value.validate.empty ) {
                    value.validate.required = true;
                }
            }

            if(custom.service) {
                console.log( `🚀 ${key} 处理中...` );
                engine( path.join(tpl_path, 'api/custom_service.tpl.ejs'), {"DB":db_data,"REQ":custom}, options );
            }
            console.log( `🚀 ${key} 处理中...` );
            engine( path.join(tpl_path, 'api/custom_controller.tpl.ejs'), {"DB":db_data,"REQ":custom}, options );
        }
    }
} 
else if (command_name?.toLocaleLowerCase() == 'custom_route') {
    // 加载自定义接口配置文件
    let custom_data = {};
    const custom_file_name = options['custom-file'] || 'metatoy.custom.jsonc';
    const custom_file = path.join(__dirname, custom_file_name);
    if( fs.existsSync(custom_file) )
    {
        const parsed =  JSONC.parse(fs.readFileSync(custom_file, 'utf8'));
        if( parsed ) custom_data = parsed;
    }

    // 循环每一个自定义接口配置
    for( let key in custom_data ) {
        const custom = custom_data[key];
        if( custom )
        {
            console.log( `🚀 ${key} 处理中...` );
            engine( path.join(tpl_path, 'api/custom_route.tpl.ejs'), {"DB":db_data,"REQ":custom}, options );
        }
    }
}
else if (command_name?.toLocaleLowerCase() == 'permission') {
    for (let [key, value] of Object.entries(db_data.tables)) {
        if( value )
        {
            console.log( `🚀 ${key} 处理中...` );
            engine( path.join(tpl_path, 'api/permission.tpl.ejs'), {"DB":db_data,"PER":value}, options );
        }
    }
}
else if( command_name?.toLocaleLowerCase() == 'new' ) 
{
    if( !options.name )
    {
        console.log('缺少参数 name');
        process.exit();
    }

    // 对name进行处理
    options.theName = helper.lc(options.name);
    options.TheName = helper.bc(options.name);
    options.the_name = helper.ul(options.name);
    options.cn_name = options.cn_name ?? options.TheName;

    // console.log( options );
    // console.log( tpl_path );
    const tpl_file = path.join(tpl_path) + '.tpl.ejs';
    const tpl_files = [];

    if( fs.existsSync( tpl_file ) && fs.statSync(tpl_file).isFile()  )
    {
        tpl_files.push(tpl_file);
        // console.log('文件模板');
    }
    else
    {
        if( fs.existsSync(tpl_path) && fs.statSync(tpl_path).isDirectory() )
        {
            // console.log('目录模板');
            // 读取目录下所有的 .tpl.ejs 文件并 push 到 tpl_files
            const files = fs.readdirSync(tpl_path);
            for( let file of files )
            {
                if( file.endsWith('.tpl.ejs') )
                {
                    tpl_files.push( path.join(tpl_path, file));
                }
            }
        }
    }

    if( tpl_files.length > 0 )
    {
        // 循环每一个模板文件，并进行代码生成
        for( let t_file of tpl_files )
        {
            console.log( `🚀 ${t_file} 处理中...` );
            engine( t_file, {"DB":db_data}, options );
        }
        console.log( `🎈 完成` );
    }
}


// 终止命令
process.exit();


