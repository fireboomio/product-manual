# 性能测试

针对不同数据库进行了压力测，详情见下文：

## Mysql

### 表100 字段10 无数据

<table><thead><tr><th width="99">api数量</th><th width="127">cpu</th><th width="90">内存</th><th width="166">配置文件大小</th><th>编译</th></tr></thead><tbody><tr><td>10</td><td>2.8%～3.3%</td><td>99.3MB</td><td>graphql:37KB json:122KB</td><td>"totalCost": "123.519209ms", "startCost": "8.292917ms"</td></tr><tr><td>50</td><td>2.9%~3.3%</td><td>172MB</td><td>graphql:37KB json:214KB</td><td>"totalCost": "149.230167ms", "startCost": "32.58175ms"</td></tr><tr><td>100</td><td>4.1%～4.6%</td><td>210.3MB</td><td>graphql:37KB json:322KB</td><td>"totalCost": "199.703125ms", "startCost": "68.734ms"</td></tr><tr><td>200</td><td>6.0%~6.6%</td><td>225.1MB</td><td>graphql:103KB json:703KB</td><td>"totalCost": "206.980375ms", "startCost": "84.656084ms"</td></tr><tr><td>500</td><td>15.9%~17.5%</td><td>260MB</td><td>graphql:123KB json:1.5MB</td><td>"totalCost": "563.586292ms", "startCost": "290.274083ms"</td></tr><tr><td>1000</td><td>30.7%~34.1%</td><td>299.9MB</td><td>graphql:855KB json:4.6MB</td><td>"totalCost": "1.012925125s", "startCost": "601.4245ms"</td></tr></tbody></table>

PS：<mark style="color:orange;">totalCost</mark>: build+start耗时；<mark style="color:orange;">startCost</mark>：start耗时

### 表200 字段10 无数据

<table><thead><tr><th width="98">api数量</th><th width="135">cpu</th><th width="110">内存</th><th width="150">配置文件大小</th><th>编译30s</th></tr></thead><tbody><tr><td>10</td><td>6.3%~6.8%</td><td>144.6MB</td><td>graphql:1.7MB json:4.4MB</td><td>"totalCost": "292.562125ms", "startCost": "47.156125ms"</td></tr><tr><td>50</td><td>7.7%%~9.3%</td><td>171.0MB</td><td>graphql:1.7MB json:4.5MB</td><td>"totalCost": "328.937834ms", "startCost": "69.415417ms"</td></tr><tr><td>100</td><td>10.3%~10.8%</td><td>192.8MB</td><td>graphql:1.7MB json:4.7MB</td><td>"totalCost": "373.724792ms", "startCost": "115.373541ms"</td></tr><tr><td>200</td><td>13.3%～15.1%</td><td>253.2MB</td><td>graphq:1.7MB json:4.8MB</td><td>"totalCost": "486.693042ms", "startCost": "181.506583ms"</td></tr><tr><td>500</td><td>21.8%~22.1%</td><td>257.5MB</td><td>graphql:1.7MB json:5.4MB</td><td>"totalCost": "609.782792ms", "startCost": "316.42125ms"</td></tr><tr><td>1000</td><td>37.8%~39.2%</td><td>322.1MB</td><td>graphql:1.7MB json:6.7MB</td><td>"totalCost": "1.159405833s", "startCost": "724.692917ms"</td></tr></tbody></table>

### 表500 字段10 无数据



<table><thead><tr><th width="85">api数量</th><th width="147">cpu</th><th width="107">内存</th><th width="144">配置文件大小</th><th>编译30s</th></tr></thead><tbody><tr><td>10</td><td>15.9%~16.8%</td><td>307.4MB</td><td>graphql:4.2MB json:10.9MB</td><td>"totalCost": "951.672958ms", "startCost": "151.391708ms"</td></tr><tr><td>50</td><td>17.7%%~19.3%</td><td>340.4MB</td><td>graphql:4.2MB json:11MB</td><td>"totalCost": "879.343583ms", "startCost": "144.660917ms"</td></tr><tr><td>100</td><td>21.1%～22.3%</td><td>380.9MB</td><td>graphql:4.2MB json:11.1MB</td><td>"totalCost": "913.015583ms", "startCost": "208.3255ms"</td></tr><tr><td>200</td><td>26.9%~29.3%</td><td>423.5MB</td><td>graphql:4.2MB json:11.3MB</td><td>"totalCost": "1.396843375s", "startCost": "508.460292ms"</td></tr><tr><td>500</td><td>38.1%~40.7%</td><td>462.0MB</td><td>graphql:4.2MB json:12MB</td><td>"totalCost": "1.436549125s", "startCost": "624.08525ms"</td></tr><tr><td>1000</td><td>58.5%~60.8%</td><td>849.1MB</td><td>graphql:4.2MB json:13.1MB</td><td>"totalCost": "2.565050667s", "startCost": "1.556542334s"</td></tr></tbody></table>

