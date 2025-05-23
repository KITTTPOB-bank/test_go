
API สำหรับจัดการข้อมูลผู้ป่วย ด้วยระบบ Backend ที่สามารถรันผ่าน Docker Compose และให้บริการ API พร้อม Swagger UI

## Start

เริ่มต้นระบบด้วยคำสั่ง:
```bash
docker-compose up --build
📚 API Documentation
หลังจากเริ่มต้นระบบแล้วสามารถเข้าถึง API docs ได้ที่:

👉 http://localhost:8010/swagger/index.html

🧪 Data Mock Endpoint
POST /staff/patients
Mock endpoint สำหรับเพิ่มข้อมูลผู้ป่วยหลายรายการ

URL:
http://localhost:8010/staff/patients

Method:
POST
Request Body (JSON):

json
 
{
  "data": [
    {
      "date_of_birth": "1998-07-15",
      "email": "sakura.hanami@example.com",
      "first_name_en": "Sakura",
      "first_name_th": "ซากุระ",
      "gender": "F",
      "last_name_en": "Hanami",
      "last_name_th": "ฮานามิ",
      "middle_name_en": "Yuki",
      "middle_name_th": "ยูคิ",
      "national_id": "1234567890123",
      "passport_id": "N1234567890123",
      "patient_hn": "HN20230523001",
      "phone_number": "0801234567"
    }
    // ... ข้อมูลผู้ป่วยเพิ่มเติม
  ]
}
```
