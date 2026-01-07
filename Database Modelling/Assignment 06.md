Question:
A query frequently runs:

SELECT * FROM bookings

WHERE doctor_id = ?

AND appointment_date >= ?

AND appointment_date <= ?;

Propose a composite index
Explain column order choice.

Solution:

Composite index that will perfectly in this scenario will be on the columns (doctor_id, appointment_date) so that doctor_id is what we're considering primary filtering criterion and then the dates come into scenario hence using composite of them in the aforementioned order will turn out to be perfect index.