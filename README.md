parking_lot
====================

Class Design
----
#### Vehicle 
- Maintains type struct Vehicle and it's getter/setter functions.
- /src/vehicle/vehicle.go

#### Slot 
- Maintains type struct Slot and it's getter/setter functions.
- Every Slot has a Vehicle attached to it when occupied
- /src/slot/slot.go

#### ParkingLot 
- Maintains type struct ParkingLot which is a slice of Slots.
- /src/parkingLot/parkingLot.go


Steps to run
----
cd <Project directory>
bin/setup
bin/parking_lot input.txt
bin/parking_lot


Assumptions
----
1. Input arguments are always in correct format. Validations on input are in place
- Length of command line arguments for each command
- Input format checking for create_parking_lot and leave commands. Also, check slotNumber < actual slots for leave command.
- Sequence of commands. eg. Log error if park is called before create_parking_lot
2. Parking lot is just single floor, can be extended to multiple floors
3. Everything is stored in memory, so no ids are generated for slots/parkings


Unit/Functional tests
----
- Tests for parkingLot.go have been added to test all functions with 100% coverage.


Git commit logs
----
Frequent commit logs have been added
```
commit ad59b13a97d191febf33ebd975a9b9baf1d73725 (HEAD -> master)
Author: Karan <karan@Karans-MacBook-Pro.local>
Date:   Thu Jan 31 18:06:02 2019 +0530

    Setup main.go to intake commands and process them

commit c3c287531a3def6d09b88c182cc51b12a3bb0ead
Author: Karan <karan@Karans-MacBook-Pro.local>
Date:   Wed Jan 30 22:35:01 2019 +0530

    Added advanced search functions with tests

commit bdc61cb680e0780f1151174edca1f7ed3dd82c1c
Author: Karan <karan@Karans-MacBook-Pro.local>
Date:   Wed Jan 30 22:05:36 2019 +0530

    Added New,Park,Leave,Status functions and their tests

commit e4aefbbf3581b1f1272fae9cdbfdd923c3c6eeb2
Author: Karan <karan@Karans-MacBook-Pro.local>
Date:   Wed Jan 30 22:04:59 2019 +0530

    Exporting functions and added helper methods

commit 1517fa21ca0c9575eae3bf201e8a5ef782060a7a
Author: Karan <karan@Karans-MacBook-Pro.local>
Date:   Wed Jan 30 19:46:37 2019 +0530

    Adding model structs for vehicle and slot

commit 1142648713fb3da257bb930b78de5978b740583c
Author: Karan <karan@Karans-MacBook-Pro.local>
Date:   Wed Jan 30 18:08:03 2019 +0530
```