### 表1000 字段10 无数据

<table><thead><tr><th width="110">api数量</th><th width="140">cpu</th><th width="103">内存</th><th width="156">配置文件大小</th><th>编译30s</th></tr></thead><tbody><tr><td>10</td><td>34.7%~39.7%</td><td>583.3</td><td>graphql:8.3MB json:21.6MB</td><td>"totalCost": "2.127657708s", "startCost": "223.030166ms"</td></tr><tr><td>50</td><td>40.4%~43.6</td><td>563.3MB</td><td>graphql:8.3MB json:21.7MB</td><td>"totalCost": "2.20061825s", "startCost": "277.581833ms"</td></tr><tr><td>100</td><td>36.1%～43.3%</td><td>588.4MB</td><td>graphql:8.3MB json:21.8MB</td><td>"totalCost": "2.499653542s", "startCost": "437.310958ms"</td></tr><tr><td>200</td><td>32%~37.5%</td><td>638.5MB</td><td>graphql:8.3MB json:22.1MB</td><td>"totalCost": "2.547601791s", "startCost": "548.155917ms"</td></tr><tr><td>500</td><td>46.2%~57.3%</td><td>740.8MB</td><td>graphql:8.3MB json:22.8MB</td><td>"totalCost": "3.548618667s", "startCost": "1.489989041s"</td></tr><tr><td>1000</td><td>52.2%~72.5%</td><td>760.8MB</td><td>graphql:8.3MB json:23.9MB</td><td>"totalCost": "3.642355459s", "startCost": "1.700005s"</td></tr></tbody></table>

### 建表语句

```sql
DROP PROCEDURE IF EXISTS `createTables`;

DELIMITER $$
CREATE PROCEDURE `createTables`()
BEGIN
    DECLARE `@i` int(11);
    DECLARE `@createSql` VARCHAR(2560);
    DECLARE `@createIndexSql1` VARCHAR(2560);
    DECLARE `@createIndexSql2` VARCHAR(2560);
    DECLARE `@createIndexSql3` VARCHAR(2560);
    DECLARE `@createIndexSql4` VARCHAR(2560);
    DECLARE `@createIndexSql5` VARCHAR(2560);
    DECLARE `@createIndexSql6` VARCHAR(2560);
    DECLARE `@createIndexSql7` VARCHAR(2560);
    DECLARE `@createIndexSql8` VARCHAR(2560);


    set `@i`=1;
    WHILE `@i`<=500 DO 
         -- `M_ID` bigint AUTO_INCREMENT PRIMARY KEY NOT NULL,
         -- createTable     
         SET @createSql = CONCAT('CREATE TABLE IF NOT EXISTS test_',`@i`,'(
            id INT(11) NOT NULL,
            name VARCHAR(255) ,
            info VARCHAR(255) ,
            info1 VARCHAR(255) ,
            info2 VARCHAR(255),
            info3 VARCHAR(255),
            info4 VARCHAR(255),
            info5 VARCHAR(255),
            info6 VARCHAR(255),
            info7 VARCHAR(255),
            PRIMARY KEY (id)  
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin');
         prepare stmt from @createSql;
         execute stmt;
         SET `@i`= `@i`+1; 
    END WHILE;
END$$
DELIMITER;


CALL createTables();
```

## Postgress

### 表100 字段10 无数据



