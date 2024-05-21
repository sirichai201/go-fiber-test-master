package colltrollers

import (
	"go-fiber-test/database"
	m "go-fiber-test/models"

	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

// // ข้อ5.0****************************************************************
func BasicAuth(c *fiber.Ctx) error {
	return c.SendString("test basicAuth ")
}

// // ข้อ5.1****************************************************************
func Factorial(c *fiber.Ctx) error {
	// รับพารามิเตอร์ "number" จาก URL และแปลงเป็นจำนวนเต็ม
	number, err := strconv.ParseInt(c.Params("number"), 10, 64)
	if err != nil || number < 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid number")
	}

	// 	// คำนวณแฟกทอเรียล
	result := big.NewInt(1)
	for i := int64(2); i <= number; i++ {
		result.Mul(result, big.NewInt(i)) //ช่วยให้สามารถคำนวณค่าคูณของจำนวนเต็มใหญ่ได้
	}

	// 	// ส่งผลลัพธ์เป็น JSON
	return c.JSON(fiber.Map{"number": number, "factorial": result.String()})
}

// // ข้อ5.2****************************************************************

func QueryParam(c *fiber.Ctx) error {
	taxID := c.Query("tax_id")
	var asciiValues []string

	for _, char := range taxID {
		asciiValues = append(asciiValues, strconv.Itoa(int(char)))
	}

	str := "ASCII values: " + strings.Join(asciiValues, " ")
	return c.JSON(fiber.Map{
		"message": str,
	})
}

// // ข้อ5.3 ****************************************************************
func Controller(c *fiber.Ctx) error {
	return c.SendString("controller.TestParams → c.TestParams")
}

// // ข้อ6**********////////////////////////////////////////////////////////////////////////////////////////
func RegisterValidate(c *fiber.Ctx) error {

	Register := new(m.Register)
	if err := c.BodyParser(Register); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	if errors := validate.Struct(Register); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errors.Error(),
		})
	}
	return c.JSON(Register)
}

// // ทดลอง
func RegisterValidate2(c *fiber.Ctx) error {
	// สร้าง instance ใหม่ของ struct Register สำหรับเก็บข้อมูลที่รับมาจาก HTTP request body
	Register2 := new(m.Register2)

	// อ่านและแปลงข้อมูลที่รับมาจาก HTTP request body ให้อยู่ในรูปแบบของ struct Register
	if err := c.BodyParser(Register2); err != nil {
		// หากมีข้อผิดพลาดในการอ่านและแปลงข้อมูล ให้ส่งคำตอบกลับให้ client ด้วยสถานะ HTTP 500 (Internal Server Error) และข้อความของข้อผิดพลาด
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// สร้าง instance ของ validator สำหรับ validate ข้อมูล
	validate := validator.New()

	// Register the custom regexp validation
	// ลงทะเบียนฟังก์ชันการ validate แบบ custom โดยใช้ Regular Expression
	validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		regex := regexp.MustCompile(fl.Param())
		return regex.MatchString(fl.Field().String())
	})

	// 	// ทำการ validate ข้อมูลที่อยู่ใน struct Register
	if errors := validate.Struct(Register2); errors != nil {
		// หากข้อมูลไม่ผ่านการ validate ให้ส่งคำตอบกลับให้ client ด้วยสถานะ HTTP 400 (Bad Request) และข้อความของข้อผิดพลาด
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errors.Error(),
		})
	}

	// 	// หากข้อมูลผ่านการ validate ให้ส่งข้อมูลที่ถูก validate กลับไปยัง client
	return c.JSON(Register2)
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadGateway).SendString("failed to parse")
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	str := p.Name + p.Pass
	return c.JSON(fiber.Map{
		"message": str,
	})
}

func ParamsTest(c *fiber.Ctx) error {
	str := "hello ==> " + c.Params("name")
	return c.JSON(fiber.Map{
		"message": str,
	})
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is  " + a
	return c.JSON(fiber.Map{
		"message": str,
	})
}

func ValidateTest(c *fiber.Ctx) error {

	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	if errors := validate.Struct(user); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errors.Error(),
		})
	}
	return c.JSON(user)
}

