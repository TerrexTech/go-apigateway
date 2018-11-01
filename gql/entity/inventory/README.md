## Usage examples
---

### Mutations

* #### InventoryInsert
  ```graphql
  mutation{
    InventoryInsert(
      itemID: "c2167f7a-1eeb-4c6e-8605-6456dbccc2a7",
      dateArrived: 1539222685400,
      dateSold: 1539222685400,
      deviceID: "5d79d6f6-3181-4fec-a474-0a5b0020c6cb",
      donateWeight: 23.23,
      expiryDate: 15392226996949,
      lot: "as-123-453",
      name: "test_item",
      origin: "test_origin",
      price: 23.65,
      quantity: 3
      rsCustomerID: "a571181c-06c3-4436-a79d-21580cef1086",
      salePrice: 23.45,
      sku: "test-sku"
      soldWeight: 23.54,
      timestamp: 1539222685400,
      totalWeight: 92.45,
      upc: 102345678912,
      wasteWeight: 45.56,
    )
    {
      _id,
      itemID,
      dateArrived,
      dateSold,
      deviceID,
      donateWeight,
      expiryDate,
      lot,
      name,
      origin,
      price,
      quantity,
      rsCustomerID,
      salePrice,
      sku,
      soldWeight,
      timestamp,
      totalWeight,
      upc,
      wasteWeight
    }
  }
  ```