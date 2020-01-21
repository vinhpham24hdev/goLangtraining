-- Create a new table called 'Envelope' in schema 'SchemaName'
-- Drop the table if it already exists
IF OBJECT_ID('SchemaName.Envelope', 'U') IS NOT NULL
DROP TABLE SchemaName.Envelope
GO
-- Create the table in the specified schema
CREATE TABLE SchemaName.Envelope
(
    Dates [NVARCHAR](50) NOT NULL PRIMARY KEY, -- primary key column
    Currency [NVARCHAR](50) NOT NULL,
    Rate [NVARCHAR](50) NOT NULL
    -- specify more columns here
);
GO