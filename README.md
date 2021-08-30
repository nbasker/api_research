# api_research
Using gRPC, go and google API recommendations to create services

# PR1 Adding filter for visibile
* Maintained the existing filter of meeting_ids and added additional field for visible
* In the service backend added the filter to the SQL query

# PR2 Adding OrderBy
* Added a new field in the existing request structure to get orderby in the body
* Followed the recommendations provided at https://cloud.google.com/apis/design/design_patterns#sorting_order for specifying one or more orderby fields with asc/desc
* Avoided any syntax check of orderby input as the SQL would fail and throw appropriate error. Could be added to improve better interaction with users.

# PR3 Adding status field
* Decided not add a field in the db as it clearly can be derived.
* Considered adding an additional field just in the response, but it would cause problems with filter, orderby etc.
* Based on reading on SQL decided to use the computed columns with case statement. This works well with filter and orderby. Faced some challenge in specifying time with Now() and then upon reading figured could use CURRENT_TIMESTAMP in sqlite3.

# PR4 Adding a new GET method 
* Followed the recommendations of https://cloud.google.com/apis/design/standard_methods#get and decided to go with races/raceN syntax.
* The syntax check for race/* is already done by the tools and used it.
* Just looked for numeral in the string and used the first number to search the data base. The string like race/race53hello21 would not throw an error but return the record with id 53.
* Error check for atleast one number is done
* Error check for syntactically valid id but not present in database is done.
* Rest of the error check done by sqlite3 is used

# PR5 Adding new service
* Defined the relevant proto files, go-generate commnds, go-mods to create the additional service on port 9001 on top of the existing service on 9000
* Took a shortcut at the back-end of the service to return to event records instead of implementing a db-service.
* Treated this as an objective to add a new-service as the database can be very similar to existing service.
* Another option for consideration is to make database as a separate "service". Define gRPC from racing and sports to talk to the database service. This helps in enhancement, maintainance and scaling of database independent of the racing or sports. 
