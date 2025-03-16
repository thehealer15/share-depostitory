
create schema if not exists platform ;
set search_path to platform; 

create table if not exists companies (
    ticker text  ,
    face_value int , 
    company_name text ,
    PRIMARY KEY (ticker)
);

create table if not exists investor (
    investor_name text  ,
    govt_ID text ,
    PRIMARY KEY (govt_ID)
);


INSERT INTO platform.companies (ticker, face_value, company_name) VALUES
('MSFT', 200, 'Microsoft Inc.'), 
('AAPL', 150, 'Apple Inc.'),
('GOOGL', 300, 'Google LLC'),
('AMZN', 250, 'Amazon Inc.'),
('TSLA', 180, 'Tesla Inc.'),
('META', 220, 'Meta Platforms'),
('NVDA', 270, 'NVIDIA Corp'),
('INTC', 190, 'Intel Corp'),
('CSCO', 210, 'Cisco Systems'),
('ORCL', 230, 'Oracle Corp');