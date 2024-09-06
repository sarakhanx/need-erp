# NEED ERP | RESOURCE | LOG
---

### FAKE DATA FOR DEVELOPMENT 📜
> ข้อมูลทั้งหมดตรงนี้ คือข้อมูลสำหรับ Development เท่านั้น ข
> ข้อมูลที่ใช้งานใน Production กรุณาดูที่ Obsidian ใน Vault ของ Need-Shopping
.
.
##### Backend ENV 🚀
- ENV ที่ใช้สำหรับ Go Backend เท่านั้น
```bash
# This is for local
#i don't wanna make the new data in database because i still working with it 'Code first' not finish yet
DB_URI = postgresql://admin:admin@localhost:5433/whalewks
# This is for docker (Uncomment it when development or debug with docker)
# DB_URI = postgresql://admin:admin@database:5432/whalewks
# This will use with JWT but i will consider it again what should i do with the secret keys
SECRET_KEYS = secretkeys
```
.
.
##### Docker Compose ENV 🐳
- ENV ที่ใช้สำหรับขั้นตอน Docker-Compose เพื่อ Debug
```bash
#Use for env that using with docker-compose
PGADMIN_LISTEN_PORT = 5001
```
.
.

---
### NEED ERP | DEVELOPER NOTE 🗂️
> ทุกๆการพัฒนาในส่วนของ Backend Monolith จะถูกบันทึกลงที่ตรงนี้

---
|date | title | developer | period | score(full :10)|NOTE|
|---|---|---|---|---|---|
|20aug24|purchasing's and warehouse's routers & controller development. Docker File Changed COMMIT: `:construction:pur and wh controller` |amp | 3week|3:10|work too long than expected
|21aug24|...purchasing order and sale_order service...|...amp...|...time...|...score...|...note...|
|29aug24|auth on roles|amp|5H|7:10|Cursor help to make auth|
|3sep24|products with query categ and pagination|amp|2D|9:10|...|
|6sep24|Docker compose and go docker images refactoring|amp|1D|7:10|try to make it easy when need to migrate to microservice|
| | | | | | | | | | | |
| | | | | | | | | | | |
|---|---|---|---|---|---|---|---|---|---|---|
---
