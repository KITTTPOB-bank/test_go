
API สำหรับจัดการข้อมูลผู้ป่วย ด้วยระบบ Backend ที่สามารถรันผ่าน Docker Compose และให้บริการ API พร้อม Swagger UI

## Start

เริ่มต้นระบบด้วยคำสั่ง:
```bash
docker-compose up --build

API Documentation
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
    },
    {
      "date_of_birth": "2000-03-21",
      "email": "kazuki.takahashi@example.com",
      "first_name_en": "Kazuki",
      "first_name_th": "คาซุกิ",
      "gender": "M",
      "last_name_en": "Takahashi",
      "last_name_th": "ทาคาฮาชิ",
      "middle_name_en": "Ren",
      "middle_name_th": "เร็น",
      "national_id": "9876543210987",
      "passport_id": "P9876543210987",
      "patient_hn": "HN20230523002",
      "phone_number": "0812345678"
    },
    {
      "date_of_birth": "1995-12-05",
      "email": "aiko.miyazaki@example.com",
      "first_name_en": "Aiko",
      "first_name_th": "ไอโกะ",
      "gender": "F",
      "last_name_en": "Miyazaki",
      "last_name_th": "มิยาซากิ",
      "middle_name_en": "Haru",
      "middle_name_th": "ฮารุ",
      "national_id": "4567891234567",
      "passport_id": "Q4567891234567",
      "patient_hn": "HN20230523003",
      "phone_number": "0898765432"
    }
  ]
}
```
