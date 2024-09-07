create database account;
use account;
#登录账号
CREATE TABLE `accounts` (
      `id` bigint(20) NOT NULL AUTO_INCREMENT,
      `account` VARCHAR(255) NOT NULL,
      `password` VARCHAR(255) NOT NULL,
      `cate` tinyint NOT NULL COMMENT '账号类型',
      `nickname` VARCHAR(255) NOT NULL COMMENT '账号昵称',
      `avatar` VARCHAR(255) DEFAULT '' COMMENT '头像',
      `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `is_delete` INT NOT NULL DEFAULT 0,
      PRIMARY KEY (`id`),
      UNIQUE KEY `unique_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#登录事件
CREATE TABLE `login` (
       `id` bigint(20) NOT NULL AUTO_INCREMENT,
       `account` VARCHAR(255) NOT NULL,
       `event` tinyint NOT NULL COMMENT '事件,1:登入  2:登出',
       `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
       `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
       `is_delete` INT NOT NULL DEFAULT 0,
       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#学校
CREATE TABLE `schools` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `region` VARCHAR(255) NOT NULL,
    `account` VARCHAR(255) NOT NULL COMMENT '园长',
    `custom_id` VARCHAR(255) COMMENT '自定义id',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_delete` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `u_region_name` (`region`, `name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#校区
CREATE TABLE `regions` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_delete` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#班级
CREATE TABLE `classes` (
     `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
     `class_name` VARCHAR(255) NOT NULL COMMENT '班级名称',
     `school_id` BIGINT(20) NOT NULL COMMENT '所属学校ID',
     `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_delete` INT NOT NULL DEFAULT 0,
     PRIMARY KEY (`id`),
     UNIQUE KEY `u_school_class` (`school_id`, `class_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



#学生
CREATE TABLE `students` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `login_number` BIGINT(20) NOT NULL COMMENT '登录账号',
    `password` VARCHAR(255) NOT NULL DEFAULT '123456' COMMENT '登录密码',
    `student_name` VARCHAR(255) NOT NULL COMMENT '学生姓名',
    `parent_name` VARCHAR(255) COMMENT '家长姓名',
    `phone_number` VARCHAR(20) COMMENT '电话号码',
    `class_id` BIGINT(20) NOT NULL COMMENT '所属班级ID',
    `school_id` BIGINT(20) NOT NULL COMMENT '所属学校ID',
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_delete` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `u_login_number` (`login_number`)
) AUTO_INCREMENT=1000001 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
# 学生积分
CREATE TABLE `student_points` (
     `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
     `student_id`  BIGINT(20) NOT NULL COMMENT '学生id',
     `points`      BIGINT(20) NOT NULL DEFAULT 0,
     `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_delete` INT NOT NULL DEFAULT 0,
     PRIMARY KEY (`id`),
     UNIQUE KEY `u_student` (`student_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


#老师
CREATE TABLE teachers (
     `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
     `login_number` BIGINT(20) NOT NULL COMMENT '登录账号',
     `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '登录密码',
     `teacher_name` VARCHAR(50) NOT NULL COMMENT '老师姓名',
     `phone_number` VARCHAR(20) COMMENT '电话号码',
     `role` tinyint NOT NULL COMMENT '角色,园长1 教师2',
     `school_id` BIGINT(20) NOT NULL COMMENT '所属学校',
     `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_delete` INT NOT NULL DEFAULT 0,
     PRIMARY KEY (`id`),
     UNIQUE KEY `u_login_number` (`login_number`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#老师班级关联表
CREATE TABLE teacher_class_assignments (
     `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
     `teacher_id` BIGINT(20) NOT NULL,
     `class_id` BIGINT(20) NOT NULL,
     `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_delete` INT NOT NULL DEFAULT 0,
     PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

#资源
CREATE TABLE resources (
     `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
     `resource_name` VARCHAR(255) NOT NULL,
     `age_group` VARCHAR(50) NOT NULL,
     `course` VARCHAR(50) NOT NULL,
     `level_1` VARCHAR(50) NOT NULL,
     `level_2` VARCHAR(50) NOT NULL,
     `path` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '可访问路径',
     `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `is_delete` INT NOT NULL DEFAULT 0,
     PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# 排课
CREATE TABLE `schedules` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `resource_id` BIGINT(20)  NOT NULL,
    `school_id` BIGINT(20)  NOT NULL,
    `class_id` BIGINT(20)  NOT NULL,
    `begin_time` DATETIME  NOT NULL,
    `end_time` DATETIME  NOT NULL,
    `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_delete` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# 积分行为
CREATE TABLE `activity_logs` (
   `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
   `student_id` BIGINT(20) NOT NULL,
   `resource_id` BIGINT(20) NOT NULL,
   `activity_date` DATE NOT NULL,
   `points_award` INT NOT NULL DEFAULT 0,
   `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
   `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   `is_delete` INT NOT NULL DEFAULT 0,
   PRIMARY KEY (`id`),
   UNIQUE KEY `u_student_resource_date` (`student_id`, `resource_id`, `activity_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;






