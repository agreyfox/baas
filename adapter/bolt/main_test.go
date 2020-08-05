package bolt

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/agreyfox/baas"
	"github.com/jinzhu/now"

	"github.com/asdine/storm/v3"
	"syreclabs.com/go/faker"

	uuid "github.com/satori/go.uuid"
)

var (
	testApplicationStorage    *ApplicationStorage
	testFileStorage           *FileStorage
	testUsageStorage          *UsageStorage
	testTransformationStorage *TransformStorage

	testUsage                *baas.Usage
	testApplication          *baas.Application
	testDeletableApplication *baas.Application
	testFile                 *baas.File
	testDeletableFile        *baas.File
)

func TestMain(m *testing.M) {
	db, err := Open(os.Getenv("TEST_MAHI_BOLT_DIR"))
	if err != nil {
		log.Println("failed setting up db connection", err)
		return
	}
	defer db.Close()

	setup(db)

	t := m.Run()

	cleanup(db)

	os.Exit(t)
}

func cleanup(db *storm.DB) {
	if err := os.RemoveAll(os.Getenv("TEST_MAHI_BOLT_DIR")); err != nil {
		log.Println("could not remove dir", err)
	}
}

func setup(db *storm.DB) {
	testApplicationStorage = &ApplicationStorage{
		DB: db,
	}

	testFileStorage = &FileStorage{
		DB: db,
	}

	testUsageStorage = &UsageStorage{
		DB: db,
	}

	testTransformationStorage = &TransformStorage{
		DB: db,
	}

	createTestApplication(db)
	testApplication = createTestApplication(db)
	testDeletableApplication = createTestApplication(db)

	createTestFile(db)
	testFile = createTestFile(db)
	testDeletableFile = createTestFile(db)

	testUsage = createTestUsage(db)

}

func createTestUsage(db *storm.DB) *baas.Usage {
	a := usage{
		ID:                    uuid.NewV4().String(),
		ApplicationID:         testApplication.ID,
		Transformations:       10,
		UniqueTransformations: 10,
		Bandwidth:             49494,
		Storage:               23232323,
		FileCount:             12,
		StartDate:             now.BeginningOfDay().Format(baas.DateLayout),
		EndDate:               now.EndOfDay().Add(2 * time.Hour).Format(baas.DateLayout),
	}

	if err := db.Save(&a); err != nil {
		panic(err)
	}

	mahiUsage, err := sanitizeUsage(a)
	if err != nil {
		panic(err)
	}

	return &mahiUsage
}

func createTestApplication(db *storm.DB) *baas.Application {
	a := application{
		ID:               uuid.NewV4().String(),
		Name:             faker.Name().String(),
		Description:      "",
		StorageEngine:    testStorageEngine,
		StorageAccessKey: testStorageAccessKey,
		StorageSecretKey: testStorageSecretKey,
		StorageBucket:    testStorageBucket,
		StorageEndpoint:  testStorageEndpoint,
		StorageRegion:    testStorageRegion,
		DeliveryURL:      testDeliveryURL,
	}

	if err := db.Save(&a); err != nil {
		panic(err)
	}

	mahiApp := sanitizeApp(a)

	return &mahiApp
}

func createTestFile(db *storm.DB) *baas.File {
	n := file{
		ID:            uuid.NewV4().String(),
		ApplicationID: testApplication.ID,
		Filename:      faker.Name().String(),
		FileBlobID:    faker.Name().String(),
		Size:          50,
		MIMEType:      "test",
		MIMEValue:     "test",
		Extension:     "test",
		URL:           faker.Name().String(),
		Hash:          "test",
		Width:         23,
		Height:        60,
	}

	if err := db.Save(&n); err != nil {
		panic(err)
	}

	mahiFile := sanitizeFile(n)

	return &mahiFile
}
