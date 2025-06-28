package seeders

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"restaurant_os/internal/models"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) SeedAll() error {

	log.Println("Starting database seeding...")

	// Seed in order due to foreign key constraints
	if err := s.seedRestaurants(); err != nil {
		return err
	}
	if err := s.seedUsers(); err != nil {
		return err
	}
	if err := s.seedBranches(); err != nil {
		return err
	}
	if err := s.seedSuppliers(); err != nil {
		return err
	}
	if err := s.seedCustomers(); err != nil {
		return err
	}
	if err := s.seedTables(); err != nil {
		return err
	}
	if err := s.seedMenuCategories(); err != nil {
		return err
	}
	if err := s.seedMenuItems(); err != nil {
		return err
	}
	if err := s.seedInventory(); err != nil {
		return err
	}
	if err := s.seedReservations(); err != nil {
		return err
	}
	if err := s.seedQRSessions(); err != nil {
		return err
	}
	if err := s.seedOrders(); err != nil {
		return err
	}
	if err := s.seedNotifications(); err != nil {
		return err
	}

	log.Println("Database seeding completed successfully!")
	return nil
}

func (s *Seeder) seedRestaurants() error {
	log.Println("Seeding restaurants...")

	restaurants := []models.Restaurant{
		{
			ID:        1,
			Name:      "The Golden Spoon",
			Email:     "info@goldenspoon.com",
			Phone:     "+91-9876543210",
			Address:   "MG Road, Kochi, Kerala, India",
			Website:   "https://goldenspoon.com",
			TaxNumber: "32AABCU9603R1ZM",
			Currency:  "INR",
			TimeZone:  "Asia/Kolkata",
			IsActive:  true,
		},
		{
			ID:        2,
			Name:      "Spice Garden Restaurant",
			Email:     "contact@spicegarden.in",
			Phone:     "+91-9876543211",
			Address:   "Marine Drive, Kochi, Kerala, India",
			Website:   "https://spicegarden.in",
			TaxNumber: "32AABCU9603R1ZN",
			Currency:  "INR",
			TimeZone:  "Asia/Kolkata",
			IsActive:  true,
		},
	}

	return s.db.Create(&restaurants).Error
}

func (s *Seeder) seedUsers() error {
	log.Println("Seeding users...")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	users := []models.User{
		{
			ID:           1,
			Name:         "Admin User",
			Email:        "admin@goldenspoon.com",
			Password:     string(hashedPassword),
			Phone:        "+91-9876543210",
			UserType:     models.UserTypeSuperAdmin,
			RestaurantID: nil,
			BranchID:     nil,
			IsActive:     true,
		},
		{
			ID:           2,
			Name:         "Restaurant Owner",
			Email:        "owner@goldenspoon.com",
			Password:     string(hashedPassword),
			Phone:        "+91-9876543211",
			UserType:     models.UserTypeRestaurant,
			RestaurantID: uintPtr(1),
			BranchID:     nil,
			IsActive:     true,
		},
		{
			ID:           3,
			Name:         "John Doe",
			Email:        "john.doe@goldenspoon.com",
			Password:     string(hashedPassword),
			Phone:        "+91-9876543212",
			UserType:     models.UserTypeEmployee,
			Role:         rolePtr(models.RoleManager),
			RestaurantID: uintPtr(1),
			BranchID:     uintPtr(1),
			IsActive:     true,
		},
		{
			ID:           4,
			Name:         "Sarah Wilson",
			Email:        "sarah.wilson@goldenspoon.com",
			Password:     string(hashedPassword),
			Phone:        "+91-9876543213",
			UserType:     models.UserTypeEmployee,
			Role:         rolePtr(models.RoleWaiter),
			RestaurantID: uintPtr(1),
			BranchID:     uintPtr(1),
			IsActive:     true,
		},
		{
			ID:           5,
			Name:         "Chef Ravi",
			Email:        "ravi.chef@goldenspoon.com",
			Password:     string(hashedPassword),
			Phone:        "+91-9876543214",
			UserType:     models.UserTypeEmployee,
			Role:         rolePtr(models.RoleChef),
			RestaurantID: uintPtr(1),
			BranchID:     uintPtr(1),
			IsActive:     true,
		},
	}

	return s.db.Create(&users).Error
}

