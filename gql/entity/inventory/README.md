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
      lot: "as-123-453",
      name: "test_item",
      origin: "test_origin",
      price: 23.65,
      rsCustomerID: "a571181c-06c3-4436-a79d-21580cef1086",
      salePrice: 23.45,
      sku: "test-sku"
      soldWeight: 23.54,
      timestamp: 1539222685400,
      totalWeight: 92.45,
      upc: "102345678912",
      wasteWeight: 45.56,
    )
    {
      _id,
      itemID,
      dateArrived,
      dateSold,
      deviceID,
      donateWeight,
      lot,
      name,
      origin,
      price,
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

  Sample Output:

  ```JSON
  {
    "data": {
      "InventoryInsert": {
        "_id": "5bec63858a9026a28af5b859",
        "dateArrived": 1539222685400,
        "dateSold": 1539222685400,
        "deviceID": "5d79d6f6-3181-4fec-a474-0a5b0020c6cb",
        "donateWeight": 23.23,
        "itemID": "c2167f7a-1eeb-4c6e-8605-6456dbccc2a7",
        "lot": "as-123-453",
        "name": "test_item",
        "origin": "test_origin",
        "price": 23.65,
        "rsCustomerID": "a571181c-06c3-4436-a79d-21580cef1086",
        "salePrice": 23.45,
        "sku": "test-sku",
        "soldWeight": 23.54,
        "timestamp": 1539222685400,
        "totalWeight": 92.45,
        "upc": "102345678912",
        "wasteWeight": 45.56
      }
    }
  }
  ```

* #### InventoryDelete
  ```graphql
  mutation{
    InventoryDelete(
      itemID: "c2167f7a-1eeb-4c6e-8605-6456dbccc2a7",
    ){
        deletedCount
    }
  }
  ```

  Sample Output:

  ```JSON
  {
    "data": {
      "InventoryDelete": {
        "deletedCount":1
      }
    }
  }
  ```

* #### InventoryUpdate
  ```graphql
  mutation{
    InventoryUpdate(
      filter: {
        itemID: "c2167f7a-1eeb-4c6e-8605-6456dbccc2a7"
      },
      update: {
        origin: "new-origin"
      },
    ){
        matchedCount, modifiedCount
    }
  }
  ```

  Sample Output:

  ```JSON
  {
    "data": {
      "InventoryUpdate": {
        "matchedCount": 1,
        "modifiedCount": 1
      }
    }
  }
  ```

### Queries

* #### InventoryQueryItem
  ```graphql
  {
    InventoryQueryItem(
      itemID: "d06e734e-7b6c-40ce-b3a6-2ca4537ebdd7",
    ){
      _id,
      itemID,
      dateArrived,
      dateSold,
      deviceID,
      donateWeight,
      lot,
      name,
      origin,
      price,
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

  Sample Output:

  ```JSON
  {
    "data": {
      "InventoryQueryItem": [
        {
          "_id": "5bec5b348a9026a28af5b857",
          "dateArrived": 1542216500,
          "dateSold": 0,
          "deviceID": "cba45c98-dadf-4f7d-b29d-d2c6a83cb371",
          "donateWeight": 0,
          "itemID": "d06e734e-7b6c-40ce-b3a6-2ca4537ebdd7",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "some-origin2",
          "price": 13.4,
          "rsCustomerID": "6441cd8f-8324-4c60-830b-3359cc293d36",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542216500,
          "totalWeight": 300,
          "upc": "123456789012",
          "wasteWeight": 12
        }
      ]
    }
  }
  ```

* #### InventoryQueryTimestamp
  ```graphql
  {
    InventoryQueryTimestamp(
      start: 1542219382824,
      end: 1542219334742
    ){
      _id,
      itemID,
      dateArrived,
      dateSold,
      deviceID,
      donateWeight,
      lot,
      name,
      origin,
      price,
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

  Sample Output:

  ```JSON
  {
    "data": {
      "InventoryQueryTimestamp": [
        {
          "_id": "5bec5b348a9026a28af5b857",
          "dateArrived": 1542216500,
          "dateSold": 0,
          "deviceID": "cba45c98-dadf-4f7d-b29d-d2c6a83cb371",
          "donateWeight": 0,
          "itemID": "d06e734e-7b6c-40ce-b3a6-2ca4537ebdd7",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "some-origin2",
          "price": 13.4,
          "rsCustomerID": "6441cd8f-8324-4c60-830b-3359cc293d36",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542216500,
          "totalWeight": 300,
          "upc": "123456789012",
          "wasteWeight": 12
        },
        {
          "_id": "5bec5a6c8a9026a28af5b855",
          "dateArrived": 1542216300,
          "dateSold": 0,
          "deviceID": "ddbe9846-bae8-45c2-96bf-a45de3871c52",
          "donateWeight": 0,
          "itemID": "a8eb4d7d-d5f2-49bb-bc7c-3451c1b9b561",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "test-origin",
          "price": 13.4,
          "rsCustomerID": "e5c7f8c7-e947-495a-934f-dffd750c4076",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542216300,
          "totalWeight": 300,
          "upc": "test-upc",
          "wasteWeight": 12
        },
        {
          "_id": "5bec35218a9026a28af5b854",
          "dateArrived": 1542206880,
          "dateSold": 0,
          "deviceID": "9f08c6b1-c86b-486c-956f-5644b649665d",
          "donateWeight": 0,
          "itemID": "0312473c-aa01-4b42-8ead-4ec631cb74ce",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "test-origin",
          "price": 13.4,
          "rsCustomerID": "c5f1940d-ba24-4be1-8a1b-6bd6398a3118",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542206880,
          "totalWeight": 300,
          "upc": "test-upc",
          "wasteWeight": 12
        }
      ]
    }
  }
  ```

* #### InventoryQueryCount
  ```graphql
  {
    InventoryQueryCount(
      count: 10
    ){
      _id,
      itemID,
      dateArrived,
      dateSold,
      deviceID,
      donateWeight,
      lot,
      name,
      origin,
      price,
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

  Sample Output:

  ```JSON
  {
    "data": {
      "InventoryQueryCount": [
        {
          "_id": "5bec5b348a9026a28af5b857",
          "dateArrived": 1542216500,
          "dateSold": 0,
          "deviceID": "cba45c98-dadf-4f7d-b29d-d2c6a83cb371",
          "donateWeight": 0,
          "itemID": "d06e734e-7b6c-40ce-b3a6-2ca4537ebdd7",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "test-origin",
          "price": 13.4,
          "rsCustomerID": "6441cd8f-8324-4c60-830b-3359cc293d36",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542216500,
          "totalWeight": 300,
          "upc": "123456789012",
          "wasteWeight": 12
        },
        {
          "_id": "5bec5a6c8a9026a28af5b855",
          "dateArrived": 1542216300,
          "dateSold": 0,
          "deviceID": "ddbe9846-bae8-45c2-96bf-a45de3871c52",
          "donateWeight": 0,
          "itemID": "a8eb4d7d-d5f2-49bb-bc7c-3451c1b9b561",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "test-origin",
          "price": 13.4,
          "rsCustomerID": "e5c7f8c7-e947-495a-934f-dffd750c4076",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542216300,
          "totalWeight": 300,
          "upc": "test-upc",
          "wasteWeight": 12
        },
        {
          "_id": "5bec35218a9026a28af5b854",
          "dateArrived": 1542206880,
          "dateSold": 0,
          "deviceID": "9f08c6b1-c86b-486c-956f-5644b649665d",
          "donateWeight": 0,
          "itemID": "0312473c-aa01-4b42-8ead-4ec631cb74ce",
          "lot": "test-lot",
          "name": "test-name",
          "origin": "test-origin",
          "price": 13.4,
          "rsCustomerID": "c5f1940d-ba24-4be1-8a1b-6bd6398a3118",
          "salePrice": 12.23,
          "sku": "test-sku",
          "soldWeight": 0,
          "timestamp": 1542206880,
          "totalWeight": 300,
          "upc": "test-upc",
          "wasteWeight": 12
        }
      ]
    }
  }
  ```
