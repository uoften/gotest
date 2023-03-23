ALTER TABLE `os`.`aw_spider_product`
ADD COLUMN `handler` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '商品的唯一名称' AFTER `share_u_id`,
ADD COLUMN `tags` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '商品标签' AFTER `handler`,
ADD COLUMN `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '' COMMENT '商品分类名称' AFTER `tags`;

CREATE TABLE `aw_purchase_mb` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `groupId` int(11) NOT NULL COMMENT '马帮采购单号',
     `providerId` int(11) DEFAULT NULL COMMENT '供应商编号',
     `amount` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '实付总金额',
     `originAmount` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '原总金额',
     `currency` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '币种',
     `currencyRate` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '汇率',
     `expressMoney` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '运费',
     `taxAmount` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '税金',
     `discountAmount` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '折扣金额',
     `paymentStatus` tinyint(1) DEFAULT NULL COMMENT '采购单支付状态:1、未支付 2、已申请 3、部分支付 4、已完成',
     `checkStatus` tinyint(1) DEFAULT NULL COMMENT '审核状态:1、待确认 2、待审核 3、审核打回 4、审核通过',
     `lastStorageTime` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '最后一次入库操作时间',
     `createTime` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '创建时间',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='马帮采购单列表';

CREATE TABLE `aw_purchase_mb_sku` (
      `group_id` int(11) NOT NULL COMMENT '马帮采购单号',
      `stock_sku` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '库存sku',
      `name_cn` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '中文名',
      `purchase_num` int(6) DEFAULT '0' COMMENT '采购量',
      `sell_price` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '采购价格',
      `origin_sell_price` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '原采购价格',
      `per_express_money` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '单个sku平均运费',
      `arrive_total` int(6) DEFAULT '0' COMMENT '到货总数',
      `confirm_num` int(6) DEFAULT '0' COMMENT '供应商确认的数量',
      `confirm_price` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '供应商确认的价格',
      `refund_num` int(6) DEFAULT '0' COMMENT '退货数量',
      UNIQUE KEY `idx_oid_sku` (`group_id`,`stock_sku`) USING HASH
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='马帮采购单sku列表';