func (s *Seeder) seedBranches() error {
	log.Println("Seeding branches...")

	openingHours := map[string]interface{}{
		"monday":    map[string]string{"open": "09:00", "close": "22:00"},
		"tuesday":   map[string]string{"open": "09:00", "close": "22:00"},
		"wednesday": map[string]string{"open": "09:00", "close": "22:00"},
		"thursday":  map[string]string{"open": "09:00", "close": "22:00"},
		"friday":    map[string]string{"open": "09:00", "close": "23:00"},
		"saturday":  map[string]string{"open": "09:00", "close": "23:00"},
		"sunday":    map[string]string{"open": "10:00", "close": "22:00"},
	}
	hoursJSON, _ := json.Marshal(openingHours)

	branches := []models.Branch{
		{
			ID:           1,
			RestaurantID: 1,
			Name:         "Main Branch - MG Road",
			Location:     "MG Road, Kochi, Kerala",
			Phone:        "+91-9876543220",
			Email:        "mgroad@goldenspoon.com",
			ManagerID:    uintPtr(3),
			IsActive:     true,
			OpeningHours: string(hoursJSON),
		},
		{
			ID:           2,
			RestaurantID: 1,
			Name:         "Marine Drive Branch",
			Location:     "Marine Drive, Kochi, Kerala",
			Phone:        "+91-9876543221",
			Email:        "marinedrive@goldenspoon.com",
			ManagerID:    uintPtr(3),
			IsActive:     true,
			OpeningHours: string(hoursJSON),
		},
	}

	return s.db.Create(&branches).Error
}

func (s *Seeder) seedSuppliers() error {
	log.Println("Seeding suppliers...")

	suppliers := []models.Supplier{
		{
			ID:       1,
			Name:     "Fresh Vegetables Supplier",
			Contact:  "Rajesh Kumar",
			Phone:    "+91-9876543230",
			Email:    "rajesh@freshveggies.com",
			Address:  "Vegetable Market, Ernakulam",
			IsActive: true,
		},
		{
			ID:       2,
			Name:     "Meat & Seafood Supplier",
			Contact:  "Mohammed Ali",
			Phone:    "+91-9876543231",
			Email:    "ali@meatseafood.com",
			Address:  "Fish Market, Fort Kochi",
			IsActive: true,
		},
		{
			ID:       3,
			Name:     "Dairy Products Supplier",
			Contact:  "Priya Nair",
			Phone:    "+91-9876543232",
			Email:    "priya@dairyproducts.com",
			Address:  "Dairy Farm, Thrissur",
			IsActive: true,
		},
	}

	return s.db.Create(&suppliers).Error
}

func (s *Seeder) seedCustomers() error {
	log.Println("Seeding customers...")

	customers := []models.Customer{
		{
			ID:            1,
			Name:          "Amit Sharma",
			Phone:         "+91-9876543240",
			Email:         "amit.sharma@email.com",
			Address:       "Kakkanad, Kochi",
			BirthDate:     timePtr(time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC)),
			TotalOrders:   15,
			TotalSpent:    25000.50,
			LoyaltyPoints: 250,
		},
		{
			ID:            2,
			Name:          "Priya Menon",
			Phone:         "+91-9876543241",
			Email:         "priya.menon@email.com",
			Address:       "Edapally, Kochi",
			BirthDate:     timePtr(time.Date(1985, 8, 22, 0, 0, 0, 0, time.UTC)),
			Anniversary:   timePtr(time.Date(2015, 12, 10, 0, 0, 0, 0, time.UTC)),
			TotalOrders:   8,
			TotalSpent:    12000.00,
			LoyaltyPoints: 120,
		},
		{
			ID:            3,
			Name:          "Rajesh Kumar",
			Phone:         "+91-9876543242",
			Email:         "rajesh.kumar@email.com",
			Address:       "Panampilly Nagar, Kochi",
			BirthDate:     timePtr(time.Date(1982, 3, 8, 0, 0, 0, 0, time.UTC)),
			TotalOrders:   22,
			TotalSpent:    35000.75,
			LoyaltyPoints: 350,
		},
	}

	return s.db.Create(&customers).Error
}

func (s *Seeder) seedTables() error {
	log.Println("Seeding tables...")

	var tables []models.Table

	// Branch 1 tables
	for i := 1; i <= 15; i++ {
		table := models.Table{
			ID:         uint(i),
			Number:     fmt.Sprintf("T%02d", i),
			BranchID:   1,
			Capacity:   2 + rand.Intn(6), // 2-7 capacity
			Status:     models.TableAvailable,
			Location:   getTableLocation(i),
			QRCode:     fmt.Sprintf("QR_%d_%02d", 1, i),
			QRToken:    fmt.Sprintf("TOKEN_%d_%02d_%d", 1, i, time.Now().Unix()),
			QRMenuURL:  fmt.Sprintf("https://goldenspoon.com/menu/qr/%d/%02d", 1, i),
			IsQRActive: true,
		}
		tables = append(tables, table)
	}

	// Branch 2 tables
	for i := 16; i <= 25; i++ {
		table := models.Table{
			ID:         uint(i),
			Number:     fmt.Sprintf("T%02d", i-15),
			BranchID:   2,
			Capacity:   2 + rand.Intn(6),
			Status:     models.TableAvailable,
			Location:   getTableLocation(i - 15),
			QRCode:     fmt.Sprintf("QR_%d_%02d", 2, i-15),
			QRToken:    fmt.Sprintf("TOKEN_%d_%02d_%d", 2, i-15, time.Now().Unix()),
			QRMenuURL:  fmt.Sprintf("https://goldenspoon.com/menu/qr/%d/%02d", 2, i-15),
			IsQRActive: true,
		}
		tables = append(tables, table)
	}

	return s.db.Create(&tables).Error
}

