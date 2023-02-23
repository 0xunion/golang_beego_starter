import deepAssign from './mini-deep-assign.mjs';

export class Helper {
    
    // uc first 别名函数
    up1(string)
    {
        return string.charAt(0).toUpperCase() + string.slice(1);
    }

    low1(string)
    {
        return string.charAt(0).toLowerCase() + string.slice(1);
    }

    // add space between camel case
    as(string)
    {
        return string.replace(/\B([A-Z])/g, ' $1');
    }

    low(string)
    {
        return string.toLocaleLowerCase();
    }

    up(string)
    {
        return string.toLocaleUpperCase();
    }

    // 驼峰转下划线 under_line
    ul(string)
    {
        return string.replace(/\B([A-Z])/g, '_$1').toLowerCase();
    }

    // 下划线转小驼峰 littleCamel
    lc(string)
    {
        return this.low1(string.replace(/_(\w)/g,  (all, letter) =>letter.toUpperCase()));
    }

    // 下划线转大驼峰 bigCamel
    bc(string)
    {
        return this.up1(this.lc(string));
    }

    // 去除/并转大驼峰
    u2bc(string)
    {
        return this.bc(string.replace(/\//g, '_'));
    }

    // 去除/并转小驼峰
    u2lc(string)
    {
        return this.lc(string.replace(/\//g, '_'));
    }

    // 去除/并转下划线
    u2ul(string)
    {
        let tmp = this.ul(string.replace(/\//g, '_'));
        // 如果是以_开头，那么就去掉
        if( tmp.charAt(0) == '_' ) {
            tmp = tmp.substr(1);
        }
        return tmp;
    }

    custom_method( REQ ) 
    {
        return REQ.method
    }

    is_native_type( type )
    {
        return ['string', 'int', 'float', 'bool', 'int64', 'float64'].includes(type);
    }

    is_array_type( type )
    {
        return type.startsWith('[]');
    }

    is_id_type( type )
    {
        return type.indexOf('PrimaryId') != -1;
    }

    is_complex_type( type )
    {
        return type.indexOf('master_types') != -1;
    }

    is_$type( type )
    {
        return type.startsWith('$');
    }

    beego_validate( type, rules ) 
    {
        const validate_string = [];
        for (let [key, value] of Object.entries(rules)) {
            if( key == 'required' ) {
                validate_string.push(`Required`);
            }

            if( key == 'min' ) {
                if (type == 'string') {
                    validate_string.push(`MinSize(${value})`);
                } else {
                    validate_string.push(`Min(${value})`);
                }
            }

            if( key == 'max' ) {
                if (type == 'string') {
                    validate_string.push(`MaxSize(${value})`);
                } else {
                    validate_string.push(`Max(${value})`);
                }
            }

            if( key == 'length' ) {
                validate_string.push(`Length(${value})`);
            }

            if( key == 'range' ) {
                validate_string.push(`Range(${value[0]}, ${value[1]})`);
            }

            if( key == 'regexp' ) {
                validate_string.push(`Match(${value})`);
            }
        }

        return validate_string.join(';');
    }

    table( meta, table )
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        return tables[table];
    }

    // 获取所有字段的数组
    fields( meta, table, scenario = null )
    {
        let required_fields = [];
        let fields = [];
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        // 首先取得字段
        fields = tables[table].fields;
        if( !fields ) return false;

        // 然后取得场景
        if( scenario != null ) {
            // 检索目标场景并获取字段
            scenario = tables[table].scenarios[scenario];
            if( !scenario ) return false;
            for (let [key, _] of Object.entries(scenario)) {
                required_fields.push({
                    name: key,
                    value: fields[key]
                });
            }                
        }

        // 如果没有场景，那么就取得所有字段
        if( scenario == null ) {
            for (let [key, _] of Object.entries(fields)) {
                required_fields.push({
                    name: key,
                    value: fields[key]
                });
            }
        }

        return required_fields;
    }

    // 如A类型的Time字段需要每次自动更新，我们需要知道这个字段应该生成什么样的数据
    field_generate_type( meta, table, field, scenario )
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        return tables[table].scenarios[scenario][field];
    }

    // 检测某个场景下，某个类型的字段有多少个
    count_generate_type( meta, table, scenario, type )
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        let fields = tables[table].scenarios[scenario];
        let count = 0;
        for (let [key, value] of Object.entries(fields)) {
            if( value == type ) count++;
        }

        return count;
    }

    field( meta, table, field )
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        // 首先取得字段
        let fields = tables[table].fields;
        if( field ) field = this.ul(field);
        if( !fields[field] ) return false;

        return fields[field];
    }

    field_validate( meta, table, field )
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        // 首先取得字段
        let fields = tables[table].fields;
        if( field ) field = this.ul(field);
        if( !fields[field] ) return false;

        field = fields[field];

        const validate_string = [];
        if( field.form_validate ) {
            for (let [key, value] of Object.entries(field.form_validate)) {
                let field_type = field.type;
                if( field_type == 'int' || field_type == 'float' 
                    || field_type == 'int64' || field_type == 'float64' ) {
                    field_type = 'number';
                } else if ( field_type == 'string') {
                    field_type = 'string';
                }

                if( key == 'required' ) {
                    validate_string.push(`Required`);
                }

                if( key == 'min' ) {
                    if ( field_type == 'number' ) {
                        validate_string.push(`Min(${value})`);
                    } else {
                        validate_string.push(`MinSize(${value})`);
                    }
                }

                if( key == 'max' ) {
                    if ( field_type == 'number' ) {
                        validate_string.push(`Max(${value})`);
                    } else {
                        validate_string.push(`MaxSize(${value})`);
                    }
                }

                if( key == 'regx' ) {
                    validate_string.push(`Match(${value})`);
                }

                if (key == 'tel') {
                    validate_string.push(`Tel`);
                }

                if (key == 'email') {
                    validate_string.push(`Email`);
                }

                if (key == 'len') {
                    validate_string.push(`Length(${value})`);
                }
            }
        }

        return validate_string.join(';');
    }

    // 获取一个模型（表）在某个场景下的访问权限
    table_permissions( meta, table, scenario ) 
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        return tables[table].permissions[scenario];
    }

    permission_field( meta, table)
    {
        let tables = meta.DB.tables;
        if( table ) table = this.ul(table);
        if( !tables[table] ) return false;
        
        for (let [_, value] of Object.entries(tables[table].fields)) {
            if( value.is_permission_id ) {
                return value;
            }
        }

        return false;
    }

    // 布尔表达式计算 TODO: 去除对beego的依赖
    bool_expression( type, condition, v1, v2 )
    {
        let condition_map = {
            'eq': '==',
            'ne': '!=', 
            'gt': '>',
            'ge': '>=',
            'lt': '<',
            'le': '<=',
        };

        if (type.indexOf('PrimaryId') >= 0) {
            return `${v1}.String() ${condition_map[condition]} ${v2}.String()`
        }

        if (type.indexOf('int') >= 0) {
            return `${v1} ${condition_map[condition]} ${v2}`
        }

        if (type.indexOf('float') >= 0) {
            return `${v1} ${condition_map[condition]} ${v2}`
        }
    }

}

export default Helper;