//********************************************************************************************************************************
//7.0.1

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var Company []m.Company

	db.Find(&Company) //delelete = null
	return c.Status(200).JSON(Company)
}
func AddCompany(c *fiber.Ctx) error {

	db := database.DBConn
	var Company m.Company

	if err := c.BodyParser(&Company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&Company)
	return c.Status(201).JSON(Company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var Company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&Company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&Company)
	return c.Status(200).JSON(Company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var Company m.Company

	result := db.Delete(&Company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// 7.0.2////////////////////////////////// ****************************************************************************************************************
// *7.0.2 สร้าง api GET ใน group dogs โชว์ข้อมูลที่ถูกลบไปแล้ว ตารางdogs
func IncludeDeleted(db *gorm.DB) *gorm.DB {
	return db.Unscoped().Where("deleted_at IS NOT NULL")
	// โดยเฉพาะมันจะบอก Gorm ให้ไม่สนใจฟิลด์ DeletedAt เมื่อค้นหาฐานข้อมูล
	// และเพิ่มเงื่อนไขกรองเฉพาะเรคอร์ดที่ฟิลด์ DeletedAt ไม่เป็นค่าว่าง
}

func GetDogsDelete(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	// ใช้สโคป IncludeDeleted เพื่อรวมเฉพาะเรคอร์ดที่ถูกลบแบบ soft-delete
	db.Scopes(IncludeDeleted).Find(&dogs)
	return c.Status(200).JSON(dogs)
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 7.1
func DogIDGreater(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ? AND dog_id < ?", 50, 100)
}

func GetDogIDGreater(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogIDGreater).Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DogIDGreaterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogIDGreaterThan100).Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

// 7.2
func GetDogsColor(c *fiber.Ctx) error {
	db := database.DBConn
	SumRed := 0
	SumGreen := 0
	SumPink := 0
	SumNoColor := 0
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			SumRed += 1
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			SumGreen += 1
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			SumPink += 1
		} else {
			typeStr = "no color"
			SumNoColor += 1
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr,

			//green

		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:       dataResults,
		Name:       "golang-testColor",
		Count:      len(dogs), //หาผลรวม,
		SumRed:     SumRed,
		SumGreen:   SumGreen,
		SumPink:    SumPink,
		SumNoColor: SumNoColor,
	}
	return c.Status(200).JSON(r)
}

// workshopGO_lang

func GetProFile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile []m.ProFile

	db.Find(&profile)
	return c.Status(200).JSON(profile)
}

func AddProFile(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var profile m.ProFile

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&profile)
	return c.Status(201).JSON(profile)
}

func UpdateProFile(c *fiber.Ctx) error {
	db := database.DBConn
	var profile m.ProFile
	id := c.Params("id")

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&profile)
	return c.Status(200).JSON(profile)
}

func RemoveProFile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var profile m.ProFile

	result := db.Delete(&profile, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func ProfileSummary(c *fiber.Ctx) error {
	db := database.DBConn
	SumGenZ := 0
	SumGenY := 0
	SumGenX := 0
	SumBabyBoomer := 0
	SumGIGeneration := 0
	SumNoAge := 0
	var profiles []m.ProFile

	db.Find(&profiles)

	var dataResults []m.ProfileRes
	for _, v := range profiles {
		typeStr := ""
		if v.Age >= 1 && v.Age < 24 {
			typeStr = "SumGenZ"
			SumGenZ += 1
		} else if v.Age >= 24 && v.Age <= 41 {
			typeStr = "SumGenY"
			SumGenY += 1
		} else if v.Age >= 42 && v.Age <= 56 {
			typeStr = "SumGenX"
			SumGenX += 1
		} else if v.Age >= 57 && v.Age <= 75 {
			typeStr = "SumBabyBoomer"
			SumBabyBoomer += 1
		} else if v.Age > 75 {
			typeStr = "SumGIGeneration"
			SumGIGeneration += 1
		} else {
			typeStr = "No Age"
			SumNoAge += 1
		}

		d := m.ProfileRes{
			Name:      v.Name,
			ProfileID: int(v.ID),
			Type:      typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultProfileData{
		Data:            dataResults,
		Name:            "golang-testSummary",
		Count:           len(profiles),
		SumGenZ:         SumGenZ,
		SumGenY:         SumGenY,
		SumGenX:         SumGenX,
		SumBabyBoomer:   SumBabyBoomer,
		SumGIGeneration: SumGIGeneration,
	}
	return c.Status(200).JSON(r)
}

func SearchProfile(c *fiber.Ctx) error {
	db := database.DBConn
	search := c.Query("search")
	var profile []m.ProFile
	log.Print("--->" + search)
	db.Where("employee_id = ? OR age = ? OR name LIKE ? OR last_name LIKE ?", search, search, "%"+search+"%", "%"+search+"%").Find(&profile)
	count := len(profile)
	response := struct {
		Profiles []m.ProFile `json:"profiles"`
		Count    int         `json:"count"`
	}{
		Profiles: profile,
		Count:    count,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