func (s *Seeder) seedMenuCategories() error {
	log.Println("Seeding menu categories...")

	categories := []models.MenuCategory{
		{
			ID:          1,
			BranchID:    1,
			Name:        "Appetizers",
			Description: "Start your meal with our delicious appetizers",
			SortOrder:   1,
			IsActive:    true,
		},
		{
			ID:          2,
			BranchID:    1,
			Name:        "Main Course",
			Description: "Our signature main course dishes",
			SortOrder:   2,
			IsActive:    true,
		},
		{
			ID:          3,
			BranchID:    1,
			Name:        "Desserts",
			Description: "Sweet endings to your perfect meal",
			SortOrder:   3,
			IsActive:    true,
		},
		{
			ID:          4,
			BranchID:    1,
			Name:        "Beverages",
			Description: "Refreshing drinks and beverages",
			SortOrder:   4,
			IsActive:    true,
		},
		{
			ID:          5,
			BranchID:    1,
			Name:        "Indian Curries",
			Description: "Traditional Indian curry dishes",
			SortOrder:   5,
			IsActive:    true,
		},
	}

	return s.db.Create(&categories).Error
}

func (s *Seeder) seedMenuItems() error {
	log.Println("Seeding menu items...")

	menuItems := []models.MenuItem{
		// Appetizers
		{
			ID:           1,
			BranchID:     1,
			CategoryID:   uintPtr(1),
			Name:         "Chicken Tikka",
			Description:  "Marinated chicken pieces grilled to perfection",
			Price:        280.00,
			CostPrice:    150.00,
			Available:    true,
			IsVegetarian: false,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    3,
			PrepTime:     20,
			ImageURL:     "https://example.com/images/chicken-tikka.jpg",
			Ingredients:  "Chicken, Yogurt, Spices, Lemon",
			Allergens:    "Dairy",
			SortOrder:    1,
		},
		{
			ID:           2,
			BranchID:     1,
			CategoryID:   uintPtr(1),
			Name:         "Vegetable Samosa",
			Description:  "Crispy fried pastry with spiced vegetable filling",
			Price:        120.00,
			CostPrice:    60.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      true,
			IsGlutenFree: false,
			Spiciness:    2,
			PrepTime:     15,
			ImageURL:     "https://example.com/images/samosa.jpg",
			Ingredients:  "Flour, Potatoes, Peas, Spices",
			Allergens:    "Gluten",
			SortOrder:    2,
		},
		// Main Course
		{
			ID:           3,
			BranchID:     1,
			CategoryID:   uintPtr(2),
			Name:         "Butter Chicken",
			Description:  "Creamy tomato-based curry with tender chicken",
			Price:        450.00,
			CostPrice:    250.00,
			Available:    true,
			IsVegetarian: false,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    2,
			PrepTime:     25,
			ImageURL:     "https://example.com/images/butter-chicken.jpg",
			Ingredients:  "Chicken, Tomatoes, Cream, Butter, Spices",
			Allergens:    "Dairy",
			SortOrder:    1,
		},
		{
			ID:           4,
			BranchID:     1,
			CategoryID:   uintPtr(2),
			Name:         "Paneer Makhani",
			Description:  "Rich and creamy cottage cheese curry",
			Price:        380.00,
			CostPrice:    200.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    2,
			PrepTime:     20,
			ImageURL:     "https://example.com/images/paneer-makhani.jpg",
			Ingredients:  "Paneer, Tomatoes, Cream, Cashews, Spices",
			Allergens:    "Dairy, Nuts",
			SortOrder:    2,
		},
		// Desserts
		{
			ID:           5,
			BranchID:     1,
			CategoryID:   uintPtr(3),
			Name:         "Gulab Jamun",
			Description:  "Soft milk dumplings in sweet syrup",
			Price:        150.00,
			CostPrice:    75.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      false,
			IsGlutenFree: false,
			Spiciness:    0,
			PrepTime:     10,
			ImageURL:     "https://example.com/images/gulab-jamun.jpg",
			Ingredients:  "Milk powder, Sugar, Cardamom",
			Allergens:    "Dairy, Gluten",
			SortOrder:    1,
		},
		// Beverages
		{
			ID:           6,
			BranchID:     1,
			CategoryID:   uintPtr(4),
			Name:         "Mango Lassi",
			Description:  "Refreshing yogurt drink with mango",
			Price:        120.00,
			CostPrice:    50.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    0,
			PrepTime:     5,
			ImageURL:     "https://example.com/images/mango-lassi.jpg",
			Ingredients:  "Yogurt, Mango, Sugar",
			Allergens:    "Dairy",
			SortOrder:    1,
		},
		{
			ID:           7,
			BranchID:     1,
			CategoryID:   uintPtr(4),
			Name:         "Masala Chai",
			Description:  "Traditional spiced tea",
			Price:        60.00,
			CostPrice:    20.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    1,
			PrepTime:     5,
			ImageURL:     "https://example.com/images/masala-chai.jpg",
			Ingredients:  "Tea, Milk, Spices, Sugar",
			Allergens:    "Dairy",
			SortOrder:    2,
		},
	}

	return s.db.Create(&menuItems).Error
}