<table><thead><tr><th width="98">api数量</th><th width="107">cpu</th><th width="98">内存</th><th width="135">配置文件大小</th><th>编译30s（编译成功有打印）内极限表，api数量</th></tr></thead><tbody><tr><td>10</td><td>1.4%~1.9%</td><td>50.7MB</td><td>graphql:37KB json:122KB</td><td>"totalCost": "157.595917ms", "startCost": "8.414834ms"</td></tr><tr><td>50</td><td>2.6%~2.7%</td><td>94.9MB</td><td>graphql:37KB json:254KB</td><td>"totalCost": "185.09025ms", "startCost": "31.112083ms"</td></tr><tr><td>100</td><td>4.4%~4.7%</td><td>122.6MB</td><td>graphql:37KB json:419KB</td><td>"totalCost": "257.391583ms", "startCost": "57.365375ms"</td></tr><tr><td>200</td><td>7.0%~9.9%</td><td>180.1MB</td><td>graphql:37KB json:666KB</td><td>"totalCost": "306.08625ms", "startCost": "116.093166ms"</td></tr><tr><td>500</td><td>15.2%~15.5%</td><td>198.1MB</td><td>graphql:37KB json:1.3MB</td><td>"totalCost": "476.344792ms", "startCost": "255.8845ms"</td></tr><tr><td>1000</td><td>32.5%~33.4%</td><td>313.2MB</td><td>graphql:37KB json:2.7MB</td><td>"totalCost": "907.860042ms", "startCost": "495.054958ms"</td></tr></tbody></table>

### 表200 字段10 无数据



<table><thead><tr><th width="119">api数量</th><th width="114">cpu</th><th width="132">内存</th><th width="150">配置文件大小</th><th>编译30s</th></tr></thead><tbody><tr><td>10</td><td>1.6%~1.7%</td><td>43.1MB</td><td>graphql:37KB json:122KB</td><td>"totalCost": "195.71125ms", "startCost": "6.9135ms"</td></tr><tr><td>50</td><td>2.6%~2.7%</td><td>82.4MB</td><td>graphql:37KB json:204KB</td><td>"totalCost": "205.815917ms", "startCost": "22.947959ms"</td></tr><tr><td>100</td><td>3.3%~3.7%</td><td>110.9MB</td><td>graphql:37KB json:319KB</td><td>"totalCost": "278.785083ms", "startCost": "45.571208ms"</td></tr><tr><td>200</td><td>5.8%~6.2%</td><td>121.4MB</td><td>graphql:37KB json:550KB</td><td>"totalCost": "322.394291ms", "startCost": "79.350083ms"</td></tr><tr><td>500</td><td>20.2%~21.1%</td><td>171.1MB</td><td>graphql:37KB json:1.7MB</td><td>"totalCost": "615.397042ms", "startCost": "285.051209ms"</td></tr><tr><td>1000</td><td>29.4%~30.7%</td><td>284.3MB</td><td>graphql:37KB json:2.4MB</td><td>"totalCost": "898.1295ms", "startCost": "457.280666ms"</td></tr></tbody></table>

### 表500 字段10 无数据



| api数量 | cpu          | 内存      | 配置文件大小                  | 编译30s                                                   |
| ----- | ------------ | ------- | ----------------------- | ------------------------------------------------------- |
| 10    | 1.9%\~2.2%   | 71.5MB  | graphql:37KB json:138KB | "totalCost": "355.816208ms", "startCost": "14.314625ms" |
| 50    | 3.1%～3.2%    | 96.0MB  | graphql:37KB json:246KB | "totalCost": "345.894042ms", "startCost": "27.083792ms  |
| 100   | 4.2%\~4.4%   | 123.8MB | graphql:37KB json:479KB | "totalCost": "417.617625ms", "startCost": "52.235709ms" |
| 200   | 8.0%\~8.8%   | 206.6MB | graphql:37KB json:714KB | "totalCost": "502.792166ms", "startCost": "126.14525ms" |
| 500   | 13.0%\~17.5% | 266.4MB | graphql:37KB json:1.4MB | totalCost": "642.748833ms", "startCost": "230.391667ms  |
| 1000  | 28.3%\~30.7% | 326.7MB | graphql:37KB json:2.4MB | "totalCost": "989.846667ms", "startCost": "427.6195ms"  |

### 表1000 字段10 无数据

