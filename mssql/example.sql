
CREATE DATABASE [example]
GO
use [example]
GO

CREATE TABLE [table1] (
    [id] int  IDENTITY(1,1) NOT NULL,
    [name] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS  NULL,
    [json] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS  NULL,
    [ref] int  NULL,
    PRIMARY KEY CLUSTERED ([id])
    WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON)
    ON [PRIMARY]
    )
    ON [PRIMARY]
    GO

CREATE TABLE [table2] (
    [id] int  IDENTITY(1,1) NOT NULL,
    [name] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS  NULL,
    [age] int  NULL,
    [ref] int  NULL,
    PRIMARY KEY CLUSTERED ([id])
    WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON)
    ON [PRIMARY]
    )
    ON [PRIMARY]
    GO

CREATE TABLE [table3] (
    [id] int  IDENTITY(1,1) NOT NULL,
    [name] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS  NULL,
    [info] varchar(255) COLLATE SQL_Latin1_General_CP1_CI_AS  NULL,
    [ref] int  NULL,
    PRIMARY KEY CLUSTERED ([id])
    WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON)
    ON [PRIMARY]
    )
    ON [PRIMARY]
    GO