func (s *Seeder) seedInventory() error {
	log.Println("Seeding inventory...")

	inventory := []models.Inventory{
		{
			ID:           1,
			BranchID:     1,
			ItemName:     "Chicken Breast",
			ItemCode:     "CHK001",
			Category:     "Meat",
			Unit:         models.UnitKG,
			CurrentStock: 25.5,
			ReorderLevel: 10.0,
			MaxLevel:     50.0,
			UnitCost:     280.00,
			SupplierName: "Meat & Seafood Supplier",
			LastOrdered:  timePtr(time.Now().AddDate(0, 0, -5)),
			ExpiryDate:   timePtr(time.Now().AddDate(0, 0, 3)),
		},
		{
			ID:           2,
			BranchID:     1,
			ItemName:     "Basmati Rice",
			ItemCode:     "RIC001",
			Category:     "Grains",
			Unit:         models.UnitKG,
			CurrentStock: 100.0,
			ReorderLevel: 20.0,
			MaxLevel:     200.0,
			UnitCost:     120.00,
			SupplierName: "Grains Supplier",
			LastOrdered:  timePtr(time.Now().AddDate(0, 0, -10)),
		},
		{
			ID:           3,
			BranchID:     1,
			ItemName:     "Fresh Milk",
			ItemCode:     "MIL001",
			Category:     "Dairy",
			Unit:         models.UnitLiter,
			CurrentStock: 50.0,
			ReorderLevel: 20.0,
			MaxLevel:     100.0,
			UnitCost:     65.00,
			SupplierName: "Dairy Products Supplier",
			LastOrdered:  timePtr(time.Now().AddDate(0, 0, -2)),
			ExpiryDate:   timePtr(time.Now().AddDate(0, 0, 2)),
		},
		{
			ID:           4,
			BranchID:     1,
			ItemName:     "Tomatoes",
			ItemCode:     "VEG001",
			Category:     "Vegetables",
			Unit:         models.UnitKG,
			CurrentStock: 15.0,
			ReorderLevel: 5.0,
			MaxLevel:     30.0,
			UnitCost:     40.00,
			SupplierName: "Fresh Vegetables Supplier",
			LastOrdered:  timePtr(time.Now().AddDate(0, 0, -3)),
			ExpiryDate:   timePtr(time.Now().AddDate(0, 0, 2)),
		},
		{
			ID:           5,
			BranchID:     1,
			ItemName:     "Paneer",
			ItemCode:     "PAN001",
			Category:     "Dairy",
			Unit:         models.UnitKG,
			CurrentStock: 8.0,
			ReorderLevel: 3.0,
			MaxLevel:     15.0,
			UnitCost:     350.00,
			SupplierName: "Dairy Products Supplier",
			LastOrdered:  timePtr(time.Now().AddDate(0, 0, -1)),
			ExpiryDate:   timePtr(time.Now().AddDate(0, 0, 3)),
		},
	}

	return s.db.Create(&inventory).Error
}

func (s *Seeder) seedReservations() error {
	log.Println("Seeding reservations...")

	reservations := []models.Reservation{
		{
			ID:            1,
			BranchID:      1,
			TableID:       uintPtr(1),
			CustomerName:  "Amit Sharma",
			CustomerPhone: "+91-9876543240",
			CustomerEmail: "amit.sharma@email.com",
			GuestCount:    4,
			ReservedDate:  time.Now().AddDate(0, 0, 1),
			ReservedTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 19, 30, 0, 0, time.Local),
			Duration:      120,
			Status:        models.ReservationConfirmed,
			Notes:         "Anniversary dinner",
			CreatedBy:     uintPtr(4),
		},
		{
			ID:            2,
			BranchID:      1,
			TableID:       uintPtr(5),
			CustomerName:  "Priya Menon",
			CustomerPhone: "+91-9876543241",
			CustomerEmail: "priya.menon@email.com",
			GuestCount:    2,
			ReservedDate:  time.Now().AddDate(0, 0, 2),
			ReservedTime:  time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+2, 20, 0, 0, 0, time.Local),
			Duration:      90,
			Status:        models.ReservationPending,
			Notes:         "Window side table preferred",
			CreatedBy:     uintPtr(4),
		},
	}

	return s.db.Create(&reservations).Error
}

