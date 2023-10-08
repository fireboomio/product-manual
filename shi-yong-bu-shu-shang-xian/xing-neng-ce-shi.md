# 性能测试

原文：[https://y1zoqrgk8wb.feishu.cn/docx/LmOJdiSE7oYlMgxOGBKc4Itundg](https://y1zoqrgk8wb.feishu.cn/docx/LmOJdiSE7oYlMgxOGBKc4Itundg)



针对不同数据库进行了压力测，详情见下文：



### mysql

<table data-header-hidden><thead><tr><th width="88">表数量</th><th>cpu</th><th width="100">内存</th><th>配置文件大小</th><th>编译30s（编译成功有打印）内极限表，api数量</th><th>备注</th><th data-hidden></th></tr></thead><tbody><tr><td>10</td><td>1.5%～2.0%</td><td>124.3MB</td><td>graphql:103KB json:271KB</td><td>120～140ms</td><td>无数据 无api</td><td>mysql</td></tr><tr><td>50</td><td>2.5%~2.9%</td><td>198.9MB</td><td>graphql:440KB json:1.1MB</td><td>150~180ms</td><td>无数据 无api</td><td>mysql</td></tr><tr><td>200</td><td>7.4%</td><td>314.3MB</td><td>graphql:1.7MB json:4.3MB</td><td>409.942375ms~483.471583ms</td><td>无数据 无api</td><td>mysql</td></tr><tr><td>500</td><td>20.3%~21.2%</td><td>426.2MB</td><td>graphql:4.1MB json:10.8MB</td><td>1.080828958s~1.144414791s</td><td>无数据 无api</td><td>mysql</td></tr><tr><td>1000</td><td>39.3%~47.8%</td><td>1001.2MB</td><td>graphql:8.3MB json:21.5MB</td><td>3.098373458s~3.121010792s</td><td>无数据 无api</td><td>mysql</td></tr></tbody></table>





###

### 1





### 2









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