<table><thead><tr><th width="107">api数量</th><th width="125">cpu</th><th width="116">内存</th><th width="145">配置文件大小</th><th>编译30s</th></tr></thead><tbody><tr><td>10</td><td>2.6%~3.3%</td><td>88.4MB</td><td>graphql:37KB json:138KB</td><td>"totalCost": "681.454792ms", "startCost": "13.541083ms"</td></tr><tr><td>50</td><td>4.0%～4.2%</td><td>96.0MB</td><td>graphql:37KB json:262KB</td><td>"totalCost": "345.894042ms", "startCost": "27.083792ms</td></tr><tr><td>100</td><td>4.2%~4.4%</td><td>139.8MB</td><td>graphql:37KB json:284KB</td><td>"totalCost": "685.456084ms", "startCost": "46.579792ms"</td></tr><tr><td>200</td><td>7.2%~9.8%</td><td>177.4MB</td><td>graphql:37KB json:630KB</td><td>"totalCost": "770.267417ms", "startCost": "116.968292ms</td></tr><tr><td>500</td><td>25.3%~28.9%</td><td>261.4MB</td><td>graphql:37KB json:1.9MB</td><td>"totalCost": "1.115410583s", "startCost": "380.595459ms"</td></tr><tr><td>1000</td><td>33.4%~36.5%</td><td>370.3MB</td><td>graphql:37KB json:2.4MB</td><td>"totalCost": "1.327615458s", "startCost": "548.964167ms"</td></tr></tbody></table>

### 建表语句

```sql
DO $$  
DECLARE
  i INT := 0;
BEGIN
  WHILE i < 100 LOOP
    EXECUTE format('CREATE TABLE table_%s (
      col1 VARCHAR(50), 
      col2 VARCHAR(50),
      col3 VARCHAR(50),
                        col4 VARCHAR(50),
                        col5 VARCHAR(50),
                        col6 VARCHAR(50),
                        col7 VARCHAR(50),
                        col8 VARCHAR(50),
                        col9 VARCHAR(50),
      col10 VARCHAR(50)
    );', i);

    i := i + 1; 
  END LOOP;
END;
$$
```

## Mongo

### 表100 字段10 无数据



| api数量 | cpu          | 内存      | 配置文件大小                   | 编译30s（编译成功有打印）内极限表，api数量                                 |
| ----- | ------------ | ------- | ------------------------ | -------------------------------------------------------- |
| 10    | 2.2%\~2.5%   | 87.1MB  | graphql:859KB json:3.8MB | "totalCost": "142.325208ms", "startCost": "14.094541ms"  |
| 50    | 3.2%～5.6%    | 96.0MB  | graphql:859KB json:4.1MB | "totalCost": "169.01375ms", "startCost": "39.202083ms"   |
| 100   | 4.1%\~4.8%   | 141.7MB | graphql:859KB json:4.2MB | totalCost": "211.196166ms", "startCost": "77.196ms"      |
| 200   | 9.1%\~9.4%   | 225.5MB | graphql:859KB json:4.4MB | "totalCost": "314.434708ms", "startCost": "160.06875ms"  |
| 500   | 19.9%\~20.2% | 296.4MB | graphql:859KB json:4.5MB | "totalCost": "557.709833ms", "startCost": "353.181709ms" |
| 1000  | 28.5%\~32.4% | 322.4MB | graphql:859KB json:4.6MB | "totalCost": "832.446125ms", "startCost": "567.292291ms" |

### 表200 字段10 无数据



| api数量 | cpu          | 内存      | 配置文件大小                   | 编译30s（编译成功有打印）内极限表，api数量                                 |
| ----- | ------------ | ------- | ------------------------ | -------------------------------------------------------- |
| 10    | 4.3%6\~.7%   | 197.6MB | graphql:1.7MB json:4.5MB | "totalCost": "437.502667ms", "startCost": "50.722708ms"  |
| 50    | 9.8%\~10.1   | 296.1MB | graphql:1.7MB json:4.7MB | "totalCost": "508.138042ms", "startCost": "106.677167ms" |
| 100   | 12.4%\~13.7% | 343.8MB | graphql:1.7MB json:4.8MB | "totalCost": "622.330917ms", "startCost": "216.177875ms" |
| 200   | 20.0%\~21.0% | 435.4MB | graphql:1.7MB json:5.3MB | "totalCost": "951.689666ms", "startCost": "362.947542ms" |
| 500   | 32.0%\~32.4% | 455.7MB | graphql:1.7MB json:5.9MB | "totalCost": "1.362056167s", "startCost": "753.840083ms" |
| 1000  | 31.2%～48.1%  | 585.5MB | graphql:1.7MB json:6.8MB | "totalCost": "1.526851792s", "startCost": "976.812042ms" |