func (s *Seeder) seedQRSessions() error {
	log.Println("Seeding QR sessions...")

	sessions := []models.QRSession{
		{
			ID:             1,
			SessionToken:   "SESSION_TOKEN_123456789",
			TableID:        1,
			BranchID:       1,
			CustomerName:   "Walk-in Customer",
			CustomerPhone:  "+91-9876543250",
			GuestCount:     2,
			Status:         models.QRSessionActive,
			StartedAt:      time.Now().Add(-30 * time.Minute),
			LastActivityAt: time.Now().Add(-5 * time.Minute),
			ExpiresAt:      time.Now().Add(2 * time.Hour),
			DeviceInfo:     `{"browser": "Chrome", "os": "Android"}`,
			IPAddress:      "192.168.1.100",
		},
		{
			ID:             2,
			SessionToken:   "SESSION_TOKEN_987654321",
			TableID:        3,
			BranchID:       1,
			CustomerName:   "Tech Visitor",
			CustomerPhone:  "+91-9876543251",
			GuestCount:     1,
			Status:         models.QRSessionCompleted,
			StartedAt:      time.Now().Add(-2 * time.Hour),
			LastActivityAt: time.Now().Add(-1 * time.Hour),
			ExpiresAt:      time.Now().Add(1 * time.Hour),
			DeviceInfo:     `{"browser": "Safari", "os": "iOS"}`,
			IPAddress:      "192.168.1.101",
		},
	}

	return s.db.Create(&sessions).Error
}

func (s *Seeder) seedOrders() error {
	log.Println("Seeding orders...")

	orders := []models.Order{
		{
			ID:               1,
			OrderNumber:      "ORD-2025-001",
			TableID:          uintPtr(1),
			BranchID:         1,
			UserID:           uintPtr(4),
			CustomerName:     "Walk-in Customer",
			CustomerPhone:    "+91-9876543250",
			OrderType:        models.OrderTypeDineIn,
			OrderSource:      models.OrderSourceQR,
			Status:           models.OrderPreparing,
			PaymentStatus:    models.PaymentPending,
			QRSessionID:      uintPtr(1),
			IsQROrder:        true,
			Subtotal:         830.00,
			TaxAmount:        149.40,
			DiscountAmount:   0.00,
			ServiceCharge:    41.50,
			Total:            1020.90,
			Notes:            "Less spicy please",
			EstimatedTime:    25,
			AssignedWaiterID: uintPtr(4),
		},
		{
			ID:            2,
			OrderNumber:   "ORD-2025-002",
			TableID:       uintPtr(3),
			BranchID:      1,
			UserID:        uintPtr(4),
			CustomerName:  "Rajesh Kumar",
			CustomerPhone: "+91-9876543242",
			OrderType:     models.OrderTypeDineIn,
			OrderSource:   models.OrderSourceStaff,
			Status:        models.OrderCompleted,
			PaymentStatus: models.PaymentPaid,
			IsQROrder:     false,
			Subtotal:      950.00,
			TaxAmount:     171.00,
			Total:         1121.00,
			Notes:         "Regular order",
			EstimatedTime: 30,
		},
	}

	if err := s.db.Create(&orders).Error; err != nil {
		return err
	}

	// Seed order items
	orderItems := []models.OrderItem{
		{
			ID:         1,
			OrderID:    1,
			MenuItemID: 3, // Butter Chicken
			Quantity:   1,
			UnitPrice:  450.00,
			TotalPrice: 450.00,
			Status:     models.OrderItemPreparing,
			Notes:      "Less spicy",
		},
		{
			ID:         2,
			OrderID:    1,
			MenuItemID: 1, // Chicken Tikka
			Quantity:   1,
			UnitPrice:  280.00,
			TotalPrice: 280.00,
			Status:     models.OrderItemReady,
		},
		{
			ID:         3,
			OrderID:    1,
			MenuItemID: 6, // Mango Lassi
			Quantity:   1,
			UnitPrice:  120.00,
			TotalPrice: 120.00,
			Status:     models.OrderItemServed,
		},
		{
			ID:         4,
			OrderID:    2,
			MenuItemID: 4, // Paneer Makhani
			Quantity:   1,
			UnitPrice:  380.00,
			TotalPrice: 380.00,
			Status:     models.OrderItemServed,
		},
		{
			ID:         5,
			OrderID:    2,
			MenuItemID: 2, // Vegetable Samosa
			Quantity:   2,
			UnitPrice:  120.00,
			TotalPrice: 240.00,
			Status:     models.OrderItemServed,
		},
		{
			ID:         6,
			OrderID:    2,
			MenuItemID: 5, // Gulab Jamun
			Quantity:   2,
			UnitPrice:  150.00,
			TotalPrice: 300.00,
			Status:     models.OrderItemServed,
		},
	}

	if err := s.db.Create(&orderItems).Error; err != nil {
		return err
	}

	// Seed payments
	payments := []models.Payment{
		{
			ID:            1,
			OrderID:       2,
			Amount:        1121.00,
			Method:        models.PaymentUPI,
			Status:        models.PaymentPaid,
			TransactionID: "UPI123456789",
			Reference:     "Payment via Google Pay",
			ProcessedBy:   uintPtr(4),
		},
		{
			ID:            2,
			OrderID:       1,
			Amount:        500.00,
			Method:        models.PaymentCard,
			Status:        models.PaymentPending,
			TransactionID: "CARD987654321",
			Reference:     "Partial payment",
			ProcessedBy:   uintPtr(4),
		},
	}

	return s.db.Create(&payments).Error
}

