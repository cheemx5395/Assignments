Suggested appropriate indexes:   

1. user_id

2. (user_id, created_at)    

Appropriate indexes can be considered on either column user_id since this one seems like it'll be an integer so we can use logical comparisons.
OR we can use complex index comprising of both columns (user_id, created_at) so it'll also be effective to use for time-logic operations according to the user_id.