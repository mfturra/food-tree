# food-tree
go language capabilities were explored to create a REST API that connects to a PostgreSQL database that collects the nutrient properties of different food groups and outputs a nutricious grocery list for the upcoming week.

## Workflow Considerations
1. Connection should be established to local PostgreSQL and successful output message should generate when connection is established.
   - If connection is attempted but isn't made to DB, then an error message should be generated.
2. JSON file was read and it's contents were inserted as a new entries into the DB's table.
   - If the JSON payload wasn't successfully entered into the table an error message would be generated. 
4. Two separate workflows were created to account for duplicate entries and to allow for querying the database.
5. DB connection will close after performing the necessary actions.