func (s *Seeder) seedNotifications() error {
	log.Println("Seeding notifications...")

	notifications := []models.Notification{
		{
			ID:       1,
			BranchID: 1,
			UserID:   uintPtr(5), // Chef Ravi
			Type:     models.NotificationNewOrder,
			Status:   models.NotificationSent,
			Title:    "New Order Received",
			Message:  "Order #ORD-2025-001 has been placed. Table T01 - Butter Chicken, Chicken Tikka",
			Data:     `{"order_id": 1, "table_number": "T01", "items_count": 3}`,
			OrderID:  uintPtr(1),
			TableID:  uintPtr(1),
			SentAt:   timePtr(time.Now().Add(-10 * time.Minute)),
		},
		{
			ID:       2,
			BranchID: 1,
			UserID:   uintPtr(4), // Sarah Wilson (Waiter)
			Type:     models.NotificationOrderReady,
			Status:   models.NotificationRead,
			Title:    "Order Ready for Service",
			Message:  "Order #ORD-2025-001 is ready to be served to Table T01",
			Data:     `{"order_id": 1, "table_number": "T01"}`,
			OrderID:  uintPtr(1),
			TableID:  uintPtr(1),
			ReadAt:   timePtr(time.Now().Add(-5 * time.Minute)),
			SentAt:   timePtr(time.Now().Add(-8 * time.Minute)),
		},
		{
			ID:       3,
			BranchID: 1,
			UserID:   uintPtr(3), // Manager
			Type:     models.NotificationInventoryLow,
			Status:   models.NotificationPending,
			Title:    "Low Inventory Alert",
			Message:  "Tomatoes stock is running low. Current: 15kg, Reorder level: 5kg",
			Data:     `{"item_name": "Tomatoes", "current_stock": 15, "reorder_level": 5}`,
		},
		{
			ID:          4,
			BranchID:    1,
			UserID:      uintPtr(4), // Waiter
			Type:        models.NotificationQRSession,
			Status:      models.NotificationSent,
			Title:       "New QR Session Started",
			Message:     "Customer started QR session at Table T01",
			Data:        `{"table_number": "T01", "guest_count": 2}`,
			TableID:     uintPtr(1),
			QRSessionID: uintPtr(1),
			SentAt:      timePtr(time.Now().Add(-30 * time.Minute)),
		},
		{
			ID:       5,
			BranchID: 1,
			UserID:   uintPtr(3), // Manager
			Type:     models.NotificationReservation,
			Status:   models.NotificationSent,
			Title:    "New Reservation",
			Message:  "Reservation confirmed for Amit Sharma - 4 guests, tomorrow 7:30 PM",
			Data:     `{"customer_name": "Amit Sharma", "guest_count": 4, "date": "2025-06-29", "time": "19:30"}`,
			SentAt:   timePtr(time.Now().Add(-2 * time.Hour)),
		},
	}

	return s.db.Create(&notifications).Error
}

// Seed QR Cart Items
func (s *Seeder) seedQRCartItems() error {
	log.Println("Seeding QR cart items...")

	cartItems := []models.QRCartItem{
		{
			ID:          1,
			QRSessionID: 1,
			MenuItemID:  7, // Masala Chai
			Quantity:    2,
			UnitPrice:   60.00,
			TotalPrice:  120.00,
			Notes:       "Extra sugar",
			AddedAt:     time.Now().Add(-15 * time.Minute),
		},
		{
			ID:          2,
			QRSessionID: 1,
			MenuItemID:  2, // Vegetable Samosa
			Quantity:    1,
			UnitPrice:   120.00,
			TotalPrice:  120.00,
			Notes:       "Extra chutney",
			AddedAt:     time.Now().Add(-10 * time.Minute),
		},
	}

	return s.db.Create(&cartItems).Error
}

