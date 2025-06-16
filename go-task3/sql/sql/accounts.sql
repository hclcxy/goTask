
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for accounts
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `balance` double(255, 3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of accounts
-- ----------------------------
INSERT INTO `accounts` VALUES (1, 190.000);
INSERT INTO `accounts` VALUES (2, 400.000);

SET FOREIGN_KEY_CHECKS = 1;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `from_account_id` int(11) NULL DEFAULT NULL,
  `to_account_id` int(11) NULL DEFAULT NULL,
  `amount` double(255, 3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
CREATE DEFINER=`root`@`localhost` PROCEDURE `transfer_money`(
    IN from_account INT,                  -- 输入参数：转出账户ID
    IN to_account INT,                    -- 输入参数：转入账户ID
    IN transfer_amount DECIMAL(10,2)      -- 输入参数：转账金额
)
BEGIN
    -- 声明变量：用于存储当前账户余额
    DECLARE current_balance DECIMAL(10,2);
    
    -- 声明异常处理器：捕获任何SQL错误
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;  -- 发生错误时回滚事务
        SELECT '交易失败：系统发生错误' AS message;
    END;

    -- 开始事务
    START TRANSACTION;
    
    -- 检查转出账户的余额并锁定该记录
    SELECT balance INTO current_balance 
    FROM accounts 
    WHERE id = from_account 
    FOR UPDATE;  -- 锁定记录，防止并发操作
    
    -- 检查余额是否充足
    IF current_balance < transfer_amount THEN
        -- 余额不足，回滚事务
        ROLLBACK;
        SELECT '交易失败：余额不足' AS message;
    ELSE
        -- 余额充足，执行转账操作
        
        -- 步骤1：从转出账户扣除金额
        UPDATE accounts 
        SET balance = balance - transfer_amount 
        WHERE id = from_account;
        
        -- 步骤2：向转入账户增加金额
        UPDATE accounts 
        SET balance = balance + transfer_amount 
        WHERE id = to_account;
        
        -- 步骤3：记录交易信息
        INSERT INTO transactions (from_account_id, to_account_id, amount)
        VALUES (from_account, to_account, transfer_amount);
        
        -- 提交事务
        COMMIT;
        SELECT '交易成功' AS message;
    END IF;
END



------------- 调用
CALL transfer_money(1, 2, 100.00);  -- 调用转账存储过程，将 100 元从账户 1 转到账户 2
-- 检查账户余额
SELECT * FROM accounts;