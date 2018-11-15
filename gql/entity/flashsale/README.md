## Usage examples
---

### Mutations

* #### InventoryInsert
  ```graphql
  mutation{
    SaleInsert(
      saleID: "cdc7a14c-19e3-488e-8c4e-22d91fd42ef1",
      items: [
    	{
    	  itemID: "39322979-d33b-4504-ba90-f2e427bdd72b",
    	  weight: 12.40,
    	  lot: "test-lot",
    	  upc: "test-upc",
    	  sku: "test-sku"
    	}
      ]
      timestamp: 1539222685400,
    )
    {
     timestamp
    }
  }
  ```