// Seed QR Code Scans for analytics
func (s *Seeder) seedQRCodeScans() error {
	log.Println("Seeding QR code scans...")

	scans := []models.QRCodeScan{
		{
			ID:                 1,
			TableID:            1,
			BranchID:           1,
			QRSessionID:        uintPtr(1),
			IPAddress:          "192.168.1.100",
			UserAgent:          "Mozilla/5.0 (Linux; Android 10; SM-G973F) AppleWebKit/537.36",
			DeviceType:         "mobile",
			ScanTime:           time.Now().Add(-35 * time.Minute),
			ConvertedToSession: true,
			ConvertedToOrder:   true,
		},
		{
			ID:                 2,
			TableID:            2,
			BranchID:           1,
			IPAddress:          "192.168.1.102",
			UserAgent:          "Mozilla/5.0 (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/605.1.15",
			DeviceType:         "mobile",
			ScanTime:           time.Now().Add(-2 * time.Hour),
			ConvertedToSession: false,
			ConvertedToOrder:   false,
		},
		{
			ID:                 3,
			TableID:            3,
			BranchID:           1,
			QRSessionID:        uintPtr(2),
			IPAddress:          "192.168.1.103",
			UserAgent:          "Mozilla/5.0 (iPad; CPU OS 14_6 like Mac OS X) AppleWebKit/605.1.15",
			DeviceType:         "tablet",
			ScanTime:           time.Now().Add(-3 * time.Hour),
			ConvertedToSession: true,
			ConvertedToOrder:   true,
		},
	}

	return s.db.Create(&scans).Error
}

