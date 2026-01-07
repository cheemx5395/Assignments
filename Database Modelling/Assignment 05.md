1. Soft Deletes rerequired:   

Soft deletees are required in project scenarios where we don't want to permanently delete the data, like remove it's access from the user but don't delete it, in that case we will use a `deleted` field in our table which will be boolean to store if the data is actually deleted or not and then allowing the user access to it or not.   

Social Media deletion is best example for using soft deletes, let's say for example our user want to delete their account, but it was just a mood swing so they can revert their idea and get access t their account back.

2. Hard Deletes:

Hard deletes should be used in case of permanent deletion scnearios, like in the cases where user is not going to want it back ever so they won't have any problems and no issues from the database storage limit cross too, hence permanent deletions or Hard Deletes come into scenario in this cases.

Real world examples are removing past history logs that will not be necessary anymore releasing a lot of memory for usage, and increasing performance speed.