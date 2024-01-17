resource "epilot-order_order" "my_order" {
  expires_at   = "2022-12-12T15:05:30.467Z"
  order_number = "...my_order_number..."
  source_type  = "manual"
  status       = "draft"
}