// Helper functions
func uintPtr(i uint) *uint {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func rolePtr(r models.EmployeeRole) *models.EmployeeRole {
	return &r
}

func getTableLocation(tableNum int) string {
	locations := []string{
		"Window side",
		"Corner table",
		"Center area",
		"Near entrance",
		"Quiet corner",
		"Garden view",
		"Main hall",
		"Private section",
	}
	return locations[tableNum%len(locations)]
}

// Additional seeding methods for comprehensive data

func (s *Seeder) SeedAdditionalData() error {
	log.Println("Seeding additional data...")

	if err := s.seedQRCartItems(); err != nil {
		return err
	}
	if err := s.seedQRCodeScans(); err != nil {
		return err
	}
	if err := s.seedMoreMenuItems(); err != nil {
		return err
	}
	if err := s.seedMoreOrders(); err != nil {
		return err
	}

	return nil
}

func (s *Seeder) seedMoreMenuItems() error {
	log.Println("Seeding additional menu items...")

	additionalItems := []models.MenuItem{
		// More Indian Curries
		{
			ID:           8,
			BranchID:     1,
			CategoryID:   uintPtr(5),
			Name:         "Dal Makhani",
			Description:  "Creamy black lentils cooked overnight",
			Price:        320.00,
			CostPrice:    180.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    2,
			PrepTime:     30,
			ImageURL:     "https://example.com/images/dal-makhani.jpg",
			Ingredients:  "Black lentils, Cream, Butter, Tomatoes, Spices",
			Allergens:    "Dairy",
			SortOrder:    1,
		},
		{
			ID:           9,
			BranchID:     1,
			CategoryID:   uintPtr(5),
			Name:         "Fish Curry",
			Description:  "Traditional Kerala fish curry with coconut",
			Price:        420.00,
			CostPrice:    220.00,
			Available:    true,
			IsVegetarian: false,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    4,
			PrepTime:     25,
			ImageURL:     "https://example.com/images/fish-curry.jpg",
			Ingredients:  "Fish, Coconut milk, Curry leaves, Spices",
			Allergens:    "Fish",
			SortOrder:    2,
		},
		// More Appetizers
		{
			ID:           10,
			BranchID:     1,
			CategoryID:   uintPtr(1),
			Name:         "Mutton Seekh Kebab",
			Description:  "Spiced minced mutton grilled on skewers",
			Price:        350.00,
			CostPrice:    200.00,
			Available:    true,
			IsVegetarian: false,
			IsVegan:      false,
			IsGlutenFree: true,
			Spiciness:    3,
			PrepTime:     25,
			ImageURL:     "https://example.com/images/seekh-kebab.jpg",
			Ingredients:  "Mutton, Onions, Spices, Herbs",
			Allergens:    "",
			SortOrder:    3,
		},
		// More Beverages
		{
			ID:           11,
			BranchID:     1,
			CategoryID:   uintPtr(4),
			Name:         "Fresh Lime Water",
			Description:  "Refreshing lime juice with mint",
			Price:        80.00,
			CostPrice:    30.00,
			Available:    true,
			IsVegetarian: true,
			IsVegan:      true,
			IsGlutenFree: true,
			Spiciness:    0,
			PrepTime:     3,
			ImageURL:     "https://example.com/images/lime-water.jpg",
			Ingredients:  "Lime, Water, Mint, Salt/Sugar",
			Allergens:    "",
			SortOrder:    3,
		},
	}

	return s.db.Create(&additionalItems).Error
}

func (s *Seeder) seedMoreOrders() error {
	log.Println("Seeding additional orders...")

	moreOrders := []models.Order{
		{
			ID:            3,
			OrderNumber:   "ORD-2025-003",
			TableID:       uintPtr(2),
			BranchID:      1,
			UserID:        uintPtr(4),
			CustomerName:  "Priya Menon",
			CustomerPhone: "+91-9876543241",
			OrderType:     models.OrderTypeTakeaway,
			OrderSource:   models.OrderSourceStaff,
			Status:        models.OrderReady,
			PaymentStatus: models.PaymentPaid,
			IsQROrder:     false,
			Subtotal:      540.00,
			TaxAmount:     97.20,
			Total:         637.20,
			Notes:         "Takeaway order",
			EstimatedTime: 20,
		},
		{
			ID:               4,
			OrderNumber:      "ORD-2025-004",
			TableID:          uintPtr(4),
			BranchID:         1,
			CustomerName:     "QR Customer",
			CustomerPhone:    "+91-9876543252",
			OrderType:        models.OrderTypeDineIn,
			OrderSource:      models.OrderSourceQR,
			Status:           models.OrderPending,
			PaymentStatus:    models.PaymentPending,
			IsQROrder:        true,
			QRSessionID:      uintPtr(2),
			Subtotal:         400.00,
			TaxAmount:        72.00,
			ServiceCharge:    20.00,
			Total:            492.00,
			Notes:            "QR Order from mobile",
			EstimatedTime:    25,
			AssignedWaiterID: uintPtr(4),
		},
	}

	if err := s.db.Create(&moreOrders).Error; err != nil {
		return err
	}

	// Additional order items
	moreOrderItems := []models.OrderItem{
		{
			ID:         7,
			OrderID:    3,
			MenuItemID: 8, // Dal Makhani
			Quantity:   1,
			UnitPrice:  320.00,
			TotalPrice: 320.00,
			Status:     models.OrderItemReady,
		},
		{
			ID:         8,
			OrderID:    3,
			MenuItemID: 7, // Masala Chai
			Quantity:   2,
			UnitPrice:  60.00,
			TotalPrice: 120.00,
			Status:     models.OrderItemReady,
		},
		{
			ID:         9,
			OrderID:    3,
			MenuItemID: 2, // Vegetable Samosa
			Quantity:   1,
			UnitPrice:  120.00,
			TotalPrice: 120.00,
			Status:     models.OrderItemReady,
		},
		{
			ID:         10,
			OrderID:    4,
			MenuItemID: 9, // Fish Curry
			Quantity:   1,
			UnitPrice:  420.00,
			TotalPrice: 420.00,
			Status:     models.OrderItemPending,
		},
	}

	if err := s.db.Create(&moreOrderItems).Error; err != nil {
		return err
	}

	// Additional payment
	morePayments := []models.Payment{
		{
			ID:          3,
			OrderID:     3,
			Amount:      637.20,
			Method:      models.PaymentCash,
			Status:      models.PaymentPaid,
			Reference:   "Cash payment for takeaway",
			ProcessedBy: uintPtr(4),
		},
	}

	return s.db.Create(&morePayments).Error
}

// Clear all data (useful for re-seeding)
func (s *Seeder) ClearAllData() error {
	log.Println("Clearing all data...")

	// Delete in reverse order of dependencies
	tables := []interface{}{
		&models.QRCodeScan{},
		&models.QRCartItem{},
		&models.Notification{},
		&models.Payment{},
		&models.OrderItem{},
		&models.Order{},
		&models.QRSession{},
		&models.Reservation{},
		&models.Inventory{},
		&models.MenuItem{},
		&models.MenuCategory{},
		&models.Table{},
		&models.Customer{},
		&models.Supplier{},
		&models.Branch{},
		&models.User{},
		&models.Restaurant{},
	}

	for _, table := range tables {
		if err := s.db.Unscoped().Where("1 = 1").Delete(table).Error; err != nil {
			return err
		}
	}

	log.Println("All data cleared successfully!")
	return nil
}
