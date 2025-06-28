package models


type OrderItem struct {
    ID         uint   `gorm:"primaryKey"`
    OrderID    uint
    MenuItemID uint
    Quantity   int
    Price      float64
    Notes      string
    MenuItem   MenuItem
}
