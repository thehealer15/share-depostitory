
create table if not exists investor_details (
    govt_ID text PRIMARY KEY    
);

create table if not exists Holdings (
    ticker text  PRIMARY KEY, 
    quantity int
);