### 表500 字段10 无数据



<table><thead><tr><th width="87">api数量</th><th width="137">cpu</th><th width="104">内存</th><th width="148">配置文件大小</th><th>编译30s</th></tr></thead><tbody><tr><td>10</td><td>15.6%~18.4%</td><td>535.8MB</td><td>graphql:4.2MB json:11.1MB</td><td>"totalCost": "1.029646042s", "startCost": "129.197333ms"</td></tr><tr><td>50</td><td>18.5%~21.3%</td><td>574.2MB</td><td>graphql:4.2MB json:11.2MB</td><td>"totalCost": "1.095261291s", "startCost": "196.224666ms"</td></tr><tr><td>100</td><td>21.6%~22.8%</td><td>615.6MB</td><td>graphql:4.2MB json:11.4MB</td><td>"totalCost": "622.330917ms", "startCost": "216.177875ms"</td></tr><tr><td>200</td><td>31.8%~32.8%</td><td>626.3MB</td><td>graphql:4.2MB json:11.8MB</td><td>"totalCost": "1.53551s", "startCost": "599.480458ms"</td></tr><tr><td>500</td><td>28.0%~87.8%</td><td>779.6MB</td><td>graphql:4.2MB json:12.7MB</td><td>"totalCost": "2.486889583s", "startCost": "1.358723584s"</td></tr><tr><td>1000</td><td>76.2%~139.4%</td><td>785.3MB</td><td>graphql:4.2MB json:13.4MB</td><td>"totalCost": "3.03070325s", "startCost": "1.821126083s"</td></tr></tbody></table>

### 表1000 字段10 无数据

| api数量 | cpu           | 内存    | 配置文件大小                    | 编译30s（编译成功有打印）内极限表，api数量                                 |
| ----- | ------------- | ----- | ------------------------- | -------------------------------------------------------- |
| 10    | 12.6%\~34.0%  | 1.03G | graphql:8.3MB json:22.1MB | "totalCost": "2.101982167s", "startCost": "216.063ms"    |
| 50    | 34.7%\~43.3   | 1.02G | graphql:8.3MB json:22.2MB | "totalCost": "2.328474125s", "startCost": "370.499125ms" |
| 100   | 44.8%\~89.4%  | 1.03G | graphql:8.3MB json:22.4MB | "totalCost": "2.520597084s", "startCost": "571.611958ms" |
| 200   | 29.5%\~61.5%  | 1.13G | graphql:8.3MB json:22.9MB | "totalCost": "3.226986208s", "startCost": "1.238466958s" |
| 500   | 32.2%\~142.1% | 1.25G | graphql:8.3MB json:23.6MB | "totalCost": "4.242829167s", "startCost": "2.161365666s" |
| 1000  | 27.6%\~95%    | 1.29G | graphql:8.3MB json:24.4MB | "totalCost": "5.393288417s", "startCost": "3.294156125s" |

### 建表语句

```javascript
const schema1 = {
  properties: {
    field1: {bsonType: 'string'},
    field2: {bsonType: 'string'},
    field3: {bsonType: 'string'},
    field4: {bsonType: 'string'},
    field5: {bsonType: 'string'},
    field6: {bsonType: 'string'},
    field7: {bsonType: 'string'},
    field8: {bsonType: 'string'},
    field9: {bsonType: 'string'},
    field10: {bsonType: 'string'}
  }
}

for (i = 1; i <= 10; i++) {
  db.createCollection("table_" + i, {
    validator: {
      $jsonSchema: schema1
    }
  }) 
}

const doc = {
  field1: "value1", 
  field2: "value2",
  field3: "value2",
  field4: "value2",
  field5: "value2",
  field6: "value2",
  field7: "value2",
  field8: "value2",
  field9: "value2",
  field10: "value10" 
}

for (i = 1; i <= 10; i++) {
  collName = "table_" + i
  
  db[collName].insertMany([doc, doc, doc]) 
}
```
